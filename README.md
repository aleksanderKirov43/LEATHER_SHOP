# leather-shop
Ручки для интернет-магазина кожевенных изделий


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
