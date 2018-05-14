package push

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type StopTaskResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
}

func StopTask(conf entity.IAppConfig, taskid string) (*StopTaskResult, error) {

	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/stop_task/" + taskid

	reply := new(StopTaskResult)
	err := util.Delete(url, conf.GetToken(), nil, reply)

	return reply, err
}
