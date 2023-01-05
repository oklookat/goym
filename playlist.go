package goym

import (
	"github.com/oklookat/goym/holly"
)

// Получение списка плейлистов пользователя.
func (c *Client) GetUserPlaylists(userId int64) (data *TypicalResponse[[]*Playlist], err error) {
	data = &TypicalResponse[[]*Playlist]{}

	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", "list"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Получить плейлист пользователя по ID.
func (c *Client) GetUserPlaylist(userId int64, playListId int64) (data *TypicalResponse[*Playlist], err error) {
	data = &TypicalResponse[*Playlist]{}

	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", i2s(playListId)})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Создать плейлист.
func (c *Client) CreatePlaylist(playlistName string, isPublic bool) (data *TypicalResponse[*Playlist], err error) {
	data = &TypicalResponse[*Playlist]{}

	var endpoint = genApiPath([]string{"users", c.UserId, "playlists", "create"})
	var visibility = Visibility_Private
	if isPublic {
		visibility = Visibility_Public
	}

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(map[string]string{
		"title":      playlistName,
		"visibility": visibility,
	}).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Переименовать плейлист.
func (c *Client) RenamePlaylist(playListId int64, newName string) (data *TypicalResponse[*Playlist], err error) {
	data = &TypicalResponse[*Playlist]{}

	var endpoint = genApiPath([]string{"users", c.UserId, "playlists", i2s(playListId), "name"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(map[string]string{
		"value": newName,
	}).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Удалить плейлист.
func (c *Client) DeletePlaylist(playListId int64) (err error) {
	var endpoint = genApiPath([]string{"users", c.UserId, "playlists", i2s(playListId), "delete"})

	var hErr = &TypicalResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(hErr).Delete(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, hErr)
	}

	return
}

// Получить рекомендации на основе плейлиста.
func (c *Client) GetPlaylistRecommendations(playListId int64) (data *TypicalResponse[*PlaylistRecommendations], err error) {
	data = &TypicalResponse[*PlaylistRecommendations]{}

	var endpoint = genApiPath([]string{"users", c.UserId, "playlists", i2s(playListId), "recommendations"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Изменить видимость плейлиста.
//
// makePublic: true = сделать публичным, false = приватным.
func (c *Client) ChangePlaylistVisibility(playListId int64, makePublic bool) (data *TypicalResponse[*Playlist], err error) {
	data = &TypicalResponse[*Playlist]{}

	var endpoint = genApiPath([]string{"users", c.UserId, "playlists", i2s(playListId), "visibility"})
	var visibility = Visibility_Private
	if makePublic {
		visibility = Visibility_Public
	}

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(map[string]string{
		"value": visibility,
	}).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}
