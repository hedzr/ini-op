/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"github.com/hedzr/cmdr"
	"github.com/sirupsen/logrus"
)

func Entry() {

	if err := cmdr.Exec(rootCmd,
		cmdr.WithEnvPrefix("IO"),
		cmdr.WithLogex(logrus.InfoLevel),
		cmdr.WithWatchMainConfigFileToo(true),
	); err != nil {
		logrus.Errorf("Error: %v", err)
	}

}
