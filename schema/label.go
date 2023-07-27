package schema

type (
	// Лейбл.
	Label struct {
		// ID лейбла.
		ID ID `json:"id"`

		// Имя лейбла.
		Name string `json:"name"`
	}
)
