package push

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type PushListResult struct {
	Result      string            `json:"result"`
	TaskId      string            `json:"taskid"`        //	任务标识号
	Desc        string            `json:"desc"`          //	错误信息描述
	CidDetail   map[string]string `json:"cid_details"`   //目标cid用户推送结果详情
	AliasDetail map[string]string `json:"alias_details"` //目标别名用户推送结果详情
}

type PushListParam struct {
	TaskId     string   `json:"taskid"`                //	任务号，取save_list_body返回的taskid
	Cid        []string `json:"cid,omitempty"`         //cid为cid list，与alias list二选一
	Alias      []string `json:"alias,omitempty"`       //	alias为alias list，与cid list二选一
	NeedDetail bool     `json:"need_detail,omitempty"` //默认值:false，是否需要返回每个CID的状态
}

func PushList(conf entity.IAppConfig, param *PushListParam) (*PushListResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/push_list"

	reply := new(PushListResult)
	err := util.Post(url, conf.GetToken(), param, reply)

	return reply, err
}
