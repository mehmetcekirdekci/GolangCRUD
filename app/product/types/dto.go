package types

type (
	CreateProductDto struct {
		Name     string
		Price    float32
		Currency CurrencyTypeEnum
	}

	UpdateProductDto struct {
		Id       int
		Name     string
		Price    float32
		Currency CurrencyTypeEnum
	}
)
