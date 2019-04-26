package main

import (
	"fmt"
)

type QueueManager struct {
	queues []queue
}

func (pm *QueueManager) qlist() ([]string) {
	qlen := len(pm.queues)
	qlist := []string{}
	for i := 0; i < qlen; i++ {
		qlist = append(qlist[:i], pm.queues[i].name)
	}
	return qlist
}

type queue struct {
	name string
	datas []data
}

type data struct {
	name string
	data string
}

func main() {
	fmt.Println("start")

	d := data{name: "Programing language", data: "Python"}

	// 初期化
	q := queue{name: "", datas: []data{}}
	qm := QueueManager{queues: []queue{}}

	q.name = d.name
	q.datas = append(q.datas, d)

	qm.queues = append(qm.queues, q)

	d = data{name: "Programing language", data: "Ruby"}
	q.datas = append(q.datas, d)
	
	fmt.Println(q)
	fmt.Println(qm.queues)
	fmt.Println(qm.qlist())
}
