package controller

type Discount struct {
	Percentage   float32 `json:"percentage"`
	ValueInCents int32   `json:"value_in_cents"`
}
