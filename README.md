# calc2
становка
 Склонируйте репозиторий:
git clone https://github.com/RaskolnikovAndTwoGrandmothers/calc2.git
для агента:
export TIME_ADDITION_MS=100
export TIME_SUBTRACTION_MS=100
export TIME_MULTIPLICATIONS_MS=100
export TIME_DIVISIONS_MS=100
export COMPUTING_POWER=4 # Количество горутин для агента
go mod tidy
запуск: go run ./cmd/calc_service/...
API Эндпоинты
1. Добавление вычисления арифметического выражения
curl --location 'localhost/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
Ответ:
json
{
    "id": <уникальный идентификатор выражения>
}
Коды ответа:
201 - выражение принято для вычисления
422 - невалидные данные
500 - что-то пошло не так
2. Получение списка выражений
Запрос:
curl --location 'localhost/api/v1/expressions'
Ответ:
json
{
    "expressions": [
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
    ]
}
Коды ответа:
200 - успешно получен список выражений
500 - что-то пошло не так
3. Получение выражения по его идентификатору
Запрос:
curl --location 'localhost/api/v1/expressions/:id'
Ответ:
{
    "expression":
        {
            "id": <идентификатор выражения>,
            "status": <статус вычисления выражения>,
            "result": <результат выражения>
        }
}
Коды ответа:
200 - успешно получено выражение
404 - нет такого выражения
500 - что-то пошло не так
4. Получение задачи для выполнения
Запрос:
curl --location 'localhost/internal/task'
Ответ:
json
{
    "task":
        {
            "id": <идентификатор задачи>,
            "arg1": <имя первого аргумента>,
            "arg2": <имя второго аргумента>,
            "operation": <операция>,
            "operation_time": <время выполнения операции>
        }
}
Коды ответа:
200 - успешно получена задача
404 - нет задачи
500 - что-то пошло не так
5. Прием результата обработки данных
Запрос:
curl --location 'localhost/internal/task' \
--header 'Content-Type: application/json' \
--data '{
  "id": <идентификатор задачи>,
  "result": <результат>
}'
Коды ответа:
200 - успешно записан результат
404 - нет такой задачи
422 - невалидные данные
500 - что-то пошло не так
