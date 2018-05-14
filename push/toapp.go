package push

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type ToAppResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type ToAppParam struct {
	Message      *entity.Message      `json:"message"` //消息内容
	Notification *entity.Notification `json:"notification,omitempty"`
	Link         *entity.Link         `json:"link,omitempty"`
	Notypopload  *entity.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *entity.Transmission `json:"transmission,omitempty"`
	Condition    []*entity.Condition  `json:"condition,omitempty"` //	筛选目标用户条件，参考下面的condition说明  可选
	Requestid    string               `json:"requestid"`           //请求唯一标识
}

func ToApp(conf entity.IAppConfig, param *ToAppParam) (*ToAppResult, error) {

	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/push_app"

	reply := new(ToAppResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
