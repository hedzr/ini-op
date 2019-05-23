# ini-op

[![Build Status](https://travis-ci.org/hedzr/ini-op.svg?branch=master)](https://travis-ci.org/hedzr/ini-op)
[![Go Report Card](https://goreportcard.com/badge/github.com/hedzr/ini-op)](https://goreportcard.com/report/github.com/hedzr/ini-op)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/hedzr/ini-op.svg?label=release)

Read/Write inifile.



```bash
$ bin/ini-op
ini-op is an effective tool for read/write inifile by Hedzr Yeh <hedzrz@gmail.com> - v0.2.3

Usages:
    ini-op [Commands]  [Options] [Parent/Global Options]

Commands:
  e, entry                                        get/put entry.
  s, section, sec                                 get/put section
  [Misc]
  g, generate, gen                                generators for this app.
  version, ver                                    Show the version of this app.

Options:
  [Misc]
       --config=[Location of config file]         load config files from where you specified
  -q,  --quiet                                    No more screen output. (default=false)
  -v,  --verbose, --vv, --vvv                     Show this help screen (default=false)
  -V,  --version                                  Show the version of this app. (default=false)

Type '-h' or '--help' to get command help screen.
```

## Commands

### command `section`

```bash
$ bin/ini-op section
ini-op is an effective tool for read/write inifile by Hedzr Yeh <hedzrz@gmail.com> - v0.2.3

Usages:
    ini-op section [Sub-Commands]  [Options] [Parent/Global Options]

Description:
    get/put section

Sub-Commands:
  g, get, rd, read                                get a section
  r, rm, remove, del, erase, delete               remove a scrtion

Global Options:
  [Misc]
       --config=[Location of config file]         load config files from where you specified
  -q,  --quiet                                    No more screen output. (default=false)
  -v,  --verbose, --vv, --vvv                     Show this help screen (default=false)
  -V,  --version                                  Show the version of this app. (default=false)

Type '-h' or '--help' to get command help screen.
```

Usages:

```bash
# print `server` section
$ ini-op s get server $HOME/abc.ini
# delete `server` section
$ ini-op s rm server $HOME/abc.ini

```

### command `entry`

```bash
$ bin/ini-op entry
ini-op is an effective tool for read/write inifile by Hedzr Yeh <hedzrz@gmail.com> - v0.2.3

Usages:
    ini-op entry [Sub-Commands]  [Options] [Parent/Global Options]

Description:
    get/put entry.

Sub-Commands:
  g, get, rd, read                                get an entry.
  p, put                                          put value to an entry.
  r, rm, remove, del, erase, delete               remove an entry.

Global Options:
  [Misc]
       --config=[Location of config file]         load config files from where you specified
  -q,  --quiet                                    No more screen output. (default=false)
  -v,  --verbose, --vv, --vvv                     Show this help screen (default=false)
  -V,  --version                                  Show the version of this app. (default=false)

Type '-h' or '--help' to get command help screen.
```

Usages:

```bash
# print section `server`, entry `port`
$ ini-op e get server port $HOME/abc.ini
$ ini-op e put server port 1313 $HOME/abc.ini
$ ini-op e rm server port $HOME/abc.ini
```



## **TODO**

- [x] rewrite with new `cmdr`;
- [ ] add `yaml`, `json`, and `toml` supports;
- [ ] clean codes



## Contrib

Feel free to issue.



## LICENSE

MIT

