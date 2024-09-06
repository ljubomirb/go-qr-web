package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Serve the HTML file on GET request
		http.ServeFile(w, r, "index.html")
		return
	}

	if r.Method == http.MethodPost {
		// Handle the QR code generation on POST request
		text := r.FormValue("text")
		if text == "" {
			http.Error(w, "Text is required", http.StatusBadRequest)
			return
		}

		// Create a QR code with the provided text
		qrCode, err := qr.Encode(text, qr.M, qr.Auto)
		if err != nil {
			http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
			return
		}

		// Scale the barcode to 150x150 pixels
		qrCode, err = barcode.Scale(qrCode, 150, 150)
		if err != nil {
			http.Error(w, "Failed to scale QR code", http.StatusInternalServerError)
			return
		}

		// Send the QR code image as a PNG response
		w.Header().Set("Content-Type", "image/png")
		png.Encode(w, qrCode)
	}
}

func main() {
	// Default port
	port := "8005"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/", handler)

	fmt.Printf("Starting server at port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
