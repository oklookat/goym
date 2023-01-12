package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить лайкнутые треки.
func (c Client) GetLikedTracks(ctx context.Context) (*schema.TracksLibrary, error) {
	return c.getLikedDislikedTracks(ctx, true)
}

// Получить дизлайкнутые треки.
func (c Client) GetDislikedTracks(ctx context.Context) (*schema.TracksLibrary, error) {
	return c.getLikedDislikedTracks(ctx, false)
}

func (c Client) getLikedDislikedTracks(ctx context.Context, liked bool) (*schema.TracksLibrary, error) {
	// GET /users/{userId}/likes/tracks
	// ||
	// GET /users/{userId}/dislikes/tracks
	var ld = "likes"
	if !liked {
		ld = "dislikes"
	}

	var endpoint = genApiPath([]string{"users", c.userId, ld, "tracks"})
	var data = &schema.TypicalResponse[*schema.TracksLibrary]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Лайкнуть трек.
func (c Client) LikeTrack(ctx context.Context, track *schema.Track) error {
	// POST /users/{userId}/likes/tracks/add
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
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Лайкнуть треки.
//
// Используйте LikeTrack() для лайка одного трека.
func (c Client) LikeTracks(ctx context.Context, tracks []*schema.Track) error {
	return c.likeUnlikeTracks(ctx, tracks, true)
}

// Убрать лайки с трека(ов).
func (c Client) UnlikeTracks(ctx context.Context, tracks []*schema.Track) error {
	return c.likeUnlikeTracks(ctx, tracks, false)
}

func (c Client) likeUnlikeTracks(ctx context.Context, tracks []*schema.Track, like bool) error {
	// POST /users/{userId}/likes/tracks/add-multiple
	// ||
	// POST /users/{userId}/likes/tracks/remove
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
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return err
}

// Получить трек по id.
func (c Client) GetTrackById(ctx context.Context, trackId int64) ([]*schema.Track, error) {
	// GET /tracks/{trackId}
	var endpoint = genApiPath([]string{"tracks", i2s(trackId)})
	var data = &schema.TypicalResponse[[]*schema.Track]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить треки по id.
func (c Client) GetTracksByIds(ctx context.Context, trackIds []int64) ([]*schema.Track, error) {
	// POST /tracks
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
	resp, err := c.self.R().SetError(data).SetFormUrlValues(vals).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить информацию о загрузке трека.
func (c Client) GetTrackDownloadInfo(ctx context.Context, t *schema.Track) ([]*schema.TrackDownloadInfo, error) {
	// GET /tracks/{trackId}/download-info
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "download-info"})
	var data = &schema.TypicalResponse[[]*schema.TrackDownloadInfo]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить дополнительную информацию о треке (текст песни, видео, etc).
func (c Client) GetTrackSupplement(ctx context.Context, t *schema.Track) (*schema.Supplement, error) {
	// GET /tracks/{trackId}/supplement
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "supplement"})
	var data = &schema.TypicalResponse[*schema.Supplement]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}

// Получить похожие треки.
func (c Client) GetSimilarTracks(ctx context.Context, t *schema.Track) (*schema.SimilarTracks, error) {
	// GET /tracks/{trackId}/similar
	if t == nil {
		return nil, ErrNilTrack
	}
	var endpoint = genApiPath([]string{"tracks", i2s(t.ID), "similar"})
	var data = &schema.TypicalResponse[*schema.SimilarTracks]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkTypicalResponse(resp, data)
	}
	return data.Result, err
}
