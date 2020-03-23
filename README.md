# go-union-payment

中文 | [English](/README_EN.md)

---


UnionPayment 是一个整合各种第三方支付网关的工具, 计划支持 Wechat（微信支付）、Alipay（支付宝）、Qpay（QQ钱包）、CMB（招商银行）等等。

之所以存在这个项目，是因为公司最近要做核心服务迁移，目前在精力PHP到GO的一个重构过程，鉴于之前使用的PHP工具包 [https://github.com/helei112g/payment] 还是比较方便实用，于是借鉴原项目开发者的思想，去重新设计一个属于golang领域的聚合支付工具，方便大家使用。

由于项目属于初期启动阶段，更新较快，测试稳定后会进行release发布

欢迎提PR申请一起参与开发

**支付网关支持列表**

- Wechat（微信支付） [开发中]
- Alipay [计划开发]
- Qpay [计划开发]
- China Merchant Bank [计划开发]
- Palpal [计划开发]

---

### 使用说明

#### 第一步

-- 使用任意一种方式加载 支付网关的配置

**a) json**

```golang
configJson, err: = ioutil.ReadFile("./config.json")
if err != nil {
    t.Error(err)
}
wechatConfig, err: = wechat.NewConfigWithJson(configJson)
if err != nil {
    t.Error(err)
}
```

**b) yaml **

```golang
configJson, err: = ioutil.ReadFile("./config.json")
if err != nil {
    t.Error(err)
}
wechatConfig, err: = wechat.NewConfigWithYaml(configJson)
if err != nil {
    t.Error(err)
}
```

**c) struct **

```golang
wechatConfig: = wechat.Config {
    UseSandbox: false,
    UseBackup: false,
    AppId: "",
    SubAppId: "",
    SubMchId: "",
    MchId: "",
    Md5Key: "",
    AppCertPem: "",
    AppKeyPem: "",
    SignType: "",
    LimitPay: nil,
    FeeType: "",
    ReturnRaw: false,
    NotifyUrl: "",
    RedirectUrl: "",
}
```


#### 第二步

-- 调用需要的网关API即可获取返回的结果，简单就完成了


```golang
// call the gateway channel api
payData: = map[string] string {
    "body": "goods information",
    "subject": "debug goods app payment",
    "trade_no": "xxx" + strconv.Itoa(rand.Intn(99999999)),
    "time_expire": strconv.Itoa(int(time.Now().Unix()) + 600),
    "amount": "0.01",
    "return_param": "anything_you_want",
    "goods_type": "1", // 0 虚拟，1 实物,
    "store_id": "",
}

result: = payment.Invoke(payloads.AliApiPayApp, payData)

if !result.State {
    t.Error(result.Msg)
}
m, e: = result.Data.ToMap()
fmt.Println(m, e)

j, e: = result.Data.ToJson()
fmt.Println(j, e)

x, e: = result.Data.ToXml()
fmt.Println(x, e)
```

### 第三步

-- 设置回调通知

#### 接管回调的服务，需要实现IPaymentNotify接口

```golang
type AliPaymentService struct {}
func (a *AliPaymentService) PayNotify(gateway payloads.UnionPaymentGateway, notifyData string) {
	if gateway == payloads.AlipayGateway {
		// @todo parse notifyData
		// and maybe change the order status
	} 
}
```


**Iris Framework - Route Context**

```golang
package notify

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type PaymentController struct {
}

func (c *PaymentController) BeforeActivation(b mvc.BeforeActivation) {
	b.HandleMany("POST", "/payment/callback", "Callback")
}

func (c *PaymentController) AfterActivation(a mvc.AfterActivation) {
}

func (c *PaymentController) Callback(ctx iris.Context) mvc.Result {
	req := ctx.Request()
    service := &AliPaymentService{}
    payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, aliConfig)
    payment.ParserNotify(req,service)
}
```

**Build-in net.http Server - http.Request **

```golang
http.HandleFunc("/your_notify_url", func(w http.ResponseWriter, r *http.Request) {        
    service := &AliPaymentService{}
    payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, aliConfig)
    payment.ParserNotify(r, service)
})
```


### 更多例子

单元测试都会同步在以下目录，供参考

> /examples/

#### 第三方依赖包

暂无