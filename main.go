// Package main
// @Description: CLD项目main函数
package main

import (
	"fmt"
	"go_dlc_project/dao/mysql"
	"go_dlc_project/dao/redis"
	"go_dlc_project/logger"
	"go_dlc_project/setting"
	"os"
)

func main() {
	//启动项目时，需要再命令行加入参数读取配置文件
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: project config.yaml")
		return
	}

	// 加载配置(viper工具)
	if err := setting.Init(os.Args[1]); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	// 加载日志
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// 加载mysql
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	// 加载redis
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	fmt.Println("配置文件已加载，项目启动成功")
}
