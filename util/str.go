package util

import (
	"biligo/log"
	"github.com/gofrs/uuid"
)

func UUID() string {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Error("failed to generate UUID: %v", err)
	}
	return u4.String()
}
