package response

type Pagination struct {
	Limit      int `json:"effective_limit"`
	Offset     int `json:"effective_offset"`
	Page       int `json:"effective_page"`
	NextOffset int `json:"next_offset"`
	NextPage   int `json:"next_page"`
}
