package model

type Crop struct {
	URI    string `json:"uri"`
	Depth  int    `json:"depth"`
	KeyNum int    `json:"keynum"`
}

type Harvest struct {
	URI      string   `json:"uri"`
	Keywords []string `json:"keywords"`
}
