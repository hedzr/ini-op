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
	"bufio"
	"fmt"
	"github.com/hedzr/ini-op/misc"
	"os"
	"path"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/hedzr/ini-op/cli_common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "API",
	//	Short: "A brief description of your application",
	//	Long: `A longer description that spans multiple lines and likely contains
	//examples and usage of using your application. For example:
	//
	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PersistentPreRun with args: %v, args: %v\n", cmd, args)
		//initLogger()
	},
	//PreRun: func(cmd *cobra.Command, args []string) {
	//	fmt.Printf("Inside rootCmd PreRun with args: %v, args: %v\n", cmd, args)
	//},
	Run: func(command *cobra.Command, args []string) {
		//if cmd1 == RootCmd {
		if command.HasPersistentFlags() {
			if ok, _ := command.PersistentFlags().GetBool("version"); ok {
				if cli_common.PrintVersion != nil {
					cli_common.PrintVersion()
				} else {
					fmt.Printf("%s version: %s [%s] | built at %s with %s | by @hedzr13\n",
						cli_common.AppName, cli_common.Version, cli_common.Githash,
						cli_common.Buildstamp, cli_common.GoVersion)
				}
				return
			}
		}
		//}
		//fmt.Printf("Inside rootCmd Run with args: %v, args: %v\n", cmd1, args)
		command.Help()
	},
	//PostRun: func(cmd *cobra.Command, args []string) {
	//	fmt.Printf("Inside rootCmd PostRun with args: %v, args: %v\n", cmd, args)
	//},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		//fmt.Printf("Inside rootCmd PersistentPostRun with args: %v, args: %v\n", cmd, args)
	},
}

func SetAppName(name string, version string) {
	cli_common.AppName = name
	RootCmd.Use = name
	cli_common.Version = version
	cli_common.ServerTag = fmt.Sprintf("%s/%s", cli_common.AppName, cli_common.Version)

	serviceName := viper.GetString("server.serviceName")
	if len(serviceName) == 0 {
		viper.Set("server.serviceName", name)
	}

	logDir := viper.GetString("server.logger.dir")
	if len(logDir) == 0 {
		ok := false
		for _, ptn := range []string{"/var/log/%s", os.Getenv("HOME") + "/.%s/log", "/tmp/%s"} {
			logDir = fmt.Sprintf(ptn, name)
			if _, err := os.Stat(logDir); os.IsNotExist(err) {
				if err = os.MkdirAll(logDir, 0775); err != nil {
					//
				} else {
					viper.Set("server.logger.dir", logDir)
					ok = true
					break
				}
			} else {
				viper.Set("server.logger.dir", logDir)
				ok = true
				break
			}
		}
		if !ok {
			log.Fatalf("cannot create log dir at: %s", logDir)
		}
	}
}

// SetRealServerStart to customize the http2 server start function and registrar deregistering function.
// <p/>
// realStart() sample:
// ```go
// // `MyAuthService`: implements github.com/univer/cli/cli_common::CoolServer
// func realStart() {
//	   e := server.New(&impl.MyAuthService{})
//	   defer e.Start()() //will block
//	   //fmt.Println(" END")
// }
// ```
// deregister()
func SetRealServerStart(realStart func(), deregister func()) {
	cli_common.RealStart = realStart
	cli_common.Deregister = deregister
}

func SetPrintVersion(printer func()) {
	cli_common.PrintVersion = printer
}

func SetXX() {
	//
}

