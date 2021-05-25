/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
)

func Pre(cmd *cmdr.Command, args []string) (err error) {
	if cmdr.GetDebugMode() {
		log.SetLevel(log.DebugLevel)
	}
	log.Debug("app starting...")
	return
}

func Post(cmd *cmdr.Command, args []string) {
	log.Debug("app stopping...")
	log.Debug("app stopped.")
}
