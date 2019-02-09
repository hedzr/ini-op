/*
 * Copyright © 2019 Hedzr Yeh
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package cmd

import (
	"fmt"
	"github.com/go-ini/ini"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

// The iniEntryCmd command describes the service and defaults to printing the
// help message.
var iniEntryCmd = &cobra.Command{
	Use:     "entry",
	Aliases: []string{"e"},
	Short:   "get/put entry",
	Long:    `.`,
}

var iniEntryReadCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"r", "rd", "read"},
	Short:   "get an entry",
	Long:    `.`,
	Run: func(cmd *cobra.Command, args []string) {
		//misc.Serve()
		if len(args) == 3 {
			cfg, err := ini.Load(args[2])
			if err != nil {
				log.Fatalf("Fail to read file: %v", err)
				os.Exit(1)
			}
			log.Infof("loading section '%s' from '%s'", args[0], args[1])
			section := cfg.Section(args[0])
			//for i, k := range section.Keys() {
			//	fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
			//}
			fmt.Println(section.Key(args[1]))
		} else {
			cmd.Help()
		}
	},
}

var iniEntryWriteCmd = &cobra.Command{
	Use:     "put",
	Aliases: []string{"write", "set"},
	Short:   "put string entry",
	Long:    `.`,
	Run: func(cmd *cobra.Command, args []string) {
		//misc.Serve()
		if len(args) == 4 {
			cfg, err := ini.Load(args[3])
			if err != nil {
				log.Fatalf("Fail to read file: %v", err)
				os.Exit(1)
			}
			log.Infof("loading section '%s' from '%s'", args[0], args[1])
			section := cfg.Section(args[0])
			//for i, k := range section.Keys() {
			//	fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
			//}

			//fmt.Println(section.Key(args[1]))
			section.Key(args[1]).SetValue(args[2])
			log.Infof("set section '%s.%s' to '%s'", args[0], args[1], args[2])

			if err := cfg.SaveTo(args[3]); err != nil {
				log.Fatalf("Fail to write file: %v", err)
				os.Exit(1)
			}

		} else {
			cmd.Help()
		}
	},
}

var iniEntryRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "erase", "del", "delete"},
	Short:   "delete an entry",
	Long:    `.`,
	Run: func(cmd *cobra.Command, args []string) {
		//misc.Serve()
		//fmt.Println(args)
		if len(args) == 3 {
			// read ini file
			cfg, err := ini.Load(args[2])
			if err != nil {
				log.Fatalf("Fail to read file: %v", err)
				os.Exit(1)
			}

			cfg.Section(args[0]).DeleteKey(args[1])

			if err := cfg.SaveTo(args[2]); err != nil {
				log.Fatalf("Fail to write file: %v", err)
				os.Exit(1)
			}

			//} else if len(args) == 1 {
			//	// read from stdin
		} else {
			cmd.Help()
		}
	},
}

func test123(ff string) {
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

func init() {
	iniEntryCmd.AddCommand(iniEntryReadCmd)
	iniEntryCmd.AddCommand(iniEntryWriteCmd)
	iniEntryCmd.AddCommand(iniEntryRmCmd)

	RootCmd.AddCommand(iniEntryCmd)
}
