package main

import (
	"testing"
	"time"
)

func TestGetDateGroups(t *testing.T) {
	got := GetDateGroups("2018-02-01", 30)
	want := DatePair{"2018-02-01", "2018-03-03"}

	if got[0] != want {
		t.Errorf("GetDateGroups(2018-02-01, 30); got: %s want: {2018-02-01 2018-03-03}", got[0])
	}
}

func TestGetRecentDates(t *testing.T) {
	got := GetRecentDates(60)
	want := DatePair{time.Now().Format("2006-01-02"), time.Now().AddDate(0, 0, -60).Format("2006-01-02")}

	if got != want {
		t.Errorf("GetRecentDates(60); got: %s", got)
	}
}
