package goym

import (
	"encoding/json"
	"strconv"
)

const (
	ApiUrl = "https://api.music.yandex.net"

	VisibilityPrivate = "private"
	VisibilityPublic  = "public"

	SearchTypeArtist   = "artist"
	SearchTypeAlbum    = "album"
	SearchTypeTrack    = "track"
	SearchTypePodcast  = "podcast"
	SearchTypePlaylist = "playlist"
	SearchTypeAll      = "all"
)

// Обычно ответ выглядит так.
type TypicalResponse[T any] struct {
	InvocationInfo *InvocationInfo `json:"invocationInfo"`

	// Если не nil, то поле result будет nil.
	Error *Error `json:"error"`

	Result T `json:"result"`
}

// Настройки пользователя.
type UserSettings struct {
	// ID.
	Uid int64 `json:"uid"`

	// Включен ли скробблинг last.fm?
	LastFmScrobblingEnabled   bool `json:"lastFmScrobblingEnabled"`
	FacebookScrobblingEnabled bool `json:"facebookScrobblingEnabled"`

	// (?) Включено ли рандомное воспроизведение треков?
	ShuffleEnabled bool `json:"shuffleEnabled"`

	// Добавлять новый трек в начало плейлиста?
	AddNewTrackOnPlaylistTop bool `json:"addNewTrackOnPlaylistTop"`

	// Громкость в процентах (example: 75).
	VolumePercents int `json:"volumePercents"`

	// Видимость музыкальной библиотеки.
	//
	// Используйте константы Visibility.
	UserMusicVisibility string `json:"userMusicVisibility"`

	// ???
	//
	// Используйте константы Visibility.
	UserSocialVisibility string `json:"userSocialVisibility"`

	AdsDisabled bool `json:"adsDisabled"`

	// example: 2019-04-14T14:55:50+00:00
	Modified string `json:"modified"`

	RbtDisabled string `json:"rbtDisabled"`

	// Тема оформления.
	//
	// black | default
	//
	// Example: black.
	Theme string `json:"theme"`

	AutoPlayRadio    bool `json:"autoPlayRadio"`
	SyncQueueEnabled bool `json:"syncQueueEnabled"`
}

// Что-то техническое.
type InvocationInfo struct {
	// (?) Время выполнения запроса в миллисекундах.
	//
	// string | int
	ExecDurationMillis any `json:"exec-duration-millis"`

	// Адрес какого-то сервера Яндекс.Музыки.
	Hostname string `json:"hostname"`

	// ID запроса.
	ReqID string `json:"req-id"`
}

// Ошибка. Ошибка валидации, например.
type Error struct {
	// example: validate.
	Name string `json:"name"`

	// example: Parameters requirements are not met.
	Message string `json:"message"`
}

// Лейбл звукозаписи.
type Label struct {
	// ID лейбла.
	ID int64 `json:"id"`

	// Имя лейбла.
	Name string `json:"name"`
}

type Account struct {
	// Текущая дата и время
	//
	// example: 2021-03-17T18:13:40+00:00.
	Now string `json:"now"`

	// Уникальный идентификатор.
	UID int64 `json:"uid"`

	// Виртуальное имя (обычно e-mail).
	Login string `json:"login"`

	// Полное имя (имя и фамилия).
	FullName string `json:"fullName"`

	// Фамилия.
	SecondName string `json:"secondName"`

	// Имя.
	FirstName string `json:"firstName"`

	// Отображаемое имя.
	DisplayName string `json:"displayName"`

	// Доступен ли сервис.
	ServiceAvailable bool `json:"serviceAvailable"`

	// Является ли пользователем чьим-то другим.
	HostedUser bool `json:"hostedUser"`

	// Мобильные номера.
	PassportPhones []struct {
		Phone string `json:"phone"`
	} `json:"passport-phones"`
}

