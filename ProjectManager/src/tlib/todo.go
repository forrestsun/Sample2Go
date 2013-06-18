package tlib

//蕃茄定义
type TodoEntry struct {
	Todo_Id    string //唯一标识
	Todo_Date  string //日期
	Start_Time string //起始时间
	End_Time   string //截止时间
	Content    string //工作内容
	Project_Id int    //项目编号
	User_Id    int    //用户编号
}
