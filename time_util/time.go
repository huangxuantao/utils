package time_util

import (
	"errors"
	"time"
)

const (
	LayoutNanoSecond = "2006-01-02 15:04:05.000000"
	LayoutSecond     = "2006-01-02 15:04:05"
	LayoutDay        = "2006-01-02"
	LayoutDay2       = "20060102"
	LayoutMonthDay   = "01-02"
	LocationShanghai = "Asia/Shanghai"
)

func GetNowTime() string {
	now := time.Now().UnixNano()
	sec := now / 1000000000
	nSec := now % 1000000000
	timeLayout := LayoutNanoSecond
	return time.Unix(sec, nSec).Format(timeLayout)
}

func GetTimeWithDateOffsetByFormatString(years int, months int, days int, format string) string {
	now := time.Now().AddDate(years, months, days).UnixNano()
	sec := now / 1000000000
	nSec := now % 1000000000
	return time.Unix(sec, nSec).Format(format)
}

func GetNowTimeByFormatString(format string) string {
	now := time.Now().UnixNano()
	sec := now / 1000000000
	nSec := now % 1000000000
	return time.Unix(sec, nSec).Format(format)
}

func GetTimeByFormatString(format string, timeString string) (time.Time, error) {
	loc, _ := time.LoadLocation(LocationShanghai)
	if timeLocation, err := time.ParseInLocation(format, timeString, loc); err != nil {
		return time.Time{}, err
	} else {
		return timeLocation, nil
	}
}

func StringToTime(timeString string, format string) (time.Time, error) {
	var rst time.Time
	loc, err := time.LoadLocation(LocationShanghai)
	if err != nil {
		return rst, err
	}

	rst, err = time.ParseInLocation(format, timeString, loc)
	if err != nil {
		return rst, err
	}

	return rst, nil
}

func StampToString(stamp int64, stampLen int, format string) (string, error) {
	if format == "" {
		format = LayoutNanoSecond
	}

	switch stampLen {
	case 10:
		return time.Unix(stamp, 0).Format(format), nil
	case 13:
		return time.Unix(stamp/1000, 0).Format(format), nil
	default:
		return "", errors.New("stamp len error")
	}
}
