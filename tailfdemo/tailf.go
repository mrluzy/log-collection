package tailfdemo

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func tailf() {
	filename := "./tail.log"
	cfg := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	// 打开文件开始读取数据
	tailf, err := tail.TailFile(filename, cfg)
	if err != nil {
		panic(err)
	}

	var (
		msg *tail.Line
		ok  bool
	)

	for {
		msg, ok = <-tailf.Lines
		if !ok {
			fmt.Println("tail closed")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("msg:%v\n", msg.Text)
	}

}
