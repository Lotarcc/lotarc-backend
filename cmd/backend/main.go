package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/joho/godotenv"
    "github.com/go-chi/chi"
)

func main() {
    err := godotenv.Load()
    if err != nil {
    log.Fatal("Error loading .env file")
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    r := chi.NewRouter()
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello World"))
    })

    srv := &http.Server{
        Handler:      r,
        Addr:         fmt.Sprintf("0.0.0.0:%s", port),
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Printf("Server listening on port %s", port)
    log.Fatal(srv.ListenAndServe())
}
