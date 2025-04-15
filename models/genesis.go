package models

type Genesis struct {
	AppState AppState `json:"app_state"`
}

type AppState struct {
	Bank Bank `json:"bank"`
}

type Bank struct {
	Balances []Balance `json:"balances"`
}

type Balance struct {
	Address string `json:"address"`
	Coins []Coin `json:"coins"`
}

type Coin struct {
	Denom string `json:"denom"`
	Amount string `json:"amount"`
}