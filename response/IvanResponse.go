package response

type IvanError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type IvanSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PingIvanSuccess struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}