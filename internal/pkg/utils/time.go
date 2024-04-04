package utils

import (
	"fmt"
	"time"
)

func StrAgo(timestamp time.Time) string {
	total := int64(time.Since(timestamp).Seconds())
	days := total / (60 * 60 * 24)
	hours := (total % (60 * 60 * 24)) / (60 * 60)
	minutes := (total % (60 * 60)) / 60
	seconds := total % 60

	if days == 1 {
		return fmt.Sprintf("%d day %02d h", days, hours)
	} else if days > 1 {
		return fmt.Sprintf("%d days", days)
	} else if days < 1 && hours > 1 {
		return fmt.Sprintf("%02d hours %02d minutes", hours, minutes)
	} else {
		return fmt.Sprintf("%02d minutes %02d seconds", minutes, seconds)
	}
}
