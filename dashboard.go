package main

import (
	"encoding/json"
	"fmt"
)

func DashboardJson() string {
	json, err := json.Marshal(map[string]interface{}{
		"contestant_times": contestantTimes(),
		"most_recent_heat": mostRecentHeat(),
		"upcoming_heats":   upcomingHeats(),
		"notice":           notice(),
		"device_status":    device_status(),
	})
	if err != nil {
		return fmt.Sprintf("\"Error: %s\"", err)
	}
	return string(json)
}

func contestantTimes() string {
	return ""
}

func mostRecentHeat() string {
	return ""
}

func upcomingHeats() string {
	return ""
}

func notice() string {
	return ""
}

func device_status() string {
	return ""
}
