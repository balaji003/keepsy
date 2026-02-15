import os
import sys
import json
import google.generativeai as genai
from dotenv import load_dotenv
from pathlib import Path

# Load environment variables
load_dotenv()

# Configure the Gemini API
API_KEY = os.getenv("GEMINI_API_KEY")
if not API_KEY:
    print("Error: GEMINI_API_KEY not found in environment variables.")
    print("Please set it in your .env file or export it.")
    sys.exit(1)

genai.configure(api_key=API_KEY)

# Define the model
model = genai.GenerativeModel('gemini-1.5-flash')

def extract_bill_data(image_path):
    """
    Extracts structured data from a bill/invoice image using Gemini 1.5 Flash.
    """
    if not os.path.exists(image_path):
        print(f"Error: Image not found at {image_path}")
        return

    # Prepare the prompt
    prompt = """
    You are an expert invoice extraction AI. 
    Analyze the provided image of a bill/invoice and extract the following information in strict JSON format:
    
    {
      "vendor_name": "Name of the store or vendor",
      "invoice_date": "Date of the invoice (YYYY-MM-DD format if possible)",
      "invoice_number": "The invoice or bill number",
      "total_amount": 0.00,
      "currency": "Currency symbol or code (e.g. INR, USD)",
      "line_items": [
        {
          "description": "Item name/description",
          "quantity": 1,
          "unit_price": 0.00,
          "total_price": 0.00
        }
      ]
    }
    
    If a field is missing, use null. ensure the output is valid JSON text only, no markdown formatting.
    """

    try:
        # Load the image
        img_data = Path(image_path).read_bytes()
        
        # transform data to be compatible with gemini
        image_part = {
            "mime_type": "image/jpeg", # Assuming jpeg/png, gemini handles both usually with this or auto-detect
            "data": img_data
        }
        
        # For simplicity in this script using the file path directly if supported or bytes
        # The python SDK supports passing the path directly for some file types or usually PIL images
        # Let's use the safer way: read bytes and pass as blob.
        
        # Actually simplest way with current SDK:
        # Generate content
        response = model.generate_content([prompt, {"mime_type": "image/jpeg", "data": img_data}])
        
        # Clean up response text (remove markdown code blocks if present)
        text = response.text.strip()
        if text.startswith("```json"):
            text = text[7:]
        if text.endswith("```"):
            text = text[:-3]
            
        print(text)

    except Exception as e:
        print(f"Error during extraction: {e}")

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python scan_bill.py <path_to_bill_image>")
        sys.exit(1)
        
    image_path = sys.argv[1]
    extract_bill_data(image_path)
