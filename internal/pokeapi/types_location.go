package pokeapi

type RespLocation struct {
	ID 					int 		`json:"id"`
	Name 				string 		`json:"name"`
	Pokemon_encounters	[]struct {
		Pokemon struct {
			Name string	`json:"name"`
			URL  string	`json:"url"`
		} `json:"pokemon"`
		Version_details []struct {
			Encounter_details []struct {
				Chance 		int		`json:"chance"`
			}		 `json:"encounter_details"`
			Version struct {
				Name string	`json:"name"`
				URL  string	`json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}