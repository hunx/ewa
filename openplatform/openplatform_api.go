package openplatform

import (
	"net/url"

	"gitee.com/wallesoft/ewa/kernel/http"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gutil"
)

//StartPushTicket 启动ticket推送服务
func (op *OpenPlatform) StartPushTicket() *http.ResponseData {
	return &http.ResponseData{
		Json: op.getClient().RequestJson("POST", "cgi-bin/component/api_start_push_ticket", g.Map{"component_appid": op.config.AppID, "component_secret": op.config.AppSecret}),
	}
}

//GetPreAuthorizationUrl 获取授权页网址
func (op *OpenPlatform) GetPreAuthorizationUrl(callback string, optional ...map[string]interface{}) (string, error) {

	val := url.Values{}
	authCode, err := op.GetPreAuthCode()
	if err != nil {
		return "", err
	}
	val.Add("pre_auth_code", authCode)
	if len(optional) > 0 {
		options := optional[0]
		if v, ok := options["auth_type"]; ok {
			val.Add("auth_type", gvar.New(v).String())
		}
		if v, ok := options["biz_appid"]; ok {
			val.Add("biz_appid", gvar.New(v).String())
		}
	}

	val.Add("component_appid", op.config.AppID)
	val.Add("redirect_uri", callback)

	return "https://mp.weixin.qq.com/cgi-bin/componentloginpage?" + val.Encode(), nil
}

//GetMobilePreAuthorizationUrl
func (op *OpenPlatform) GetMobilePreAuthorizationUrl(callback string, optional ...map[string]interface{}) (string, error) {
	val := url.Values{}
	authCode, err := op.GetPreAuthCode()
	if err != nil {
		return "", err
	}
	val.Add("pre_auth_code", authCode)
	if len(optional) > 0 {
		options := optional[0]
		if v, ok := options["auth_type"]; ok {
			val.Add("auth_type", gvar.New(v).String())
		}
		if v, ok := options["biz_appid"]; ok {
			val.Add("biz_appid", gvar.New(v).String())
		}
	}

	val.Add("component_appid", op.config.AppID)
	val.Add("redirect_uri", callback)
	val.Add("action", "bindcomponent")
	val.Add("no_scan", "1")

	return "https://mp.weixin.qq.com/safe/bindcomponent?" + val.Encode() + "#wechat_redirect", nil
}

//HandleAuthorize
func (op *OpenPlatform) HandleAuthorize(code string) *http.ResponseData {

	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/api_query_auth", map[string]string{
			"component_appid":    op.config.AppID,
			"authorization_code": code,
		}),
	}

}

//GetAuthorizer get authorizer info type as gjson.Json
func (op *OpenPlatform) GetAuthorizer(appid string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/api_get_authorizer_info", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
		}),
	}

}

//GetAuthorizers get authorizer list
func (op *OpenPlatform) GetAuthorizers(offset int, count int) *http.ResponseData {
	if count > 500 {
		count = 500
	}
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/api_get_authorizer_list", map[string]interface{}{
			"component_appid": op.config.AppID,
			"offset":          offset,
			"count":           count,
		}),
	}
}

//GetAuthorizerOption get authorizer option info
func (op *OpenPlatform) GetAuthorizerOption(appid string, name string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/api_get_authorizer_option", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
			"option_name":      name,
		}),
	}
}

//SetAuthorizerOption set authorizer option
func (op *OpenPlatform) SetAuthorizerOption(appid string, name string, value string) *http.ResponseData {
	client := op.getClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson("POST", "cgi-bin/component/api_set_authorizer_option", map[string]string{
			"component_appid":  op.config.AppID,
			"authorizer_appid": appid,
			"option_name":      name,
			"option_value":     value,
		}),
	}
}

//GetVerifyTicket
func (op *OpenPlatform) GetVerifyTicket() string {
	return op.verifyTicket.GetTicket()
}

//GetAccessToken
func (op *OpenPlatform) GetAccessToken() string {
	return op.accessToken.GetToken()
}

func (op *OpenPlatform) GetPreAuthCode() (string, error) {
	var code string
	var err error
	gutil.TryCatch(func() {
		client := op.getClientWithToken()
		v := client.RequestJson("POST", "cgi-bin/component/api_create_preauthcode", map[string]string{
			"component_appid": op.config.AppID,
		})

		if have := v.Contains("errcode"); have {
			panic(v.MustToJsonString())
		}
		if have := v.Contains("pre_auth_code"); have {
			code = v.GetString("pre_auth_code")
		} else {
			panic("Request pre_auth_code fail:" + v.MustToJsonString())
		}
	}, func(e error) {
		err = e
		op.Logger.File(op.Logger.ErrorLogPattern).Stdout(op.Logger.LogStdout).Print(err.Error())
	})

	return code, err
}
