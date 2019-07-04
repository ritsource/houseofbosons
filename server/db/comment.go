package db

// Comment ...
type Comment struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Text      string `json:"text"`
	CreatedAt int32  `json:"created_at"`
}
