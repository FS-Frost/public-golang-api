package main

import (
	"log"
	"os"
	"strconv"
)

func GetEnvOrDefaultString(key string, def string) string {
	val := os.Getenv(key)

	if val != "" {
		return val
	}

	return def
}

func GetEnvOrDefaultInt(key string, def int) int {
	val := os.Getenv(key)

	if val != "" {
		intVal, err := strconv.Atoi(val)

		if err != nil {
			log.Fatalf("error parsing environment variable <%s> with valule <%s>: %v", key, val, err)
		}

		return int(intVal)
	}

	return def
}

func GetEnvOrDefaultFloat(key string, def float64) float64 {
	val := os.Getenv(key)

	if val != "" {
		floatVal, err := strconv.ParseFloat(val, 64)

		if err != nil {
			log.Fatalf("error parsing environment variable <%s> with valule <%s>: %v", key, val, err)
		}

		return floatVal
	}

	return def
}

func GetEnvOrDefaultBool(key string, def bool) bool {
	val := os.Getenv(key)

	if val != "" {
		boolVal, err := strconv.ParseBool(val)

		if err != nil {
			log.Fatalf("error parsing environment variable <%s> with valule <%s>: %v", key, val, err)
		}

		return boolVal
	}

	return def
}
