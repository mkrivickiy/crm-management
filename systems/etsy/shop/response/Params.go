package response

type Params struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	UserID int `json:"user_id"`
}
