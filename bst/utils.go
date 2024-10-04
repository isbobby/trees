package bst

import (
	"fmt"
	"math"
)

func (root *Node) PrintSubtreeLevelOrder() {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}

	whitespaces := []byte(" ")
	nilString := []byte("_")

	res := []string{}
	queue := []*Node{root}
	level := 1

	for len(queue) > 0 {
		i := 0
		currLevel := []byte{}
		for i < level && len(queue) > 0 {
			top := queue[0]
			queue = queue[1:]
			currLevel = append(currLevel, []byte(top.String())...)
			currLevel = append(currLevel, whitespaces...)
			i++

			if top != nil {
				queue = append(queue, top.left)
				queue = append(queue, top.right)
			}
		}

		for i < level {
			currLevel = append(currLevel, nilString...)
			currLevel = append(currLevel, whitespaces...)
			i++
		}

		level = level << 1
		res = append(res, string(currLevel))
	}

	fmt.Println("Root:", root.val, "Size:", root.Size(), "Level:", math.Log2(float64(level)))
	for i := range res {
		toPrint := []byte{}
		for j := 0; j < (level-len(res[i]))/2; j++ {
			toPrint = append(toPrint, []byte(" ")...)
		}
		toPrint = append(toPrint, res[i]...)
		fmt.Println(string(toPrint))
	}
	line := []byte{}
	for i := 0; i < level; i++ {
		line = append(line, '=')
	}
	fmt.Println(string(line))
}
