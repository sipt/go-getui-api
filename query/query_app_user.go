package query

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type QueryAppUserResult struct {
	Result           string `json:"result"`           //操作结果 无效用户返回no_user
	Data             string `json:"data"`             //查询数据对象
	AppId            string `json:"appId"`            //	请求的AppId
	Date             string `json:"date"`             //查询的日期（格式：yyyy-MM-dd）
	NewRegistCount   int64  `json:"newRegistCount"`   //	新注册用户数
	RegistTotalCount int64  `json:"registTotalCount"` //累计注册用户数
	ActiveCount      int64  `json:"activeCount"`      //活跃用户数
	OnlineCount      int64  `json:"onlineCount"`      //在线用户数使用方法
}

func QueryAppUser(conf entity.IAppConfig, timeStr string) (*QueryAppUserResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/query_app_user/" + timeStr //20160404

	reply := new(QueryAppUserResult)
	err := util.Get(url, conf.GetToken(), reply)

	return reply, err
}
