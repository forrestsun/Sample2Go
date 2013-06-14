package main

import (
	"fmt"
)

//蕃茄定义
type Tomato struct {
	date       string //日期
	start_time string //起始时间
	end_time   string //截止时间
	content    string //工作内容
	project_id int    //项目编号
	user_id    int    //用户编号
}

func AddTomato(list []Tomato, value Tomato) []Tomato {
	list = append(list, value)
	return list
}

func main() {
	todo_list := []Tomato{}

	list_count := 1
	for list_count < 10 {
		todo_list = AddTomato(todo_list, Tomato{"2010-6-6", "9:30", "10:00", "一个测试用的番茄", list_count, 1})
		list_count++
	}

	for _, v := range todo_list {
		fmt.Println("=====================================================")
		fmt.Println("Test tomato date is ", v.date)
		fmt.Println("Test tomato start time is ", v.start_time)		fmt.Println("Test tomato end time is ", v.end_time)
		fmt.Println("Test tomato content is", v.content)
		fmt.Println("Test tomato project id is ", v.project_id)
	}
}
