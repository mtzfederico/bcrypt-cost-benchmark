package main

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	var target time.Duration = 250 * time.Millisecond
	cost := 8

	for {
		start := time.Now()
		_, err := hash("correctHorseBatteryStaple", cost)
		elapsed := time.Since(start)

		log.Printf("cost: %v, elapsed: %v, target: %v", cost, elapsed, target)

		if err != nil {
			log.Fatal(err)
			return
		}

		if elapsed >= target {
			return
		}

		if elapsed < target {
			cost++
		}
	}

}

// Returnes the hashed input as a string
func hash(input string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
