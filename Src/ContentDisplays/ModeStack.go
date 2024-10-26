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

// Returns a new stack identical to a given one
func (baseNode *ModeStackNode) Clone() ModeStackNode {
	var outp ModeStackNode = ModeStackNode{}
	var modeSlice = baseNode.Slice(baseNode.depth + 1)
	for i := range modeSlice {
		var mode int = modeSlice[len(modeSlice) - i - 1].mode
		if(i == 0) {
			outp.mode = mode
			continue
		}
		outp = outp.Push(mode)
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

// Returns the biggest difference between stacks 
// The first array represents the elements that have to be pulled from the first stack so that there is no difference since the deepest node
// The second array represents the elements that have to be pulled from the first stack so that there is no difference since the deepest node
// Returns (nil, nil) if there is no difference between stacks
func (baseNode *ModeStackNode) CalculateBiggestDifference (anotherNode ModeStackNode) ([]ModeStackNode, []ModeStackNode) {
	var len int = baseNode.depth 
	if(len > anotherNode.depth) {
		len = anotherNode.depth
	}
	len++
	var currentNodeSlice = baseNode.Slice(baseNode.depth + 1)
	var anotherNodeSlice = anotherNode.Slice(anotherNode.depth + 1)
	for i := range len {
		var currentNodeSliceIndex = baseNode.depth - i
		var anotherNodeSliceIndex = anotherNode.depth - i
		// fmt.Printf("%v - %v\n", currentNodeSlice[currentNodeSliceIndex].mode, anotherNodeSlice[anotherNodeSliceIndex].mode)
		if(currentNodeSlice[currentNodeSliceIndex].mode != anotherNodeSlice[anotherNodeSliceIndex].mode) {
			return baseNode.Slice(currentNodeSliceIndex + 1), anotherNode.Slice(anotherNodeSliceIndex + 1)
		}
	}
	// If the second stack is shorter
	if(baseNode.depth > anotherNode.depth) {
		return baseNode.Slice(baseNode.depth - anotherNode.depth), nil
	}
	// If the second stack is longer
	if(baseNode.depth < anotherNode.depth) {
		return nil, anotherNode.Slice(anotherNode.depth - baseNode.depth)
	}
	return nil, nil
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
	node1 = node1.Push(3)
	fmt.Printf("Are the stacks equal: %v\n", node.EqualsTo(node1))

	// Dif calculating test
	var dif, dif1 = node.CalculateBiggestDifference(node1)
	fmt.Print("Difference in 1:")
	TEMP_printSlice(dif)
	fmt.Print("Difference in 2:")
	TEMP_printSlice(dif1)

	// Clone test
	var clone = node.Clone()
	fmt.Println("The clones are identical:", node.EqualsTo(clone))

	fmt.Println("===========================")
}

func TEMP_printSlice (slice []ModeStackNode) {
	fmt.Print("[")
	for _, val := range slice {
		fmt.Printf("%v,", val.mode)
	}
	fmt.Println("]")
}
