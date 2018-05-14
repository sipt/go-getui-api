package alias

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type QueryAliasResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
	Alias  string `json:"alias"`
}

func QueryAlias(conf entity.IAppConfig, cid string) (*QueryAliasResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/query_alias/" + cid

	reply := new(QueryAliasResult)
	err := util.Get(url, conf.GetToken(), reply)

	return reply, err
}
