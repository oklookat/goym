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
func (c Client) GetAlbumById(ctx context.Context, id schema.UniqueID, withTracks bool) (*schema.Album, error) {
	// GET /albums/{albumId}
	// ||
	// GET /albums/{albumId}/with-tracks
	var endP = []string{"albums", id.String()}
	if withTracks {
		endP = append(endP, "with-tracks")
	}
	var endpoint = genApiPath(endP)

	var data = &schema.Response[*schema.Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Получить альбомы по id.
func (c Client) GetAlbumsByIds(ctx context.Context, albumIds []schema.UniqueID) ([]*schema.Album, error) {
	// POST /albums
	if albumIds == nil {
		return nil, nil
	}

	var body = schema.GetAlbumsByIdsRequestBody{
		AlbumIds: albumIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"albums"})
	var data = &schema.Response[[]*schema.Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return data.Result, err
}

// Лайкнуть альбом.
func (c Client) LikeAlbum(ctx context.Context, al *schema.Album) error {
	// POST /users/{userId}/likes/albums/add
	if al == nil {
		return nil
	}

	var body = schema.LikeAlbumRequestBody{
		AlbumId: al.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", "add"})
	var data = &schema.Response[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).
		Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Убрать лайк с альбома.
func (c Client) UnlikeAlbum(ctx context.Context, al *schema.Album) error {
	// POST /users/{userId}/likes/albums/{albumId}/remove
	if al == nil {
		return nil
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", al.ID.String(), "remove"})
	var data = &schema.Response[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Получить лайкнутые альбомы.
func (c Client) GetLikedAlbums(ctx context.Context) ([]*schema.AlbumShort, error) {
	// GET /users/{userId}/likes/albums
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums"})
	var data = &schema.Response[[]*schema.AlbumShort]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
