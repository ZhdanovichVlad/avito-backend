# avito-backend

Тестовое задание Avito (осень 2024, стажировка Golang): HTTP API для тендеров и предложений.

Это учебный take-home, не коммерческий сервис. Два более поздних задания — `service-podof` и `zhvAvitoPrivat`.

## Что умеет

REST API на Go и PostgreSQL: тендеры, статусы, предложения. Текст задания и OpenAPI лежат в `задание/` и в `README_FROM_AVITO.md`.

Честно про текущее состояние:

- точка входа `backend/cmd/main.go` поднимает ping и тендеры из `backend/internal`
- в дереве есть второй набросок `backend/internal2` (в том числе bids); в `main` он не подключён
- автотестов нет

## Запуск

Нужны Go 1.23+ и PostgreSQL. Переменные — в `.env` в корне (файл в git не хранится):

```
SERVER_ADDRESS=0.0.0.0:8080
POSTGRES_USERNAME=
POSTGRES_PASSWORD=
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=
```

В базе должны быть таблицы `employee`, `organization` и связь `organization_responsible` — это часть условия задания.

```sh
go mod download
go run ./backend/cmd/main.go -env=local
```

Сервер слушает `0.0.0.0:8080`.

Через Docker (порт 8080, как требовало задание):

```sh
docker build -t avito-backend .
docker run --env-file .env -p 8080:8080 avito-backend
```

## Структура

```
backend/cmd/main.go      точка входа
backend/internal/        тендеры, которые реально стартуют из main
backend/internal2/       более полный набросок handlers, в main не используется
задание/                 исходное ТЗ и OpenAPI
```
