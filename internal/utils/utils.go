package utils

import (
	"regexp"
	"strings"
)

// ImageURLCleaner to clean mal image cover url.
func ImageURLCleaner(str string) string {
	match, _ := regexp.MatchString("(questionmark)|(qm_50)|(na.gif)", str)

	if match {
		return ""
	}

	str = strings.ReplaceAll(str, "v.jpg", ".jpg")
	str = strings.ReplaceAll(str, "t.jpg", ".jpg")
	str = strings.ReplaceAll(str, "_thumb.jpg", ".jpg")
	str = strings.ReplaceAll(str, "userimages/thumbs", "userimages")

	r := regexp.MustCompile(`r\/\d{1,3}x\d{1,3}\/`)
	str = r.ReplaceAllString(str, "")
	r = regexp.MustCompile(`\?.+`)
	str = r.ReplaceAllString(str, "")

	return str
}
