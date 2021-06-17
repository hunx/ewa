package payment

//营销
type Markting struct {
	payment *Payment
}

//红包
func (m *Markting) Redpack() *Redpack {
	r := &Redpack{}

	r.config.Set("mch_id", m.payment.config.MchID)
	r.config.Set("wxappid", m.payment.config.AppID) // 默认appid，小程序可在配置处修改

	return r
}