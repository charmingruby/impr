package helper

import "time"

func BirthdateToString(t time.Time) string {
	return t.Format("01/02/2006")
}

func StringToBirthdate(dateStr string) (time.Time, error) {
	return time.Parse("01/02/2006", dateStr)
}
