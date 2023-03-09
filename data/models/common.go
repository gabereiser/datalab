package models

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID(id *string) ID {
	if id == nil {
		return uuid.Must(uuid.NewUUID())
	}
	return uuid.MustParse(*id)
}
