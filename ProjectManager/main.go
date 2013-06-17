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

type TasksManager struct {
	tms []Tomato
}

func (t *TasksManager) Add(tm *Tomato) {
	t.tms = append(t.tms, *tm)
}

func main() {
	ts := &TasksManager{make([]Tomato, 0)}
	t0 := &Tomato{"2010-6-6", "9:30", "10:00", "一个测试用的番茄", 1, 1}
	ts.Add(t0)

	for _, v := range ts.tms {
		fmt.Println("=====================================================")
		fmt.Println("Test tomato date is ", v.date)
		fmt.Println("Test tomato start time is ", v.start_time)
		fmt.Println("Test tomato end time is ", v.end_time)
		fmt.Println("Test tomato content is", v.content)
		fmt.Println("Test tomato project id is ", v.project_id)
	}
}
