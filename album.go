package goym

import (
	"errors"
)

// Получить альбом по id.
//
// withTracks - получить альбом с треками?
//
// Если да, то треки будут в Volumes и Duplicates.
func (c *Client) GetAlbum(albumId int64, withTracks bool) (*TypicalResponse[Album], error) {
	var endP = []string{"albums", i2s(albumId)}
	if withTracks {
		endP = append(endP, "with-tracks")
	}
	var endpoint = genApiPath(endP)

	var data = &TypicalResponse[Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получить альбомы по id.
func (c *Client) GetAlbums(albumIds []int64) (*TypicalResponse[[]Album], error) {
	if albumIds == nil {
		return nil, errors.New("nil albumIds")
	}
	var endpoint = genApiPath([]string{"albums"})

	var data = &TypicalResponse[[]Album]{}
	resp, err := c.self.R().SetError(data).SetResult(data).
		SetFormData(formAlbumIds(albumIds)).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Лайкнуть альбом.
func (c *Client) LikeAlbum(albumId int64) error {
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", "add"})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormData(formAlbumId(albumId)).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Убрать лайк с альбома.
func (c *Client) UnlikeAlbum(albumId int64) error {
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "albums", i2s(albumId), "remove"})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}
