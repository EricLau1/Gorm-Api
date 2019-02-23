package main

import (
  "fmt"
  "log"
  "net/http"
  _"api/models"
  "api/routes"
)

func main() {
  //models.AutoMigrations()
  listen(9000)
}

func listen(p int) {
  fmt.Printf("\n\nListening on port %d...", p)
  port := fmt.Sprintf(":%d", p)
  r := routes.NewRouter()
  log.Fatal(http.ListenAndServe(port, r))
}
