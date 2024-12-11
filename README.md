# vertex-api
vertex-api - это координирующий сервер для Vertex. 

# Структура бд
Структура базы данных и первые тестовые данные хранятся в SQL файлах в директории migrations/init.

# Начало работы

### Для запуска этого репозитория (локально) потребуется
- Установленный Go
- .env файл для загрузки переменных окружений (в корне проекта)
- config.local.json

### .env
```
BUSINESS_DB_PASSWORD= ваш пароль
SERVER_MODE=DEVELOPMENT
DOMAIN=http://localhost:3000
LOG_DIRECTORY=logs
CONFIG_PATH=configs/config.local.json
SERVER_VERSION=local
IS_VERIFY_DEPENDENCIES=true
```

### config.local.json
```
{
  "server": {
    "port": "25504"
  },
  "business-database": {
    "host": "localhost",
    "port": "5432",
    "username": "",
    "db_name": "",
    "ssl_mode": "disable"
  },
  "jwt": {
    "secret_key": "ваш секретный ключ (можно сгенерировать через openssl rand -hex 32)"
  }
}

```

### goose (изменять в баш файлах)
```
export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING=postgresql://логин:пароль@localhost:порт/название_бд?sslmode=disable
```

# Быстрый старт
- Создаем .env
- Создаем config.local.json
- запускаем ./goose_init.sh
- запускаем ./swagger_init_and_start.sh
- запускаем main