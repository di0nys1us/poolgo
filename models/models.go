package models

import (
	"time"
)

type Model struct {
	ID      int               `json:"id" db:"id"`
	Deleted bool              `json:"deleted" db:"deleted"`
	Version int               `json:"version" db:"version"`
	Links   map[string]string `json:"links"`
}

type Session struct {
	Model
	UserID            string        `json:"userId" db:"user_id"`
	Date              time.Time     `json:"date" db:"date"`
	PoolLength        int           `json:"poolLength" db:"pool_length"`
	Calories          int           `json:"calories" db:"calories"`
	TotalSets         int           `json:"totalSets" db:"total_sets"`
	TotalSwimmingTime time.Duration `json:"totalSwimmingTime" db:"total_swimming_time"`
	TotalRestTime     time.Duration `json:"totalRestTime" db:"total_rest_time"`
	TotalTime         time.Duration `json:"totalTime" db:"total_time"`
	TotalLaps         int           `json:"totalLaps" db:"total_laps"`
	TotalDistance     int           `json:"totalDistance" db:"total_distance"`
}

type SessionSet struct {
	Model
	SessionID       int           `json:"sessionId" db:"session_id"`
	SetNumber       int           `json:"number" db:"set_number"`
	SwimmingTime    time.Duration `json:"swimmingTime" db:"swimming_time"`
	RestTime        time.Duration `json:"restTime" db:"rest_time"`
	TotalTime       time.Duration `json:"totalTime" db:"total_time"`
	Laps            int           `json:"laps" db:"laps"`
	Distance        int           `json:"distance" db:"distance"`
	AverageStrokes  int           `json:"averageStrokes" db:"average_strokes"`
	Speed           int           `json:"speed" db:"speed"`
	EfficiencyIndex int           `json:"efficiencyIndex" db:"efficiency_index"`
}
