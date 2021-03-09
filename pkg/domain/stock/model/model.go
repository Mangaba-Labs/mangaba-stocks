package model

// Share model
type Share struct {
	ID          int
	CompanyName string
	Symbol      string
	BuyValue    float32
	NowValue    float32
}

// ResponseHG getHG response
type ResponseHG struct {
	Results map[string]ResultFields
}

// ResultFields in ResponseHG
type ResultFields struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Price  float32 `json:"price"`
}

// ParseResponse from HG to our Share
func (s *Share) ParseResponse(hg *ResponseHG, symbol string) {
	s.Symbol = hg.Results[symbol].Symbol
	s.CompanyName = hg.Results[symbol].Name
	s.NowValue = hg.Results[symbol].Price
}
