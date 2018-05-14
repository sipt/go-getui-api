package alias

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type BindAliasRequest struct {
	AliasList []string `json:"alias_list"`
	Cid       string   `json:"cid"`   //需要绑定别名的cid列表
	Alias     string   `json:"alias"` //需要绑定的别名
}

type BindAliasResult struct {
	Result string `json:"result"`
	Desc   string `json:"desc"`
}

func BindAlias(conf entity.IAppConfig, req *BindAliasRequest) (*BindAliasResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/bind_alias"

	reply := new(BindAliasResult)
	err := util.Post(url, conf.GetToken(), req, reply)

	return reply, err
}
