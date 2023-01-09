package goym

// Получение списка плейлистов пользователя.
func (c *Client) GetUserPlaylists(userId int64) (*TypicalResponse[[]*Playlist], error) {
	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", "list"})

	var data = &TypicalResponse[[]*Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получить плейлист пользователя по ID.
func (c *Client) GetUserPlaylist(userId int64, kind int64) (*TypicalResponse[*Playlist], error) {
	var endpoint = genApiPath([]string{"users", i2s(userId), "playlists", i2s(kind)})

	var data = &TypicalResponse[*Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Создать плейлист.
func (c *Client) CreatePlaylist(name string, public bool) (*TypicalResponse[*Playlist], error) {
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", "create"})

	var data = &TypicalResponse[*Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormData(formTitleVisibility(name, public)).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Переименовать плейлист.
func (c *Client) RenamePlaylist(kind int64, newName string) (*TypicalResponse[*Playlist], error) {
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(kind), "name"})

	var data = &TypicalResponse[*Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormData(formValue(newName)).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Удалить плейлист.
func (c *Client) DeletePlaylist(kind int64) error {
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(kind), "delete"})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Получить рекомендации на основе плейлиста.
//
// Только для плейлистов, созданных пользователем.
func (c *Client) GetPlaylistRecommendations(kind int64) (*TypicalResponse[*PlaylistRecommendations], error) {
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(kind), "recommendations"})

	var data = &TypicalResponse[*PlaylistRecommendations]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Изменить видимость плейлиста.
//
// makePublic: true = сделать публичным, false = приватным.
func (c *Client) ChangePlaylistVisibility(kind int64, public bool) (*TypicalResponse[*Playlist], error) {
	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(kind), "visibility"})

	var data = &TypicalResponse[*Playlist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormData(formValue(visibilityToString(public))).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// // Добавить треки в плейлист.
// func (c *Client) AddPlaylistTracks(pl *Playlist, tr []*Track) (*TypicalResponse[*Playlist], error) {
// 	if pl == nil {
// 		return nil, errors.New("nil playlist")
// 	}
// 	if len(tr) == 0 {
// 		return nil, errors.New("empty tracks")
// 	}

// 	var op = playlistChange{}
// 	op.New(pl)
// 	var diff = &playlistDiff{}
// 	err := diff.NewInsert(pl, tr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	op.AddDiff(diff)
// 	form, err := op.GetForm()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "change"})
// 	var data = &TypicalResponse[*Playlist]{}
// 	resp, err := c.self.R().SetError(data).SetResult(data).
// 		SetFormData(form).Post(endpoint)
// 	if err == nil {
// 		err = checkTypicalResponse(resp, data)
// 	}

// 	return data, err
// }

// // Удалить треки из плейлиста.
// //
// // from - позиция в плейлисте, откуда начать
// //
// // to - позиция в плейлисте, где закончить
// func (c *Client) RemovePlaylistTracks(pl *Playlist, from int, to int) (*TypicalResponse[*Playlist], error) {
// 	if pl == nil {
// 		return nil, errors.New("nil playlist")
// 	}

// 	var op = playlistChange{}
// 	op.New(pl)
// 	var diff = &playlistDiff{}
// 	if err := diff.NewDelete(pl, from, to); err != nil {
// 		return nil, err
// 	}
// 	op.AddDiff(diff)
// 	form, err := op.GetForm()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var endpoint = genApiPath([]string{"users", c.userId, "playlists", i2s(pl.Kind), "change"})
// 	var data = &TypicalResponse[*Playlist]{}
// 	resp, err := c.self.R().SetError(data).SetResult(data).
// 		SetFormData(form).Post(endpoint)
// 	if err == nil {
// 		err = checkTypicalResponse(resp, data)
// 	}

// 	return data, err
// }
