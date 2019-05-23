/*
 * Copyright © 2019 Hedzr Yeh.
 */

package ini_op

const (
//APP//NAME = "ini-op"
//VERSION = "0.3.5"
)

var (
	//Verbose    bool = false
	//VerboseVV  bool = false
	//VerboseVVV bool = false
	//Debug      bool = false

	AppExitCh = make(chan bool)

	// Version    = ""
	// Buildstamp = ""
	// Githash    = ""

	//ServerTag = fmt.Sprintf("GwkGateway/%s", Version)
	//ServerID = ""
	//
	//Version = "0.3.1"
	//Buildstamp = ""
	//Githash = ""

	apiVer    = "v1"
	apiPrefix = "/" + apiVer + "/api"
)

func GetApiVersion() string {
	return apiVer
}
func GetApiPrefix() string {
	return apiPrefix
}
