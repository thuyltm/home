package main

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	values := []string{"d", "b", "g", "g", "c", "e", "a", "h", "f", "i", "j", "l", "k"}
	data := []string{"delta", "bravo", "golang", "golf", "charlie", "echo", "alpha", "hotel", "foxtrot", "india", "juliett", "lima", "kilo"}
	tree := &Tree{}
	for i := 0; i < len(values); i++ {
		fmt.Println("Insert " + values[i] + ": " + data[i])
		tree.Insert(values[i], data[i])
		tree.Dump()
		fmt.Println()
	}
	fmt.Print("Sorted values: | ")
	tree.Traverse(tree.Root, func(n *Node) {
		fmt.Print(n.Value, ": ", n.Data, " | ")
	})
	fmt.Println()
}
