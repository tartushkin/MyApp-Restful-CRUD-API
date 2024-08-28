#  MyApp

**MyApp** - это веб-приложение на Go для добавления клиентов в БД. Приложение использует SQLite для хранения клиентов и предоставляет API для взаимодействия с клентами.

## Установка

### Требования

- Go 1.16 или новее
- SQLite3

### Установка

1. Склонируйте репозиторий:
    ```bash
    git clone https://github.com/tartushkin/go_final_project.git
    cd go_final_project
    ```
2. Установите зависимости:
    ```bash
    go mod tidy
    ```

## Настройка

### Инициализация базы данных

Приложение автоматически создаст файл базы данных и необходимые таблицы при первом запуске, если файл базы данных не существует.

## Запуск приложения

### Запуск сервера

Для запуска сервера выполните следующую команду:

```bash

go run main.go
```

Сервер будет доступен по адресу `http://localhost:8080`

### Примеры использования API

Добавление клиента/пользователя
```bash
 curl -X POST "http://localhost:8080/api/task" -H "Content-Type: application/json" -d '{
"Email": "jonwick@ya.ru",
"Phone": "1234567",
"FirstName": "Джон",
"LastName": "Уик"
}'
```

### Получение всех пользователей

```bash
curl -X GET "http://localhost:8080/api/tasks"
```
### Получение пользователей по идентификатору
```bash
curl -X GET "http://localhost:8080/api/task?id=185"
```
### Обновление пользователя
```bash
curl -X PUT "http://localhost:8080/api/task" -H "Content-Type: application/json" -d '{
"id": 185,
"Email": "jonwick@ya.ru",
"Phone": "12345678",
"FirstName": "Джон",
"LastName": "Уик"
}'
```


### Удаление задачи
```bash
curl -X DELETE "http://localhost:8080/api/task?id=185"
```


## Структура проекта
- `main.go` : Главный файл приложения, точка входа сервера.
- `MyDB/`: Пакет для инициализации базы данных.
- `http/`: Пакет с обработчиками API запросов.
- `app/`: Пакет с моделями данных.
- `logic/`: Пакет с логикой.