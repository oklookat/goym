package goym

import (
	"errors"

	"github.com/oklookat/goym/holly"
)

// Получить альбом по id.
//
// withTracks - получить альбом с треками?
//
// Если да, то треки будут в Volumes и Duplicates.
func (c *Client) GetAlbum(albumId int64, withTracks bool) (data *TypicalResponse[*Album], err error) {
	data = &TypicalResponse[*Album]{}

	var par = []string{"albums", i2s(albumId)}
	if withTracks {
		par = append(par, "with-tracks")
	}
	var endpoint = genApiPath(par)

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Получить альбомы по id.
func (c *Client) GetAlbums(albumIds []int64) (data *TypicalResponse[[]*Album], err error) {
	if albumIds == nil {
		err = errors.New("nil albumIds")
		return
	}
	data = &TypicalResponse[[]*Album]{}

	var endpoint = genApiPath([]string{"albums"})

	var form = make(map[string]string)
	form["album-ids"] = i64Join(albumIds)

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(form).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Лайкнуть альбом.
func (c *Client) LikeAlbum(albumId int64) (err error) {
	var endpoint = genApiPath([]string{"users", c.UserId, "likes", "albums", "add"})

	var data = &TypicalResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).SetFormData(map[string]string{
		"album-id": i2s(albumId),
	}).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}

// Убрать лайк с альбома.
func (c *Client) UnlikeAlbum(albumId int64) (err error) {
	var endpoint = genApiPath([]string{"users", c.UserId, "likes", "albums", i2s(albumId), "remove"})

	var data = &TypicalResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Post(endpoint)

	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return
}
