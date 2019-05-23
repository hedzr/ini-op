/*
 * Copyright © 2019 Hedzr Yeh.
 */

package cli_common

var (
	CfgFile string
	AppName string

	Verbose    bool = false // verbose logging
	VerboseVV  bool = false
	VerboseVVV bool = false
	Debug      bool = false // debug mode
	DebugBuild bool = false // Debug Build

	Version    = "0.2.1"
	Buildstamp = ""
	Githash    = ""
	GoVersion  = ""

	ServerTag = ""
	ServerID  = ""

	RealStart    func() = nil
	Deregister   func() = nil
	PrintVersion func() = nil
)
