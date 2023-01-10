package goym

import (
	"errors"

	"github.com/oklookat/goym/schema"
)

// Получить альбом по id.
//
// withTracks - получить альбом с треками?
//
// Если да, то треки будут в Volumes и Duplicates.
//
// GET /albums/{albumId} | /albums/{albumId}/with-tracks
func (c Client) GetAlbumById(id int64, withTracks bool) (*schema.Album, error) {
	var endP = []string{"albums", i2s(id)}
	if withTracks {
		endP = append(endP, "with-tracks")
	}
	var endpoint = genApiPath(endP)

	var data = &schema.TypicalResponse[*schema.Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data.Result, err
}

// Получить альбомы по id.
//
// POST /albums
func (c Client) GetAlbumsByIds(albumIds []int64) ([]*schema.Album, error) {
	if albumIds == nil {
		return nil, errors.New("nil albumIds")
	}

	var body = schema.GetAlbumsByIdsRequestBody{
		AlbumIds: albumIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"albums"})
	var data = &schema.TypicalResponse[[]*schema.Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data.Result, err
}

// Лайкнуть альбом.
//
// POST /users/{userId}/likes/albums/add
func (c Client) LikeAlbum(al *schema.Album) error {
	if al == nil {
		return ErrNilAlbum
	}

	var body = schema.LikeAlbumRequestBody{
		AlbumId: al.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", "add"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormUrlValues(vals).
		Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Убрать лайк с альбома.
//
// POST /users/{userId}/likes/albums/{albumId}/remove
func (c Client) UnlikeAlbum(al *schema.Album) error {
	if al == nil {
		return ErrNilAlbum
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", i2s(al.ID), "remove"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}
