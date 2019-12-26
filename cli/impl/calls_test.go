/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl_test

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/hedzr/logex"
	"os"
	"testing"
)

func TestFine(t *testing.T) {
	// test123("dsds=1")
}

func test123(ff string) {
	defer logex.CaptureLog(t).Release()

	// read ini file
	cfg, err := ini.Load(ff)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())

	// 我们可以做一些候选值限制的操作
	fmt.Println("Server Protocol:",
		cfg.Section("server").Key("protocol").In("http", []string{"http", "https"}))
	// 如果读取的值不在候选列表内，则会回退使用提供的默认值
	fmt.Println("Email Protocol:",
		cfg.Section("server").Key("protocol").In("smtp", []string{"imap", "smtp"}))

	// 试一试自动类型转换
	fmt.Printf("Port Number: (%[1]T) %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Printf("Enforce Domain: (%[1]T) %[1]v\n", cfg.Section("server").Key("enforce_domain").MustBool(false))

	// 差不多了，修改某个值然后进行保存
	cfg.Section("").Key("app_mode").SetValue("production")
	if err := cfg.SaveTo(ff); err != nil {
		fmt.Printf("Fail to write file: %v", err)
		os.Exit(1)
	}
}
