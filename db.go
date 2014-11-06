package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db gorm.DB

func ConnectDB() {
	dbFile := "db.sqlite3"
	var err error
	db, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("Failed to open the database " + dbFile)
	}
	db.SetLogger(log.New(os.Stdout, "\n", 0))
	db.LogMode(true)
	db.AutoMigrate(&Contestant{}, &Heat{}, &Run{})
}

type Contestant struct {
	Id          int32
	Name        string
	Retired     bool `sql:"not null;default false"`
	Runs        []Run
	AverageTime float32
}

type Heat struct {
	Id           int32
	Sequence     int32
	Status       string
	ContestantId int32
	Runs         []Run
}

type Run struct {
	Id           int32
	Time         float32
	Lane         int32
	ContestantId int32
	HeatId       int32
}

func Ranked(db *gorm.DB) *gorm.DB {
	return db.
		Select("contestants.*, avg(runs.time) AS average_time").
		Joins("LEFT JOIN runs ON contestants.id = runs.contestant_id").
		Joins("LEFT JOIN heats ON runs.heat_id = heats.id").
		Where("heats.sequence >= 0").
		Group("contestants.id").
		Order("average_time")
}
