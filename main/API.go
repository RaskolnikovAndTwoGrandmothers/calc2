package orchestrator

import (
    "net/http"

    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    
    router.HandleFunc("/api/v1/calculate", CalculateExpression).Methods("POST")
    router.HandleFunc("/api/v1/expressions", GetExpressions).Methods("GET")
    router.HandleFunc("/api/v1/expressions/{id}", GetExpressionByID).Methods("GET")
    router.HandleFunc("/internal/task", GetTask).Methods("GET")
    router.HandleFunc("/internal/task", PostResult).Methods("POST")

    return router
}
