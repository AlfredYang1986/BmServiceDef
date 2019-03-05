package BmConfig

import "os"

type BmXmppConfig struct {
	Host          string
	Port          string
	HostName      string
	LoginUser     string
	LoginUserPwd  string
	ReportUser    string
	ReportUserPwd string
	ListenUser    string
	ListenUserPwd string
}

func (xc *BmXmppConfig) GenerateConfig(envHome string) {
	bmHome := os.Getenv(envHome)
	configPath := bmHome + "resource/xmppconfig.json"
	profileItems := BmGetConfigMap(configPath)

	xc.Host = profileItems["Host"].(string)
	xc.Port = profileItems["Port"].(string)
	xc.HostName = profileItems["HostName"].(string)
	xc.LoginUser = profileItems["LoginUser"].(string)
	xc.LoginUserPwd = profileItems["LoginUserPwd"].(string)
	xc.ReportUser = profileItems["ReportUser"].(string)
	xc.ReportUserPwd = profileItems["ReportUserPwd"].(string)
	xc.ListenUser = profileItems["ListenUser"].(string)
	xc.ListenUserPwd = profileItems["ListenUserPwd"].(string)
}
