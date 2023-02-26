package types

import (
	"github.com/google/uuid"
)

type (
	CreateUserDto struct {
		CustomerId uuid.UUID
		Age        int
		Name       string
		LastName   string
	}

	UpdateUserDto struct {
		CustomerId string
		Age        int
		Name       string
		LastName   string
	}
)
