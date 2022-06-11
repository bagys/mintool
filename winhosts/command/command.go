package command

import (
	"bufio"
	"fmt"
	"github.com/bagys/ctb"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func t() string {

	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command] {{end}}{{if .HasExample}}

command:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasExample}}

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
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}
`
}

type cmds struct {
	RootCmd *cobra.Command
}

func NewCmds() *cmds {
	c := &cmds{}
	c.RootCmd = &cobra.Command{
		Use:              "hosts ",
		TraverseChildren: true,
		Example:          "hosts ls",
	}
	// 使用自己的模板
	c.RootCmd.SetUsageTemplate(t())
	c.RootCmd.CompletionOptions.DisableDefaultCmd = true
	// 不排序命令
	cobra.EnableCommandSorting = false

	help := &cobra.Command{}
	c.RootCmd.SetHelpCommand(help)

	c.ls()
	c.add()
	return c
}

func (c *cmds) ls() {
	lscd := &cobra.Command{
		Use:                   "ls",
		DisableFlagsInUseLine: true,
		Long:                  "List all hosts",
		Run: func(cmd *cobra.Command, args []string) {
			ls()
		},
	}
	c.RootCmd.AddCommand(lscd)
}

func (c *cmds) add() {
	addcd := &cobra.Command{
		Use:                   "add",
		DisableFlagsInUseLine: true,
		Example:               "hosts add --ip=127.0.0.1 --host=localhost \n  hosts add -i 127.0.0.1 -h localhost",
		Run: func(cmd *cobra.Command, args []string) {

			ip, _ := cmd.Flags().GetString("ip")
			host, _ := cmd.Flags().GetString("host")

			if net.ParseIP(ip) == nil || host == "" {
				cmd.Help()
				fmt.Println()
				return
			}
			hostadd(ip, host)
		},
	}
	// 不排序
	addcd.Flags().SortFlags = false

	// 覆盖默认的 help
	addcd.Flags().Bool("help", false, "help")
	addcd.Flags().MarkHidden("help")

	addcd.Flags().StringP("ip", "i", "", "register ip")
	_ = addcd.MarkFlagRequired("ip")

	addcd.Flags().StringP("host", "h", "", "register host")
	_ = addcd.MarkFlagRequired("host")

	addcd.SetUsageTemplate(z())
	c.RootCmd.AddCommand(addcd)

}

func ls() {
	T := ctb.NewTable()
	T.SetPrefixDisable(false)
	f, errf := os.Open(`C:\Windows\System32\drivers\etc\hosts`)
	defer f.Close()
	if errf != nil {
		log.Fatal(errf)
	}
	read := bufio.NewReader(f)

	for {
		line, err := read.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}

		if line[:1] == "#" {
			continue
		}

		re, _ := regexp.Compile(`\s+`)
		line = strings.Trim(re.ReplaceAllString(line, " "), " ")

		hostArr := strings.Split(line, " ")
		if len(hostArr) <= 1 {
			continue
		}

		ipOb := net.ParseIP(hostArr[0])
		if ipOb == nil {
			continue
		}
		ipstr := ipOb.String()
		for _, domain := range hostArr[1:] {
			T.SetDataOne([]ctb.LineData{
				{Data: ipstr, Color: color.Green},
				{Data: domain, Color: color.Green},
			})
		}
	}

	T.SetTab([]ctb.LineData{
		{Data: "IP"},
		{Data: "HOST"},
	})
	T.Print()
}

func hostadd(ip, host string) {
	f, errf := os.OpenFile(`C:\Windows\System32\drivers\etc\hosts`, os.O_APPEND, 0644)
	defer f.Close()
	if errf != nil {
		log.Fatal(errf)
	}

	str := fmt.Sprintf("%s %s\r\n", ip, host)
	_, errWrite := f.Write([]byte(str))
	if errWrite != nil {
		log.Fatal(errWrite)
	}

	fmt.Printf("ip:%s - host:%s\n", ip, host)
	color.Greenln("Success create")
}

func (this *cmds) Execute() {
	this.RootCmd.Execute()
}
