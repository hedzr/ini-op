/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/hedzr/cmdr"
	"github.com/hedzr/log"
	"os"
)

//
//
//

//
//
//

func entryGet(cmd *cmdr.Command, args []string) (err error) {
	if len(args) == 3 {
		var cfg *ini.File
		var fn = os.ExpandEnv(args[2])
		cfg, err = ini.Load(fn)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
			os.Exit(1)
		}
		log.Debugf("loading section '%s' from '%s'", args[0], args[1])
		section := cfg.Section(args[0])
		// for i, k := range section.Keys() {
		// 	fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
		// }
		fmt.Println(section.Key(args[1]))
	} else {
		cmd.PrintHelp(false)
	}
	return
}

func entryPut(cmd *cmdr.Command, args []string) (err error) {
	if len(args) == 4 {
		var cfg *ini.File
		var fn = os.ExpandEnv(args[3])
		cfg, err = ini.Load(fn)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
			os.Exit(1)
		}
		log.Debugf("loading section '%s' from '%s'", args[0], args[1])
		section := cfg.Section(args[0])
		// for i, k := range section.Keys() {
		//	fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
		// }

		// fmt.Println(section.Key(args[1]))
		section.Key(args[1]).SetValue(args[2])
		log.Debugf("set section '%s.%s' to '%s'", args[0], args[1], args[2])

		if err := cfg.SaveTo(fn); err != nil {
			log.Fatalf("Fail to write file: %v", err)
			os.Exit(1)
		}

	} else {
		cmd.PrintHelp(false)
	}
	return
}

func entryRemove(cmd *cmdr.Command, args []string) (err error) {
	if len(args) == 3 {
		var cfg *ini.File
		var fn = os.ExpandEnv(args[2])
		// read ini file
		cfg, err := ini.Load(fn)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
			os.Exit(1)
		}

		cfg.Section(args[0]).DeleteKey(args[1])

		if err := cfg.SaveTo(fn); err != nil {
			log.Fatalf("Fail to write file: %v", err)
			os.Exit(1)
		}

		//} else if len(args) == 1 {
		//	// read from stdin
	} else {
		cmd.PrintHelp(false)
	}
	return
}

//
//
//

func sectionGet(cmd *cmdr.Command, args []string) (err error) {
	if len(args) == 2 {
		var cfg *ini.File
		var fn = os.ExpandEnv(args[1])
		cfg, err = ini.Load(fn)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
			os.Exit(1)
		}
		log.Debugf("loading section '%s' from '%s'", args[0], args[1])
		section := cfg.Section(args[0])
		for i, k := range section.Keys() {
			fmt.Printf("%s\t= %s  ; %d\n", k.Name(), section.Key(k.Name()), i)
		}
	} else {
		cmd.PrintHelp(false)
	}
	return
}

func sectionRemove(cmd *cmdr.Command, args []string) (err error) {
	if len(args) == 2 {
		var cfg *ini.File
		var fn = os.ExpandEnv(args[1])
		// delete section from ini file
		cfg, err := ini.Load(fn)
		if err != nil {
			log.Fatalf("Fail to read file: %v", err)
			os.Exit(1)
		}

		cfg.DeleteSection(args[0])

		if err := cfg.SaveTo(fn); err != nil {
			log.Fatalf("Fail to write file: %v", err)
			os.Exit(1)
		}

		// } else if len(args) == 1 {
		// read from stdin
	} else {
		cmd.PrintHelp(false)
	}
	return
}
