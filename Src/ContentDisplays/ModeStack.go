package parser

import (
	"fmt"
)

type ModeStackNode struct {
	mode int
	depth int
	nextNode *ModeStackNode
}

// Creates and returnes the next node in the stack. 
// New node has to be assigned explicitly
func (baseNode ModeStackNode) Push(mode int) ModeStackNode{
	var newNode ModeStackNode
	newNode.mode = mode
	newNode.depth = baseNode.depth + 1
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
	// baseNode.depth = baseNode.depth - 1
	return *baseNode
}

// Returns an array of modes in the stack starting from the top one
func (baseNode *ModeStackNode) Slice(depth int) []ModeStackNode {
	// We add 1 since depth is zero-based iterator
	if(depth > baseNode.depth + 1) {
		depth = baseNode.depth + 1
	}
	var outp = make([]ModeStackNode, depth)
	var floartingReference *ModeStackNode
	for i := range depth {
		if(floartingReference == nil) {
			floartingReference = baseNode
		} else {
			floartingReference = floartingReference.nextNode
		}
		outp[i] = *floartingReference
	}
	return outp
}

// Returns true if the content of a stack is equal
func (baseNode *ModeStackNode) EqualsTo(anotherNode ModeStackNode, checkDepth ...int) bool {
	// The simplest ones to check
	if(anotherNode.depth != baseNode.depth || baseNode.mode != anotherNode.mode) {
		return false
	}
	// At this point we are sure that these arrays are the same depth
	var currentNodeSlice = baseNode.Slice(baseNode.depth + 1)
	var anotherNodeSlice = anotherNode.Slice(anotherNode.depth + 1)
	// Checking equality of each element
	for i := range currentNodeSlice {
		if(currentNodeSlice[i].mode != anotherNodeSlice[i].mode) {
			return false
		}
	}
	return true
}

// Returns the biggest difference between stacks as a slice
func (baseNode *ModeStackNode) CalculateBiggestDifference (anotherNode ModeStackNode) []ModeStackNode {
	// TODO
	return nil
}

func (baseNode ModeStackNode) String() string {
	return fmt.Sprintf("[value = %v depth = %v]", baseNode.mode, baseNode.depth)
}

func Test() {
	var node ModeStackNode
	node.mode = 1
	node = node.Push(2)
	node = node.Push(3)
	node = node.Push(4)

	// Slices test
	var slice = node.Slice(6)
	TEMP_printSlice(slice)

	// Comparison test
	var node1 ModeStackNode
	node1.mode = 1
	node1 = node1.Push(2)
	node1 = node1.Push(3)
	node1 = node1.Push(4)
	fmt.Printf("Are the stacks equal: %v\n", node.EqualsTo(node1))
}

func TEMP_printSlice (slice []ModeStackNode) {
	fmt.Print("[")
	for _, val := range slice {
		fmt.Printf("%v,", val.mode)
	}
	fmt.Println("]")
}
