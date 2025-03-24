func TestCalculator(t *testing.T) {
    tests := []struct {
        input    string
        expected float64
        hasError bool
    }{
        {"2+2", 4, false},
        {"2/0", 0, true},
        {"2++2", 0, true},
    }
    
    for _, test := range tests {
        result, err := Calculate(test.input)
        if test.hasError && err == nil {
            t.Errorf("Expected error for input %s", test.input)
        }
        if !test.hasError && result != test.expected {
            t.Errorf("For input %s expected %f but got %f", test.input, test.expected, result)
        }
    }
}
