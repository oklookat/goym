package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Лайкнуть артиста.
func (c Client) LikeArtist(ctx context.Context, a *schema.Artist) error {
	// POST /users/{userId}/likes/artists/add
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
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Убрать лайк с артиста.
func (c Client) UnlikeArtist(ctx context.Context, a *schema.Artist) error {
	// POST /users/{userId}/likes/artists/{artistId}/remove
	if a == nil {
		return ErrNilArtist
	}
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", i2s(a.ID), "remove"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}
