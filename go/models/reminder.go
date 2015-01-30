package models

import (
	"strconv"
	"time"
)

type Reminder struct {
	Id        int
	Reminder  string
	Timestamp int32
	Error     string
}

func (t *Reminder) GetId() string {
	return strconv.Itoa(t.Id)
}

func (t *Reminder) RootURL() string {
	return "http://localhost:8080/reminders"
}

func (t *Reminder) GetDate() string {
	year, month, day := time.Unix(int64(t.Timestamp), 0).Date()
	readableDate := strconv.Itoa(day) + " " + month.String() + ", " + strconv.Itoa(year)
	return readableDate
}
