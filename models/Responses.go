package models

type MenuResponses struct {
	Menu []DetailResponses `json:"menu"`
}

type DetailResponses struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Endpoint string `json:"endpoint"`
}
