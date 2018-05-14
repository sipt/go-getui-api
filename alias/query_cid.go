package alias

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type QueryCidResult struct {
	Result string   `json:"result"`
	Desc   string   `json:"desc"`
	Cid    []string `json:"cid"`
}

func QueryCid(conf entity.IAppConfig, alias string) (*QueryCidResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/query_cid/" + alias

	reply := new(QueryCidResult)
	err := util.Get(url, conf.GetToken(), reply)

	return reply, err
}
