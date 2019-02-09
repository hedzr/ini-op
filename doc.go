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

	Version    = ""
	Buildstamp = ""
	Githash    = ""

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
