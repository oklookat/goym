package schema

type Status struct {
	Account      *Account      `json:"account"`
	Subscription *Subscription `json:"subscription"`
}

// Настройки пользователя.
type AccountSettings struct {
	// ID.
	UID int64 `json:"uid" url:"-"`

	// Включен ли скробблинг last.fm?
	LastFmScrobblingEnabled   bool `json:"lastFmScrobblingEnabled" url:"lastFmScrobblingEnabled,omitempty"`
	FacebookScrobblingEnabled bool `json:"facebookScrobblingEnabled" url:"lastFmScrobblingEnabled,omitempty"`

	// (?) Включено ли рандомное воспроизведение треков?
	ShuffleEnabled bool `json:"shuffleEnabled" url:"shuffleEnabled,omitempty"`

	// Добавлять новый трек в начало плейлиста?
	AddNewTrackOnPlaylistTop bool `json:"addNewTrackOnPlaylistTop" url:"addNewTrackOnPlaylistTop,omitempty"`

	// Громкость в процентах (example: 75).
	VolumePercents uint8 `json:"volumePercents" url:"volumePercents,omitempty"`

	// Видимость музыкальной библиотеки.
	UserMusicVisibility Visibility `json:"userMusicVisibility" url:"userMusicVisibility,omitempty"`

	// ???
	UserSocialVisibility Visibility `json:"userSocialVisibility" url:"userSocialVisibility,omitempty"`

	AdsDisabled bool `json:"adsDisabled" url:"adsDisabled,omitempty"`

	// example: 2019-04-14T14:55:50+00:00
	Modified string `json:"modified" url:"-"`

	RbtDisabled bool `json:"rbtDisabled" url:"-"`

	// Тема оформления.
	Theme Theme `json:"theme" url:"theme,omitempty"`

	AutoPlayRadio    bool `json:"autoPlayRadio" url:"autoPlayRadio,omitempty"`
	SyncQueueEnabled bool `json:"syncQueueEnabled" url:"syncQueueEnabled,omitempty"`
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

type PromocodeStatus struct {
	// Статус активации промо-кода.
	//
	// Например: "code-not-exists".
	Status string `json:"status"`

	// Описание статуса.
	//
	// Например: "Gift code does not exist".
	StatusDesc string `json:"statusDesc"`

	AccountStatus *Status `json:"accountStatus"`
}

// Информация о подписках пользователя
type Subscription struct {
	HadAnySubscription bool `json:"hadAnySubscription"`
}

// Владелец. Владелец плейлиста, например.
type Owner struct {
	// id.
	UID int64 `json:"uid"`

	// Логин.
	Login string `json:"login"`

	// Имя.
	Name string `json:"name"`

	// Пол.
	Sex string `json:"sex"`

	// (?) Плейлист от редакции.
	Verified bool `json:"verified"`
}

// POST /account/consume-promo-code
type AccountConsumePromocodeRequestBody struct {
	// Промокод.
	Code string `url:"code"`

	// Язык *чего-то*.
	Language string `url:"language"`
}

// POST /account/settings
//
// Используйте метод Change().
type ChangeAccountSettingsRequestBody struct {
	AccountSettings
}

// Изменить настройки.
//
// Настройку нельзя изменить, если в поле структуры есть url:"-".
func (c *ChangeAccountSettingsRequestBody) Change(a AccountSettings) {
	c.AccountSettings = a
}
