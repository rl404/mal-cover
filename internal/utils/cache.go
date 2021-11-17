package utils

import (
	"fmt"
	"strings"
)

// GetKey to generate cache key.
func GetKey(params ...interface{}) string {
	strParams := []string{"mc"}
	for _, p := range params {
		strParams = append(strParams, fmt.Sprintf("%v", p))
	}
	return strings.Join(strParams, ":")
}
