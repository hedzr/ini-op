/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"github.com/hedzr/cmdr"
	"github.com/hedzr/ini-op"
)

var (
	rootCmd = &cmdr.RootCommand{
		Command: cmdr.Command{
			BaseOpt: cmdr.BaseOpt{
				Name: ini_op.APPNAME,
			},
			Flags: []*cmdr.Flag{
				// global options here.
			},
			PreAction:  Pre,
			PostAction: Post,
			SubCommands: []*cmdr.Command{
				// dnsCommands,
				// playCommand,
				// generatorCommands,
				// serverCommands,
				entryCommands,
				sectionCommands,
			},
		},

		AppName:    ini_op.APPNAME,
		Version:    ini_op.Version,
		VersionInt: ini_op.VersionInt,
		Copyright:  "ini-op is an effective tool for read/write inifile",
		Author:     "Hedzr Yeh <hedzrz@gmail.com>",
	}

	entryCommands = &cmdr.Command{
		BaseOpt: cmdr.BaseOpt{
			Name:  "Entry",
			Short: "e",
			Full:  "entry",
			// Aliases:     []string{"serve", "svr",},
			Description: "get/put entry.",
		},
		SubCommands: []*cmdr.Command{
			{
				BaseOpt: cmdr.BaseOpt{
					Short:       "g",
					Full:        "get",
					Aliases:     []string{"rd", "read"},
					Description: "get an entry.",
					Action:      entryGet,
				},
				// PreAction: impl.ServerStartPre,
				// PostAction: impl.ServerStartPost,
			},
			{
				BaseOpt: cmdr.BaseOpt{
					Short: "p",
					Full:  "put",
					// Aliases:     []string{"stp", "halt", "pause",},
					Description: "put value to an entry.",
					Action:      entryPut,
				},
			},
			{
				BaseOpt: cmdr.BaseOpt{
					Short:       "r",
					Full:        "rm",
					Aliases:     []string{"remove", "del", "erase", "delete"},
					Description: "remove an entry.",
					Action:      entryRemove,
				},
			},
		},
	}

	sectionCommands = &cmdr.Command{
		BaseOpt: cmdr.BaseOpt{
			Name:        "Section",
			Short:       "s",
			Full:        "section",
			Aliases:     []string{"sec"},
			Description: "get/put section",
		},
		SubCommands: []*cmdr.Command{
			{
				BaseOpt: cmdr.BaseOpt{
					Short:       "g",
					Full:        "get",
					Aliases:     []string{"rd", "read"},
					Description: "get a section",
					Action:      sectionGet,
				},
			},
			{
				BaseOpt: cmdr.BaseOpt{
					Short:       "r",
					Full:        "rm",
					Aliases:     []string{"remove", "del", "erase", "delete"},
					Description: "remove a scrtion",
					Action:      sectionRemove,
				},
			},
		},
	}
)
