package main

import (
    "log"
    "net/http"
    "os"

    "calc_service/internal/orchestrator"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    router := orchestrator.NewRouter()
    log.Printf("Сервер запущен на порту %s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        log.Fatal(err)
    }
}
