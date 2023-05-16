package schema

import (
	"encoding/json"
	"net/url"
	"time"
)

// Тип станции.
type RotorStationType string

const (
	// Жанровая станция.
	RotorStationTypeGenre = "genre"

	// Локальная станция. Типа "что слушают в Москве?".
	RotorStationTypeLocal = "local"
)

// Ограничение станции по языку.
type RotorLanguageRestriction string

const (
	// Любой.
	RotorLanguageRestrictionAny RotorLanguageRestriction = "any"

	// Русский.
	RotorLanguageRestrictionRussian RotorLanguageRestriction = "russian"

	// Иностранный.
	RotorLanguageRestrictionNotRussian RotorLanguageRestriction = "not-russian"
)

// Ограничение станции по характеру.
type RotorDiversityRestriction string

const (
	// Все (a.k.a Любое).
	RotorDiversityRestrictionDefault RotorDiversityRestriction = "default"

	// Редкие.
	RotorDiversityRestrictionDiverse RotorDiversityRestriction = "diverse"

	// Незнакомое.
	RotorDiversityRestrictionDiscover RotorDiversityRestriction = "diverse"

	// Любимые (a.k.a Любимое).
	RotorDiversityRestrictionFavorite RotorDiversityRestriction = "favorite"

	// Популярные (a.k.a Популярное).
	RotorDiversityRestrictionPopular RotorDiversityRestriction = "popular"
)

// Ограничение станции по настроению.
type RotorMoodEnergyRestriction string

const (
	// Борое.
	RotorMoodEnergyRestrictionActive = "active"

	// Веселое.
	RotorMoodEnergyRestrictionFun = "fun"

	// Спокойное.
	RotorMoodEnergyRestrictionCalm = "calm"

	// Грустное.
	RotorMoodEnergyRestrictionSad = "sad"

	// Любое.
	RotorMoodEnergyRestrictionAll = "all"
)

// Сообщение о начале прослушивания радио, начале и конце трека, его пропуска.
type RotorStationFeedbackType string

const (
	// Запустили радио.
	RotorStationFeedbackTypeRadioStarted RotorStationFeedbackType = "radioStarted"

	// Начался трек.
	RotorStationFeedbackTypeTrackStarted RotorStationFeedbackType = "trackStarted"

	// Трек закончился.
	RotorStationFeedbackTypeTrackFinished RotorStationFeedbackType = "trackFinished"

	// Трек пропущен.
	RotorStationFeedbackTypeSkip RotorStationFeedbackType = "skip"
)

