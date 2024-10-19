package parser

import (
	"fmt"
)

type ModeStackNode struct {
	mode int
	nextNode *ModeStackNode
}

// Creates and returnes the next node in the stack. 
// New node has to be assigned explicitly
func (baseNode ModeStackNode) Push(mode int) ModeStackNode{
	var newNode ModeStackNode
	newNode.mode = mode
	newNode.nextNode = &baseNode
	return newNode
}

// Pulls a node from a stack and returns it. 
// If there are no more nodes in stack the base none becames a default value
func (baseNode *ModeStackNode) Pull() ModeStackNode {
	// Null pointer handling
	if(baseNode.nextNode == nil) {
		*baseNode = ModeStackNode{}
		return *baseNode
	}
	*baseNode = *baseNode.nextNode
	return *baseNode
}

func Test() {
	var node ModeStackNode
	node.mode = 11
	node = node.Push(12)
	// Println calls functions 
	var val int = node.mode
	fmt.Println(val, " -> ", node.Pull().mode)
}
