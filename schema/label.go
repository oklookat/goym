package schema

type (
	// Лейбл звукозаписи.
	Label struct {
		// ID лейбла.
		ID ID `json:"id"`

		// Имя лейбла.
		Name string `json:"name"`
	}
)
