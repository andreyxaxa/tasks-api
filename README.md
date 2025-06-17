# HTTP API для работы с задачами.

Длительность выполнения каждой задачи = 3-5 минут.
- При создании, задача получает статус "Created";
- Когда задача начинает "работать" (в данном случае, каждая созданная задача сразу же запускается), она получает статус "In progress";
- Когда задача отработала(прошел ее Duration, типа мы дождались ответа 3-5 минут), она получает статус "Done". Также записывается захардкоженный результат - "Completed successfully".

## Функционал:
1. Добавление новой задачи (POST /v1/tasks)

Запрос:
```
curl -X POST http://localhost:8080/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "name": "fix bug #12"
}'
```
Ответ (генерируется рандомный uuid, у вас будет другой):
```
{
  "created":"2025-06-17T00:44:45.6640525+03:00",
  "duration":240000000000,
  "id":"94c56cdf-7001-4447-816b-75bba15dd461",
  "name":"fix bug #12",
  "status":"Created"
}

```
2. Получение даты создания, времени работы и статуса задачи (GET /v1/tasks/{id})

Запрос:
```
curl -X GET http://localhost:8080/v1/tasks/94c56cdf-7001-4447-816b-75bba15dd461
```
Ответ:
```
{
  "created_at":"2025-06-17T00:44:45.6640525+03:00",
  "in_work":45653628600,
  "status":"In progress"
}

```
3. Удаление задачи по ID (POST /v1/tasks/{id})

Запрос:
```
curl -X POST http://localhost:8080/v1/tasks/94c56cdf-7001-4447-816b-75bba15dd461
```
Ответ:
Status OK

## Детали 

- У сервиса есть конфиг - [config/config.go](https://github.com/andreyxaxa/tasks-api/blob/main/config/config.go); Читается из `.env` файла. В рамках тестового задания .env прямо в репозитории, очевидно в проде он должен быть заигнорен.
- В слое хэндлеров применяется версионирование - [internal/controller/http/v1](https://github.com/andreyxaxa/tasks-api/tree/main/internal/controller/http/v1).
  Для версии v2 нужно будет просто добавить папку `http/v2` с таким же содержимым, в файле [internal/controller/http/router.go](https://github.com/andreyxaxa/tasks-api/blob/main/internal/controller/http/router.go) добавить строку:
```go
{
    v1.NewTasksRoutes(apiV1Group, t)
}

{
    v2.NewTasksRoutes(apiV1Group, t)
}
```
- Используется dependency injection - [internal/controller/http/v1/controller.go](https://github.com/andreyxaxa/tasks-api/blob/main/internal/controller/http/v1/controller.go).
- Реализован graceful shutdown - [internal/app/app.go](https://github.com/andreyxaxa/tasks-api/blob/main/internal/app/app.go).
- Удобная и гибкая конфигурация HTTP сервера - [pkg/httpserver/options.go](https://github.com/andreyxaxa/tasks-api/blob/main/pkg/httpserver/options.go).
  Позволяет конфигурировать сервер в конструкторе таким образом:
```go
httpServer := httpserver.New(httpserver.Port(cfg.HTTP.Port))
```

## Запуск

### Local:
Клонируем репозиторий, выполняем:
```
make run
```

### Docker:
Клонируем репозиторий, выполняем:
```
make compose-up
```

## Прочие `make` команды
Зависимости:
```
make deps
```
docker compose down:
```
make compose-down
```
