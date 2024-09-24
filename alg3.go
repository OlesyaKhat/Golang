package main

import (
 "container/heap"
 "fmt"
 "sort"
)

type Task struct {
 name     string
 priority int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int {
 return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
 return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
 pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
 task := x.(*Task)
 *pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() interface{} {
 old := *pq
 n := len(old)
 task := old[n-1]
 *pq = old[0 : n-1]
 return task
}

type TaskManager struct {
 pq       PriorityQueue
 taskDict map[string]int
}

func NewTaskManager() *TaskManager {
 return &TaskManager{
  taskDict: make(map[string]int),
 }
}

func (tm *TaskManager) AddTask(name string, priority int) {
 task := &Task{name: name, priority: priority}
 heap.Push(&tm.pq, task)
 tm.taskDict[name] = priority
}

func (tm *TaskManager) SearchByName(name string) (int, bool) {
 priority, found := tm.taskDict[name]
 return priority, found
}

func (tm *TaskManager) SearchByPriority(priority int) (string, bool) {
 for _, task := range tm.pq {
  if task.priority == priority {
   return task.name, true
  }
 }
 return "", false
}

func main() {
 tm := NewTaskManager()
 tasks := []struct {
  name     string
  priority int
 }{
  {"реализация интерфейса", 3},
  {"реализация функции аутентификации", 1},
  {"реализация запросов к БД", 3},
  {"подключение к БД", 2},
  {"тестирование программы", 5},
  {"реализация основной функции", 0},
  {"ревью кода", 4},
 }

 for _, task := range tasks {
  tm.AddTask(task.name, task.priority)
 }

 
 fmt.Println("Задачи по приоритету :")
 for tm.pq.Len() > 0 {
  task := heap.Pop(&tm.pq).(*Task)
  fmt.Printf(" %d: %s\n", task.priority, task.name)
 }


 sort.Sort(tm.pq)
 fmt.Println("Задачи по приоритету :")
 for _, task := range tm.pq {
  fmt.Printf(" %d: %s\n", task.priority, task.name)
 }

 taskName := "реализация функции аутентификации"
 if priority, found := tm.SearchByName(taskName); found {
  fmt.Printf("Приоритет задачи '%s': %d\n", taskName, priority)
 } else {
  fmt.Printf("Задача '%s' не найдена\n", taskName)
 }

 priority := 3
 if task, found := tm.SearchByPriority(priority); found {
  fmt.Printf("Задача с приоритетом %d: %s\n", priority, task)
 } else {
  fmt.Printf("Задача с приоритетом %d не найдена\n", priority)
 }
}