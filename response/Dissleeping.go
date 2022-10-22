package response

type Dissleeping struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
	Nama       string `json:"nama"`
}

type DissleepingResponses struct {
	Title      string   `json:"title"`
	Thumbnail  string   `json:"thumbnail"`
	Alter      string   `json:"alter"`
	Status     string   `json:"status"`
	Author     string   `json:"author"`
	Ilustrator string   `json:"ilustrator"`
	Grafis     string   `json:"grafis"`
	Genre      []string `json:"genre"`
	Synopsis   string   `json:"synopsis"`
	Chapter    []string `json:"chapter"`
}

type KomikActionChapter struct {
	Title  string   `json:"title"`
	Thumb  string   `json:"thumb"`
	Images []string `json:"images"`
}