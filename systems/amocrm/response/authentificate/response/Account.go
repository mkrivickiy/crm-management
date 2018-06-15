package response

type Account struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Subdomain string `json:"subdomain"`
	Language  string `json:"language"`
	Timezone  string `json:"timezone"`
}
