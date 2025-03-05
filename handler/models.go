package pkg

type CalculateRequest struct {
    Expression string `json:"expression"`
}

type Expression struct {
    ID     int    `json:"id"`
    Status string `json:"status"`
    Input  string `json:"input"`
    Result float64 `json:"result,omitempty"`
}

type Result struct {
    ID     int     `json:"id"`
    Result float64 `json:"result"`
}
