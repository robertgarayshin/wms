# WMS - Warehouses Management System
## THIS IS BAD FEATURE
Для запуска проекта:  
```
git clone https://github.com/robertgarayshin/wms.git && cd wms
```
```
make up
```
Команда запускает БД PostgreSQL и приложение в Docker-контейнерах, применяет миграции и наполняет БД тестовыми данными.  
Для доступа к Swagger документации нужно перейти по ссылке:  
[SwaggerDocs URL](http://localhost:8080/swagger/index.html)  
После запуска приложения запросы к API отправляются по эндпоинту:
`http://localhost:8080/v1/`

### Запросы к API:  
1. Создание товара
```
curl -X 'PUT' \
  'http://localhost:8080/v1/items' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "items": [
    {
      "name": "TEST_SIZE",
      "quantity": 5,
      "size": "TEST_SIZE",
      "unique_id": "TEST0001",
      "warehouse_id": 0
    }
  ]
}'
```
**Ответ**: 
```
{
  "status": 201,
  "status_message": "Created",
  "message": "items successfully created",
  "error": ""
}
```  
2. Создание склада
```
curl -X 'POST' \
  'http://localhost:8080/v1/warehouses' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "warehouse": {
    "availability": true,
    "name": "TEST_warehouse"
  }
}'
```  
**Ответ**:
```
{
  "status": 201,
  "status_message": "Created",
  "message": "warehouse created successfully",
  "error": ""
}
```  
3. Создание резервации
```
curl -X 'POST' \
  'http://localhost:8080/v1/reserve' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "ids": [
    "A0001",
    "A0002"
  ]
}'
```  
**Ответ**:
```
{
  "status": 201,
  "status_message": "Created",
  "message": "reservation successfully created",
  "error": ""
}
```  
4. Создание резервации (информаиця о товаре не записана в БД)
```
curl -X 'POST' \
  'http://localhost:8080/v1/reserve' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "ids": [
    "THIS_ITEM_IS_NOT_PRESENTED"  
]
}'
```
**Ответ**:
```
{
  "status": 201,
  "status_message": "Created",
  "message": "reservation successfully created",
  "error": ""
}
```  
5. Удаление резервации
```
curl -X 'DELETE' \
  'http://localhost:8080/v1/reserve' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "ids": [
    "A0001"
  ]
}'
```  
**Ответ**:
```
{
  "status": 200,
  "status_message": "OK",
  "message": "reservation successfully cancelled",
  "error": ""
}
```  
6. Получение количества товаров на складе по идентификационному номеру
```
curl -X 'GET' \
  'http://localhost:8080/v1/items/0/quantity' \
  -H 'accept: application/json'
```
**Ответ**:
```
{
  "status": 200,
  "status_message": "OK",
  "message": {
    "TEST0001": 10,
    "THIS_ITEM_IS_NOT_PRESENTED": -1
  },
  "error": ""
}
```  
7. Запрос на резервацию к недопступному складу:
```
curl -X 'POST' \
  'http://localhost:8080/v1/reserve' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "ids": [
    "F0001"
  ]
}'
```
**Ответ**:
```
{
  "status": 403,
  "status_message": "Forbidden",
  "message": null,
  "error": "warehouse is unavailable"
}
```  
8. Поступление товара на несуществующий склад
```
curl -X 'PUT' \
  'http://localhost:8080/v1/items' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "items": [
    {
      "name": "string",
      "quantity": 0,
      "size": "string",
      "unique_id": "string",
      "warehouse_id": 12
    }
  ]
}'
```
**Ответ**:
```
{
  "status": 404,
  "status_message": "Not Found",
  "message": null,
  "error": "warehouse is not exist"
}
```  

### Особенности:  
1. **Резервация товара при отсутствии информации о нем в БД:**  
> **Usecase**:  
> Информация о товаре есть в системе, но не содержится в данном сервисе.
> Возможно при предварительной резервации товара, который еще не поступил на склад.  

Товар будет создан, при этом в БД будет отсутствовать информация о его названии и размере, будет только уникальный код
и код склада == 0 ("Unstored"). При этом количество будет равно 0 - количество зарезервированного товара.  
При поступлении товара на склад приходит запрос на его добавление в БД, при этом обновятся поля имени, размера и склада товара.  
Количество товара будет равно **К(поступившее) - К(зарезервированное)**    

2. **Доступ к недоступному складу**  
> **Usecase**:  
> Склад временно закрылся по техническим причинам. Сделать резерв или удалить резерв со склада невозможно.
> На складе есть некоторое количество товаров.

На недоступном складе остается некоторое количество товара, которое можно узнать с помощью запроса на получение 
количества товаров.  
На недоступном складе невозможно зарезервировать или снять резервацию товара. Будет выведена 403 ошибка и текст, 
сообщающий, что склад недоступен.    

3. **Удаление резерваций при их количестве == 0**
> **Usecase**  
> К API приходит запрос на удаление резервации от другого сервиса, который передал информацию с опозданием
> по некоторым причинам (например, сетевой сбой)

При попытке удалить резервацию, когда их нет, приложение не позволит это сделать и выведет ошибку, что у данного товара
нет резерваций. Но при этом создать резервацию для товара, которого нет, все еще возможно.
