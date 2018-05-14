package alias

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type UnbindAliasRequest struct {
	Cid   string `json:"cid"`
	Alias string `json:"alias"`
}
type UnbindAliasResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func UnbindAlias(conf entity.IAppConfig, req *UnbindAliasRequest) (*UnbindAliasResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/unbind_alias"

	reply := new(UnbindAliasResult)
	err := util.Post(url, conf.GetToken(), req, reply)

	return reply, err
}
