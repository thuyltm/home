package main

import (
	"fmt"
	"strings"
)

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
	bal   int // height(n.Right) - height(n.Left)
}

func (n *Node) Insert(value, data string) bool {
	switch {
	case value == n.Value:
		n.Data = data
		return false
	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			if n.Right == nil {
				n.bal = -1
			} else {
				n.bal = 0
			}
		} else {
			if n.Left.Insert(value, data) {
				if n.Left.bal < -1 || n.Left.bal > 1 {
					n.rebalance(n.Left)
				} else {
					n.bal--
				}
			}
		}
	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			if n.Left == nil {
				n.bal = 1
			} else {
				n.bal = 0
			}
		} else {
			if n.Right.Insert(value, data) {
				if n.Right.bal < -1 || n.Right.bal > 1 {
					n.rebalance(n.Right)
				} else {
					n.bal++
				}
			}
		}
	}
	if n.bal != 0 {
		return true
	}
	return false
}

func (n *Node) rotateLeft(c *Node) {
	fmt.Println("rotateLeft " + c.Value)
	r := c.Right
	c.Right = r.Left
	r.Left = c
	if c == n.Left {
		n.Left = r
	} else {
		n.Right = r
	}
	c.bal = 0
	r.bal = 0
}

func (n *Node) rotateRight(c *Node) {
	fmt.Println("rotateRight " + c.Value)
	l := c.Left
	c.Left = l.Right
	l.Right = c
	if c == n.Left {
		n.Left = l
	} else {
		n.Right = l
	}
	c.bal = 0
	l.bal = 0
}

// `rotateRightLeft` first rotates the right child of `c` to the right, then `c` to the left
func (n *Node) rotateRightLeft(c *Node) {
	c.Right.Left.bal = 1
	c.rotateRight(c.Right)
	c.Right.bal = 1
	n.rotateLeft(c)
}

// first rotates the left child of `c` to the left, then `c` to the right
func (n *Node) rotateLeftRight(c *Node) {
	c.Left.Right.bal = -1
	c.rotateLeft(c.Left)
	c.Left.bal = -1
	n.rotateRight(c)
}

func (n *Node) rebalance(c *Node) {
	fmt.Println("rebalance " + c.Value)
	c.Dump(0, "")
	switch {
	case c.bal == -2 && c.Left.bal == -1:
		n.rotateRight(c)
	case c.bal == 2 && c.Right.bal == 1:
		n.rotateLeft(c)
	case c.bal == -2 && c.Left.bal == 1:
		n.rotateLeftRight(c)
	case c.bal == 2 && c.Right.bal == -1:
		n.rotateRightLeft(c)
	}
}

func (n *Node) Find(s string) (string, bool) {
	if n == nil {
		return "", false
	}
	switch {
	case s == n.Value:
		return n.Data, true
	case s < n.Value:
		return n.Left.Find(s)
	default:
		return n.Right.Find(s)
	}
}

func (n *Node) Dump(i int, lr string) {
	if n == nil {
		return
	}
	indent := ""
	if i > 0 {
		indent = strings.Repeat(" ", (i-1)*4) + "+" + lr + "--"
	}
	fmt.Printf("%s%s[%d]\n", indent, n.Value, n.bal)
	n.Left.Dump(i+1, "L")
	n.Right.Dump(i+1, "R")
}
