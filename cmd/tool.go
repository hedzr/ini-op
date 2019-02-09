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
	log "github.com/sirupsen/logrus"
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
