package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить альбом по id.
//
// withTracks - получить альбом с треками?
//
// Если да, то треки будут в Volumes и Duplicates.
func (c Client) Album(ctx context.Context, id schema.ID, withTracks bool) (*schema.Album, error) {
	// GET /albums/{albumId}
	// ||
	// GET /albums/{albumId}/with-tracks
	endP := []string{"albums", id.String()}
	if withTracks {
		endP = append(endP, "with-tracks")
	}
	endpoint := genApiPath(endP...)

	data := &schema.Response[*schema.Album]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Получить альбомы по id.
func (c Client) Albums(ctx context.Context, albumIds []schema.ID) ([]*schema.Album, error) {
	// POST /albums
	if albumIds == nil {
		return nil, nil
	}

	body := schema.GetAlbumsByIdsRequestBody{
		AlbumIds: albumIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("albums")
	data := &schema.Response[[]*schema.Album]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Лайкнуть альбом по ID.
func (c Client) LikeAlbum(ctx context.Context, id schema.ID) error {
	// POST /users/{userId}/likes/albums/add
	body := schema.LikeAlbumRequestBody{
		AlbumId: id,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endpoint := genApiPath("users", c.userId, "likes", "albums", "add")
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).
		Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Убрать лайк с альбома по ID.
func (c Client) UnlikeAlbum(ctx context.Context, id schema.ID) error {
	// POST /users/{userId}/likes/albums/{albumId}/remove
	endpoint := genApiPath("users", c.userId, "likes", "albums", id.String(), "remove")
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Получить лайкнутые альбомы.
func (c Client) LikedAlbums(ctx context.Context) ([]*schema.AlbumShort, error) {
	// GET /users/{userId}/likes/albums
	endpoint := genApiPath("users", c.userId, "likes", "albums")
	data := &schema.Response[[]*schema.AlbumShort]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
