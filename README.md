# Test_Seller

Метод создания объявления:

● принимает поля: название, описание, цена

● возвращает ID созданного объявления либо ошибку

POST http://localhost:8080/api/v1/announcements
Content-Type: application/json

{
"name": "name",
"description": "description",
"price": 132
}

======================================================================

Метод для добавления изображения в
объявление:

В формате multipart/form-data принимает изображение

● Идентификатор изображения принимается в path параметре запроса

PATCH  http://localhost:8080/api/v1/announcements/?uuid=1&path=3
Content-Type: multipart/form-data; boundary=WebAppBoundary

--WebAppBoundary
Content-Disposition: form-data; name="file"; filename="file44.txt"

< ./handler.go

======================================================================

Метод получения конкретного объявления:

● обязательные поля в ответе: название объявления, цена

● опциональные поля (можно запросить, передав query параметр fields):

описание, ссылки на все фото

GET http://localhost:8080/api/v1/announcements/?uuid=1

GET http://localhost:8080/api/v1/announcements/?uuid=1&fields=description&fields=images

======================================================================

Метод получения списка объявлений:

● возможность сортировки: по цене (возрастание/убывание) и по
дате создания (возрастание/убывание)

● поля в ответе: название объявления, ссылка на главное фото (первое в
списке), цена

GET http://localhost:8080/api/v1/announcements/

GET http://localhost:8080/api/v1/announcements/?page=1&limit=3

GET http://localhost:8080/api/v1/announcements/?page=1&limit=2&sort=desc(asc)&sortby=data

GET http://localhost:8080/api/v1/announcements/?page=1&limit=2&sort=desc(asc)&sortby=price