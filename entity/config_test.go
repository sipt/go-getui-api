package entity

import (
	"testing"
	"time"
)

var (
	AppID        = "fpY3MeiAjm6A46OQhbbWm9"
	AppSecret    = "vof6tggfrE6b8dNb0QpLj5"
	AppKey       = "hUO99mPghLAdIT6oJUc8z4"
	MasterSecret = "H7uEHgcGeA7NNAKMfgvi14"
)

func TestAppConfig(t *testing.T) {
	conf := NewAppConfig(AppID, AppSecret, AppKey, MasterSecret)
	t.Error("app_id:", conf.GetAppID(), ", token: ", conf.GetToken())
	timer := time.NewTimer(20 * time.Second)
	<-timer.C
	t.Error("app_id:", conf.GetAppID(), ", token: ", conf.GetToken())
	timer.Reset(10 * time.Second)
}
