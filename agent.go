package agent

import (
    "bytes"
    "encoding/json"
    "log"
    "net/http"
    "os"
    "sync"
    "time"

    "calc_service/pkg"
)

var (
    wg sync.WaitGroup
)

func StartAgent() {
    numWorkers := getNumWorkers()
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go worker(i)
    }
    wg.Wait()
}

func worker(workerID int) {
    defer wg.Done()

    for {
        task, err := getTask()
        if err != nil {
            log.Printf("Worker %d: Ошибка получения задачи: %s", workerID, err)
            time.Sleep(1 * time.Second) // Задержка на случай ошибки
            continue
        }

        log.Printf("Worker %d: Выполнение задачи %d: %s", workerID, task.ID, task.Operation)
        result, err := performOperation(task)
        if err != nil {
            log.Printf("Worker %d: Ошибка при выполнении задачи %d: %s", workerID, task.ID, err)
            continue
        }

        err = sendResult(task.ID, result)
        if err != nil {
            log.Printf("Worker %d: Ошибка отправки результата задачи %d: %s", workerID, task.ID, err)
        } else {
            log.Printf("Worker %d: Успешно отправлен результат задачи %d: %f", workerID, task.ID, result)
        }
    }
}

func getTask() (*pkg.Task, error) {
    resp, err := http.Get("http://localhost/internal/task")
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusNotFound {
        return nil, nil // Нет активных задач
    } else if resp.StatusCode != http.StatusOK {
        return nil, err
    }

    var task pkg.Task
    if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
        return nil, err
    }
    return &task, nil
}

func performOperation(task *pkg.Task) (float64, error) {
    var result float64
    var err error

    switch task.Operation {
    case "+":
        result, err = pkg.EvaluateExpression(task.Arg1 + "+" + task.Arg2)
    case "-":
        result, err = pkg.EvaluateExpression(task.Arg1 + "-" + task.Arg2)
    case "*":
        result, err = pkg.EvaluateExpression(task.Arg1 + "*" + task.Arg2)
    case "/":
        result, err = pkg.EvaluateExpression(task.Arg1 + "/" + task.Arg2)
    default:
        return 0, pkg.ErrInvalidOperation
    }

    // Задержка на выполнение операции с учетом времени, установленного в переменных окружения
    time.Sleep(time.Duration(pkg.GetOperationTime(task.Operation)) * time.Millisecond)
    return result, err
}

func sendResult(taskID int, result float64) error {
    reqBody := pkg.Result{
        ID:     taskID,
        Result: result,
    }
    body, err := json.Marshal(reqBody)
    if err != nil {
        return err
    }

    resp, err := http.Post("http://localhost/internal/task", "application/json", bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return pkg.ErrUnableToSendResult
    }
    return nil
}

func getNumWorkers() int {
    if numWorkers := os.Getenv("COMPUTING_POWER"); numWorkers != "" {
        if val, err := strconv.Atoi(numWorkers); err == nil {
            return val
        }
    }
    return 4 //умолчание
}
