package push

import (
	"github.com/sipt/go-getui-api/entity"
	"github.com/sipt/go-getui-api/util"
)

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSingleParam struct {
	Message      *entity.Message      `json:"message"`
	Notification *entity.Notification `json:"notification,omitempty"`
	Link         *entity.Link         `json:"link,omitempty"`
	Notypopload  *entity.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *entity.Transmission `json:"transmission,omitempty"`
	Cid          string               `json:"cid,omitempty"`
	Alias        string               `json:"alias,omitempty"`
	RequestId    string               `json:"requestid"`
}

type PushSingleResult struct {
	Result string `json:"result"` //ok 鉴权成功
	TaskId string `json:"taskid"` //任务标识号
	Desc   string `json:"desc"`   //错误信息描述
	Status string `json:"status"` //推送结果successed_offline 离线下发successed_online 在线下发successed_ignore 非活跃用户不下发
}

func PushSingle(conf entity.IAppConfig, param *PushSingleParam) (*PushSingleResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/push_single"

	reply := new(PushSingleResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
