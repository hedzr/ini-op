/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"github.com/hedzr/logex/build"
)

func Entry() {

	if err := cmdr.Exec(rootCmd,
		cmdr.WithEnvPrefix("IO"),
		cmdr.WithLogx(build.New(log.NewLoggerConfigWith(true, "sugar", "debug"))),
		cmdr.WithWatchMainConfigFileToo(true),
	); err != nil {
		log.Errorf("Error: %v", err)
	}

}
