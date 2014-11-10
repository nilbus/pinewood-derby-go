package main

import (
	"encoding/json"
	"fmt"
)

func DashboardJson() string {
	json, err := json.Marshal(map[string]interface{}{
		"contestant_times": contestantTimes(),
		"most_recent_heat": mostRecentHeatRuns(),
		"upcoming_heats":   upcomingHeats(),
		"notice":           notice(),
		"device_status":    device_status(),
	})
	if err != nil {
		return fmt.Sprintf("\"Error: %s\"", err)
	}
	return string(json)
}

func contestantTimes() []float32 {
	var contestants []Contestant
	times := make([]float32, 0)
	db.Scopes(Ranked).Find(&contestants)
	for _, contestant := range contestants {
		times = append(times, contestant.AverageTime)
	}
	return times
}

func mostRecentHeatRuns() []map[string]interface{} {
	runData := make([]map[string]interface{}, 0)
	var heat Heat
	var runs []Run
	db.Scopes(MostRecentHeat).Find(&heat)
	db.Model(&heat).Related(&runs)
	for _, run := range runs {
		var contestant Contestant
		db.Model(&run).Related(&contestant)
		runDatum := map[string]interface{}{
			"name": contestant.Name,
			"time": run.Time,
			"lane": run.Lane,
		}
		runData = append(runData, runDatum)
	}
	return runData
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
