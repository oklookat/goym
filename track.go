package goym

import (
	"errors"
)

// Получить (диз)лайкнутые треки.
func (c *Client) GetLikedDislikedTracks(liked bool) (*TypicalResponse[TracksLibrary], error) {
	var endP = "likes"
	if !liked {
		endP = "dislikes"
	}
	var endpoint = genApiPath([]string{"users", c.userId, endP, "tracks"})

	var data = &TypicalResponse[TracksLibrary]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Поставить/снять лайк у трека/треков.
//
// like = true - поставить лайк
//
// like = false - убрать лайк
func (c *Client) LikeUnlikeTracks(trackIds []int64, like bool) error {
	if len(trackIds) == 0 {
		return errors.New("nil trackIds")
	}

	// Интересный факт:
	// Если в add-multiple будет один трек, типа {"track-ids": "idтрека"}, то ничего не произойдет.
	// Чтобы что-то прозошло, в track-ids надо указать не просто id трека,
	// а id трека и альбома через двоеточие. Например {"track-ids": "idтрека:idальбома"}.
	// Но так как у нас есть метод add, то заморачиваться не надо.

	var endP string

	var form map[string]string
	if !like || len(trackIds) > 1 {
		if !like {
			// если убрать лайки
			endP = "remove"
		} else {
			// если треков много
			endP = "add-multiple"
		}
		form = formTrackIds(trackIds)
	} else if len(trackIds) < 2 {
		// если трек один
		endP = "add"
		form = formTrackId(trackIds[0])
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "tracks", endP})

	var data = &TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetFormData(form).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return err
}

// Получить трек по id.
func (c *Client) GetTrackById(trackId int64) (*TypicalResponse[[]Track], error) {
	var endpoint = genApiPath([]string{"tracks", i2s(trackId)})

	var data = &TypicalResponse[[]Track]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получить треки по id.
//
// key - track id
//
// value - album id.
func (c *Client) GetTracksById(trackIds []int64) (*TypicalResponse[[]Track], error) {
	if trackIds == nil {
		return nil, errors.New("nil trackIds")
	}

	var endpoint = genApiPath([]string{"tracks"})

	var data = &TypicalResponse[[]Track]{}
	resp, err := c.self.R().SetError(data).SetFormData(formTrackIds(trackIds)).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получить (диз)лайкнутые треки.
func (c *Client) GetTrackDownloadInfo(trackId int64) (*TypicalResponse[[]TrackDownloadInfo], error) {
	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "download-info"})

	var data = &TypicalResponse[[]TrackDownloadInfo]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получение дополнительной информации о треке (Текст песни, видео, и т.д.).
func (c *Client) GetTrackSupplement(trackId int64) (*TypicalResponse[Supplement], error) {
	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "supplement"})

	var data = &TypicalResponse[Supplement]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}

// Получение похожих треков.
func (c *Client) GetSimilarTracks(trackId int64) (*TypicalResponse[SimilarTracks], error) {
	var endpoint = genApiPath([]string{"tracks", i2s(trackId), "similar"})

	var data = &TypicalResponse[SimilarTracks]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}

	return data, err
}
