package types

type (
	Product struct {
		Id       int              `gorm:"column:Id;primary_key"`
		Name     string           `gorm:"column:Name"`
		Price    float32          `gorm:"Price"`
		Currency CurrencyTypeEnum `gorm:"Currency"`
	}
)