type (
	// Список радио.
	RotorDashboard struct {
		DashboardID string `json:"dashboardId"`

		// Станции.
		Stations []struct {
			Station *RotorStation `json:"station"`

			Settings RotorSettings `json:"settings"`

			Settings2 RotorSettings2 `json:"settings2"`

			Explanation string `json:"explanation"`

			AdParams *RotorAdParams `json:"adParams"`

			// Пример: "Моя волна".
			RupTitle string `json:"rupTitle"`

			// Пример: "Волна подстраивается под жанр и вас. Слушайте только то, что нравится!".
			RupDescription string `json:"rupDescription"`
		} `json:"stations"`
		Pumpkin bool `json:"pumpkin"`
	}

	// Настройки станции.
	RotorSettings struct {
		Language  RotorLanguageRestriction  `json:"language"`
		Mood      uint8                     `json:"mood"`
		Energy    uint8                     `json:"energy"`
		Diversity RotorDiversityRestriction `json:"diversity"`
	}

	// Настройки станции 2 (вторая версия?).
	RotorSettings2 struct {
		Language   RotorLanguageRestriction   `json:"language"`
		MoodEnergy RotorMoodEnergyRestriction `json:"moodEnergy"`
		Diversity  RotorDiversityRestriction  `json:"diversity"`
	}

	// Станция радио.
	RotorStation struct {
		ID       RotorStationID  `json:"id"`
		ParentID *RotorStationID `json:"parentId"`

		// Название. Например: "Прогрессив-метал".
		Name string `json:"name"`

		// Иконка станции.
		Icon struct {
			// Цвет фона в HEX формате. Например: #9D65A9.
			BackgroundColor string `json:"backgroundColor"`

			// Ссылка на иконку в avatars.yandex.net.
			ImageURL string `json:"imageUrl"`
		} `json:"icon"`

		// см. Icon.
		MtsIcon struct {
			BackgroundColor string `json:"backgroundColor"`
			ImageURL        string `json:"imageUrl"`
		} `json:"mtsIcon"`

		// Ссылка на какую-то картинку в avatars.yandex.net.
		FullImageURL string `json:"fullImageUrl"`

		// см. FullImageURL.
		MtsFullImageURL string `json:"mtsFullImageUrl"`

		// Пример: "genre-metal_progmetal"
		IDForFrom string `json:"idForFrom"`

		// Доступные настройки станции.
		//
		// Обратите внимание: это не сами настройки, а лишь доступные значения.
		//
		// То есть если вы хотите изменить настройки станции, обращайте внимания на
		// доступные значения, которые обозначены тут.
		Restrictions struct {
			// По языку.
			Language RotorEnum[RotorLanguageRestriction] `json:"language"`

			// По настроению.
			Mood RotorDiscreteScale `json:"mood"`

			// По энергии.
			Energy RotorDiscreteScale `json:"energy"`

			// По характеру.
			Diversity RotorEnum[RotorDiversityRestriction] `json:"diversity"`
		} `json:"restrictions"`
		Restrictions2 struct {
			Diversity  RotorEnum[RotorDiversityRestriction]  `json:"diversity"`
			MoodEnergy RotorEnum[RotorMoodEnergyRestriction] `json:"moodEnergy"`
			Language   RotorEnum[RotorLanguageRestriction]   `json:"language"`
		} `json:"restrictions2"`
	}

	RotorEnum[T RotorLanguageRestriction | RotorDiversityRestriction | RotorMoodEnergyRestriction] struct {
		// "enum".
		Type string `json:"type"`

		// Пример: "По характеру", "По языку".
		Name string `json:"name"`

		PossibleValues []struct {
			// Само ограничение.
			Value T `json:"value"`

			// Имя ограничения (используется в UI).
			Name string `json:"name"`

			// Ссылка на avatars.mds.yandex.net.
			//
			// (может быть) Доступно в Restrictions2.
			ImageURL *string `json:"imageUrl"`

			// Пример: "settingDiversity:favorite", "settingMoodEnergy:fun".
			//
			// (может быть) Доступно в Restrictions2.
			SerializedSeed *string `json:"serializedSeed"`

			// (может быть) Доступно в Restrictions2.
			Unspecified *bool `json:"unspecified"`
		} `json:"possibleValues"`
	}

	RotorDiscreteScale struct {
		// "discrete-scale".
		Type string `json:"type"`

		// Пример: "Энергичность", "Под настроение".
		Name string `json:"name"`

		Min struct {
			// Пример: 1.
			Value uint8 `json:"value"`

			// Пример: "Спокойнее", "Грустнее".
			Name string `json:"name"`
		} `json:"min"`
		Max struct {
			// Пример: 4.
			Value uint8 `json:"value"`

			// Пример: "Бодрее", "Веселее".
			Name string `json:"name"`
		} `json:"max"`
	}

	// Треки станции.
	RotorStationTracks struct {
		ID       RotorStationID `json:"id"`
		Sequence []struct {
			Track *Track `json:"track"`
			// Параметры трека.
			TrackParameters struct {
				// Кол-во ударов в минуту.
				BPM uint16 `json:"bpm"`
				Hue uint16 `json:"hue"`
				// Какой-то супер-точный показатель энергии (?).
				Energy float64 `json:"energy"`
			} `json:"trackParameters"`
			// Трек лайкнут?
			Liked bool `json:"liked"`
		} `json:"sequence"`
		BatchID string `json:"batchId"`
		Pumpkin bool   `json:"pumpkin"`
	}

	// Статус аккаунта в Радио.
	RotorAccountStatus struct {
		Account      Account            `json:"account"`
		Permissions  AccountPermissions `json:"permissions"`
		Subscription struct {
			End time.Time `json:"end"`
		} `json:"subscription"`
		SkipsPerHour  int  `json:"skipsPerHour"`
		StationExists bool `json:"stationExists"`
		Plus          Plus `json:"plus"`
		PremiumRegion int  `json:"premiumRegion"`
	}

	// Одна станция из RotorStationsList.
	RotorStationList struct {
		Station *RotorStation `json:"station"`
		Data    struct {
			Title       string    `json:"title"`
			Description string    `json:"description"`
			ImageURI    string    `json:"imageUri"`
			Artists     []*Artist `json:"artists"`
		} `json:"data"`
		Settings       RotorSettings  `json:"settings"`
		Settings2      RotorSettings2 `json:"settings2"`
		AdParams       *RotorAdParams `json:"adParams"`
		RupTitle       string         `json:"rupTitle"`
		RupDescription string         `json:"rupDescription"`
		// Свое имя станции.
		CustomName *string `json:"customName"`
	}

	// Аудиореклама.
	RotorAdParams struct {
		PartnerID   string `json:"partnerId"`
		CategoryID  string `json:"categoryId"`
		PageRef     string `json:"pageRef"`
		TargetRef   string `json:"targetRef"`
		GenreID     int    `json:"genreId"`
		GenreName   string `json:"genreName"`
		OtherParams string `json:"otherParams"`
		AdVolume    int    `json:"adVolume"`
	}

	// Информация о станции.
	RotorStationInfo struct {
		Station *RotorStation
		Data    *struct {
			Artists []*Artist `json:"artists"`
		} `json:"data"`
		Settings       RotorSettings  `json:"settings"`
		Settings2      RotorSettings2 `json:"settings2"`
		AdParams       *RotorAdParams `json:"adParams"`
		RupTitle       string         `json:"rupTitle"`
		RupDescription string         `json:"rupDescription"`
	}

	// GET /rotor/stations/list
	GetRotorStationsListQueryParams struct {
		// Язык, на котором будет информация о станциях (ISO 639-1).
		Language *string `url:"language,omitempty"`
	}
)

