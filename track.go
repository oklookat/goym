package goym

import (
	"errors"

	"github.com/oklookat/goym/holly"
)

// Получить (диз)лайкнутые треки.
func (c *Client) GetLikedDislikedTracks(liked bool) (data *GetResponse[*TracksLibrary], err error) {
	data = &GetResponse[*TracksLibrary]{}

	var endP = "likes"
	if !liked {
		endP = "dislikes"
	}
	var endpoint = genApiPath([]string{"users", c.UserId, endP, "tracks"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Поставить/снять лайк у трека/треков.
//
// like = true - поставить лайк
//
// like = false - убрать лайк
func (c *Client) LikeUnlikeTracks(trackIds []int64, like bool) (err error) {
	if len(trackIds) == 0 {
		err = errors.New("nil trackIds")
		return
	}

	// Интересный факт:
	// Если в add-multiple будет один трек, типа {"track-ids": "idтрека"}, то ничего не произойдет.
	// Чтобы что-то прозошло, в track-ids надо указать не просто id трека,
	// а id трека и альбома через двоеточие. Например {"track-ids": "idтрека:idальбома"}.
	// Но так как у нас есть метод add, то заморачиваться не надо.

	var form = make(map[string]string)
	var endP string

	if !like {
		endP = "remove"
		form["track-ids"] = i64Join(trackIds)
	} else {
		if len(trackIds) < 2 {
			endP = "add"
			form["track-id"] = i2s(trackIds[0])
		} else {
			endP = "add-multiple"
			form["track-ids"] = i64Join(trackIds)
		}
	}

	var endpoint = genApiPath([]string{"users", c.UserId, "likes", "tracks", endP})

	var data = &GetResponse[any]{}
	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetFormData(form).Post(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Получить трек по id.
func (c *Client) GetTrackById(trackId int64) (data *GetResponse[[]*Track], err error) {
	data = &GetResponse[[]*Track]{}
	var endpoint = genApiPath([]string{"tracks", i2s(trackId)})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Получить треки по id.
func (c *Client) GetTracksById(trackIds []int64) (data *GetResponse[[]*Track], err error) {
	data = &GetResponse[[]*Track]{}
	var endpoint = genApiPath([]string{"tracks"})

	var form = make(map[string]string)
	form["track-ids"] = i64Join(trackIds)

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Post(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Получить (диз)лайкнутые треки.
func (c *Client) GetTrackDownloadInfo(trackId int64) (data *GetResponse[[]*TrackDownloadInfo], err error) {
	data = &GetResponse[[]*TrackDownloadInfo]{}

	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "download-info"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Получение дополнительной информации о треке (Текст песни, видео, и т.д.).
func (c *Client) GetTrackSupplement(trackId int64) (data *GetResponse[*Supplement], err error) {
	data = &GetResponse[*Supplement]{}

	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "supplement"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}

// Получение похожих треков.
func (c *Client) GetSimilarTracks(trackId int64) (data *GetResponse[*SimilarTracks], err error) {
	data = &GetResponse[*SimilarTracks]{}

	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "similar"})

	var resp *holly.Response
	resp, err = c.self.R().SetError(data).SetResult(data).Get(endpoint)

	if err == nil {
		err = checkGetResponse(resp, data)
	}

	return
}
