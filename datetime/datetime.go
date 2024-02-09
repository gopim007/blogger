package datetime

import (
	"strconv"
	"time"
)

func GetCurrentTimestampAsInt() int {
	currentTime := time.Now().UTC()
	intDate, _ := strconv.Atoi(currentTime.Format("20060102150405"))
	return intDate
}
