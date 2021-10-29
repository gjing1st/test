// Created by dolitTeam
//@Author : GJing
//@Time : 2020/10/23 11:46
//@File : log
package log

import (
	//"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"strings"
	"time"
)

// 用于应用初始化。
func initLog() {
	timeString := time.Now()
	path :=  "./" + timeString.Format("2006-01") + "/" + timeString.Format("2006-01-02")
	//设置日志路径，自动创建目录
	glog.SetPath(path)
	//开启异步日志记录
	glog.SetAsync(true)
	//关闭控制台输出
	glog.SetStdoutPrint(false)
}

// LogFile
// @description: 日志记录到文件
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2020/10/23 11:35
// @return:
func LogFile(fileName string, v ...interface{}) {
	initLog()
	//使用对象设置方法，高并发时容易导致日志写入其他文件，改用链式操作
	//glog.SetFile(fileName)
	pathArr := strings.Split(fileName, "/")
	//字符串切割，如果是路由则取第一个路径为文件名
	if len(pathArr) > 1 {
		fileName = pathArr[1]
	}
	//如果文件名为空，则默认使用common
	if len(fileName) == 0 {
		fileName = "common"
	}

	//使用回溯值记录调用日志文件名和行号
	glog.Skip(1).Line(true).File(fileName).Println(v)
}

// LogInfo
// @description:
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2021/9/3 10:22
func LogInfo(fileName string, v ...interface{}) {
	//使用协程去记录日志
	//TODO 需验证，是否是否关闭，协程会不会文件流冲突
	go LogFile(fileName, v)
}
