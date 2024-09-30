# Biophilia API

Бэкенд-часть приложения для работы с биомолекулами.

Для запуска:
- Создайте файл .env в корневой директории проекта, используя .env.example, находящегося в директории env/local/docker/.env.example
- Соберите и запустите приложение командой docker compose --env-file .env -f env/local/docker/docker-compose.yml up -d --build

После успешного запуска Swagger можно найти по следующему маршруту:
- http://localhost:8080/swagger/index.html