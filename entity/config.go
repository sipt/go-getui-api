package entity

import (
	"time"
	"github.com/sipt/go-getui-api/util"
	"github.com/sipt/go-getui-api/token"
	"errors"
)

const (
	tokenExpire     = 24 * time.Hour
	tokenTTL        = 23 * time.Hour
	tokenRetry      = 30 * time.Second
	tokenCloseAfter = 1 * time.Minute
)

type IAppConfig interface {
	GetAppID() string
	GetToken() string
}

type ITokenGetter interface {
	//获取token
	Get() string
	//到时自动刷新token
	AutoRefresh(timeSpacing time.Duration) error
}

func NewAppConfig(AppID string,
	AppSecret string,
	AppKey string,
	MasterSecret string) IAppConfig {
	return &AppConfig{
		AppID:        AppID,
		AppKey:       AppKey,
		AppSecret:    AppSecret,
		MasterSecret: MasterSecret,
	}
}

type AppConfig struct {
	AppID        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	AppKey       string `json:"app_key"`
	MasterSecret string `json:"master_secret"`
	TokenGetter  ITokenGetter
}

func (a *AppConfig) GetAppID() string {
	return a.AppID
}

func (a *AppConfig) GetToken() string {
	if a.TokenGetter == nil {
		a.TokenGetter = NewTokenGetter(a)
	}
	return a.TokenGetter.Get()
}

func NewTokenGetter(conf *AppConfig) ITokenGetter {
	getter := &tokenGetter{
		appId:        conf.AppID,
		appKey:       conf.AppKey,
		masterSecret: conf.MasterSecret,
	}
	getter.AutoRefresh(tokenTTL)
	return getter
}

type tokenGetter struct {
	token string

	appId        string
	appKey       string
	masterSecret string

	timer *time.Timer
}

func (t *tokenGetter) Get() string {
	return t.token
}

func (t *tokenGetter) AutoRefresh(timeSpacing time.Duration) error {
	if t.timer != nil {
		t.timer.Stop()
	}
	if err := t.tokenRequest(); err != nil {
		return err
	}
	if t.timer != nil {
		t.timer.Reset(timeSpacing)
	} else {
		t.timer = time.NewTimer(timeSpacing)
	}
	go func() {
		for {
			<-t.timer.C
			oldToken := t.token
			if err := t.tokenRequest(); err != nil {
				// 打印log
				t.timer.Reset(tokenRetry)
			} else {
				// 如果token申请成功，tokenCloseAfter 时间后关闭老token
				go func() {
					<-time.NewTimer(tokenCloseAfter).C
					reply, err := token.SetAuthClose(t.appId, oldToken)
					if err == nil && reply.Result != "ok" {
						err = errors.New(util.RESULT_MAP[reply.Result])
					}
					if err != nil {
						// 打印log
					}
				}()
				t.timer.Reset(timeSpacing)
			}
		}
	}()
	return nil
}

func (t *tokenGetter) tokenRequest() error {
	reply, err := token.GetAuthSign(t.appId, t.appKey, t.masterSecret)
	if err != nil {
		return err
	}
	if reply.Result == "ok" {
		t.token = reply.AuthToken
	} else {
		return errors.New(util.RESULT_MAP[reply.Result])
	}
	return nil
}
