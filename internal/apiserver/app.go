package apiserver

import (
	"github.com/yshujie/questionnaire-scale/pkg/app"
	"github.com/yshujie/questionnaire-scale/pkg/log"
)

// commandDesc 命令描述
const commandDesc = `The Questionnaire Scale API server validates and configures data
for the api objects which include users, policies, secrets, and
others. The API Server services REST operations to do the api objects management.`

// NewApp 创建 App
func NewApp(basename string) *app.App {
	opts := NewOptions()
	application := app.NewApp("Questionnaire Scale API Server",
		basename,
		app.WithDescription(commandDesc),
		app.WithDefaultValidArgs(),
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)

	return application
}

func run(opts *Options) app.RunFunc {
	return func(basename string) error {
		// 初始化日志（使用从配置文件加载的配置）
		log.Init(opts.Log)
		defer log.Flush()

		log.Info("Starting questionnaire-scale ...")

		// 打印配置信息
		log.Infof("Server mode: %s", opts.Server.Mode)
		log.Infof("Health check enabled: %v", opts.Server.Healthz)

		return nil
	}
}
