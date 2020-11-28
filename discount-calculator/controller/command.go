package controller

type ProductCommand struct {
	Id           string   `json:"id"`
	PriceInCents int      `json:"price_in_cents"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Discount     Discount `json:"discount"`
}

type Discount struct {
	Percentage   float32 `json:"percentage"`
	ValueInCents int     `json:"value_in_cents"`
}
