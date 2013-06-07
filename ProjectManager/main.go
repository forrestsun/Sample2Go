package main

import (
	"fmt"
)

//项目基本信息定义
type Project_Stand_Info struct {
	project_id         string //项目编号
	project_name       string //项目名称
	project_start_date string //项目启动日期
	project_end_date   string //项目截止日期
	project_content    string //项目描述
	project_url        string //项目发布网址
	project_logo       string //项目LOGO
	project_version    string //项目最新版本号
}

type Opt_System_Info struct {
	opt_system_name    string //操作系统名称 2k3 2k8..
	opt_system_version string //2003 2008..
}

type Server_Info struct {
	Opt_System_Info        //操作系统
	server_name     string //服务器名称
	server_ip       string //服务器IP
	server_port     string //服务器端口
}

type DB_Server_Info struct {
	Opt_System_Info        //操作系统
	Server_Info            //服务器信息
	db_type         string //数据库类型 sqlserver mysql ..
	db_version      string //数据库版本 2000 2005 5.0..
}

type Web_Server_Info struct {
	Server_Info        //服务器信息
	web_type    string //iis apache ngix..
	web_version string //6.0 7.0 ..
}

type Project_Detail_Info struct {
	Project_Stand_Info //项目基本信息
	Web_Server_Info    //web服务器信息
	DB_Server_Info     //DB服务器信息
}

//蕃茄定义
type Tomato struct {
	date       string //日期
	start_time string //起始时间
	end_time   string //截止时间
	content    string //工作内容
	project_id int    //项目编号
	user_id    int    //用户编号
}

func main() {

	//初始化一个番茄
	tomato := Tomato{"2010-6-6", "9:30", "10:00", "一个测试用的番茄", 1,1}
	//显示这个番茄的内容
	fmt.Println("Test tomato date is ", tomato.date)
	fmt.Println("Test tomato start time is ", tomato.start_time)
	fmt.Println("Test tomato end time is ", tomato.end_time)
	fmt.Println("Test tomato content is", tomato.content)
	fmt.Println("Test tomato project id is ", tomato.project_id)

	todo_list := make(map[int]Tomato)
	list_count := 1
	for list_count < 10 {
		todo_list[list_count] = Tomato{"2010-6-6", "9:30", "10:00", "一个测试用的番茄", list_count,1}
		list_count++
	}

	for _, v := range todo_list {
		fmt.Println(v.project_id)
	}
}
