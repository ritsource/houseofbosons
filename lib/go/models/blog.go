package models

// Blog ...
type Blog struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Author        string `json:"author"`
	FormattedDate string `json:"formatted_date"`
	DocType       string `json:"doc_type"`
	MDSrc         string `json:"md_src"`
	HTMLSrc       string `json:"html_src"`
	CreatedAt     int32  `json:"created_at"`
	IsPublic      bool   `json:"is_public"`
	IsDeleted     bool   `json:"is_deleted"`
}
