package bills

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) UploadBill(w http.ResponseWriter, r *http.Request) {
	// 1. Parse Multipart Form
	// Limit upload size to 10MB
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large or invalid form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// 2. Parse Other Fields
	userIDStr := r.FormValue("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		name = header.Filename // Default to filename if not provided
	}

	var categoryID *int
	if catStr := r.FormValue("category_id"); catStr != "" {
		if id, err := strconv.Atoi(catStr); err == nil {
			categoryID = &id
		}
	}

	var amount *float64
	if amtStr := r.FormValue("amount"); amtStr != "" {
		if val, err := strconv.ParseFloat(amtStr, 64); err == nil {
			amount = &val
		}
	}

	req := CreateBillRequest{
		UserID:     userID,
		CategoryID: categoryID,
		Name:       name,
		Amount:     amount,
	}

	// 3. Call Service
	bill, err := h.service.UploadBill(r.Context(), file, header.Filename, header.Header.Get("Content-Type"), req)
	if err != nil {
		http.Error(w, "Failed to upload bill: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bill)
}

func (h *Handler) ListBills(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	if userIdStr == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	bills, err := h.service.ListUserBills(r.Context(), userID)
	if err != nil {
		http.Error(w, "Failed to list bills", http.StatusInternalServerError)
		return
	}

	if bills == nil {
		bills = []*Bill{}
	}
	json.NewEncoder(w).Encode(bills)
}

func (h *Handler) DownloadBill(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.URL.Query().Get("user_id")
	if userIdStr == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Missing bill id", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid bill id", http.StatusBadRequest)
		return
	}

	url, err := h.service.GetBillDownloadURL(r.Context(), id, userID)
	if err != nil {
		if err.Error() == "unauthorized access to bill" {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else {
			http.Error(w, "Failed to get download URL", http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
