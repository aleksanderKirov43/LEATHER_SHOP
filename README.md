# leather-shop
Ручки для интернет-магазина кожевенных изделий

*Перед тем как собрать и запустить сервис в Docker создайте файл .env по пути /config/env/ со следующими данными:*
```
LEATHER_APP_PORT: 8089

LEATHER_DB_HOST: db
LEATHER_DB_PORT: 5432
LEATHER_DB_USER: leather
LEATHER_DB_PASSWORD: ef4t_A7yyU
LEATHER_DB_DATABASE: leather_base
JWT_SECRET: zoVqwG_6p7
JWT_ACCESS_TTL: 20
JWT_REFRESH_TTL: 90

```

*Далее, что-бы с сервисом можно было работать запускайте docker-compose.dev.yaml, после билда образа, командой:* 

```
docker-compose -f docker-compose.dev.yaml up                 
```
*При подключении к БД черз свой клиент используйте localhost*

## API Маршруты

### Авторизация и обновление токенов

1. **Вход пользователя:**

    - **URL:** `POST /api/auth/login`
    - **Headers:**
        - `Content-Type`: `application/json`
    - **Body (raw JSON):**
```json
{
  "username": "Admin",
  "password": "admin"
}
```
- Ответ:
```json
{
  "Аккес токен": "access_token_value",
  "Рефреш токен": "refresh_token_value"
}
```

2. **Обновление токена:**

- **URL:** `POST /api/auth/refresh`
- **Headers:**
    - `Authorization-Refresh`: `Bearer {refresh_token}`

- Ответ:
```json
{
  "Аккес токен": "new_access_token_value"
}
```

### Управление пользователями

1. **Создание пользователя:**

    - **URL:** `POST /api/users`
    - **Headers:**
        - `Content-Type`: `application/json`
    - **Body (raw JSON):**
```json
{
  "firstname": "Александр",
  "lastname": "Лобжа",
  "username": "Admin",
  "password": "admin",
  "type": 1,
  "email": "region-manager43@yandex.ru",
  "phone": "89005295557",
  "wishlist": 0,
  "cart": 0
}
```

- Ответ:
```json
{
 "id": 1,
 "firstname": "Александр",
 "lastname": "Лобжа",
 "username": "Admin",
 "type": 1,
 "email": "region-manager43@yandex.ru",
 "phone": "89005295557",
 "password": "$2a$10$PjtlbkUK/VgmYF.ydjdBZOMQe0JIiwiw1GuKL9OcGOCcYcCMSn/Be",
 "wishlist": 0,
 "cart": 0
}
```

2. **Получение пользователя по ID:**

- **URL:** `GET /api/users/{id}`
- **Headers:**
    - `Authorization`: `Bearer {access_token}`
- **Ответ:**
```json
{
  "id": 1,
  "firstname": "Александр",
  "lastname": "Лобжа",
  "username": "Admin",
  "type": 1,
  "email": "region-manager43@yandex.ru",
  "phone": "89005295557",
  "wishlist": 0,
  "cart": 0
}
```
3. **Получение всех пользователей:**

- **URL:** `GET /api/users`
- **Headers:**
    - `Authorization`: `Bearer {access_token}`
- **Ответ:**
```json
[
  {
    "id": 1,
    "firstname": "Александр",
    "lastname": "Лобжа",
    "username": "Admin",
    "type": 1,
    "email": "region-manager43@yandex.ru",
    "phone": "89005295557",
    "wishlist": 0,
    "cart": 0
  }
]
```
4. Редактирование пользователя:

- **URL:** `PUT /api/users/{id}`
- **Headers:**
    - `Authorization`: `Bearer {access_token}`
    - `Content-Type`: `application/json`

