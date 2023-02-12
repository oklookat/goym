package schema

import (
	"encoding/json"
	"time"
)

type (
	// Нормализация.
	//
	// https://en.wikipedia.org/wiki/EBU_R_128
	R128 struct {
		I float32 `json:"i"`

		// True Peak.
		Tp float32 `json:"tp"`
	}

	TrackItem struct {
		ID            UniqueID  `json:"id"`
		Track         *Track    `json:"track"`
		Timestamp     time.Time `json:"timestamp"`
		OriginalIndex uint16    `json:"originalIndex"`
		Recent        bool      `json:"recent"`
	}

	// Дополнительная информация о треке.
	Supplement struct {
		// Уникальный идентификатор дополнительной информации.
		ID     string           `json:"id"`
		Lyrics *Lyrics          `json:"lyrics"`
		Videos *VideoSupplement `json:"videos"`

		// Доступно ли радио.
		RadioIsAvailable bool `json:"radioIsAvailable"`

		// Полное описание эпизода подкаста.
		Description string `json:"description"`
	}

	// Текст трека.
	Lyrics struct {
		// Уникальный идентификатор текста трека.
		ID UniqueID `json:"id"`

		// Первые строки текст песни.
		Lyrics string `json:"lyrics"`

		// Есть ли права.
		HasRights bool `json:"hasRights"`

		// Текст песни.
		FullLyrics string `json:"fullLyrics"`

		// Язык текста.
		TextLanguage string `json:"textLanguage"`

		// Доступен ли перевод.
		ShowTranslation bool `json:"showTranslation"`

		// Ссылка на источник перевода. Обычно genius.com.
		Url string `json:"url"`
	}

	// Список треков.
	TracksLibrary struct {
		Library struct {
			// Уникальный идентификатор пользователя.
			Uid UniqueID `json:"uid"`

			Revision RevisionID `json:"revision"`

			// Список треков в укороченной версии.
			Tracks []*TrackShort `json:"tracks"`
		} `json:"library"`
	}

	// Список похожих треков на другой трек.
	SimilarTracks struct {
		Track *Track `json:"track"`
		// Похожие треки.
		SimilarTracks []*Track `json:"similarTracks"`
	}

	// Информация о вариантах загрузки трека.
	TrackDownloadInfo struct {
		// Кодек аудиофайла (mp3, aac).
		Codec string `json:"codec"`

		// Усиление.
		Gain bool `json:"gain"`

		// Предварительный просмотр.
		Preview bool `json:"preview"`

		// Ссылка на XML документ содержащий данные для загрузки трека
		//
		// При переходе по этому
		// URL также необходимо иметь auth header. Без него или будет 401, или будет массив с mp3/128.
		//
		// Если собираетесь сделать загрузку mp3, смотрите в эту сторону:
		// https://github.com/MarshalX/yandex-music-api/blob/main/yandex_music/download_info.py
		DownloadInfoUrl string `json:"downloadInfoUrl"`

		// Прямая ли ссылка.
		Direct bool `json:"direct"`

		// Битрейт аудиофайла в кбит/с.
		BitrateInKbps uint16 `json:"bitrateInKbps"`
	}

	// POST /users/{userId}/likes/tracks/add
	LikeTrackRequestBody struct {
		// ID трека.
		TrackId UniqueID `url:"track-id"`
	}

	// POST ​/tracks​
	GetTracksByIdsRequestBody struct {
		// ID треков.
		TrackIds []UniqueID `url:",track-ids"`

		// С позициями?
		WithPositions bool `url:"with-positions"`
	}

	// POST /play-audio
	//
	// (!) Я не проверял эти параметры.
	PlayAudioRequestBody struct {
		// Уникальный идентификатор трека.
		TrackId UniqueID `url:"track-id,omitempty"`

		// Проигрывается ли трек с кеша.
		FromCache bool `url:"from-cache,omitempty"`

		// Наименования клиента с которого происходит прослушивание.
		From string `url:"from"`

		// Уникальный идентификатор проигрывания.
		PlayId string `url:"play-id,omitempty"`

		// Уникальный идентификатор пользователя.
		Uid UniqueID `url:"uid,omitempty"`

		// Текущая дата и время в ISO.
		Timestamp time.Time `url:"timestamp,omitempty"`

		// Продолжительность трека в секундах.
		TrackLengthSeconds uint16 `url:"track-length-seconds,omitempty"`

		// Продолжительность трека в секундах.
		TotalPlayedSeconds uint16 `url:"total-played-seconds,omitempty"`

		// Продолжительность трека в секундах.
		EndPositionSeconds uint16 `url:"end-position-seconds,omitempty"`

		// Уникальный идентификатор альбома.
		AlbumId UniqueID `url:"album-id,omitempty"`

		// Уникальный идентификатор проигрывания.
		PlaylistId UniqueID `url:"playlist-id,omitempty"`

		// Текущая дата и время клиента в ISO.
		ClientNow string `url:"client-now,omitempty"`
	}
)

