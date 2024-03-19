package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New(`nombre vacío`)
	} else {
		message := fmt.Sprintf(randomFormat(), name)
		return message, nil
	}

}

//Hellos returns a mpa that associates each of the named people with a greeting message
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {

	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno saludarte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))]
}
