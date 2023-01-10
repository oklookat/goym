# goym

Неофициальное API десктопного приложения Яндекс.Музыки для Go.

- Методы: [отсюда](https://www.cherkashin.dev/yandex-music-open-api).
- Полезные вещи: [можно найти тут](https://github.com/MarshalX/yandex-music-api).
- Убрал, исправил некоторые поля, структуры.
- Реализованы не все методы.
- Могут быть ошибки. API может перестать работать, если сменятся client_id, client_secret у приложения ЯМ для Windows.
- Вход по методу "[Пользователь вводит код на Яндекс.OAuth](https://yandex.ru/dev/id/doc/dg/oauth/reference/simple-input-client.html)".
- Реализовано (но не протестировано) [обновление токенов](https://yandex.ru/dev/id/doc/dg/oauth/reference/refresh-client.html).

Требования:
- Go 1.19+


# Гайд

1. ```go get github.com/oklookat/goym```

2. Вызовите `New()` из пакета [auth](./auth), передайте контекст с отменой, логин (или почту),
и функцию, в которой вам надо будет направить пользователя по URL из параметра `url`, чтобы он ввел там код из параметра `code`. 

Пользователю надо сделать это в течение пяти минут, потом токены истекут, а вам вернется ошибка.

3. Каждые 7 секунд горутина будет отправлять запрос, проверяя, ввел ли пользователь код. До первой ошибки, или успеха. 

Поэтому важно передать контекст с отменой, чтобы вы смогли отменить бесконечный цикл, в случае чего.

4. После того, как пользователь введет код, при следующем запросе к api, функция вернет токены или ошибку.

Чтоб не авторизироваться по 100 раз, полученные токены сохраните куда-нибудь.

5. Отдайте токены функции [New() в пакете goym](./main.go), в ответ вернется клиент для запросов к API.

# Обновление токенов
В токенах есть поле "RefreshAfter", которое обозначает время, когда эти токены надо обновить. Если токены получены только что, то это время будет +- 3 месяца. 

У токенов есть метод `Refresh()`, который будет выполнять обновление. 

Если токены обновятся, то **метод вернет вам новые токены, а старые перестанут действовать**. **Если обновление токенов не требуется, то вернется просто nil**. 

Вызывайте время от времени `Refresh()`, получайте новые токены, и перезаписывайте старые.
