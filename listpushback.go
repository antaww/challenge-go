package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := NodeL{}
	n.Data = ""
	m := NodeL{}
	m.Data = ""
	n.Next = &m
}
