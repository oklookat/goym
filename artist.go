package goym

// Лайкнуть артиста.
func (c *Client) LikeArtist(artistId int64) error {
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", "add"})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormData(formArtistId(artistId)).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Убрать лайк с артиста.
func (c *Client) UnlikeArtist(artistId int64) error {
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", i2s(artistId), "remove"})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}
