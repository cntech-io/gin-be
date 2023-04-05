package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetBooleanEnv(key string, panicFlag bool) bool {
	result, ok := os.LookupEnv(key)
	if panicFlag && !ok {
		panic(fmt.Sprintf("%v not found.", key))
	}
	if !ok {
		log := fmt.Sprintf("%v not found.", key)
		fmt.Println(log)
		return false
	}
	value, err := strconv.ParseBool(result)
	if err != nil {
		panic(fmt.Sprintf("%v has invalid data type => %v", key, err))
	}
	return value
}

func GetNumberEnv(key string, panicFlag bool) int {
	result, ok := os.LookupEnv(key)
	if panicFlag && !ok {
		panic(fmt.Sprintf("%v not found.", key))
	}
	if !ok {
		log := fmt.Sprintf("%v not found.", key)
		fmt.Println(log)
		return 0
	}
	value, err := strconv.ParseInt(result, 10, 32)
	if err != nil {
		panic(fmt.Sprintf("%v has invalid data type => %v", key, err))
	}
	return int(value)
}

func GetStringEnv(key string, panicFlag bool) string {
	result, ok := os.LookupEnv(key)
	if panicFlag && !ok {
		panic(fmt.Sprintf("%v not found.", key))
	}
	if !ok {
		log := fmt.Sprintf("%v not found.", key)
		fmt.Println(log)
		return ""
	}
	return result
}

func GetStringArrayEnv(key string, seperator string, panicFlag bool) []string {
	result, ok := os.LookupEnv(key)
	if panicFlag && !ok {
		panic(fmt.Sprintf("%v not found.", key))
	}
	if !ok {
		log := fmt.Sprintf("%v not found.", key)
		fmt.Println(log)
		return nil
	}

	return strings.Split(result, seperator)
}
