package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

// 自定义时间格式
type Time time.Time

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t *Time) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	formatted := fmt.Sprintf("\"%s\"", tTime.Format(timeFormat))
	return []byte(formatted), nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// Value insert timestamp into mysql need this function.
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time

	//判断给定时间是否和默认零时间的时间戳相同
	tTime := time.Time(t)
	if tTime.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tTime, nil
}

// Scan valueof time.Time
func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("custom time can not convert %v to timestamp", v)
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}
