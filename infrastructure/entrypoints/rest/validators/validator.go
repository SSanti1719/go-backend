package validators

import "regexp"

func EmailValidator(value string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	exp := regexp.MustCompile(regex)
	return exp.MatchString(value)
}

func MessageIdValidator(value string) bool {
	regex := `^<[0-9]+.*@[a-zA-Z0-9]+>$`
	exp := regexp.MustCompile(regex)
	return exp.MatchString(value)
}

func IsNumber(value string) bool {
	regex := `^[0-9]+$`
	exp := regexp.MustCompile(regex)
	return exp.MatchString(value)
}
