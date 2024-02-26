package services

import (
	"encoding/json"
	"time"
)

type JSONTime struct {
	CURRENT_TIME string `json:"time"`
}

func GetCurrentTime() string {
	currentTime := JSONTime{time.Now().Format(time.RFC3339)}
	if json_data, err := json.Marshal(currentTime); err == nil {
		return string(json_data)
	}
	return "Oops...Something went wrong"
}
