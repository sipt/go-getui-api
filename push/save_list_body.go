package push

import (
	"github.com/sipt/go-getui-api/entity"
	"github.com/sipt/go-getui-api/util"
)

type SaveListBodyResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type SaveListBodyParam struct {
	Message      *entity.Message      `json:"message"` //消息内容
	Notification *entity.Notification `json:"notification,omitempty"`
	Link         *entity.Link         `json:"link,omitempty"`
	Notypopload  *entity.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *entity.Transmission `json:"transmission,omitempty"`
	PushInfo     string               `json:"push_info,omitempty"` //json串，当手机为ios，并且为离线的时候；或者简化推送的时候，使用该参数
	TaskName     string               `json:"task_name,omitempty"` //	任务名称 可以给多个任务指定相同的task_name，后面用task_name查询推送结果能得到多个任务的结果  可选
}

func SaveListBody(conf entity.IAppConfig, param *SaveListBodyParam) (*SaveListBodyResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/save_list_body"

	reply := new(SaveListBodyResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
