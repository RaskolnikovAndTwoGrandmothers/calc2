package pkg

import (
    "os"
)

func GetOperationTime(operation string) int {
    switch operation {
    case "addition":
        return getEnv("TIME_ADDITION_MS", 100) 
    case "subtraction":
        return getEnv("TIME_SUBTRACTION_MS", 100)
    case "multiplication":
        return getEnv("TIME_MULTIPLICATIONS_MS", 100)
    case "division":
        return getEnv("TIME_DIVISIONS_MS", 100)
    }
    return 100 
}

func getEnv(key string, defaultValue int) int {
    if value, exists := os.LookupEnv(key); exists {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }
    return defaultValue
}
