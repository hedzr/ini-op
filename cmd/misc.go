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
	"github.com/hedzr/ini-op/cli_common"
	"github.com/hedzr/ini-op/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// The main command describes the service and defaults to printing the
// help message.
var mainCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "Event dispatch service.",
	Long:  `HTTP service that consumes events and dispatches them to subscribers.`,
	Run: func(cmd *cobra.Command, args []string) {
		misc.Serve()
	},
}

// The version command prints this service.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version.",
	Long:  "The version of the dispatch service.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", cli_common.Version)
	},
}

func init() {
	if cli_common.DebugBuild {
		RootCmd.AddCommand(mainCmd)
	}

	RootCmd.AddCommand(versionCmd)

	if cli_common.DebugBuild {
		flags := mainCmd.Flags()

		flags.Bool("debug", false, "Turn on debugging.")
		flags.String("addr", "localhost:5002", "Address of the service")
		flags.String("smtp-addr", "localhost:25", "Address of the SMTP server")
		flags.String("smtp-user", "", "User to authenticate with the SMTP server")
		flags.String("smtp-password", "", "Password to authenticate with the SMTP server")
		flags.String("email-from", "noreply@example.com", "The from email address.")

		viper.BindPFlag("debug", flags.Lookup("debug"))
		viper.BindPFlag("addr", flags.Lookup("addr"))
		viper.BindPFlag("smtp_addr", flags.Lookup("smtp-addr"))
		viper.BindPFlag("smtp_user", flags.Lookup("smtp-user"))
		viper.BindPFlag("smtp_password", flags.Lookup("smtp-password"))
		viper.BindPFlag("email_from", flags.Lookup("email-from"))
	}
}
