package main

type jobItem struct {
	// id
	Id int
	// company
	CompanyName    string
	CompanyArea    string
	CompanySize    string
	CompanyUrl     string
	CompanyStage   string
	CompanyAddress string
	// job
	JobName   string
	JobSalary string
	JobUpTime string
	JobExp    string
	JobDegree string
	JobDesc   string
	//hr
	HrName string
	// 7天内处理完成的简历所占比例
	HrPercent string
	// 完成简历处理的平均用时
	HrExecTime string
	LagouURL   string
}
