package initconfig

import (
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log-collection/global"
)

func Tail() {
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	filename := viper.GetString("collect.logfile_path")
	tailf, err := tail.TailFile(filename, cfg)
	if err != nil {
		panic(err)
	}
	global.Tailf = tailf
	logrus.Info("tail init success")
}
