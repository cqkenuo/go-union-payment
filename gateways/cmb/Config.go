package cmb

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UseSandbox    bool   `json:"use_sandbox"`
	BranchNo      string `json:"branch_no"`
	MchId         string `json:"mch_id"`
	MerKey        string `json:"mer_key"`
	CmbPubKey     string `json:"cmb_pub_key"`
	OpPwd         string `json:"op_pwd"`
	SignType      string `json:"sign_type"`
	LimitPay      string `json:"limit_pay"`
	NotifyUrl     string `json:"notify_url"`
	ReturnUrl     string `json:"return_url"`
	SignNotifyUrl string `json:"sign_notify_url"`
	SignReturnUrl string `json:"sign_return_url"`
}

func (c Config) ParseConfig() interface{} {
	return c
}

func NewConfigWithJson(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewConfigWithYaml(content []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
