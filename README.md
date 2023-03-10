# goym

Неофициальное API Яндекс.Музыки для Go.

- Могут быть ошибки. API может перестать работать, если сменятся client_id, client_secret у приложения ЯМ для Windows.
- Вход по методу "[Пользователь вводит код на Яндекс.OAuth](https://yandex.ru/dev/id/doc/dg/oauth/reference/simple-input-client.html)".

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
**метод не тестировался**

В токенах есть поле "RefreshAfter", которое обозначает время, когда эти токены надо обновить. Если токены получены только что, то это время будет +- 3 месяца. 

У токенов есть метод `Refresh()`, который будет выполнять обновление. 

Если токены обновятся, то **метод вернет вам новые токены, а старые перестанут действовать**. **Если обновление токенов не требуется, то вернется просто nil**. 

Вызывайте время от времени `Refresh()`, получайте новые токены, и перезаписывайте старые.

# Реализовано

## Недокументированные методы
- [ ] Методы, не указанные в [документации](https://www.cherkashin.dev/yandex-music-open-api/).
- Новые методы можно искать [тут](https://github.com/MarshalX/yandex-music-api). Или с помощью анализатора трафика смотреть приложения ЯМ.

## Account
- [x] GET /account/status
- [x] GET /account/settings
- [x] POST /account/consume-promo-code
- [x] POST /account/settings

## Album
- [x] GET /albums/{albumId}
- [x] GET /albums/{albumId}/with-tracks
- [x] GET /users/{userId}/likes/albums
- [x] POST /albums
- [x] POST /users/{userId}/likes/albums/add
- [x] POST /users/{userId}/likes/albums/{albumId}/remove

## Artist
- [x] GET /users/{userId}/likes/artists
- [x] GET /artists/{artistId}/track-ids-by-rating 
- [x] GET /artists/{artistId}/brief-info
- [x] GET /artists/{artistId}/tracks
- [x] GET /artists/{artistId}/direct-albums
- [x] POST /users/{userId}/likes/artists/add
- [x] POST /users/{userId}/likes/artists/{artistId}/remove

## Playlist
- [x] GET /users/{userId}/playlists/list
- [x] GET /users/{userId}/playlists/{kind}
- [x] GET /users/{userId}/playlists/{kind}/recommendations
- [x] POST /playlists/list
- [x] POST /users/{userId}/playlists/create
- [x] POST /users/{userId}/playlists/{kind}/name
- [x] POST /users/{userId}/playlists/{kind}/delete
- [x] POST /users/{userId}/playlists/{kind}/visibility
- [x] POST /users/{userId}/playlists/{kind}/change-relative
- [x] POST /users/{userId}/playlists/{kind}/change
- [x] POST /users/{userId}/likes/playlists/add
- [x] POST /users/{userId}/likes/playlists/{ownerUID}-{kind}/remove

## Search
- [x] GET /search
- [x] GET /search/suggest

## Track
- [x] GET /users/{userId}/likes/tracks
- [x] GET /users/{userId}/dislikes/tracks
- [x] GET /tracks/{trackId}/download-info
- [x] GET /tracks/{trackId}/supplement
- [x] GET /tracks/{trackId}/similar
- [x] GET /tracks/{trackId}
- [x] POST /users/{userId}/likes/tracks/add
- [x] POST /users/{userId}/likes/tracks/add-multiple
- [x] POST /users/{userId}/likes/tracks/remove
- [x] POST /tracks

## Tags
- [ ] /tags/{tagId}/playlist-ids

## Landing
- [ ] GET /landing3
- [ ] GET /landing3/{landingBlock}
- [ ] GET ​/landing3​/new-releases
- [ ] GET /landing3/podcasts
- [ ] GET /landing3/new-playlists
- [ ] GET /landing3/chart/{chartType}

## Rotor
- [x] GET /rotor/station/{type:tag}/tracks
- [x] GET /rotor/account/status
- [x] GET /rotor/stations/list
- [x] GET /rotor/stations/dashboard
- [x] GET /rotor/station/{type:tag}/info
- [x] POST /rotor/station/{type:tag}/feedback


## Default
- [ ] GET /settings
- [ ] GET /permission-alerts
- [ ] GET /feed/wizard/is-passed
- [ ] GET /feed
- [ ] GET /genres

# Чем можно помочь?
- Рефакторингом.
- Написанием нормальных комментариев к методам, документации.
- Написанием, исправлением тестов.
- Проверкой существующих методов.
- Поиском и реализацией новых методов.
- ???
