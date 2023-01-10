package goym

import (
	"github.com/oklookat/goym/schema"
)

// Лайкнуть артиста.
//
// POST /users/{userId}/likes/artists/add
func (c Client) LikeArtist(a *schema.Artist) error {
	if a == nil {
		return ErrNilArtist
	}

	var body = schema.LikeArtistRequestBody{
		ArtistId: a.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", "add"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Убрать лайк с артиста.
//
// POST /users/{userId}/likes/artists/{artistId}/remove
func (c Client) UnlikeArtist(a *schema.Artist) error {
	if a == nil {
		return ErrNilArtist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", i2s(a.ID), "remove"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}
