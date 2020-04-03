package helper

import "time"

var DefaultLocation = time.FixedZone("Asia/Jakarta", 7 * 60 * 60)

func GetNowTime() time.Time {
	return time.Now().In(DefaultLocation)
}