package linklist
 
type DataNode struct {
	Cmd string
	Desc string
	Handler func() int
	Next *DataNode
};