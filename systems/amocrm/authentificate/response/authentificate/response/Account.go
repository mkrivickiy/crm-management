package response

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Subdomain string `json:"subdomain"`
	Language  string `json:"language"`
	Timezone  string `json:"timezone"`
}
