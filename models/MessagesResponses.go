package models

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Success struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PingSuccess struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
