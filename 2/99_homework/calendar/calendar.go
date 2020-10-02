package main

import (
	"time"
)

type Calendar struct {
  time time.Time
}

type ICalendar interface {
  NewCalendar() ICalendar
	CurrentQuarter() int
}

func (calendar *Calendar) CurrentQuarter() int {
  month := calendar.time.Month();
  // Month's underlying type is int
  return (int(month) - 1) / 3 + 1
}


func NewCalendar(time time.Time) Calendar {
  var calendar Calendar
  calendar.time = time
  return calendar
}

func main() {

}