package response

type Nurul struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	
}

type Komik struct {
	Title  string   `json:"title"`
	Thumb  string   `json:"thumb"`
	Images []string `json:"images"`
}