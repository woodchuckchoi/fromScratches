package module

import (
	"fmt"
)

type Crop struct {
	URI    string `json:"uri"`
	Depth  int    `json:"depth"`
	KeyNum int    `json:"keynum"`
}

type CropList struct {
	CropList []Crop `json:"crop_list"`
}

func (c *Crop) DoSomething() {
	fmt.Println(c.URI)
}
