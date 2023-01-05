package goym

import "github.com/oklookat/goym/holly"

// Лайкнуть артиста.
func (c *Client) LikeArtist(artistId int64) (err error) {
	var endpoint = genApiPath([]string{"users", c.UserId, "likes", "artists", "add"})

	var data = &GetResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(map[string]string{
		"artist-id": i2s(artistId),
	}).Post(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Убрать лайк с артиста.
func (c *Client) UnlikeArtist(artistId int64) (err error) {
	var endpoint = genApiPath([]string{"users", c.UserId, "likes", "artists", i2s(artistId), "remove"})

	var data = &GetResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Post(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}
