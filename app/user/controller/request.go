package controller

type (
	CreateUserRequest struct {
		Age      int    `json:"age" validate:"required"`
		Name     string `json:"name" validate:"required"`
		LastName string `json:"lastName" validate:"required"`
	}

	UpdateUserRequest struct {
		CustomerId string `json:"customerId" validate:"required"`
		Age        int    `json:"age" validate:"required"`
		Name       string `json:"name" validate:"required"`
		LastName   string `json:"lastName" validate:"required"`
	}
)
