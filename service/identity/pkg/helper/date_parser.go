package helper

import "time"

func BirthdateParser(t time.Time) string {
	return t.Format("01/02/2006")
}
