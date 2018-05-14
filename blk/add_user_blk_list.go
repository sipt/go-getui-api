package blk

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type AddUserBlkListRequest struct {
	Cid []string `json:"cid"`
}

type AddUserBlkListResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func AddUserBlkList(conf entity.IAppConfig, req *AddUserBlkListRequest) (*AddUserBlkListResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/user_blk_list"

	reply := new(AddUserBlkListResult)
	err := util.Post(url, conf.GetToken(), req, reply)

	return reply, err
}
