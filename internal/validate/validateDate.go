package validate

import "regexp"
import "time"

func IsDate(dateStr string) bool {
	ISO8601DateRegexString := `^(19\d\d|20[0-2]\d)-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$`
	ISO8601DateRegex := regexp.MustCompile(ISO8601DateRegexString)
	return ISO8601DateRegex.MatchString(dateStr)
}

func IsDateValid(dateStr string) bool {
  date, err := time.Parse("2006-01-02", dateStr)
  if err != nil{
    return false
  }
	if date.After(time.Now()) {
		return false
	}

  if (time.Now().Year() - date.Year()) < 12 {
    return false
  }

  return true
}
