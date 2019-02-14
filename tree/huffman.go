package tree

import (
	"fmt"
)

type huffmanTreeNode struct {
	Weight int
	Left   *huffmanTreeNode
	Right  *huffmanTreeNode
	// Data interface{} here Data is useless
}

func (h *huffmanTreeNode) String() string {
	return fmt.Sprintf("(weight: %d, left)", h.Weight)
}

func (h *huffmanTreeNode) print() {
	printNode("", h, false)
}

func printNode(prefix string, node *huffmanTreeNode, isLeft bool) {
	if node == nil {
		return
	}

	if isLeft {
		fmt.Println(prefix+"|-- ", node.Weight)
		printNode(prefix+"|   ", node.Left, true)
		printNode(prefix+"|   ", node.Right, false)
		return
	}
	fmt.Println(prefix+"\\--", node.Weight)
	printNode(prefix+"    ", node.Left, true)
	printNode(prefix+"    ", node.Right, false)
}

func construtor(weights []int) *huffmanTreeNode {
	sets := make([]*huffmanTreeNode, len(weights))
	max := 0

	// create leaf nodes
	for i := 0; i < len(weights); i++ {
		max = max + weights[i]
		sets[i] = &huffmanTreeNode{
			Weight: weights[i],
			Left:   nil,
			Right:  nil,
		}
	}

	// build the tree
	for len(sets) != 1 {
		secMin, min := max, max
		secMinIdx, minIdx := -1, -1

		// fmt.Printf("%v\n", sets)
		// find two minest nodes
		for idx, node := range sets {
			if node.Weight < min {
				// println("hhhh", max, node.Weight, min)
				secMin = min
				secMinIdx = minIdx

				min = node.Weight
				minIdx = idx
				continue
			}

			if node.Weight == min {
				secMin = node.Weight
				secMinIdx = idx
				continue
			}

			if node.Weight > min && secMinIdx == -1 {
				secMin = node.Weight
				secMinIdx = idx
			}
		}

		// println(minIdx, secMinIdx, len(sets))

		newNode := &huffmanTreeNode{
			Weight: secMin + min,
			Left:   sets[minIdx],
			Right:  sets[secMinIdx],
		}

		// remove old nodes
		var (
			biggerIdx  = secMinIdx
			smallerIdx = minIdx
		)
		if minIdx > secMinIdx {
			biggerIdx = minIdx
			smallerIdx = secMinIdx
		}

		sets = append(sets[:biggerIdx], sets[biggerIdx+1:]...)
		sets = append(sets[:smallerIdx], sets[smallerIdx+1:]...)

		// add new nodes
		sets = append(sets, newNode)
	}

	return sets[0]
}