// ID станции. Отличие от ID других структур в формате "тип_станции:тег_станции".
type RotorStationID struct {
	// Тип станции.
	Type RotorStationType `json:"type"`

	// Зависит от Type.
	//
	// Название жанра. Например: "progmetal".
	//
	// Название города. Например: "saint-petersburg".
	Tag string `json:"tag"`
}

func (r RotorStationID) String() string {
	return string(r.Type) + ":" + r.Tag
}

// GET /rotor/station/{type:tag}/tracks
type GetRotorStationTracksQueryParams struct {
	// Использовать ли второй набор настроек.
	// Все официальные клиенты выполняют запросы с settings2 = True.
	Settings2 bool `url:"settings2"`

	// Уникальной идентификатор трека, который только что был.
	Queue string `url:"queue,omitempty"`
}

func (g *GetRotorStationTracksQueryParams) SetLastTrack(tr *Track) {
	if tr == nil {
		return
	}
	g.Queue = string(tr.ID)
}

// POST /rotor/station/{type:tag}/feedback
//
// Тут используются как QueryParams, так и json Body.
//
// Используйте Fill для заполнения полей, GetJson для получения Body, и GetQuery для параметров.
type RotorStationFeedbackRequestBodyQueryString struct {
	// Уникальный идентификатор партии треков. Возвращается при получении треков.
	//
	// см. RotorStationTracks.
	//
	// это поле является querystring.
	BatchID string `json:"-"`
	// Тип отправляемого фидбека.
	Type RotorStationFeedbackType `json:"type"`
	// Текущее дата и время.
	Timestamp time.Time `json:"timestamp"`
	// Источник воспроизведения радио.
	//
	// пример: "mobile-radio-user-123456789".
	From string `json:"from,omitempty"`
	// Сколько было проиграно секунд трека перед действием.
	TotalPlayedSeconds float32 `json:"totalPlayedSeconds,omitempty"`
	// "track_id:album_id"
	TrackID string `json:"trackId,omitempty"`
	// https://github.com/MarshalX/yandex-music-api/blob/main/yandex_music/client.py#L1251
}

// Заполнить поля. Потом используйте GetJson и GetQueryString.
func (r *RotorStationFeedbackRequestBodyQueryString) Fill(fType RotorStationFeedbackType, tracks *RotorStationTracks, currentTrack *Track, TotalPlayedSeconds float32) {
	if tracks == nil || currentTrack == nil {
		return
	}
	r.BatchID = tracks.BatchID
	r.Type = fType
	r.Timestamp = time.Now()
	r.TotalPlayedSeconds = TotalPlayedSeconds
	if len(currentTrack.Albums) > 0 {
		r.TrackID = string(currentTrack.ID) + ":" + string(currentTrack.Albums[0].ID)
	}
}

// Получить Body.
func (r *RotorStationFeedbackRequestBodyQueryString) GetJson() (string, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Получить query params.
func (r *RotorStationFeedbackRequestBodyQueryString) GetQuery() (url.Values, error) {
	type Query struct {
		BatchID string `url:"batch-id"`
	}
	val := Query{BatchID: r.BatchID}
	return ParamsToValues(val)
}