type Album struct {
	// Идентификатор альбома.
	ID int64 `json:"id"`

	// Название альбома.
	Title string `json:"title"`

	// Мета тип (single, podcast, music, remix).
	MetaType string `json:"metaType"`

	// например: "single".
	Type *string `json:"type"`

	// например: "Remix".
	Version *string `json:"version"`

	// Год релиза.
	Year int `json:"year"`

	// Дата релиза в формате ISO 8601.
	ReleaseDate string `json:"releaseDate"`

	// Ссылка на обложку.
	CoverUri string `json:"coverUri"`

	// Ссылка на превью Open Graph.
	OgImage string `json:"ogImage"`

	// Жанр музыки.
	Genre string `json:"genre"`

	// Количество треков.
	TrackCount int `json:"trackCount"`

	// Количество лайков.
	LikesCount int `json:"likesCount"`

	// Является ли альбом новым.
	Recent bool `json:"recent"`

	// Популярен ли альбом у слушателей.
	VeryImportant bool `json:"veryImportant"`

	// Артисты.
	Artists []*Artist `json:"artists"`

	// Лейблы.
	//
	// Может быть как слайсом строк с названиями, так и слайсом структур Label.
	//
	// (?) Слайсы строк используются при поиске, а слайсы структур в остальных случаях.
	Labels []any `json:"labels"`

	// Доступен ли альбом.
	Available bool `json:"available"`

	// Доступен ли альбом для пользователей с подпиской.
	AvailableForPremiumUsers bool `json:"availableForPremiumUsers"`

	// Доступен ли альбом из приложения для телефона.
	AvailableForMobile bool `json:"availableForMobile"`

	// Доступен ли альбом частично для пользователей без подписки.
	AvailablePartially bool `json:"availablePartially"`

	// ID лучших треков альбома.
	Bests []int64 `json:"bests"`

	// Ремиксы, и прочее. Не nil, например когда запрашивается альбом с треками.
	Duplicates []*Album `json:"duplicates"`

	// Треки альбома, разделенные по дискам.
	Volumes [][]*Track `json:"volumes"`
}

// Артист.
type Artist struct {
	ID int64 `json:"id"`

	// Имя.
	Name     string `json:"name"`
	Various  bool   `json:"various"`
	Composer bool   `json:"composer"`

	// (?) Аватар.
	Cover *Cover `json:"cover"`

	// Жанры.
	Genres []string `json:"genres"`
}

type Cover struct {
	// Если не nil, значит остальные поля структуры будут nil.
	//
	// Example: "cover doesn't exist".
	Error *string `json:"error"`

	Custom *bool `json:"custom"`

	// Существует когда поле type = "pic".
	Dir *string `json:"dir"`

	// pic | mosaic.
	Type *string `json:"type"`

	// Существует когда поле type = "mosaic".
	ItemsUri []string `json:"itemsUri"`

	// Существует когда поле type = "pic".
	Uri *string `json:"uri"`

	Version *string `json:"version"`
}

// Владелец. Владелец плейлиста, например.
type Owner struct {
	// id.
	Uid int64 `json:"uid"`

	// Логин.
	Login string `json:"login"`

	// Имя.
	Name string `json:"name"`

	// Пол.
	Sex string `json:"sex"`

	// (?) Плейлист от редакции.
	Verified bool `json:"verified"`
}

// Трек.
type Track struct {
	// В зависимости от запроса.
	//
	// Например, при получении альбома с треками, ID будет string.
	//
	// int64 | string
	ID     int64  `json:"-"`
	RealId string `json:"realId"`

	// Название.
	Title string `json:"title"`

	// OWN.
	TrackSource string `json:"trackSource"`

	// Лейбл.
	Major                    *Label `json:"major"`
	Available                bool   `json:"available"`
	AvailableForPremiumUsers bool   `json:"availableForPremiumUsers"`

	// (?) Трек могут послушать даже те, кто без подписки, или не вошел в аккаунт?
	AvailableFullWithoutPermission bool `json:"availableFullWithoutPermission"`

	// Длительность в миллисекундах.
	DurationMs int64 `json:"durationMs"`

	StorageDir        string    `json:"storageDir"`
	FileSize          int64     `json:"fileSize"`
	R128              *R128     `json:"r128"`
	PreviewDurationMs int64     `json:"previewDurationMs"`
	Artists           []*Artist `json:"artists"`
	Albums            []*Album  `json:"albums"`
	CoverUri          string    `json:"coverUri"`
	OgImage           string    `json:"ogImage"`

	// Доступен ли текст трека.
	LyricsAvailable bool `json:"lyricsAvailable"`

	Type             string `json:"type"`
	RememberPosition bool   `json:"rememberPosition"`
}

