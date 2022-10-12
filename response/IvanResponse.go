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

type Pokedex struct {
	Id     	   string          `json:"id"`
	Name       string       `json:"name"`
	Total	   string		`json:"total"`
	HP 	   string    	`json:"hp"`
	Attack	   string    	`json:"attack"`
	Defense	   string    	`json:"defense"`
	Speed      string          `json:"speed"`
	SpAttack      string          `json:"spattack"`
	SpDefense      string          `json:"spdefense"`
}