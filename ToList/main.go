package main

import "fmt"

// 添加任务
func addList(tasks []Task, title string) []Task {
	return append(tasks, Task{Title: title, Done: false})
}

// 打印任务列表
func ListTasks(tasks []Task) {
	for i, task := range tasks {
		status := "未完成"
		if task.Done {
			status = "完成"
		}
		fmt.Println(i, task.Title, status)
	}
}

// 定义任务结构体
type Task struct {
	Title string //名字
	Done  bool   //完成状态
	Time  string //完成时间
}

// 标记任务完成
func markDone(tasks []Task, index int) []Task {
	if index < 0 || index >= len(tasks) {
		fmt.Print("索引超出范围值")
		return tasks
	}
	tasks[index].Done = true
	tasks[index].Time = "20260506"
	return tasks
}

func main() {
	tasks := []Task{}

	tasks = addList(tasks, "吃饭")
	tasks = addList(tasks, "睡觉")
	tasks = addList(tasks, "写代码")
	ListTasks(tasks)
	tasks = markDone(tasks, 1)
}
