package schema

const (
	// Своя картинка. Например вы создали плейлист, и на обложку установили свою картинку.
	CoverTypePic CoverType = "pic"

	// Например когда плейлист без обложки, то в качестве обложки используется коллаж из четырех обложек первых треков.
	CoverTypeMosaic CoverType = "mosaic"

	// Когда у артиста есть официальное фото в профиле ЯМ. Используется на странице профиля артиста.
	CoverTypeFromArtistPhotos CoverType = "from-artist-photos"

	// Когда у артиста нет официального фото в профиле ЯМ, то в качестве фото артиста используется обложка одного из альбомов.
	// Используется на странице профиля артиста.
	CoverTypeFromAlbumCover CoverType = "from-album-cover"
)

type (
	CoverType string

	Cover struct {
		// Если не nil, значит остальные поля структуры будут пустыми/nil.
		//
		// Пример: "cover doesn't exist".
		Error *string `json:"error"`

		// (?) Пользовательская обложка?
		Custom *bool `json:"custom"`

		// Тип обложки.
		Type *CoverType `json:"type"`

		// Пример: "7d7e16a0.p.ЕЩЁ_КАКОЙ_ТО_ID/"
		Prefix *string `json:"prefix"`

		// Существует когда поле type = "pic".
		Dir *string `json:"dir"`

		// Существует когда поле type == mosaic.
		//
		// Видимо здесь находятся ссылки на изображения используемые для создания мозайки. Например ссылки на обложки треков.
		ItemsUri []string `json:"itemsUri"`

		// Существует, когда поле Type не mosaic.
		//
		// пример: "avatars.yandex.net/get-music-content/КАКОЙ_ТО_ID/7d7e16a0.p.ЕЩЁ_КАКОЙ_ТО_ID/%%"
		URI *string `json:"uri"`

		// ???.
		Version *string `json:"version"`
	}
)
