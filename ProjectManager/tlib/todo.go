package tlib

//蕃茄定义
type TodoEntry struct {
	todo_id    string //唯一标识
	date       string //日期
	start_time string //起始时间
	end_time   string //截止时间
	content    string //工作内容
	project_id int    //项目编号
	user_id    int    //用户编号
}
