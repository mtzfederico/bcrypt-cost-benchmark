package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	costSeconds := flag.Int("target", 250, "The target cost in millisecond")
	passLength := flag.Int("passLen", 25, "The length of the random string hashed")
	flag.Parse()

	var target time.Duration = time.Duration(*costSeconds) * time.Millisecond
	cost := bcrypt.DefaultCost

	pass := generateRandomPassword(*passLength)

	fmt.Printf("Starting Cost: %v Target: %v\nUsing password: %s\n", cost, target, pass)

	for {
		start := time.Now()
		_, err := hash(pass, cost)
		elapsed := time.Since(start)

		log.Printf("cost: %v, elapsed: %v, target: %v", cost, elapsed, target)

		if err != nil {
			log.Fatal(err)
			return
		}

		if elapsed > target {
			log.Println("Target Reached")
			return
		}

		if elapsed < target {
			cost++
		}
	}

}

// Returns the hashed input as a string
func hash(input string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
