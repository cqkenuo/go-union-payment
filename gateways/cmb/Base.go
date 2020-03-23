package cmb

import (
	"fmt"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/payloads"
)

func Factory(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) contracts.IGateway {
	cfg := config.ParseConfig().(Config)
	b := NewBase(cfg)
	fmt.Println(b)

	switch channel {
	case payloads.CmbChannelApp:
	case payloads.CmbChannelWap:
	case payloads.CmbChannelWeb:
	case payloads.CmbChannelQr:
	case payloads.CmbChannelLite:
	}
	return nil
}

type Base struct {
	Config Config
}

func NewBase(config Config) *Base {
	b := &Base{}
	b.Config = config

	return b
}