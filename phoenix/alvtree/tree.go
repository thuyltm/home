package main

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(key int, value int) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree) Remove(key int) {
	t.root = t.root.remove(key)
}

func (t *AVLTree) Update(oldKey int, newKey int, newValue int) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey, newValue)
}

func (t *AVLTree) Search(key int) (node *AVLNode) {
	return t.root.search(key)
}

func (t *AVLTree) DispalyInOrder() {
	t.root.displayNodesInOrder()
}
