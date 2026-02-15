package main

import (
	"fmt"
	"log"
	"net/http"

	"keepsy-backend/internal/auth"
	"keepsy-backend/internal/bills"
	"keepsy-backend/internal/categories"
	"keepsy-backend/internal/config"
	"keepsy-backend/internal/db"
	"keepsy-backend/internal/products"
	"keepsy-backend/internal/storage"
	"keepsy-backend/internal/users"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	database, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Initialize repositories and handlers
	userRepo := users.NewMySQLRepository(database.Conn)
	userService := users.NewService(userRepo)
	userHandler := users.NewHandler(userService)

	authRepo := auth.NewMySQLRepository(database.Conn)
	authService := auth.NewService(authRepo, userRepo)
	authHandler := auth.NewHandler(authService)

	categoryRepo := categories.NewMySQLRepository(database.Conn)
	categoryService := categories.NewService(categoryRepo)
	categoryHandler := categories.NewHandler(categoryService)

	productRepo := products.NewMySQLRepository(database.Conn)
	productService := products.NewService(productRepo)
	productHandler := products.NewHandler(productService)

	// Storage Service (Local FS)
	// Save to "uploads" directory in current working dir
	// Serve via "/uploads/" path
	storageService, err := storage.NewLocalStorage("./uploads", "http://localhost:8080/uploads")
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	billsRepo := bills.NewMySQLRepository(database.Conn)
	billsService := bills.NewService(billsRepo, storageService)
	billsHandler := bills.NewHandler(billsService)

	mux := http.NewServeMux()

	// Serve static files from uploads directory
	// STRIP /uploads prefix so requests to /uploads/file.jpg go to ./uploads/file.jpg
	fs := http.FileServer(http.Dir("./uploads"))
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", fs))

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Auth Routes
	mux.HandleFunc("POST /auth/register", authHandler.Register)
	mux.HandleFunc("POST /auth/login", authHandler.Login)

	// User Routes
	mux.HandleFunc("GET /users", userHandler.GetUser)
	mux.HandleFunc("POST /users/check", userHandler.CheckUser)

	// Category Routes
	// mux.HandleFunc("POST /categories", categoryHandler.CreateCategory) // Disabled per requirements
	mux.HandleFunc("GET /categories", categoryHandler.ListCategories)

	// Product Routes
	mux.HandleFunc("POST /products", productHandler.CreateProduct)
	mux.HandleFunc("GET /products", productHandler.GetProduct)        // ?id=...
	mux.HandleFunc("GET /products/list", productHandler.ListProducts) // ?user_id=...

	// Bills Routes
	mux.HandleFunc("POST /bills/upload", billsHandler.UploadBill)
	mux.HandleFunc("GET /bills", billsHandler.ListBills)
	mux.HandleFunc("GET /bills/download", billsHandler.DownloadBill)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
