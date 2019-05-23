/*
 * Copyright © 2019 Hedzr Yeh.
 */

package main

import "github.com/hedzr/ini-op/cli/impl"

func main() {
	// cmd.SetAppName(impl.APPNAME, impl.Version)
	// cmd.SetRealServerStart(realStart, deregister)
	// cmd.Execute()
	impl.Entry()
}

// func realStart() {
// 	// e := server.New(gwkserver.New())
// 	// defer e.Start()() 	//will block here
// 	println("real start.")
// }
//
// func deregister() {
// 	// server.Deregister()	// deregister self from consul/etcd registrar (13.registrar.yml)
// 	println("deregister")
// }
