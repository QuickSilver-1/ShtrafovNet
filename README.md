# ShtrafovNet
 
  ![Build Status](https://github.com/QuickSilver-1/ShtrafovNet/actions/workflows/go.yml/badge.svg)

 <h3>Прототип сервиса аукцинов на чистой архитектуре</h3>

 <h2>Структура проекта</h2>

<h2>Запуск</h2>

Запуск сервера
<ul>
<li>Настраиваем env file ---> <code>migrate -path /internal/infrastructure/migrations -database "postgresql://{DB_USER}:{DB_PASS}@{DB_HOST}:{DB_PORT}/{DB_NAME}?sslmode=disable" -verbose up</code> ---><code>cd cmd/ShtrafovNet</code>---><code>go run .</code></li>
<li>Собрать и запустить докер образ <code>docker build . --file Dockerfile -t app:latest</code>---><code>docker-compose up</code></li>
<li>Или просто использовать Makefile <code>make up</code></li></ul>

<h2>Общее описание</h2>
Спасибо за интересную задачу, было интересно делать. Реализовал все необходимые функции, кроме тестов (не успел). Аунтификация реализована через JWT. После входа в систему выдается токен, который затем необходимо передавать через Заголовок <code>Auntification</code>. Интересно было реализовывать полноценную чистую архитектуру: программа разделена на 4 слоя - доменный - основные структуры, методы и интерфейсы, инфраструктурны - взаимодействие со сторонними сервисами, приложение - основной слой и слой презентации - ручки для GRPC и REST API
<h3>Спецификацию api можно посмотреть в папке api</h3>

<code>
/ShtrafovNet
│
├── /api
│   └── /swagger.yml
├── /cmd
│   └── /ShtrafovNet
│       └── main.go - ОСНОВНОЙ ПАКЕТ
|       └── config.env - файл конфигурации
|
├── /internal
│   └── /domain
│   |   └── entity - пакет основных сущностей
│   |   └── errors - ошибки доменного слоя
│   |   └── interfaces
│   |   └── odt - структуры для обмена данными между слоями
│   |   └── service - ключевые методы
|   |   
│   └── /infrastructure
│   |   └── errors - ошибки инфраструктурного слоя
│   |   └── migrations - файлик с миграциями бд
│   |   └── repository - реалезации интерфейсов из доменного слоя для интеграции сторонних сервисов
|   |
│   └── /application
│   |   └── errors - ошибки слоя приложениия
│   |   └── auction - функционал, связанные напрямую с аукционами
│   |   └── config - подтягивает конфиги из env
│   |   └── database - подключение к бд
│   |   └── logger - имплементация логгера
│   |   └── usesrs - функционал для определения пользователей
│   |   └── workers - воркер
|   |
│   └── /presentation
│       └── errors - ошибки слоя презентации
│       └── server - пакет релизующий grpc и grpc Gateway
|  
├── /log
│   └── log.log - файл с логами
|
├── Dockerfile - файл создания образа докер
├── docker-compose.yml - подъем образа и базы данных
├── Makefile 
</code>
