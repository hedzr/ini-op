# ini-op

[![Build Status](https://travis-ci.org/hedzr/ini-op.svg?branch=master)](https://travis-ci.org/hedzr/ini-op)

Read/Write inifile.

```bash
$ ini-op
Usage:
  ini-op [flags]
  ini-op [command]

Available Commands:
  entry       get/put entry
  help        Help about any command
  section     get/put section
  version     Print the version.

Flags:
      --config string   config file (default is $HOME/./*.yaml)
  -d, --debug           debug mode
  -h, --help            help for ini-op
  -v, --verbose         verbose output
  -V, --version         Show versions

Use "ini-op [command] --help" for more information about a command.

```

## Commands

### command `section`

```bash
$ bin/ini-op s

Usage:
  ini-op section [command]

Aliases:
  section, s, sec

Available Commands:
  get         get section
  rm          delete section

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
$ bin/ini-op e

Usage:
  ini-op entry [command]

Aliases:
  entry, e

Available Commands:
  get         get an entry
  put         put string entry
  rm          delete an entry

```

Usages:

```bash
# print section `server`, entry `port`
$ ini-op e get server port $HOME/abc.ini
$ ini-op e put server port 1313 $HOME/abc.ini
$ ini-op e rm server port $HOME/abc.ini
```

## Contributions

welcome

## LICENSE

MIT

