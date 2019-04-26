package main

import (
	"fmt"
)

type QueueManager struct {
	queues *[]queue
}

type queue struct {
	name string
	datas *[]string
}

func main() {
	// 初期化
	dlist := []string{}
	qname := "Programing_language"
	q := queue{name: qname, datas: &dlist}
	qlist := []queue{q}
	qm := QueueManager{queues: &qlist}

	// データを格納
	d := "Python"
	dlist = append(dlist, d)

	d = "Ruby"
	dlist = append(dlist, d)

	// キューを表示
	fmt.Println(q)
	fmt.Println((*qm.queues)[0])
	fmt.Println("Queue name:", q.name)
	fmt.Println("Data:", *(*qm.queues)[0].datas)
}
