package utils

import (
	"time"
)

func DisplayNowDate() string {
	now := time.Now();
	return now.Format("Jan 2, 2006 Mon");
	
}

func DisplayBeforeDate(day int) string {

	date := time.Now().Add(-time.Hour * 24 * time.Duration(day))
	return date.Format("Jan 2, 2006 Mon");
}

func BeforeDate(day int) string {
	date := time.Now().Add(-time.Hour * 24 * time.Duration(day))
	return date.Format("20060102");
}
