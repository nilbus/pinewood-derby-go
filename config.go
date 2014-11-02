package main

var derbyConfig map[string]interface{}

func loadDerbyConfig() {
	derbyConfig = map[string]interface{}{
		"laneCount": 3,
	}
}
