/*
 * Copyright © 2019 Hedzr Yeh.
 */

package impl

import (
	"github.com/hedzr/log"
	"os"
	"path/filepath"
)

func appDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func appPath() string {
	path, err := filepath.Abs(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func appPath1() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func appName() string {
	name := filepath.Base(os.Args[0])
	return name
}

func writeTextFile(path string, text string) (err error) {
	fo, err := os.Create(path)
	if err != nil {
		return
	}
	// close fo on exit and check for its returned error
	defer func() {
		fo.Sync()
		if err = fo.Close(); err != nil {
			return
		}
	}()

	_, err = fo.WriteString(text)
	return
}
