package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)
func apiHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(map[string]interface{}{
    "name":"Lesan Kiptoo",
    "title":"Computer Technologist - Maseno University",
    "role":"Branch Supervisor at Urban Waters, Kilimani",
    "stack":[]string{"Go 1.26.4","Python","Termux","aarch64"},
    "location":"Nairobi, Kenya",
    "live":"https://lesan24.github.io/my-first-folder/",
  })
}
func cvHandler(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(map[string]interface{}{
    "name":"Lesan Kiptoo",
    "email":"lesankiptoo24@gmail.com",
    "phone":"0727 919 676",
    "cv_pdf":"/Lesan-Kiptoo-CV.pdf",
    "skills":[]string{"Go","Python","ICT & Cloud Security","Sales & Stock Management"},
  })
}
func main(){
  fs:=http.FileServer(http.Dir("."))
  http.Handle("/",fs)
  http.HandleFunc("/api",apiHandler)
  http.HandleFunc("/api/cv",cvHandler)
  fmt.Println("🚀 LIVE: http://localhost:8080")
  fmt.Println("   /api -> portfolio JSON")
  fmt.Println("   /api/cv -> CV JSON")
  http.ListenAndServe(":8080",nil)
}
