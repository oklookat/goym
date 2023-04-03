package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить лайкнутых артистов.
func (c Client) LikedArtists(ctx context.Context) ([]*schema.Artist, error) {
	// GET /users/{userId}/likes/artists
	endpoint := genApiPath("users", c.userId, "likes", "artists")
	data := &schema.Response[[]*schema.Artist]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Лайкнуть артиста по ID.
func (c Client) LikeArtist(ctx context.Context, id schema.ID) error {
	// POST /users/{userId}/likes/artists/add
	body := schema.LikeArtistRequestBody{
		ArtistId: id,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	endpoint := genApiPath("users", c.userId, "likes", "artists", "add")
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Убрать лайк с артиста по ID.
func (c Client) UnlikeArtist(ctx context.Context, id schema.ID) error {
	// POST /users/{userId}/likes/artists/{artistId}/remove
	endpoint := genApiPath("users", c.userId, "likes", "artists", id.String(), "remove")
	data := &schema.Response[any]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Получить список треков артиста по его ID.
func (c Client) ArtistTracks(ctx context.Context, id schema.ID, page uint16, pageSize uint16) (*schema.ArtistTracksPaged, error) {
	// GET /artists/{artistId}/tracks
	body := schema.GetArtistTracksQueryParams{
		Page:     page,
		PageSize: pageSize,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("artists", id.String(), "tracks")
	data := &schema.Response[*schema.ArtistTracksPaged]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить альбомы артиста по его ID.
//
// Приложение под Windows в качестве pageSize обычно использует 50.
func (c Client) ArtistAlbums(ctx context.Context, id schema.ID, page uint16, pageSize uint16, sortBy schema.SortBy, sortOrder schema.SortOrder) (*schema.ArtistAlbumsPaged, error) {
	// GET /artists/{artistId}/direct-albums
	body := schema.GetArtistAlbumsQueryParams{
		Page:      page,
		PageSize:  pageSize,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath("artists", id.String(), "direct-albums")
	data := &schema.Response[*schema.ArtistAlbumsPaged]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить лучшие треки артиста по его ID.
func (c Client) ArtistTopTracks(ctx context.Context, id schema.ID) (*schema.ArtistTopTracks, error) {
	// GET /artists/{artistId}/track-ids-by-rating
	endpoint := genApiPath("artists", id.String(), "track-ids-by-rating")
	data := &schema.Response[*schema.ArtistTopTracks]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить полную информацию об артисте по его ID.
func (c Client) ArtistInfo(ctx context.Context, id schema.ID) (*schema.ArtistBriefInfo, error) {
	// GET /artists/{artistId}/brief-info
	endpoint := genApiPath("artists", id.String(), "brief-info")
	data := &schema.Response[*schema.ArtistBriefInfo]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
