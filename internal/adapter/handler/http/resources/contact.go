package resources

type CreateContactRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}
