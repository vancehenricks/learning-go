package greetings

import (
    "errors"
    "fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {

    if name == "" {
        return "", errors.New("empty name")
    }

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func Hellos(name []string) (map[string]string, error) {

    messages := make(map[string]string)

    for _, n := range name {
        message, err := Hello(n)
        if err != nil {
            return nil, err
        }
        
        messages[n] = message
    }

    return messages, nil
}

func randomFormat() string {
    format := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hello, %v. Well met!",
    }

    return format[rand.Intn(len(format))]
}