package main

import "fmt"


func commandMapf(cfg *config) (error) {

	data, err := httpJSONGet(cfg.Next)
	if err != nil {
		return fmt.Errorf("Error fetching map data:", err)
	}
	if len(data) == 0 {
		return fmt.Errorf("No map data found")
	}

	cfg.Next, _ = data[0]["next"].(string)
	cfg.Previous, _ = data[0]["previous"].(string)

	fmt.Println(data[0]["next"])

	map_items, ok := data[0]["results"].([]interface{})
	if !ok {
		return fmt.Errorf("No results field in map data")
	}

	for _, item := range map_items {
		m, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Invalid map item format")
		}
		fmt.Println(m["name"])
	}

	return nil
}

func commandMapb(cfg *config) (error) {

	if cfg.Previous == "" {
		return fmt.Errorf("You're on the first page")
	}

	data, err := httpJSONGet(cfg.Previous)
	if err != nil {
		return fmt.Errorf("Error fetching map data:", err)
	}
	if len(data) == 0 {
		return fmt.Errorf("No map data found")
	}

	cfg.Next, _ = data[0]["next"].(string)
	cfg.Previous, _ = data[0]["previous"].(string)

	fmt.Println(data[0]["next"])

	map_items, ok := data[0]["results"].([]interface{})
	if !ok {
		return fmt.Errorf("No results field in map data")
	}

	for _, item := range map_items {
		m, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("Invalid map item format")
		}
		fmt.Println(m["name"])
	}

	return nil
}