package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

const WapChargeMethod = "pay/unifiedorder"

type WapCharge struct {
	Base *Base
}

func (w WapCharge) Request() *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl(WapChargeMethod)
	params := w.BuildParams()
	resp, err := w.Base.Request(uri, params)
	if err != nil {
		return payloads.NewUnionPaymentResult(false, fmt.Sprintf("%s", err), nil)
	}


	fmt.Println(resp, err)

	// decode xml

	return payloads.NewUnionPaymentResult(true, "", nil)
}

func (w WapCharge) BuildParams() map[string]string {
	ret := map[string]string{
		"app_id":     w.Base.Config.AppId,
		"sub_appid":  w.Base.Config.SubAppId,
		"mch_id":     w.Base.Config.MchId,
		"sub_mch_id": w.Base.Config.SubMchId,
		"nonce_str":  w.Base.NonceStr,
		"sign_type":  w.Base.SignType,
	}

	return ret
}