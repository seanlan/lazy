package init

import "go.uber.org/zap"

func init() {
	conf := zap.NewDevelopmentConfig()
	//conf.DisableCaller = true
	logger, _ := conf.Build()
	zap.ReplaceGlobals(logger)
}
