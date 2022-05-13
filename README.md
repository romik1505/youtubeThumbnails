# Youtube Thumbnails
**Сервис для загрузки превью с ютуба**

Статус последнего Деплоя:<br>
<img src="https://github.com/romik1505/youtubeThumbnails/workflows/Main-Workflow/badge.svg?branch=main"><br>

**Функциональное тестирование сервиса**
1. Подготовьте исходные данные с ссылками
2. Запишите ссылки в файл с корнем проекта (по умолчанию 1.txt)
3. Загрузите зависимости проекта `make bin-depth`
4. Произведите миграцию базы данных  `make db:up`
5. Запустите сервис `make run`
6. Запустите клиент `go run ./cmd/client/main.go -f input_file -o output directory -async`

**Параметры клиента**
`-f input_file`  -- файл со ссылками на видео записанные с новой строки
`-o directory`  -- директория с картинками (превью видео с youtube)
`-async`  -- флаг асинхронности

**Параметры по умолчанию:**

    -f 1.txt 
    -o output
    -async false


**Структура проекта**

    ├── 1.txt
    ├── api
    │   └── thumbnails
    │       └── thumbnails.proto            Прото файлы сервиса
    ├── cmd
    │   ├── client
    │   │   └── main.go                     Клиент для взаимодействия с сервисом
    │   └── server
    │       └── main.go                     
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── app
    │   │   ├── config
    │   │   │   ├── config.go               Загружает локальный конфиг
    │   │   │   └── sqlite.go               Подключение к БД
    │   │   ├── model
    │   │   │   └── thumbnail.go            Модели для БД(DTO)
    │   │   ├── server
    │   │   │   └── server.go               Конфигурация сервера
    │   │   ├── service
    │   │   │   ├── get.go                  Логика загрузки превью 
    │   │   │   ├── get_test.go
    │   │   │   └── service.go
    │   │   └── store
    │   │       ├── nullType.go             Типы для взаимодействия с БД
    │   │       ├── store.go                
    │   │       └── thumbnail               Репозиторий для работы с БД
    │   │           ├── thumbnail.go        
    │   │           └── thumbnail_test.go
    │   └── pkg
    │       └── mock
    │           ├── repository
    │           │   └── mock_rep.go         Моки репозитория
    │           └── service
    │               └── mock_service.go     Моки сервиса
    ├── Makefile
    ├── migrations                          Файлы миграций
    │   └── 20220512002617_thumbnail.sql
    ├── pkg
    │   └── api
    │       └── thumbnails
    │           ├── thumbnails_grpc.pb.go
    │           ├── thumbnails.pb.go
    │           └── thumbnails.swagger.json
    └── README.md
