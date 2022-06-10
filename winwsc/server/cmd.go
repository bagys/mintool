package server

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"sync"
)

func t() string {

	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command] [service] {{end}}{{if .HasExample}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Examples:
  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}
{{end}}
`
}

func z() string {

	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} <service>{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Service:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}
`
}

type Cmd struct {
	ActionCmd   []*cobra.Command
	ServicesCmd []*cobra.Command
	RootCmd     *cobra.Command
	Services    []map[string]string
	Services2   []string

	CurAction  string
	CurService string
}

var cmds *Cmd
var once sync.Once

func NewCmd() *Cmd {
	once.Do(func() {
		cmds = &Cmd{}

		cmds.Services = NewServiceConfig().List()
		for _, service := range cmds.Services {
			cmds.Services2 = append(cmds.Services2, service["svcname"])
		}

		cmds.RootCmd = &cobra.Command{
			Use:              "wsc",
			TraverseChildren: true,
			Example:          "wsc ls\n  wsc startall nginx\n  wsc start nginx",
		}
		// 使用自己的模板
		cmds.RootCmd.SetUsageTemplate(t())
		cmds.RootCmd.CompletionOptions.DisableDefaultCmd = true

		help := &cobra.Command{}
		cmds.RootCmd.SetHelpCommand(help)

		cmds.registerAction()
		cmds.ActionBindSvc()
	})
	return cmds
}

// 注册动作
func (this *Cmd) registerAction() {
	for _, v := range [3]string{"start", "restart", "stop"} {

		cmd := &cobra.Command{
			Use:                   v,
			DisableFlagsInUseLine: true, // svc start [flags] 不显示 [flags]
			TraverseChildren:      true,
			Run: func(cmd *cobra.Command, args []string) {
				cmd.Help()
				fmt.Println()
			},
		}
		// 使用自己的模板
		cmd.SetUsageTemplate(z())
		this.ActionCmd = append(this.ActionCmd, cmd)
		this.RootCmd.AddCommand(cmd)
	}
}

// 注册服务绑定aciton
func (this *Cmd) ActionBindSvc() {
	// 循环动作
	for _, acmd := range this.ActionCmd {

		// 循环动作，每次必须生产新的 否则覆会覆盖上个
		for _, v := range this.Services {
			svcCmd := &cobra.Command{
				Use:                   v["svcname"],
				DisableFlagsInUseLine: true, // svc start [flags] 不显示 [flags]
				TraverseChildren:      true,
				Run: func(cmd *cobra.Command, args []string) {
					this.CurAction = cmd.Parent().Name()
					this.CurService = cmd.Use
					Run(cmd.Parent().Name(), cmd.Use)
				},
			}

			acmd.AddCommand(svcCmd)
		}
	}
}

// 打印列表
func (this Cmd) PrintList() {
	T := NewTable()
	T.SetTab([]LineData{
		{
			Data: "SERVICE",
		},
		{
			Data: "DESCRIBE",
		},
		{
			Data: "STATUS",
		},
	})

	for _, v := range this.Services {
		status := Status(v["svcname"])
		var C color.Color
		switch status {
		case "stopped":
			C = color.Yellow
		case "running":
			C = color.Green
		default:
			C = color.Red
		}
		T.SetData([]LineData{
			{
				Data: v["svcname"],
			},
			{
				Data: v["describe"],
			},
			{
				Data:  status,
				Color: C,
			},
		})
	}
	T.Print()
}

func (this *Cmd) Execute() {
	// 绑定 list
	this.RootCmd.AddCommand(&cobra.Command{
		Use:                   "ls",
		DisableFlagsInUseLine: true, // svc start [flags] 不显示 [flags]
		Run: func(cmd *cobra.Command, args []string) {
			this.PrintList()
		},
	})

	this.RootCmd.AddCommand(&cobra.Command{
		Use:                   "startall",
		DisableFlagsInUseLine: true, // svc start [flags] 不显示 [flags]
		Run: func(cmd *cobra.Command, args []string) {
			this.CurAction = "start"
			_all(this.Services2)
		},
	})

	this.RootCmd.AddCommand(&cobra.Command{
		Use:                   "stopall",
		DisableFlagsInUseLine: true, // svc start [flags] 不显示 [flags]
		Run: func(cmd *cobra.Command, args []string) {
			this.CurAction = "stop"
			_all(this.Services2)
		},
	})
	_ = this.RootCmd.Execute()
}
