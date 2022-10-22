package response

type Nurul struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}