package schema

type (
	// Лейбл звукозаписи.
	Label struct {
		// ID лейбла.
		ID UniqueID `json:"id"`

		// Имя лейбла.
		Name string `json:"name"`
	}
)
