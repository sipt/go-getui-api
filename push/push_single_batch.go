package push

import (
	"github.com/sipt/go-getui-api/entity"
	"github.com/sipt/go-getui-api/util"
)

type PushSingleBatchResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

type PushSingleBatchListParam struct {
	MsgList    []*PushSingleBatchParam `json:"msg_list"`              //
	NeedDetail bool                    `json:"need_detail,omitempty"` //默认值:false，是否需要返回每个CID的状态
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSingleBatchParam struct {
	Message      *entity.Message        `json:"message"` //消息内容
	Notification *entity.Notification   `json:"notification,omitempty"`
	Link         *entity.Link           `json:"link,omitempty"`
	Notypopload  *entity.NotyPopload    `json:"notypopload,omitempty"`
	Transmission *entity.Transmission   `json:"transmission,omitempty"`
	Cid          string                 `json:"cid"`   //cid为cid list，与alias list二选一
	Alias        string                 `json:"alias"` //	alias为alias list，与cid list二选一
	RequestId    string                 `json:"requestid"`
	PushInfo     map[string]interface{} `json:"push_info,omitempty"` //APNs消息内容  可选
}

func PushSingleBatch(conf entity.IAppConfig, param *PushSingleBatchListParam) (*PushSingleBatchResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/push_single_batch"

	reply := new(PushSingleBatchResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
