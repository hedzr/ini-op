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

var iniSectionCmd = &cobra.Command{
	Use:     "section",
	Aliases: []string{"s", "sec"},
	Short:   "get/put section",
	Long:    `.`,
}

var iniSectionReadCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"r", "rd", "read"},
	Short:   "get section",
	Long:    `.`,
	Run: func(cmd *cobra.Command, args []string) {
		//misc.Serve()

		if len(args) == 2 {
			cfg, err := ini.Load(args[1])
			if err != nil {
				log.Fatalf("Fail to read file: %v", err)
				os.Exit(1)
			}
			log.Infof("loading section '%s' from '%s'", args[0], args[1])
			section := cfg.Section(args[0])
			for i, k := range section.Keys() {
				fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
			}
		} else {
			cmd.Help()
		}
	},
}

var iniSectionRmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "erase", "del", "delete"},
	Short:   "delete section",
	Long:    `.`,
	Run: func(cmd *cobra.Command, args []string) {
		//misc.Serve()
		if len(args) == 2 {
			// delete section from ini file
			cfg, err := ini.Load(args[1])
			if err != nil {
				log.Fatalf("Fail to read file: %v", err)
				os.Exit(1)
			}

			cfg.DeleteSection(args[0])

			if err := cfg.SaveTo(args[1]); err != nil {
				log.Fatalf("Fail to write file: %v", err)
				os.Exit(1)
			}

			//} else if len(args) == 1 {
			// read from stdin
		} else {
			cmd.Help()
		}
	},
}

func init() {
	iniSectionCmd.AddCommand(iniSectionReadCmd)
	iniSectionCmd.AddCommand(iniSectionRmCmd)

	RootCmd.AddCommand(iniSectionCmd)
}
