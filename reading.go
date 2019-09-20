// Package reading implements functions related to measuring the time needed to
// read written text.
package reading

import (
	"time"
)

const averageReadingSpeed = 200 // [words per minute]
const averageWordLength = 5.1   // [characters]

// Time calculates the time needed by an average human to read the provided
// text.
func Time(text string) time.Duration {
	numberOfWords := float64(len(text)) / averageWordLength
	return time.Duration(numberOfWords / averageReadingSpeed * float64(time.Minute))
}
