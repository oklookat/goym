package schema

// GET /rotor/station/{rotorID}/tracks
type GetRotorTracksQueryParams struct {
	// Использовать ли второй набор настроек.
	// Все официальные клиенты выполняют запросы с settings2 = True.
	Settings2 bool `url:"settings2"`

	// Уникальной идентификатор трека, который только что был.
	Queue string `url:"queue"`
}

// GET /rotor/stations/list
type GetRotorStationsQueryParams struct {
	// Язык, на котором будет информация о станциях (ru, например)
	Language string `url:"language"`
}

// POST /rotor/station/{station}/feedback
//
// Получите body через Get().
type RotorStationFeedbackRequestBody struct{}

// Получить Request Body.
func (r RotorStationFeedbackRequestBody) Get() string {
	// всё нормально, так и должно быть.
	return "{}"
}
