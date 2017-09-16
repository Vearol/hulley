package main

import (
  "html/template"
  "net/http"
  "path"
  "log"
  "fmt"
)

func main() {
  http.HandleFunc("/admin", renderAdminPanel)

  fmt.Println("You can access admin panel via http://localhost:8080/admin")
  
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderAdminPanel(w http.ResponseWriter, r *http.Request) {
  indexPath := path.Join("templates", "index.html")

  tmpl, err := template.ParseFiles(indexPath)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, nil); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}
