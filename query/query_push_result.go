package query

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type PushRESResult struct {
	Result string  `json:"result"`
	Data   []*Data `json:"data"` //查询数据对象
}

type Data struct {
	TaskId     string      `json:"taskId"`     //任务标识号
	MsgTotal   int64       `json:"msgTotal"`   //有效可下发总数
	MsgProcess int64       `json:"msgProcess"` //消息回执总数
	ClickNum   int64       `json:"clickNum"`   //用户点击数
	PushNum    int64       `json:"pushNum"`    //im下发总量
	GT         interface{} `json:"GT"`         //个推推送结果数据sent 成功下发数feedback 回执数clicked 点击数displayed 展示数
	APN        interface{} `json:"APN"`        //iOS推送结果数据，详细字段参考GT
}

type PushRESParam struct {
	TaskIdList []string `json:"taskIdList"` //查询的任务尖列表
}

func PushResult(conf entity.IAppConfig, param *PushRESParam) (*PushRESResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/push_result"

	reply := new(PushRESResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
