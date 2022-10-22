package response

type ItsukaHandoko struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Nama       string `json:"nama"`
	Genre	   string `json:"genre"`
}

type ItsukaSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}