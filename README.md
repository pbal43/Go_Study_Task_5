# Go_Study_Task_5

Описания изменений - в корне в файле desk.txt
Коллекция Postman v. 2.1 - в корне в файле Task5.postman_collection.json

Запуск сервера:
1. docker-compose up --build

Подключение к БД: 
1. Установить connection по данным из docker-compose

Выполнение запросов:
1. Импортировать коллекцию в Postman

При проблемах с запуском попробуй выполнить:
docker container prune
docker rm -f postgres
docker volume rm $(docker volume ls -q)

