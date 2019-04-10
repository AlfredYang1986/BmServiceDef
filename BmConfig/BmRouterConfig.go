package BmConfig

import (
	"os"
)

type BmRouterConfig struct {
	Host   string
	Port   string
	TmpDir string
}

func (br *BmRouterConfig) GenerateConfig(envHome string) {
	bmHome := os.Getenv(envHome)
	configPath := bmHome + "/resource/routerconfig.json"
	println(configPath)
	profileItems := BmGetConfigMap(configPath)

	br.Host = profileItems["Host"].(string)
	br.Port = profileItems["Port"].(string)
	br.TmpDir = bmHome + "/" + profileItems["TmpDir"].(string)
}