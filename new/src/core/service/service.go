package service

import (
	"time"

	"github.com/woodchuckchoi/fromScratches/src/constant"
	"github.com/woodchuckchoi/fromScratches/src/core/model"
)

// use functions to defind Crop's methods here or move them to where the model definition is

func RequestAnalytics(list []model.Crop) []model.Harvest {
	resp := []model.Harvest{}

	var resReceiver = make(chan model.Harvest, len(list))
	var sig chan struct{}

	for _, crop := range list {
		go func(c model.Crop, receiver chan<- model.Harvest, sig <-chan struct{}) {
			module.
			// crawl websites, should receive a signal when the parent wants it to abort the current working process
		}(crop, resReceiver, sig)
	}

	select {
	case h := <-resReceiver:
		resp = append(resp, h)
		if len(resp) == len(list) {
			break
		}
	case <-time.After(constant.TIMEOUT):
		close(sig)
		break
	}
}
