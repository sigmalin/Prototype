package uuid

import (
	satori "github.com/satori/go.uuid"
)

type SUUID struct {
}

func Generate() (string, error) {
	return satori.NewV4().String(), nil
}