// Execute adds all child commands to the root command sets flags appropriately.
//   This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cli_common.CfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cli_common.CfgFile)
	}

	// Viper supports reading from yaml, toml and/or json files. Viper can
	// search multiple paths. Paths will be searched in the order they are
	// provided. Searches stopped once Config File found.

	viper.SetConfigName(cli_common.AppName) // name of config file (without extension)

	var f = fmt.Sprintf("%s/.%s", os.Getenv("HOME"), cli_common.AppName)
	if _, err := os.Stat(f); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(f, 0755)
	}
	var ff = fmt.Sprintf("%s/%s.yml", f, cli_common.AppName)
	if _, err := os.Stat(ff); os.IsNotExist(err) {
		_, err = os.OpenFile(ff, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}

	viper.AddConfigPath(fmt.Sprintf("/etc/%s", cli_common.AppName))
	viper.AddConfigPath(f)
	//viper.AddConfigPath("/tmp")          // path to look for the config file in
	viper.AddConfigPath(".") // more path to look for the config files
	viper.AddConfigPath(fmt.Sprintf("./ci/etc/%s", cli_common.AppName))

	viper.SetEnvPrefix("AUTH")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//if log.GetLevel() > log.InfoLevel {
		//	fmt.Println("Using config file:", viper.ConfigFileUsed())
		//}

		level := viper.Get("server.logger.level")
		if level == nil {
			viper.Set("server.logger.level", "WARN")
			level = viper.Get("server.logger.level")
		}
		if cli_common.Verbose && cli_common.Debug {
			viper.Set("server.logger.level", "DEBUG")
			level = viper.Get("server.logger.level")
			log.SetLevel(log.DebugLevel)
		}

		log.Debug("=================================")
		log.Debugf("Using config file: %s", viper.ConfigFileUsed())
		log.Debugf("server.logger.level = %s ; %s", level, log.GetLevel())

		misc.InitLogger()
		//fmt.Printf("server.logger.level = %s ; %s\n", level, log.GetLevel())

		// load conf.d/*.yml
		ddir := fmt.Sprintf("%s/conf.d", path.Dir(viper.ConfigFileUsed()))
		err := filepath.Walk(ddir, visit)
		if err != nil {
			fmt.Printf("ERROR: filepath.Walk() returned %v\n", err)
		}
	}

	if log.GetLevel() > log.InfoLevel {
		log.Infof("server.logger.level = %s", viper.Get("server.logger.level"))
	}

	viper.WatchConfig()
	viper.OnConfigChange(reloadConfig)
	// TODO watch conf.d sub-directory too.
}

// 实现了配置文件的拆分，更好地编辑和维护配置文件
func visit(path string, f os.FileInfo, err error) error {
	//fmt.Printf("Visited: %s, e: %v\n", path, err)
	if f != nil && !f.IsDir() {
		// log.Infof("    path: %v, ext: %v", path, filepath.Ext(path))
		switch filepath.Ext(path) {
		case ".yml", ".yaml", "yml", "yaml":
			if cli_common.VerboseVVV {
				log.Infof("    merging from: %s", path)
			}
			file, err := os.Open(path)
			if err != nil {
				log.Warnf("ERROR: os.Open() returned %v", err)
			} else {
				viper.MergeConfig(bufio.NewReader(file))
				if cli_common.VerboseVVV {
					log.Infof("server.logger.level = %s", viper.Get("server.logger.level"))
				}
				//env := viper.Get("server.registrar.env")
				//key := fmt.Sprintf("server.registrar.consul.%s.addr", env)
				//log.Infof("%s = %s", key, viper.Get(key))
			}
		}
	}
	return nil
}

func reloadConfig(e fsnotify.Event) {
	// todo reload the configurations from file
	//log.Info()
	log.Info("\n\nConfig file changed:\n", e.Name)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cli_common.CfgFile, "config", "", fmt.Sprintf("config file (default is $HOME/.%s/*.yaml)", cli_common.AppName))

	if cli_common.DebugBuild {
		// Cobra also supports local flags, which will only run
		// when this action is called directly.
		RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	}
	RootCmd.PersistentFlags().BoolP("version", "V", false, "Show versions")

	RootCmd.PersistentFlags().BoolVarP(&cli_common.Verbose, "verbose", "v", false, "verbose output")
	//RootCmd.PersistentFlags().BoolVarP(&cli_common.VerboseVV, "vv", "vv", false, "very verbose output")
	//RootCmd.PersistentFlags().BoolVarP(&cli_common.VerboseVVV, "vvv", "vvv", false, "very very verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cli_common.Debug, "debug", "d", false, "debug mode")
}