// Трек.
type Track struct {
	// Идентификатор трека.
	ID UniqueID `json:"-"`

	// Идентификатор подменного трека.
	//
	// ID и RealID совпадают в случаях:
	//
	// 1. Трек доступен для прослушивания.
	//
	// 2. Трек недоступен и не имеет идентичного трека для автозамены.
	RealID string `json:"realId"`

	// Название трека.
	Title string `json:"title"`

	// Лейбл.
	Major *Label `json:"major"`

	// Доступен для стриминга?
	Available bool `json:"available"`

	AvailableForPremiumUsers bool `json:"availableForPremiumUsers"`

	// (?) Трек могут послушать даже те, кто без подписки, или не вошел в аккаунт?
	AvailableFullWithoutPermission bool `json:"availableFullWithoutPermission"`

	// Например: ["bookmate"].
	AvailableForOptions []string `json:"availableForOptions"`

	// Адрес каталога, в котором хранится трек.
	StorageDir string `json:"storageDir"`

	// Продолжительность трека в миллисекундах.
	DurationMs uint64 `json:"durationMs"`

	// Размер трека в байтах.
	FileSize          uint64 `json:"fileSize"`
	R128              *R128  `json:"r128"`
	PreviewDurationMs uint16 `json:"previewDurationMs"`

	// Список исполнителей трека, в минимальной информации.
	Artists []*Artist `json:"artists"`

	// Список альбомов, в которые входит трек, в базовой информации.
	Albums []*Album `json:"albums"`

	// URI обложки.
	CoverUri string `json:"coverUri"`

	// Ссылка на загруженную обложку трека.
	OgImage string `json:"ogImage"`

	// Доступен ли текст трека.
	LyricsAvailable bool `json:"lyricsAvailable"`

	Type             string `json:"type"`
	RememberPosition bool   `json:"rememberPosition"`
	TrackSharingFlag string `json:"trackSharingFlag"`
	LyricsInfo       struct {
		HasAvailableSyncLyrics bool `json:"hasAvailableSyncLyrics"`
		HasAvailableTextLyrics bool `json:"hasAvailableTextLyrics"`
	} `json:"lyricsInfo"`
	// OWN.
	TrackSource    string `json:"trackSource"`
	AvailableAsRbt bool   `json:"availableAsRbt"`
	// Трек 18+?
	Explicit bool     `json:"explicit"`
	Regions  []string `json:"regions"`

	// Версия трека.
	Version *string `json:"version,omitempty"`
}

// В некоторых запросах ID может быть как строкой, так и числом.
//
// Надо привести ID к числу.
func (t *Track) UnmarshalID(id UniqueID, data []byte) error {
	type TrackFake Track
	// демаршал в TrackFake
	var faked TrackFake
	if err := json.Unmarshal(data, &faked); err != nil {
		return err
	}
	// копирование полей из TrackFake в Track,
	// только ID ставим сами
	*t = Track(faked)
	t.ID = id
	return nil
}

func (t *Track) UnmarshalJSON(data []byte) error {
	dem := func(id UniqueID, data []byte) error {
		type TrackFake Track
		var faked TrackFake
		if err := json.Unmarshal(data, &faked); err != nil {
			return err
		}
		*t = Track(faked)
		t.ID = id
		return nil
	}
	return unmarshalID(dem, data)
}

// Укороченная версия трека с неполными данными.
type TrackShort struct {
	// Уникальный идентификатор трека.
	ID UniqueID `json:"id"`

	// Уникальный идентификатор альбома.
	AlbumId UniqueID `json:"albumId"`

	// Дата.
	Timestamp time.Time `json:"timestamp"`
}

func (t *TrackShort) UnmarshalJSON(data []byte) error {
	type real struct {
		ID        string    `json:"id"`
		AlbumId   string    `json:"albumId"`
		Timestamp time.Time `json:"timestamp"`
	}
	realVal := &real{}
	if err := json.Unmarshal(data, realVal); err != nil {
		return err
	}

	var idUid UniqueID = 0
	if err := idUid.FromString(realVal.ID); err != nil {
		return err
	}
	t.ID = idUid

	var albumId UniqueID = 0
	if err := albumId.FromString(realVal.AlbumId); err != nil {
		return err
	}
	t.AlbumId = albumId

	t.Timestamp = realVal.Timestamp
	return nil
}

// POST /users/{userId}/likes/tracks/add-multiple
//
// POST /users/{userId}/likes/tracks/remove
//
// Доступен метод GetIds().
type LikeUnlikeTracksRequestBody struct {
	// ID треков.
	TrackIds []UniqueID `url:",track-ids"`
}

// Устанавливает ID в TrackIds. Если слайс треков == nil, ничего не делает.
func (l *LikeUnlikeTracksRequestBody) SetIds(tracks []*Track) {
	if tracks == nil {
		return
	}
	if l.TrackIds == nil {
		l.TrackIds = make([]UniqueID, 0)
	}
	for i := range tracks {
		l.TrackIds = append(l.TrackIds, tracks[i].ID)
	}
}
