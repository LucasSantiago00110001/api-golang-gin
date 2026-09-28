package posts

// ListResponse is the paginated response returned by GET /posts.
type ListResponse struct {
	Total     int    `json:"total"`
	Page      int    `json:"page"`
	Registros []Post `json:"registros"`
}
