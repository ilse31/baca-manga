package response

type KhoeruError struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
}

type KhoeruSuccess struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}