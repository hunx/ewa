package openplatform

import (
	"context"

	"gitee.com/wallesoft/ewa/kernel/http"
	"gitee.com/wallesoft/ewa/miniprogram"
	"github.com/gogf/gf/v2/frame/g"
)

//代码上传 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/commit.html
//and ext_json @see https://developers.weixin.qq.com/miniprogram/dev/devtools/ext.html#%E5%B0%8F%E7%A8%8B%E5%BA%8F%E6%A8%A1%E6%9D%BF%E5%BC%80%E5%8F%91
func (mp *MiniProgram) Commit(ctx context.Context, templateId string, extJson string, version string, desc string) *http.ResponseData {
	client := mp.GetClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/commit", g.Map{
			"template_id":  templateId,
			"ext_json":     extJson,
			"user_version": version,
			"user_desc":    desc,
		}),
	}
}

//获取已上传的代码的页面列表 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/get_page.html
func (mp *MiniProgram) GetPage(ctx context.Context) *http.ResponseData {
	client := mp.GetClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "GET", "wxa/get_page"),
	}
}

//获取体验二维码 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_qrcode.html
func (mp *MiniProgram) GetQrcode(ctx context.Context, path ...string) *miniprogram.AppCode {
	var param g.Map
	if len(path) > 0 {
		param["path"] = path[0]
	}
	client := mp.GetClientWithToken()
	return &miniprogram.AppCode{
		Mp:  mp.MiniProgram,
		Raw: client.RequestRaw(ctx, "GET", "wxa/get_qrcode", param),
	}
}

//提交审核 config 参数具体查看 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/submit_audit.html
func (mp *MiniProgram) SubmitAudit(ctx context.Context, config ...g.Map) *http.ResponseData {
	client := mp.GetClientWithToken()
	var param g.Map
	if len(config) > 0 {
		param = config[0]
	}
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/submit_audit", param),
	}
}

//查询指定发布审核单的审核状态
func (mp *MiniProgram) GetAuditStatus(ctx context.Context, auditId string) *http.ResponseData {
	client := mp.GetClientWithToken()
	return &http.ResponseData{
		Json: client.RequestJson(ctx, "POST", "wxa/get_auditstatus", g.Map{"auditid": auditId}),
	}
}

//查询最新一次提交的审核状态
func (mp *MiniProgram) GetLatestAuditStatus(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/get_latest_auditstatus"),
	}
}

//小程序审核撤回
//注意： 单个帐号每天审核撤回次数最多不超过 5 次（每天的额度从0点开始生效），一个月不超过 10 次

func (mp *MiniProgram) UndoCodeAudit(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/undocodeaudit"),
	}
}

//发布已通过审核的小程序 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/release.html
//注：post的data为空，不等于不需要传data，否则会报错【errcode: 44002 "errmsg": "empty post data"】
func (mp *MiniProgram) Release(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/release", g.Map{}),
	}
}

//小程序版本回退 @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/Mini_Programs/code/revertcoderelease.html
func (mp *MiniProgram) RevertCodeRelease(ctx context.Context, version ...int) *http.ResponseData {
	var v int
	if len(version) > 0 {
		v = version[0]
	}
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/revertcoderelease", g.Map{"app_version": v}),
	}
}

//获取可回退的小程序版本
func (mp *MiniProgram) GetRevertReleaseHistory(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/revertcoderelease", g.Map{"action": "get_history_version"}),
	}
}

//分阶段发布
func (mp *MiniProgram) GrayRelease(ctx context.Context, grayPercentage int) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/grayrelease", g.Map{"gray_percentage": grayPercentage}),
	}
}

//查询当前分阶段发布详情
func (mp *MiniProgram) GetGrayRelease(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/getgrayreleaseplan"),
	}
}

//取消分阶段发布
func (mp *MiniProgram) RevertGrayRelease(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/revertgrayrelease"),
	}
}

//修改小程序服务状态  action: 'open'/'close';
func (mp *MiniProgram) ChangeVisitStatus(ctx context.Context, action string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/change_visitstatus", g.Map{"action": action}),
	}
}

//查询当前设置的最低基础库版本及各版本用户占比
func (mp *MiniProgram) GetSupportVersion(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "cgi-bin/wxopen/getweappsupportversion", g.Map{}),
	}
}

//设置最低基础库版本
func (mp *MiniProgram) SetSupoortVersion(ctx context.Context, version string) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "cgi-bin/wxopen/setweappsupportversion", g.Map{"version": version}),
	}
}

//查询服务商的当月提审限额（quota）和加急次数
func (mp *MiniProgram) QueryQuota(ctx context.Context) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "GET", "wxa/queryquota"),
	}
}

//加急审核
func (mp *MiniProgram) SpeedupAudit(ctx context.Context, auditId int) *http.ResponseData {
	return &http.ResponseData{
		Json: mp.GetClientWithToken().RequestJson(ctx, "POST", "wxa/speedupaudit", g.Map{"auditId": auditId}),
	}
}
