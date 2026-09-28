package posts

type Post struct {
	ID       uint32 `json:"id"`
	FKUserID uint32 `json:"fk_user_id"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
