package init

import "go.uber.org/zap"

func init() {
	conf := zap.NewDevelopmentConfig()
	logger, _ := conf.Build()
	zap.ReplaceGlobals(logger)
}
