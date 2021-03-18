package repository

type User struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	// Birthday time.Time `json:"birthday"`
}
