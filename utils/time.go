package utils

import (
	"time"
)

func NowDate() string {
	now := time.Now();
	return now.Format("Jan 2, 2006 Mon");
	
}
