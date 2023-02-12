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
	ld := "likes"
	if !liked {
		ld = "dislikes"
	}

	endpoint := genApiPath([]string{"users", c.userId, ld, "tracks"})
	data := &schema.Response[*schema.TracksLibrary]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Лайкнуть трек.
func (c Client) LikeTrack(ctx context.Context, track *schema.Track) error {
	// POST /users/{userId}/likes/tracks/add
	if track == nil {
		return nil
	}
	body := schema.LikeTrackRequestBody{
		TrackId: track.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endpoint := genApiPath([]string{"users", c.userId, "likes", "tracks", "add"})
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
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
		return nil
	}
	body := schema.LikeUnlikeTracksRequestBody{}
	body.SetIds(tracks)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endEndPoint := "add-multiple"
	if !like {
		endEndPoint = "remove"
	}
	endpoint := genApiPath([]string{"users", c.userId, "likes", "tracks", endEndPoint})
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return err
}

// Получить трек по id.
func (c Client) GetTrackById(ctx context.Context, trackId schema.UniqueID) ([]*schema.Track, error) {
	// GET /tracks/{trackId}
	endpoint := genApiPath([]string{"tracks", trackId.String()})
	data := &schema.Response[[]*schema.Track]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить треки по id.
func (c Client) GetTracksByIds(ctx context.Context, trackIds []schema.UniqueID) ([]*schema.Track, error) {
	// POST /tracks
	if trackIds == nil {
		return nil, nil
	}
	body := schema.GetTracksByIdsRequestBody{
		TrackIds: trackIds,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"tracks"})
	data := &schema.Response[[]*schema.Track]{}
	resp, err := c.Http.R().SetError(data).SetFormUrlValues(vals).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить информацию о загрузке трека.
func (c Client) GetTrackDownloadInfo(ctx context.Context, tr *schema.Track) ([]*schema.TrackDownloadInfo, error) {
	// GET /tracks/{trackId}/download-info
	if tr == nil {
		return nil, nil
	}
	endpoint := genApiPath([]string{"tracks", tr.ID.String(), "download-info"})
	data := &schema.Response[[]*schema.TrackDownloadInfo]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить дополнительную информацию о треке (текст песни, видео, etc).
func (c Client) GetTrackSupplement(ctx context.Context, tr *schema.Track) (*schema.Supplement, error) {
	// GET /tracks/{trackId}/supplement
	if tr == nil {
		return nil, nil
	}
	endpoint := genApiPath([]string{"tracks", tr.ID.String(), "supplement"})
	data := &schema.Response[*schema.Supplement]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить похожие треки.
func (c Client) GetSimilarTracks(ctx context.Context, tr *schema.Track) (*schema.SimilarTracks, error) {
	// GET /tracks/{trackId}/similar
	if tr == nil {
		return nil, nil
	}
	endpoint := genApiPath([]string{"tracks", tr.ID.String(), "similar"})
	data := &schema.Response[*schema.SimilarTracks]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
