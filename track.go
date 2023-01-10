package goym

import (
	"github.com/oklookat/goym/schema"
)

// Получить лайкнутые треки.
//
// GET /users/{userId}/likes/tracks
func (c Client) GetLikedTracks() (*schema.TracksLibrary, error) {
	return c.getLikedDislikedTracks(true)
}

// Получить дизлайкнутые треки.
//
// GET /users/{userId}/dislikes/tracks
func (c Client) GetDislikedTracks() (*schema.TracksLibrary, error) {
	return c.getLikedDislikedTracks(false)
}

// Получить (диз)лайкнутые треки.
func (c Client) getLikedDislikedTracks(liked bool) (*schema.TracksLibrary, error) {
	var endP = "likes"
	if !liked {
		endP = "dislikes"
	}

	var endpoint = genApiPath([]string{"users", c.userId, endP, "tracks"})
	var data = &schema.TypicalResponse[*schema.TracksLibrary]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Поставить лайк треку.
//
// POST /users/{userId}/likes/tracks/add
func (c Client) LikeTrack(track *schema.Track) error {
	if track == nil {
		return ErrNilTrack
	}
	var body = schema.LikeTrackRequestBody{
		TrackId: track.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "tracks", "add"})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Поставить лайки трекам.
//
// POST /users/{userId}/likes/tracks/add-multiple
func (c Client) LikeTracks(tracks []*schema.Track) error {
	return c.likeUnlikeTracks(tracks, true)
}

// Убрать лайки с треков.
//
// POST /users/{userId}/likes/tracks/remove
func (c Client) UnlikeTracks(tracks []*schema.Track) error {
	return c.likeUnlikeTracks(tracks, false)
}

// Поставить лайки трекам.
//
// Убрать лайки с треков.
func (c Client) likeUnlikeTracks(tracks []*schema.Track, like bool) error {
	if tracks == nil {
		return ErrNilTracks
	}
	var body = schema.LikeUnlikeTracksRequestBody{}
	body.SetIds(tracks)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endEndPoint = "add-multiple"
	if !like {
		endEndPoint = "remove"
	}
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "tracks", endEndPoint})
	var data = &schema.TypicalResponse[any]{}
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Получить трек по id.
func (c Client) GetTrackById(trackId int64) ([]*schema.Track, error) {
	var endpoint = genApiPath([]string{"tracks", i2s(trackId)})
	var data = &schema.TypicalResponse[[]*schema.Track]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить треки по id.
func (c Client) GetTracksByIds(trackIds []int64) ([]*schema.Track, error) {
	if trackIds == nil {
		return nil, ErrNilTrackIds
	}
	var body = schema.GetTracksByIdsRequestBody{
		TrackIds: trackIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"tracks"})
	var data = &schema.TypicalResponse[[]*schema.Track]{}
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).SetResult(data).Post(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить информацию о загрузке трека.
func (c Client) GetTrackDownloadInfo(t *schema.Track) ([]*schema.TrackDownloadInfo, error) {
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "download-info"})
	var data = &schema.TypicalResponse[[]*schema.TrackDownloadInfo]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получение дополнительной информации о треке (Текст песни, видео, и т.д.).
func (c Client) GetTrackSupplement(t *schema.Track) (*schema.Supplement, error) {
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "supplement"})
	var data = &schema.TypicalResponse[*schema.Supplement]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получение похожих треков.
func (c Client) GetSimilarTracks(t *schema.Track) (*schema.SimilarTracks, error) {
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "similar"})
	var data = &schema.TypicalResponse[*schema.SimilarTracks]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}
