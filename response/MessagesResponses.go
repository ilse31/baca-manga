package models

type Error struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PingSuccess struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}
