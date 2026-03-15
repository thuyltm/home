package main

type Tree struct {
	Root *Node
}

func (t *Tree) Insert(value, data string) {
	if t.Root == nil {
		t.Root = &Node{Value: value, Data: data}
		return
	}
	t.Root.Insert(value, data)
	if t.Root.bal < -1 || t.Root.bal > 1 {
		t.rebalance()
	}
}

func (t *Tree) rebalance() {
	fakeParent := &Node{Left: t.Root, Value: "fakeParent"}
	fakeParent.rebalance(t.Root)
	t.Root = fakeParent.Left
}

func (t *Tree) Find(s string) (string, bool) {
	if t.Root == nil {
		return "", false
	}
	return t.Root.Find(s)
}

func (t *Tree) Traverse(n *Node, f func(*Node)) {
	if n == nil {
		return
	}
	t.Traverse(n.Left, f)
	f(n)
	t.Traverse(n.Right, f)
}

func (t *Tree) Dump() {
	t.Root.Dump(0, "")
}
