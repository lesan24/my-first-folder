package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	data := map[string]interface{}{
		"name":     "Lesan",
		"role":     "Go & Python Developer from Narok",
		"location": "Narok, Kenya - aarch64 / Savanna Fibre",
		"server":   "Go 1.26.4 on Termux",
		"stack":    []string{"Go", "Python", "Termux", "Git", "GitHub Pages"},
		"endpoints": map[string]string{
			"/":    "portfolio site",
			"/api": "this JSON",
		},
		"status": "live",
	}
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/api", infoHandler)

	fmt.Println("🚀 Lesan's Go API running:")
	fmt.Println("  http://localhost:8080")
	fmt.Println("  http://localhost:8080/api")
	http.ListenAndServe(":8080", nil)
}
