package query

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type QueryAppPushResult struct {
	Result          string `json:"result"`          //操作结果 无效用户返回no_user
	Data            string `json:"data"`            //查询数据对象
	AppId           string `json:"appId"`           //	请求的AppId
	Date            string `json:"date"`            //查询的日期（格式：yyyy-MM-dd）
	SendCount       int64  `json:"sendCount"`       //发送总数
	SendOnlineCount int64  `json:"sendOnlineCount"` //在线发送数
	ReceiveCount    int64  `json:"receiveCount"`    //接收数
	ShowCount       int64  `json:"showCount"`       //展示数
	ClickCount      int64  `json:"clickCount"`      //点击数
}

func QueryAppPush(conf entity.IAppConfig, timeStr string) (*QueryAppPushResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/query_app_push/" + timeStr //20160404

	reply := new(QueryAppPushResult)
	err := util.Get(url, conf.GetToken(), reply)

	return reply, err
}
