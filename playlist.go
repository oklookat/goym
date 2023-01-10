package goym

import "github.com/oklookat/goym/schema"

// Получение списка плейлистов пользователя.
//
// GET /users/{userId}/playlists/list
func (c Client) GetUserPlaylists(userId int64) ([]*schema.Playlist, error) {
	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", "list"})
	var data = &schema.TypicalResponse[[]*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получение плейлиста по уникальному идентификатору.
//
// GET /users/{userId}/playlists/{kind}
func (c Client) GetUserPlaylistById(userId int64, kind int64) (*schema.Playlist, error) {
	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", i2s(kind)})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Создать плейлист.
//
// POST /users/{userId}/playlists/create
func (c Client) CreatePlaylist(name string, vis schema.Visibility) (*schema.Playlist, error) {
	var body = schema.CreatePlaylistRequestBody{
		Title:      name,
		Visibility: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", "create"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Переименовать плейлист.
//
// POST /users/{userId}/playlists/{kind}/name
func (c Client) RenamePlaylist(pl *schema.Playlist, newName string) (*schema.Playlist, error) {
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var body = schema.RenamePlaylistRequestBody{
		Value: newName,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "name"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Удалить плейлист.
//
// POST /users/{userId}/playlists/{kind}/delete
func (c Client) DeletePlaylist(pl *schema.Playlist) error {
	if pl == nil {
		return ErrNilPlaylist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "delete"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Получить рекомендации на основе плейлиста.
//
// Только для плейлистов, созданных пользователем.
//
// Если в плейлисте нет треков, рекомендаций не будет.
//
// GET /users/{userId}/playlists/{kind}/recommendations
func (c Client) GetPlaylistRecommendations(pl *schema.Playlist) (*schema.PlaylistRecommendations, error) {
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "recommendations"})
	var data = &schema.TypicalResponse[*schema.PlaylistRecommendations]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Изменить видимость плейлиста.
//
// POST /users/{userId}/playlists/{kind}/visibility
func (c Client) ChangePlaylistVisibility(pl *schema.Playlist, vis schema.Visibility) (*schema.Playlist, error) {
	if pl == nil {
		return nil, ErrNilPlaylist
	}
	var body = schema.ChangePlaylistVisibilityRequestBody{
		Value: vis,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "visibility"})
	var data = &schema.TypicalResponse[*schema.Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}
