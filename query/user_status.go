package query

import (
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/entity"
)

type UserStatusResult struct {
	Result    string `json:"result"`    //操作结果 无效用户返回no_user
	Cid       string `json:"cid"`       //查询状态的用户cid
	LastLogin string `json:"lastlogin"` //	sdk上次登陆时间戳
	Status    string `json:"status"`    //用户状态 online在线 offline离线
}

func UserStatus(conf entity.IAppConfig, cid string) (*UserStatusResult, error) {
	url := util.TOKEN_DOMAIN + conf.GetAppID() + "/user_status/" + cid

	reply := new(UserStatusResult)
	err := util.Get(url, conf.GetToken(), reply)
	
	return reply, err
}
