# goym

Неофициальный клиент API Яндекс.Музыки для Go.

- Могут быть ошибки. API может перестать работать, если сменятся client_id, client_secret у приложения ЯМ для Windows.

# Гайд

1. ```go get github.com/oklookat/goym```

2. Получите клиента вызвав `New`.

Токен можно получить например через пакет [yandexauth](https://github.com/oklookat/yandexauth).

ID и Secret (приложение для Windows):

ID: ```23cabbbdc6cd418abb4b39c32c41195d```

Secret: ```53bc75238f0c4d08a118e51fe9203300```

# Полезная информация

## Проверка на 404

Если нужна проверка на отсутствие сущности, проверяйте schema.Error.Validate, schema.Error.IsNotFound, 
и пустоту полей кроме ID (например поле Name).
