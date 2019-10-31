package main

import (
	"fmt"

	"github.com/student/1031/conf"
	"gopkg.in/ini.v1"
)

var (
	cfg = new(conf.AppConf)
)

func main() {
	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg.EdctConf.Address)
}
