package helper

import "time"

func BirthdateToString(t time.Time) string {
	return t.Format("2006-01-02")
}

func StringToBirthdate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}
