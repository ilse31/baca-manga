package response

type Pokedex struct {
	Id     	   int            	 `json:"id"`
	Name       string            `json:"name"`
	Total	   int    			 `json:"total"`
	HP		   int    			 `json:"hp"`
	Attack	   int    			 `json:"attack"`
	Defense	   int    			 `json:"defense"`
	Speed      int            	 `json:"speed"`
	Images    []string 			 `json:"images"`
}
