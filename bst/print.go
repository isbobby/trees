package bst

import (
	"fmt"
	"math"
)

func (root *Node) PrintNodeInfo() {
	if root == nil {
		fmt.Println("Node is NIL")
		return
	}

	nodeInfo := NodeInfo{
		Val:     root.val,
		Deleted: root.deleted,
	}

	if root.right != nil {
		nodeInfo.Right = root.right
	}

	if root.left != nil {
		nodeInfo.Left = root.left
	}

	if root.parent != nil {
		nodeInfo.Parent = root.parent
	}
	fmt.Println(nodeInfo)
}

func (root *Node) PrintSubtreeLevelOrder() {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}

	// whitespaces := []byte(" ")
	nilString := []byte("_")

	elements := [][][]byte{}
	queue := []*Node{root}
	level := 1
	elementLength := 1
	hasNode := true
	for len(queue) > 0 && hasNode {
		i := 0
		currLevelElements := [][]byte{}
		hasNode = false
		for i < level && len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			currElement := []byte(top.String())
			currLevelElements = append(currLevelElements, currElement)
			if len(currElement) > elementLength {
				elementLength = len(currElement)
			}

			if top != nil {
				hasNode = true
				queue = append(queue, top.left, top.right)
			} else {
				queue = append(queue, nil, nil)
			}
			i++
		}

		for i < level {
			currLevelElements = append(currLevelElements, nilString)
			i++
		}

		if hasNode {
			level = level << 1
			elements = append(elements, currLevelElements)
		}
	}

	fmt.Println("Root:", root.val, "Size:", root.Size(), "Level:", math.Log2(float64(level)))
	leadingSpace := 0
	spacedElements := [][]byte{}
	for i := len(elements) - 1; i >= 0; i-- {
		k := len(elements) - i
		gapBetweenNodes := (int(math.Pow(2.0, float64(k))) - 1) * elementLength

		spacedElementsAtLevel := []byte{}

		// append leading space
		for j := 0; j < leadingSpace; j++ {
			spacedElementsAtLevel = append(spacedElementsAtLevel, ' ')
		}

		// create gap
		gap := []byte{}
		for j := 0; j < gapBetweenNodes; j++ {
			gap = append(gap, ' ')
		}

		elementsAtLevel := elements[i]
		for _, element := range elementsAtLevel {
			spacedElementsAtLevel = append(spacedElementsAtLevel, fillElement(element, elementLength)...)
			spacedElementsAtLevel = append(spacedElementsAtLevel, gap...)
		}

		spacedElements = append(spacedElements, spacedElementsAtLevel)
		leadingSpace = gapBetweenNodes
	}

	for i := len(spacedElements) - 1; i >= 0; i-- {
		fmt.Println(string(spacedElements[i]))
	}

	line := []byte{}
	for i := 0; i < len(spacedElements[0]); i++ {
		line = append(line, '=')
	}
	fmt.Println(string(line))
}

func fillElement(element []byte, requiredLength int) []byte {
	if len(element) >= requiredLength {
		return element
	}

	diff := requiredLength - len(element)

	half := float64(diff) / 2
	left, right := math.Ceil(half), math.Floor(half)
	leadingSpace, endingSpace := []byte{}, []byte{}
	for i := 0; i < int(left); i++ {
		leadingSpace = append(leadingSpace, ' ')
	}

	for i := 0; i < int(right); i++ {
		endingSpace = append(endingSpace, ' ')
	}

	spacedElement := leadingSpace
	spacedElement = append(leadingSpace, element...)
	spacedElement = append(spacedElement, endingSpace...)

	return spacedElement
}
