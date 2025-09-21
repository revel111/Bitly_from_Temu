package link

type CreateLinkRequest struct {
	Url string `json:"url" binding:"required,url"`
}
