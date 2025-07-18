package app

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/yshujie/questionnaire-scale/pkg/errors"
	cliflag "github.com/yshujie/questionnaire-scale/pkg/flag"
	"github.com/yshujie/questionnaire-scale/pkg/log"
	"github.com/yshujie/questionnaire-scale/pkg/term"
	"github.com/yshujie/questionnaire-scale/pkg/version"
	"github.com/yshujie/questionnaire-scale/pkg/version/verflag"
)

var (
	progressMessage = color.GreenString("==>")
)

// App 应用
type App struct {
	basename    string
	name        string
	description string
	noVersion   bool
	noConfig    bool
	silence     bool
	options     CliOptions
	cmd         *cobra.Command
	args        cobra.PositionalArgs
	commands    []*Command
	runFunc     RunFunc
}

// Option 应用选项
type Option func(*App)

// RunFunc 定义应用程序的启动回调函数
type RunFunc func(basename string) error

// WithDescription 设置应用程序的描述
func WithDescription(description string) Option {
	return func(a *App) {
		a.description = description
	}
}

// WithOptions 打开应用程序的函数，从命令行或配置文件中读取参数
func WithOptions(opt CliOptions) Option {
	return func(a *App) {
		a.options = opt
	}
}

// WithRunFunc 设置应用程序的启动回调函数选项
func WithRunFunc(run RunFunc) Option {
	return func(a *App) {
		a.runFunc = run
	}
}

// WithValidArgs 设置 args
func WithValidArgs(args cobra.PositionalArgs) Option {
	return func(a *App) {
		a.args = args
	}
}

// WithDefaultValidArgs 设置默认的 args
func WithDefaultValidArgs() Option {
	return func(a *App) {
		a.args = func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}

			return nil
		}
	}
}

// NewApp 创建应用
func NewApp(name string, basename string, opts ...Option) *App {
	// 创建 App
	a := &App{
		name:     name,
		basename: basename,
	}
	// 设置应用选项
	for _, opt := range opts {
		opt(a)
	}

	// 构建命令
	a.buildCommand()

	// 返回 app
	return a
}

// buildCommand 构建命令
func (a *App) buildCommand() {
	// 使用 cobra 创建命令
	cmd := &cobra.Command{
		Use:           FormatBaseName(a.basename),
		Short:         a.name,
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}

	// 设置输出和错误输出
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)

	// 设置排序
	cmd.Flags().SortFlags = true

	// 初始化命令行参数
	cliflag.InitFlags(cmd.Flags())

	// 如果命令不为空，则添加命令
	if len(a.commands) > 0 {
		// 添加命令
		for _, command := range a.commands {
			cmd.AddCommand(command.cobraCommand())
		}
		// 设置帮助命令
		cmd.SetHelpCommand(helpCommand(FormatBaseName(a.basename)))
	}

	// 如果启动回调函数不为空，则设置启动回调函数
	if a.runFunc != nil {
		cmd.RunE = a.runCommand
	}

	// 如果选项不为空，则添加选项
	var namedFlagSets cliflag.NamedFlagSets
	if a.options != nil {
		namedFlagSets = a.options.Flags()
		fs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			fs.AddFlagSet(f)
		}
	}

	// 如果版本标志不为空，则添加版本标志
	if !a.noVersion {
		verflag.AddFlags(namedFlagSets.FlagSet("global"))
	}

	// 如果配置标志不为空，则添加配置标志
	if !a.noConfig {
		addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))
	}

	// 添加全局标志到命令标志集
	cmd.Flags().AddFlagSet(namedFlagSets.FlagSet("global"))

	// 添加命令模板
	addCmdTemplate(cmd, namedFlagSets)

	// 设置命令
	a.cmd = cmd
}

// Run 运行应用程序
func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		fmt.Printf("%v %v\n", color.RedString("Error:"), err)
		os.Exit(1)
	}
}

// runCommand 运行命令
func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	// 打印工作目录
	printWorkingDir()
	// 打印命令行参数
	cliflag.PrintFlags(cmd.Flags())

	// 打印版本信息
	if !a.noVersion {
		verflag.PrintAndExitIfRequested()
	}

	// 如果配置标志不为空，则绑定命令行标志
	if !a.noConfig {
		// 绑定命令行标志
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}

		// 打印配置信息
		fmt.Printf("Viper Config: %+v\n", viper.AllSettings())

		// 如果选项不为空，则反序列化选项
		if err := viper.Unmarshal(a.options); err != nil {
			return err
		}

		// 打印选项信息
		fmt.Printf("Options: %+v\n", a.options)

		// 打印 secure 配置
		fmt.Printf("Secure Config: %+v\n", viper.Get("secure"))
		fmt.Printf("Secure TLS Config: %+v\n", viper.Get("secure.tls"))
		fmt.Printf("Secure TLS Cert File: %+v\n", viper.Get("secure.tls.cert-file"))
		fmt.Printf("Secure TLS Private Key File: %+v\n", viper.Get("secure.tls.private-key-file"))
	}

	// 如果静默标志不为空，则打印日志
	if !a.silence {
		log.Infof("%v Starting %s ...", progressMessage, a.name)
		if !a.noVersion {
			log.Infof("%v Version: `%s`", progressMessage, version.Get().ToJSON())
		}
		if !a.noConfig {
			log.Infof("%v Config file used: `%s`", progressMessage, viper.ConfigFileUsed())
		}
	}

	// 如果选项不为空，则应用选项规则
	if a.options != nil {
		if err := a.applyOptionRules(); err != nil {
			return err
		}
	}

	// 运行应用程序
	if a.runFunc != nil {
		return a.runFunc(a.basename)
	}

	return nil
}

// applyOptionRules 应用选项规则
func (a *App) applyOptionRules() error {
	if completeableOptions, ok := a.options.(CompleteableOptions); ok {
		if err := completeableOptions.Complete(); err != nil {
			return err
		}
	}

	if errs := a.options.Validate(); len(errs) != 0 {
		return errors.NewAggregate(errs)
	}

	if printableOptions, ok := a.options.(PrintableOptions); ok && !a.silence {
		log.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}

// printWorkingDir 打印工作目录
func printWorkingDir() {
	wd, _ := os.Getwd()
	log.Infof("%v WorkingDir: %s", progressMessage, wd)
}

// addCmdTemplate 添加命令模板
func addCmdTemplate(cmd *cobra.Command, namedFlagSets cliflag.NamedFlagSets) {
	usageFmt := "Usage:\n  %s\n"
	cols, _, _ := term.TerminalSize(cmd.OutOrStdout())
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStderr(), namedFlagSets, cols)

		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
		cliflag.PrintSections(cmd.OutOrStdout(), namedFlagSets, cols)
	})
}
