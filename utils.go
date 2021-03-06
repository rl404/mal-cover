package main

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

func imageURLCleaner(str string) string {
	match, _ := regexp.MatchString("(questionmark)|(qm_50)|(na.gif)", str)

	if match {
		return ""
	}

	str = strings.Replace(str, "v.jpg", ".jpg", -1)
	str = strings.Replace(str, "t.jpg", ".jpg", -1)
	str = strings.Replace(str, "_thumb.jpg", ".jpg", -1)
	str = strings.Replace(str, "userimages/thumbs", "userimages", -1)

	r := regexp.MustCompile(`r\/\d{1,3}x\d{1,3}\/`)
	str = r.ReplaceAllString(str, "")
	r = regexp.MustCompile(`\?.+`)
	str = r.ReplaceAllString(str, "")

	return str
}

func getCacheKey(key ...string) string {
	return strings.Join(append([]string{"mc"}, key...), ":")
}

func setCache(key string, data interface{}) error {
	log.Println("setting cache", key)
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return cache.Set(key, dataByte)
}

func getCache(key string, data interface{}) error {
	dataByte, err := cache.Get(key)
	if err != nil {
		return err
	}
	return json.Unmarshal(dataByte, &data)
}
