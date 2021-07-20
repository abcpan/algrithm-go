package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name!")
	}
	template := randomFormat()
	msg := fmt.Sprintf(template,name)
	return msg, nil
}

// 产生随机问候模板
func randomFormat() string {
	formats := []string {
		"Hi, %v, Welcome",
		"Great to see you, %v!",
		"Hail, %v Well met",
	}

	return formats[rand.Intn(len(formats))]
}