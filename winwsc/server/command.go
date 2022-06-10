package server

import (
	"bytes"
	"fmt"
	"github.com/axgle/mahonia"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

func Run(action, service string) {
	if action == "start" {
		_start(service)
	}

	if action == "stop" {
		_stop(service)
	}

	if action == "restart" {
		_stop(service)
		_start(service)
	}
}

func _start(service string) {
	command, _ := exec.LookPath("net")
	cmd := exec.Command(command, "start", service)
	_exec(cmd, service)
}

func _stop(service string) {
	command, _ := exec.LookPath("net")
	cmd := exec.Command(command, "stop", service)
	_exec(cmd, service)
}

func _all(ss []string) {
	if len(ss) <= 0 {
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(ss))
	for _, s := range ss {
		go func(server string) {
			if NewCmd().CurAction == "start" {
				_start(server)
			} else {
				_stop(server)
			}
			wg.Done()
		}(s)
	}
	wg.Wait()
}

func _exec(cmd *exec.Cmd, service string) {
	var out bytes.Buffer
	cmd.Stdout = &out

	var outerr bytes.Buffer
	cmd.Stderr = &outerr

	err := cmd.Run()
	enc := mahonia.NewDecoder("gbk")

	if err != nil {
		// 读一行
		line, _ := outerr.ReadString('\r')
		errStr := enc.ConvertString(line)
		if NewCmd().CurAction == "start" {
			fmt.Println(service, errStr)
		} else {
			fmt.Println(errStr)
		}
	}
	fmt.Print(enc.ConvertString(string(out.Bytes())))
}

// 状态
func Status(service string) string {

	command, _ := exec.LookPath("sc")
	cmd := exec.Command(command, "query", service)

	var out bytes.Buffer
	cmd.Stdout = &out
	var outerr bytes.Buffer
	cmd.Stderr = &outerr
	err := cmd.Run()
	enc := mahonia.NewDecoder("gbk")

	if err != nil {
		//err := enc.ConvertString(string(outerr.Bytes()))
		return "Not found.."
	}

	msg := enc.ConvertString(string(out.Bytes()))

	expr, exprErr := regexp.Compile(`(?i)STATE[^:]+:[^a-z]+([^\s]+)`)
	if exprErr != nil {
		return "Not found."
	}

	re := expr.FindStringSubmatch(msg)
	if len(re) > 1 {
		return strings.ToLower(re[1])
	}

	return "Not found"
}
