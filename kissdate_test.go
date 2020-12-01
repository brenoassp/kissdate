package kissdate

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestClient_AddDate(t *testing.T) {
	withMock := func(runner func(t *testing.T, c Client)) func(t *testing.T) {
		return func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			runner(t, Client{})
		}
	}

	t.Run("should return date for the last day of February when substracting one month from March 30 or 31",
		withMock(func(t *testing.T, c Client) {
			date := time.Date(int(2020), time.Month(3), int(31), int(0), int(0), int(0), int(0), time.UTC)
			assert.Equal(t, "2020-02-29", c.AddDate(date, 0, -1, 0).Format("2006-01-02"))

			date = time.Date(int(2020), time.Month(3), int(30), int(0), int(0), int(0), int(0), time.UTC)
			assert.Equal(t, "2020-02-29", c.AddDate(date, 0, -1, 0).Format("2006-01-02"))
		}),
	)

	t.Run("should return date for the last day of April when adding one month to March 30 or 31",
		withMock(func(t *testing.T, c Client) {
			date := time.Date(int(2020), time.Month(3), int(31), int(0), int(0), int(0), int(0), time.UTC)
			assert.Equal(t, "2020-04-30", c.AddDate(date, 0, 1, 0).Format("2006-01-02"))

			date = time.Date(int(2020), time.Month(3), int(30), int(0), int(0), int(0), int(0), time.UTC)
			assert.Equal(t, "2020-04-30", c.AddDate(date, 0, 1, 0).Format("2006-01-02"))
		}),
	)

	t.Run("should return date for the same day and Month when subtracting twelve months to some date",
		withMock(func(t *testing.T, c Client) {
			date := time.Date(int(2020), time.Month(11), int(27), int(0), int(0), int(0), int(0), time.UTC)
			assert.Equal(t, "2019-11-27", c.AddDate(date, 0, -12, 0).Format("2006-01-02"))
		}),
	)
}