// Разбираемся с ID трека.
func (t *Track) UnmarshalJSON(data []byte) error {
	// чтобы избежать stack overflow, делаем alias на Track.
	// по сути TrackFake, это тот же Track, только без методов
	type TrackFake Track
	var unmarshal = func(id int64) error {
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

	// если ID int: окей
	var idInt = &struct {
		ID int64 `json:"id"`
	}{}
	if err := json.Unmarshal(data, idInt); err == nil {
		return unmarshal(idInt.ID)
	}

	// если ID строка: конвертируем в int
	var idString = &struct {
		ID string `json:"id"`
	}{}
	if err := json.Unmarshal(data, idString); err != nil {
		return err
	}
	var err error
	var converted int64
	if converted, err = strconv.ParseInt(idString.ID, 10, 64); err != nil {
		return err
	}
	return unmarshal(converted)
}

// Нормализация.
//
// https://en.wikipedia.org/wiki/EBU_R_128
type R128 struct {
	I float64 `json:"i"`

	// True Peak.
	Tp float64 `json:"tp"`
}

type TrackItem struct {
	Id            int64  `json:"id"`
	Track         *Track `json:"track"`
	Timestamp     string `json:"timestamp"`
	OriginalIndex int    `json:"originalIndex"`
	Recent        bool   `json:"recent"`
}

type Playlist struct {
	// Владелец плейлиста.
	Owner *Owner `json:"owner"`

	// Уникальный ID.
	PlaylistUuid string `json:"playlistUuid"`

	// Тоже какой-то ID.
	Uid int64 `json:"uid"`

	// И это похоже на ID.
	Kind int64 `json:"kind"`

	// Описание.
	Description string `json:"description"`
	Available   bool   `json:"available"`

	// Название.
	Title      string `json:"title"`
	Collective bool   `json:"collective"`

	// Обложка.
	Cover *Cover `json:"cover"`

	Created    string `json:"created"`
	Modified   string `json:"modified"`
	DurationMs int64  `json:"durationMs"`
	OgImage    string `json:"ogImage"`

	// Количество треков.
	TrackCount int `json:"trackCount"`

	// Количество лайков.
	LikesCount int `json:"likesCount"`

	// Видимость.
	//
	// public | private
	Visibility string `json:"visibility"`

	// Треки.
	Tracks []*TrackItem `json:"tracks"`

	Revision int `json:"revision"`
}

type Pager struct {
	Total   int `json:"total"`
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
}

type Status struct {
	Account      *Account      `json:"account"`
	Subscription *Subscription `json:"subscription"`
}

// Информация о подписках пользователя
type Subscription struct {
	HadAnySubscription bool `json:"hadAnySubscription"`
}

// Поисковая подсказка.
type Suggestions struct {
	// Лучший результат.
	Best *Best[any] `json:"best"`

	// Предложения на основе запроса.
	//
	// Например, запрос: "emine"
	//
	// Suggestions будут примерно такие:
	//
	// ["eminem", "mount eminest", "eminen", "eminem - encore"], и так далее.
	Suggestions []string `json:"suggestions"`
}

// Лучший результат поиска
type Best[T Track | Artist | Album | Playlist | Video | any] struct {
	// Тип лучшего результата
	//
	// track | artist | album | playlist | video
	Type string `json:"type"`

	Text string `json:"text"`

	Result T `json:"result"`
}

// Рекомендации для плейлиста
type PlaylistRecommendations struct {
	// Уникальный идентификатор партии треков
	BatchId string `json:"batch_id"`

	Tracks []*Track `json:"tracks"`
}

// Дополнительная информация о треке.
type Supplement struct {
	// Уникальный идентификатор дополнительной информации.
	Id     string           `json:"id"`
	Lyrics *Lyrics          `json:"lyrics"`
	Videos *VideoSupplement `json:"videos"`

	// Доступно ли радио.
	RadioIsAvailable bool `json:"radioIsAvailable"`

	// Полное описание эпизода подкаста.
	Description string `json:"description"`
}

// Текст трека.
type Lyrics struct {
	// Уникальный идентификатор текста трека.
	Id int64 `json:"id"`

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

// Видеоклипы.
type VideoSupplement struct {
	// URL на обложку видео.
	Cover string `json:"cover"`

	// Сервис поставляющий видео.
	Provider string `json:"provider"`

	// Название видео.
	Title string `json:"title"`

	// Уникальный идентификатор видео на сервисе.
	ProviderVideoId string `json:"providerVideoId"`

	// URL на видео.
	Url string `json:"url"`

	// URL на видео, находящегося на серверах Яндекса.
	EmbedUrl string `json:"embedUrl"`

	// HTML тег для встраивания видео.
	Embed string `json:"embed"`
}

// Видео.
type Video struct {
	// Название видео.
	Title string `json:"title"`

	// Ссылка на изображение.
	Cover string `json:"cover"`

	// Ссылка на видео.
	EmbedUrl string `json:"embedUrl"`

	// Сервис поставляющий видео.
	Provider string `json:"provider"`

	// Уникальный идентификатор видео на сервисе.
	ProviderVideoId string `json:"providerVideoId"`

	// Ссылка на видео YouTube.
	YoutubeUrl string `json:"youtubeUrl"`

	// Ссылка на изображение.
	ThumbnailUrl string `json:"thumbnailUrl"`

	// Длительность видео в секундах.
	Duration int64 `json:"duration"`

	// Текст.
	Text string `json:"text"`

	// HTML тег для встраивания в разметку страницы.
	HtmlAutoPlayVideoPlayer string `json:"htmlAutoPlayVideoPlayer"`

	// example: ["RUSSIA_PREMIUM", "RUSSIA"].
	Regions []string `json:"regions"`
}

type PlaylistId struct {
	// Уникальный идентификатор пользователя владеющим плейлистом.
	Uid int64 `json:"uid"`

	// Уникальный идентификатор плейлиста.
	Kind int64 `json:"kind"`
}

// Список похожих треков на другой трек.
type SimilarTracks struct {
	Track *Track `json:"track"`
	// Похожие треки.
	SimilarTracks []*Track `json:"similarTracks"`
}

// Список треков.
type TracksLibrary struct {
	Library struct {
		// Уникальный идентификатор пользователя.
		Uid int64 `json:"uid"`

		Revision int64 `json:"revision"`

		// Список треков в укороченной версии.
		Tracks []*TrackShort `json:"tracks"`
	} `json:"library"`
}

// Укороченная версия трека с неполными данными.
type TrackShort struct {
	// Уникальный идентификатор трека.
	Id string `json:"id"`

	// Уникальный идентификатор альбома.
	AlbumId string `json:"albumId"`

	// Дата.
	Timestamp string `json:"timestamp"`
}

// Информация о вариантах загрузки трека.
type TrackDownloadInfo struct {
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
	BitrateInKbps int `json:"bitrateInKbps"`
}

// Результаты поиска.
type Search struct {
	// Был ли исправлен запрос.
	MisspellCorrected bool `json:"misspellCorrected"`

	// Было ли отключено исправление результата.
	Nocorrect bool `json:"nocorrect"`

	// Поисковой запрос (оригинальный или исправленный).
	Text string `json:"text"`

	// Исправленный поисковой запрос.
	MisspellResult string `json:"misspellResult"`

	// Оригинальный поисковой запрос.
	MisspellOriginal string `json:"misspellOriginal"`

	// ID запроса.
	SearchResultId *string `json:"searchResultId"`

	// artist | album | track | podcast | playlist
	//
	// Default: all.
	Type string `json:"type"`

	// Текущая страница. Доступно при использовании параметра type.
	Page *int64 `json:"page"`

	// Результатов на странице. Доступно при использовании параметра type.
	PerPage *int64 `json:"perPage"`

	// Лучший результат.
	Best Best[*any] `json:"best"`

	// Найденные альбомы.
	Albums SearchResult[*Album] `json:"albums"`

	// Найденные артисты.
	Artists SearchResult[*Artist] `json:"artists"`

	// Найденные плейлисты.
	Playlists SearchResult[*Playlist] `json:"playlists"`

	// Найденные треки.
	Tracks SearchResult[*Track] `json:"tracks"`

	// Найденные видео.
	Videos SearchResult[*Video] `json:"videos"`

	// Найденные подкасты.
	Podcasts SearchResult[*any] `json:"podcasts"`

	// Найденные эписозды подкастов.
	PodcastEpisodes SearchResult[*any] `json:"podcast_episodes"`
}

type SearchResult[T any] struct {
	// Тип результата
	Type string `json:"type"`

	// Количество результатов
	Total int64 `json:"total"`

	// Максимальное количество результатов на странице.
	PerPage int64 `json:"perPage"`

	// Позиция блока
	Order int64 `json:"order"`

	Results []T `json:"results"`
}
