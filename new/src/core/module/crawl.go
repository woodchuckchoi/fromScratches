package module

import "github.com/woodchuckchoi/fromScratches/src/core/model"

func Crawl(uri string, keynum int) []string {
	keywords := []struct {
		keyword   string
		occurence int
	}{}

	ret := []string{}
	for _, keyword := range keywords {
		ret = append(ret, keyword.keyword)
	}
	return ret
}

func Crawl2(c model.Crop, receiver chan<- model.Harvest, sig <-chan struct{}) {

}
