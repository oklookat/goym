package goym

import (
	"context"

	"github.com/oklookat/goym/schema"
)

// Получить рекомендованные станции.
func (c Client) GetRotorDashboard(ctx context.Context) (schema.Response[*schema.RotorDashboard], error) {
	// GET /rotor/stations/dashboard
	endpoint := genApiPath("rotor", "stations", "dashboard")
	data := &schema.Response[*schema.RotorDashboard]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить треки со станции.
//
// lastTrack - последний трек со станции. Может быть nil.
func (c Client) GetRotorStationTracks(ctx context.Context, st *schema.RotorStation, lastTrack *schema.Track) (schema.Response[*schema.RotorStationTracks], error) {
	// GET /rotor/station/{type:tag}/tracks
	data := &schema.Response[*schema.RotorStationTracks]{}
	if st == nil {
		return *data, nil
	}
	body := schema.GetRotorStationTracksQueryParams{
		Settings2: true,
	}
	body.SetLastTrack(lastTrack)
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("rotor", "station", st.ID.String(), "tracks")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить информацию об аккаунте в радио.
func (c Client) GetRotorAccountStatus(ctx context.Context) (schema.Response[*schema.RotorAccountStatus], error) {
	// GET /rotor/account/status
	endpoint := genApiPath("rotor", "account", "status")
	data := &schema.Response[*schema.RotorAccountStatus]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить все станции с настройками пользователя.
//
// language: язык ответа (ISO 639-1). Может быть nil.
func (c Client) GetRotorStationsList(ctx context.Context, language *string) (schema.Response[[]schema.RotorStationList], error) {
	// GET /rotor/stations/list
	data := &schema.Response[[]schema.RotorStationList]{}
	body := schema.GetRotorStationsListQueryParams{
		Language: language,
	}
	vals, err := schema.ParamsToValues(body)
	if err != nil {
		return *data, err
	}

	endpoint := genApiPath("rotor", "stations", "list")
	resp, err := c.Http.R().SetError(data).SetResult(data).SetQueryParams(vals).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
}

// Получить информацию о станции.
func (c Client) GetRotorStationInfo(ctx context.Context, st *schema.RotorStation) (schema.Response[[]schema.RotorStationInfo], error) {
	// GET /rotor/station/{type:tag}/info
	data := &schema.Response[[]schema.RotorStationInfo]{}

	if st == nil {
		return *data, nil
	}
	endpoint := genApiPath("rotor", "station", st.ID.String(), "info")
	resp, err := c.Http.R().SetError(data).SetResult(data).Get(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return *data, err
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

	endpoint := genApiPath("rotor", "station", st.ID.String(), "feedback")
	data := &schema.Response[string]{}
	resp, err := c.Http.R().SetError(data).SetResult(data).SetJsonString(jsonBody).SetQueryParams(params).Post(ctx, endpoint)
	if err == nil {
		err = checkResponse(resp, data)
	}
	return data.Result, err
}
