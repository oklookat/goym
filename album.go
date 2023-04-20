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

	body := schema.AlbumIdsRequestBody{
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

// Получить лайкнутые альбомы.
func (c Client) LikedAlbums(ctx context.Context) ([]*schema.AlbumShort, error) {
	// GET /users/{userId}/likes/albums
	return likesDislikes[[]*schema.AlbumShort](ctx, &c, true, "albums")
}

// Лайкнуть альбом.
func (c Client) LikeAlbum(ctx context.Context, id schema.ID) error {
	return c.likeUnlikeAlbums(ctx, []schema.ID{id}, true)
}

// Убрать лайк с альбома.
func (c Client) UnlikeAlbum(ctx context.Context, id schema.ID) error {
	return c.likeUnlikeAlbums(ctx, []schema.ID{id}, false)
}

// Лайкнуть альбомы.
//
// Используйте LikeAlbum() для лайка одного альбома.
func (c Client) LikeAlbums(ctx context.Context, ids []schema.ID) error {
	return c.likeUnlikeAlbums(ctx, ids, true)
}

// Снять лайки с альбомов.
func (c Client) UnlikeAlbums(ctx context.Context, ids []schema.ID) error {
	return c.likeUnlikeAlbums(ctx, ids, false)
}

func (c Client) likeUnlikeAlbums(ctx context.Context, ids []schema.ID, like bool) error {
	// POST /users/{userId}/likes/albums/add-multiple
	// ||
	// POST /users/{userId}/likes/albums/remove
	body := schema.AlbumIdsRequestBody{}
	body.SetIds(ids...)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}
	return addRemoveMultiple(ctx, &c, vals, like, "albums")
}
