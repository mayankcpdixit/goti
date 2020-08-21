package mutil

import (
	"fmt"
)

/*
  #################################################
  Maths
	#################################################
*/
func Min(arr ...int) int {
	var min int
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func Max(arr ...int) int {
	var max int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func Pow(num, times int) int {
	if times == 0 {
		return num
	}
	n := 1
	for i := 0; i < times; i++ {
		n *= num
	}
	return n
}

/*
  #################################################
  Array
	#################################################
*/
func ArrayMin(arr []int) int {
	var min int
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

func ArrayMax(arr []int) int {
	var max int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func GetEmpty2DArray(iLimit, jLimit int) [][]int {
	result := [][]int{}
	for i := 0; i < iLimit; i++ {
		result = append(result, make([]int, jLimit))
	}
	return result
}

/*
  #################################################
  Linked List
	#################################################
*/
type LListIntNode struct {
	Val  int
	Next *LListIntNode
}

func GetLListInt(limit int) *LListIntNode {
	node := &LListIntNode{Val: 1}
	head := node
	for i := 2; i <= limit; i++ {
		node.Next = &LListIntNode{Val: i}
		node = node.Next
	}
	return head
}

func PrintLListInt(node *LListIntNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}

/*
  #################################################
	Queue
	#################################################
*/
type Queue struct {
	List []interface{}
}

func InitQueue() *Queue {
	return &Queue{}
}
func ArrayToQueue(arr []interface{}) *Queue {
	q := InitQueue()
	for i := 0; i < len(arr); i++ {
		q.Add(arr[i])
	}
	return q
}
func (q *Queue) Add(val interface{}) {
	q.List = append(q.List, val)
}
func (q *Queue) Size() int {
	return len(q.List)
}
func (q *Queue) Empty() bool {
	return len(q.List) == 0
}
func (q *Queue) Get() interface{} {
	if q.Empty() {
		return nil
	}
	val := q.List[0]
	q.List = q.List[1:]
	return val
}

/*
  #################################################
  Tree
	#################################################
*/

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func GetTreeNode() *TreeNode {
	return &TreeNode{}
}
func ArrayToTree(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	q := InitQueue()
	head := &TreeNode{Val: values[0]}
	q.Add(head)

	for i := 1; i < len(values); {
		node := q.Get().(*TreeNode)
		if node == nil {
			continue
		}

		if values[i] != nil {
			node.Left = &TreeNode{Val: values[i]}
			q.Add(node.Left)
		}
		i++

		if i < len(values) && values[i] != nil {
			node.Right = &TreeNode{Val: values[i]}
			q.Add(node.Right)
		}
		i++
	}
	return head
}
