package kissdate

import "time"

// Client ...
type Client struct{}

// NewClient ...
func NewClient() Client {
	return Client{}
}

// AddDate function had to be created because of an
// issue on the time.Time#AddDate() function:
//
// - https://github.com/golang/go/issues/31145
//
// where adding a month to Jan 30 would not get Feb 28
func (c Client) AddDate(date time.Time, y, m, d int) time.Time {
	expectedMonth := date.Month() + time.Month(m%12)

	date = date.AddDate(y, m, 0)

	if date.Month() == expectedMonth {
		return date.AddDate(0, 0, d)
	}

	return date.AddDate(0, 0, d-date.Day())
}