- **Body (raw JSON):**
```json
{
  "firstname": "UpdatedName",
  "lastname": "UpdatedSurname",
  "username": "UpdatedUsername",
  "password": "updated_password",
  "type": 1,
  "email": "updated_email@example.com",
  "phone": "89005555555",
  "wishlist": 0,
  "cart": 0
}
```
- Ответ:
```json
{
  "id": 1,
  "firstname": "UpdatedName",
  "lastname": "UpdatedSurname",
  "username": "UpdatedUsername",
  "type": 1,
  "email": "updated_email@example.com",
  "phone": "89005555555",
  "wishlist": 0,
  "cart": 0
}
```

5. **Удаление пользователя:**

- **URL:** `DELETE /api/users/{id}`
- **Headers:**
    - `Authorization`: `Bearer {access_token}`
- **Ответ:**
```json
{
  "message": "User deleted successfully"
}
```



### Управление товарами

1. **Создание товара:**
    - **URL:** `POST /api/products`

    - **Body (raw JSON):**
```json
{
    "name": "Тестовый товар",
    "description": "Тут должно содержаться описание тестового товара",
    "quantity": 100,
    "image": ["http://example.com/image1.jpg", "http://example.com/image2.jpg"],
    "sale": 10,
    "price": 4500,
    "status": 2,
    "category": 1,
    "property": 1
}
```

**Ответ:**
```json
{
    "id": 1,
    "name": "Тестовый товар",
    "description": "Тут должно содержаться описание тестового товара",
    "quantity": 100,
    "image": [
        "http://example.com/image1.jpg",
        "http://example.com/image2.jpg"
    ],
    "sale": 10,
    "price": 4500,
    "status": 2,
    "category": 1,
    "property": 1
}

```

**Получение товара по ID:**

- **URL:** `GET /api/products/{id}`

- **Ответ:**
```json
{
    "id": 1,
    "name": "Тестовый товар",
    "description": "Тут должно содержаться описание тестового товара",
    "quantity": 100,
    "image": [
        "http://example.com/image1.jpg",
        "http://example.com/image2.jpg"
    ],
    "sale": 10,
    "price": 4500,
    "status": 2,
    "category": 1,
    "property": 1
}
```

**Получение всех товаров:**

- **URL:** `GET /api/products`

*Тут мы просто получаем список всех существующих товаров. Пример не приведён для улучшения читабельности README, что бы не засорять пространство*

**Редактирование товара:**

- **URL:** `PUT /api/products/{id}`

Body (raw JSON):
```json
{
    "id": 1,
    "name": "Тестовый товар отредактирован",
    "description": "Описание тестового товара",
    "quantity": 30,
    "image": [
        "http://example.com/image1.jpg",
        "http://example.com/image2.jpg"
    ],
    "sale": 0,
    "price": 4500,
    "status": 2,
    "category": 1,
    "property": 1
}
```

**Удаление товара:**

- **URL:** `DELETE /api/products/{id}`

*Ответ:*

```json
{
    "message": "Товар успешно удалён"
}
```


### Управление категориями


1. **Создание категории:**
    - **URL:** `POST http://localhost:8089/category`

    - **Body (raw JSON):**
```json
{
  "id": 2,
  "name": "Кошельки",
  "props": {
    "Длина": "120 мм",
    "Ширина": "100 мм",
    "Вес": "135 грамм",
    "Отделений под карты": "3 шт",
    "Монетница": "Да"
  }
}
```

**Ответ:**
```json
{
  "id": 2,
  "name": "Кошельки",
  "props": {
    "Длина": "120 мм",
    "Ширина": "100 мм",
    "Вес": "135 грамм",
    "Отделений под карты": "3 шт",
    "Монетница": "Да"
  }
}

```

**Получение категории по ID:**

- **URL:** `GET /category/{id}`


**Получение всех категорй:**

- **URL:** `GET /category`

*Тут мы просто получаем список всех существующих категорий. Пример не приведён для улучшения читабельности README, что бы не засорять пространство*

**Редактирование категории:**

- **URL:** `PUT /category/{id}`

Body (raw JSON):

**Удаление категории:**

- **URL:** `DELETE /category/{id}`

*Ответ:*

```json
{
    "message": "Категория успешно удалена"
}
```
