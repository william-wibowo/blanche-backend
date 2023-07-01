package util

import (
	"encoding/json"
	"regexp"
	"strings"
)

func ConvertJSONToPrettyText(data any) string {
	res, _ := json.MarshalIndent(data, "", "    ")
	return string(res)
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
