module github.com/hedzr/ini-op

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/crypto v0.0.0-20180904163835-0709b304e793 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/crypto v0.0.0-20181203042331-505ab145d0a9 => github.com/golang/crypto v0.0.0-20181203042331-505ab145d0a9
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180821023952-922f4815f713
	golang.org/x/sys v0.0.0-20180905080454-ebe1bf3edb33 => github.com/golang/sys v0.0.0-20181205085412-a5c9d58dba9a
	golang.org/x/sys v0.0.0-20181205085412-a5c9d58dba9a => github.com/golang/sys v0.0.0-20181205085412-a5c9d58dba9a
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
	golang.org/x/tools v0.0.0-20180828015842-6cd1fcedba52 => github.com/golang/tools v0.0.0-20180828015842-6cd1fcedba52

	google.golang.org/genproto v0.0.0-20180817151627-c66870c02cf8 => github.com/google/go-genproto v0.0.0-20180817151627-c66870c02cf8
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/coreos/go-systemd v0.0.0-20190212144455-93d5ec2c7f76
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-ini/ini v1.42.0
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/sirupsen/logrus v1.3.0
	github.com/smartystreets/goconvey v0.0.0-20190222223459-a17d461953aa // indirect
	github.com/spf13/cobra v0.0.3
	github.com/spf13/viper v1.3.1
	gopkg.in/ini.v1 v1.42.0 // indirect
)
