package pokeapi

type RespShallowPokemon struct {
	Abilities		[]struct{
		Ability struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		}
		IsHidden 	bool	`json:"is_hidden"`
		Slot		int		`json:"slot"`
	}			`json:"abilities"`
	Base_Experience		int		`json:"base_experience"`
	Cries 	struct{
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	}		`json:"previous"`
	Forms	[]struct{
		Name string	`json:"name"`
		URL  string	`json:"url"`
	} `json:"forms"`
	Game_Indices	[]struct{
		Game_Index	int	`json:"game_index"`
		Version		struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		}
	} `json:"game_indices"`
	Height		int		`json:"height"`
	Held_Items	[]struct{
		Item struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		} `json:"item"`
		Version_Details []struct{
			Rarity	int	`json:"rarity"`
			Version	struct{
				Name string	`json:"name"`
				URL  string	`json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID		int		`json:"id"`
	Is_Default	bool	`json:"is_default"`
	Location_Area_Encounters string `json:"location_area_encounters"`
	Moves		[]struct{
		Move struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		} `json:"move"`
		Version_Group_Details []struct{
			Level_Learned_At	int	`json:"level_learned_at"`
			Move_Learn_Method	struct{
				Name string	`json:"name"`
				URL  string	`json:"url"`
			} `json:"move_learn_method"`
			Order 	int	`json:"order"`
			Version_Group	struct{
				Name string	`json:"name"`
				URL  string	`json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name		string	`json:"name"`
	Order		int		`json:"order"`
	Weight		int		`json:"weight"`
	Stats		[]struct{
		Base_Stat	int	`json:"base_stat"`
		Effort		int	`json:"effort"`
		Stat		struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types		[]struct{
		Slot	int	`json:"slot"`
		Type	struct{
			Name string	`json:"name"`
			URL  string	`json:"url"`
		} `json:"type"`
	} `json:"types"`
}