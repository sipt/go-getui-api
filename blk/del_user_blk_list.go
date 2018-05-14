package blk

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type DelUserBlkListRequest struct {
	Cid []string `json:"cid"`
}

type DelUserBlkListResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func DelUserBlkList(conf entity.IAppConfig, req *DelUserBlkListRequest) (*DelUserBlkListResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/user_blk_list"

	reply := new(DelUserBlkListResult)
	err := util.Delete(url, conf.GetToken(), req, reply)

	return reply, err
}
