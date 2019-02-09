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

package misc

import (
	"fmt"
	"github.com/coreos/go-systemd/journal"
	log "github.com/sirupsen/logrus"
	//"github.com/cihub/seelog"
	//"path"
	//"strings"
	////"github.com/op/go-logging"
	"github.com/spf13/viper"
	"os"

	"github.com/hedzr/univer/daemon"
)

func GetLevel() log.Level {
	return log.GetLevel()
}

func InitLogger() {
	var foreground = viper.GetBool("server.foreground")
	var file = daemon.DefaultLogFile()
	var lvl = viper.GetString("server.logger.level")

	var target = viper.GetString("server.logger.target")
	var format = viper.GetString("server.logger.format")
	if len(target) == 0 {
		target = "default"
	}
	if len(format) == 0 {
		format = "text"
	}
	if target == "journal" {
		format = "text"
	}
	switch format {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.TextFormatter{ForceColors: true})
	}
	// Log as JSON instead of the default ASCII formatter.

	// Only log the warning severity or above.
	l, _ := log.ParseLevel(lvl)
	log.SetLevel(l)
	log.Debugf("Using logger: format=%s, level=%s, target=%s", format, lvl, target)

	can_use_log_file := true
	// daemon mode 才会发送给 journal
	if journal.Enabled() && target == "journal" && foreground == false {
		can_use_log_file = false
		//sink = NewJournalSink()
	} else {
		//sink = NewDefaultSink()
	}

	if foreground == false && can_use_log_file {
		if len(file) == 0 {
			file = os.DevNull //"/dev/null"
		}

		logFile, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0400)
		if err != nil {
			logFile, err = os.Create(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				log.Debugf("Using new log file: %s\n\n", file)
			}
		} else {
			log.Debugf("Using exists log file: %s\n\n", file)
		}

		//log.Infof("Using log file: %s", file)
		//fmt.Printf("Using log file: %s\n\n", file)

		// Output to stdout instead of the default stderr
		// Can be any io.Writer, see below for File example
		//log.SetOutput(os.Stdout)
		log.SetOutput(logFile)
	}
	//var lvl = viper.GetString("server.logger.level")
	//var file = daemon.DefaultLogFile()
	//var err error
	//
	//logFile, err := os.OpenFile(file, os.O_WRONLY,0400)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//lf, err := logging.LogLevel(lvl)
	//backend1 := logging.NewLogBackend(logFile, "", 0)
	//backend1Leveled := logging.AddModuleLevel(backend1)
	//backend1Leveled.SetLevel(lf, "")
	//
	//if foreground {
	//	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	//	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	//	logging.SetBackend(backend1Leveled, backend2Formatter)
	//}

	//log.Debugf("debug %s", Password("secret"))
	//log.Info("info")
	//log.Notice("notice")
	//log.Warning("warning")
	//log.Error("suweia.com")
	//log.Critical("太严重了")
	//os.Exit(0)
}

//func InitLoggerLast() {
//	var logger seelog.LoggerInterface
//	var err error
//	var foreground = viper.GetBool("server.foreground")
//	var basicLogConfig string
//
//	homeDir := os.Getenv("HOME")
//	seelogPath := path.Join(homeDir, DEFAULT_SEELOG_FILENAME)
//	if _, err := os.Stat(seelogPath); err == nil {
//		logger, err = seelog.LoggerFromConfigAsFile(seelogPath)
//	} else {
//		if foreground {
//			basicLogConfig = DEFAULT_SEELOG_CONFIG
//		} else {
//			basicLogConfig = DAEMON_SEELOG_CONFIG
//		}
//
//		basicLogConfig := strings.Replace(basicLogConfig, "${dir}", daemon.DefaultLogDirectory(), -1)
//
//		lvl := viper.GetString("server.logger.level")
//		basicLogConfig = strings.Replace(basicLogConfig, "minlevel=\"info\"", fmt.Sprintf("minlevel=\"%s\"", lvl), -1)
//
//		//if foreground == false {
//		//	basicLogConfig = strings.Replace(basicLogConfig, "<format id=\"main-console\" format=\"[%LEV] %Msg%n\"/>", "", -1)
//		//}
//
//		//fmt.Printf("@@@@\n%s\n@@@@\n", basicLogConfig)
//
//		logger, err = seelog.LoggerFromConfigAsString(basicLogConfig)
//	}
//	if err != nil {
//		seelog.Critical("err parsing config log file", err)
//		return
//	}
//
//	seelog.ReplaceLogger(logger)
//	seelog.Flush()
//
//	//log.Trace("TRACE")
//	//log.Debug("DEBUG")
//	//log.Info("INFO")
//	//log.Warn("WARN")
//	//log.Error("ERROR")
//	//log.Critical("CRITICAL")
//}
