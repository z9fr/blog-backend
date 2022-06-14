package utils

import (
	"crypto/rand"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// generate a unique slug for a post
// @TODO
// do more validations and checking
func GenerateSlug(title string, addrand bool) string {
	title = strings.ToLower(title)

	title = FirstN(title, 80)
	title = strings.ToLower(title)
	title = strings.ReplaceAll(title, " ", "-")
	title = strings.ReplaceAll(title, ",", "-")
	title = strings.ReplaceAll(title, ".", "-")

	if addrand {
		randNumber, err := GenerateRandomNumber(10)
		if err != nil {
			logrus.Errorf("Error generating random number: %v", err)
		}

		title = title + strconv.Itoa(randNumber)
	}

	return title
}

//@utils
// Return first n chars of a string
// https://stackoverflow.com/a/41604514/17126147
func FirstN(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

func GenerateRandomNumber(numberOfDigits int) (int, error) {
	maxLimit := int64(int(math.Pow10(numberOfDigits)) - 1)
	lowLimit := int(math.Pow10(numberOfDigits - 1))

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(maxLimit))
	if err != nil {
		return 0, err
	}
	randomNumberInt := int(randomNumber.Int64())

	// Handling integers between 0, 10^(n-1) .. for n=4, handling cases between (0, 999)
	if randomNumberInt <= lowLimit {
		randomNumberInt += lowLimit
	}

	// Never likely to occur, kust for safe side.
	if randomNumberInt > int(maxLimit) {
		randomNumberInt = int(maxLimit)
	}
	return randomNumberInt, nil
}
