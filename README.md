# goym

Неофициальное API Яндекс.Музыки для Go.

- Могут быть ошибки. API может перестать работать, если сменятся client_id, client_secret у приложения ЯМ для Windows.

# Гайд

1. ```go get github.com/oklookat/goym```

2. Получите клиента вызвав `New`.

Токен можно получить например используя [этот пакет](https://github.com/oklookat/yandexauth).

# Реализовано

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
- [x] POST /users/{userId}/likes/albums/add-multiple
- [x] POST /users/{userId}/likes/albums/remove

## Artist
- [x] GET /users/{userId}/likes/artists
- [x] GET /artists/{artistId}/track-ids-by-rating 
- [x] GET /artists/{artistId}/brief-info
- [x] GET /artists/{artistId}/tracks
- [x] GET /artists/{artistId}/direct-albums
- [x] POST /users/{userId}/likes/artists/add
- [x] POST /users/{userId}/likes/artists/add-multiple
- [x] POST /users/{userId}/likes/artists/remove

## Playlist
- [x] GET /users/{userId}/playlists/list
- [x] GET /users/{userId}/playlists/{kind}
- [x] GET /users/{userId}/playlists/{kind}/recommendations
- [x] POST /playlists/list
- [x] POST /users/{userId}/playlists/create
- [x] POST /users/{userId}/playlists/{kind}/name
- [x] POST /users/{userId}/playlists/{kind}/delete
- [x] POST /users/{userId}/playlists/{kind}/visibility
- [x] POST /users/{userId}/playlists/{kind}/description
- [x] POST /users/{userId}/playlists/{kind}/change-relative
- [x] POST /users/{userId}/playlists/{kind}/change
- [x] GET /users/{userId}/likes/playlists
- [x] POST /users/{userId}/likes/playlists/add
- [x] POST /users/{userId}/likes/playlists/add-multiple
- [x] POST /users/{userId}/likes/playlists/remove

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

# Где искать новые методы
- С помощью анализатора трафика смотреть приложения ЯМ.
- [Или тут](https://www.cherkashin.dev/yandex-music-open-api/).
- [И тут](https://github.com/MarshalX/yandex-music-api). 
- [И здесь](https://github.com/K1llMan/Yandex.Music.Api). 

# Чем можно помочь?
- Рефакторингом.
- Написанием нормальных комментариев к методам, документации.
- Написанием, исправлением тестов.
- Проверкой существующих методов.
- Поиском и реализацией новых методов.
- ???
