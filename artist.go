package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить лайкнутых артистов.
func (c Client) GetLikedArtists(ctx context.Context) ([]*schema.Artist, error) {
	// GET /users/{userId}/likes/artists
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists"})
	var data = &schema.Response[[]*schema.Artist]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Лайкнуть артиста.
func (c Client) LikeArtist(ctx context.Context, a *schema.Artist) error {
	// POST /users/{userId}/likes/artists/add
	if a == nil {
		return nil
	}

	var body = schema.LikeArtistRequestBody{
		ArtistId: a.ID,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return err
	}

	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", "add"})
	var data = &schema.Response[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Убрать лайк с артиста.
func (c Client) UnlikeArtist(ctx context.Context, ar *schema.Artist) error {
	// POST /users/{userId}/likes/artists/{artistId}/remove
	if ar == nil {
		return nil
	}
	var endpoint = genApiPath([]string{"users", c.userId, "likes", "artists", ar.ID.String(), "remove"})
	var data = &schema.Response[any]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}

	return err
}

// Получить список треков артиста по его ID.
func (c Client) GetArtistTracks(ctx context.Context, artistId schema.UniqueID, page uint16, pageSize uint16) (*schema.ArtistTracksPaged, error) {
	// GET /artists/{artistId}/tracks
	var body = schema.GetArtistTracksQueryParams{
		Page:     page,
		PageSize: pageSize,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"artists", artistId.String(), "tracks"})
	var data = &schema.Response[*schema.ArtistTracksPaged]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить альбомы артиста по его ID.
func (c Client) GetArtistAlbums(ctx context.Context, artistId schema.UniqueID, page uint16, pageSize uint16, sortBy schema.SortBy) (*schema.ArtistAlbumsPaged, error) {
	// GET /artists/{artistId}/direct-albums
	var body = schema.GetArtistAlbumsQueryParams{
		Page:     page,
		PageSize: pageSize,
		SortBy:   sortBy,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	var endpoint = genApiPath([]string{"artists", artistId.String(), "direct-albums"})
	var data = &schema.Response[*schema.ArtistAlbumsPaged]{}
	resp, err := c.self.R().SetError(data).SetResult(data).SetFormUrlValues(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить лучшие треки артиста.
func (c Client) GetArtistTopTracks(ctx context.Context, ar *schema.Artist) (*schema.ArtistTopTracks, error) {
	// GET /artists/{artistId}/track-ids-by-rating
	if ar == nil {
		return nil, nil
	}
	var endpoint = genApiPath([]string{"artists", ar.ID.String(), "track-ids-by-rating"})
	var data = &schema.Response[*schema.ArtistTopTracks]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить полную информацию об артисте.
func (c Client) GetArtistInfo(ctx context.Context, ar *schema.Artist) (*schema.ArtistBriefInfo, error) {
	// GET /artists/{artistId}/brief-info
	if ar == nil {
		return nil, nil
	}
	var endpoint = genApiPath([]string{"artists", ar.ID.String(), "brief-info"})
	var data = &schema.Response[*schema.ArtistBriefInfo]{}
	resp, err := c.self.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
