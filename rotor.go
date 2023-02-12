package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить рекомендованные станции.
func (c Client) GetRotorDashboard(ctx context.Context) (*schema.RotorDashboard, error) {
	// GET /rotor/stations/dashboard
	endpoint := genApiPath([]string{"rotor", "stations", "dashboard"})
	data := &schema.Response[*schema.RotorDashboard]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить треки со станции.
//
// lastTrack - последний трек со станции. Может быть nil.
func (c Client) GetRotorStationTracks(ctx context.Context, st *schema.RotorStation, lastTrack *schema.Track) (*schema.RotorStationTracks, error) {
	// GET /rotor/station/{type:tag}/tracks
	if st == nil {
		return nil, nil
	}
	body := schema.GetRotorStationTracksQueryParams{
		Settings2: true,
	}
	body.SetLastTrack(lastTrack)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"rotor", "station", st.ID.String(), "tracks"})
	data := &schema.Response[*schema.RotorStationTracks]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить информацию об аккаунте в радио.
func (c Client) GetRotorAccountStatus(ctx context.Context) (*schema.RotorAccountStatus, error) {
	// GET /rotor/account/status
	endpoint := genApiPath([]string{"rotor", "account", "status"})
	data := &schema.Response[*schema.RotorAccountStatus]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить все станции с настройками пользователя.
//
// language: язык ответа (ISO 639-1). Может быть nil.
func (c Client) GetRotorStationsList(ctx context.Context, language *string) ([]*schema.RotorStationList, error) {
	// GET /rotor/stations/list
	body := schema.GetRotorStationsListQueryParams{
		Language: language,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return nil, err
	}

	endpoint := genApiPath([]string{"rotor", "stations", "list"})
	data := &schema.Response[[]*schema.RotorStationList]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Получить информацию о станции.
func (c Client) GetRotorStationInfo(ctx context.Context, st *schema.RotorStation) ([]*schema.RotorStationInfo, error) {
	// GET /rotor/station/{type:tag}/info
	if st == nil {
		return nil, nil
	}
	endpoint := genApiPath([]string{"rotor", "station", st.ID.String(), "info"})
	data := &schema.Response[[]*schema.RotorStationInfo]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}

// Отправка ответной реакции на происходящее при прослушивании радио.
//
// Вариант ответа: "ok" или сообщение об ошибке типа "unknown: omitted".
func (c Client) RotorStationFeedback(ctx context.Context, st *schema.RotorStation, fType schema.RotorStationFeedbackType, tracks *schema.RotorStationTracks, currentTrack *schema.Track, totalPlayedSeconds float32) (string, error) {
	// POST /rotor/station/{type:tag}/feedback
	if st == nil || tracks == nil || currentTrack == nil {
		return "", nil
	}
	body := schema.RotorStationFeedbackRequestBodyQueryString{}
	body.Fill(fType, tracks, currentTrack, totalPlayedSeconds)
	jsonBody, err := body.GetJson()
	if err != nil {
		return "", err
	}
	params, err := body.GetQuery()
	if err != nil {
		return "", err
	}

	endpoint := genApiPath([]string{"rotor", "station", st.ID.String(), "feedback"})
	data := &schema.Response[string]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetJsonString(jsonBody).SetQueryParams(params).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
