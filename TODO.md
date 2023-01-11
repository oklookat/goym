# Заметки

- найти метод поиска треков по ISRC?
- добавить новых методов, улучшить старые
- проверить пакет auth (после всех изменений я его еще не тестил)

# Готово
## Account
- GET /account/status

## Album
- GET /albums/{albumId}
- GET /albums/{albumId}/with-tracks
- POST /albums
- POST /users/{userId}/likes/albums/add
- POST /users/{userId}/likes/albums/{albumId}/remove

## Artist
- POST /users/{userId}/likes/artists/add
- POST /users/{userId}/likes/artists/{artistId}/remove

## Playlist
- GET /users/{userId}/playlists/list
- GET /users/{userId}/playlists/{kind}
- POST /users/{userId}/playlists/create
- POST /users/{userId}/playlists/{kind}/name
- POST /users/{userId}/playlists/{kind}/delete
- GET /users/{userId}/playlists/{kind}/recommendations
- POST /users/{userId}/playlists/{kind}/visibility
- POST /users/{userId}/playlists/{kind}/change-relative
- POST /users/{userId}/playlists/{kind}/change

## Search
- GET /search
- GET /search/suggest

## Track
- GET /users/{userId}/likes/tracks
- GET /users/{userId}/dislikes/tracks
- POST /users/{userId}/likes/tracks/add
- POST /users/{userId}/likes/tracks/add-multiple
- POST /users/{userId}/likes/tracks/remove
- GET /tracks/{trackId}
- POST /tracks
- GET /tracks/{trackId}/download-info
- GET /tracks/{trackId}/supplement
- GET /tracks/{trackId}/similar
