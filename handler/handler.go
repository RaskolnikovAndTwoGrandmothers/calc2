package orchestrator

import (
    "encoding/json"
    "net/http"
    "sync/atomic"

    "calc_service/pkg"
)

var (
    expressions = make(map[int]*pkg.Expression)
    idCounter    int32
)

func CalculateExpression(w http.ResponseWriter, r *http.Request) {
    var req pkg.CalculateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Недействительные данные", http.StatusUnprocessableEntity)
        return
    }

    atomic.AddInt32(&idCounter, 1)
    id := int(idCounter)

    expression := &pkg.Expression{
        ID:         id,
        Status:     "pending",
        Input:      req.Expression,
    }
    expressions[id] = expression

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func GetExpressions(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(expressions)
}

func GetExpressionByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    
    expression, exists := expressions[id]
    if !exists {
        http.Error(w, "Нет такого выражения", http.StatusNotFound)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(expression)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
}

func PostResult(w http.ResponseWriter, r *http.Request) {
    var result pkg.Result
    if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
        http.Error(w, "Недействительные данные", http.StatusUnprocessableEntity)
        return
    }
}
