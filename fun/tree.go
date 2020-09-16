package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

/*
 * Tree types and functions
 */
type Tree struct {
	Left *Tree
	Value int
	Right *Tree
}

const treeSize int = 10

type ErrDuplicateInTree int

func (e ErrDuplicateInTree) Error() string {
	return fmt.Sprintf("cannot add duplicate value to tree: %v already in tree", int(e))
}

func (t *Tree) Insert(v int) error {
	if t.Value == 0 {
		t.Value = v
		return nil
	} else if v < t.Value {
		if(t.Left == nil) {
			t.Left = &Tree{nil, v, nil}
			return nil
		} else {
			return t.Left.Insert(v)
		}
	} else if v > t.Value {
		if(t.Right == nil) {
			t.Right = &Tree{nil, v, nil}
			return nil
		} else {
			return t.Right.Insert(v)
		}
	} else {
		return ErrDuplicateInTree(v)
	}
}

// Build tree with values k, 2k, 3k, ... (treeSize * k)
func buildTree(k int) *Tree {
	// Get values to put in tree
	var values []int
	for i := 1; i <= treeSize; i++ {
		values = append(values, k*i)
	}
	// Create tree, adding nodes in random order
	root := Tree{nil, 0, nil}
	indicesUsed := make(map[int]bool)
	usedCount := 0
	for usedCount < treeSize {
		index := rand.Intn(treeSize)
		_, exists := indicesUsed[index]
		if !exists {
			err := root.Insert(values[index])
			if err != nil {
				log.Fatal(err)
			}
			indicesUsed[index] = true
			usedCount++
		}
	}
	// Return pointer to root
	return &root
}

// Inorder tree traversal, push values into channel
func Walk(t *Tree, ch chan int) {
	if(t == nil) {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Preorder tree traversal, push values into channel
func Preorder(t *Tree, ch chan int) {
	if(t == nil) {
		return
	}
	ch <- t.Value
	Preorder(t.Left, ch)
	Preorder(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for i := 0; i < treeSize; i++ {
		if <-c1 != <-c2 {
			return false
		}
	}
	return true
}

// Seed random number generator on init
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	testWalk()
	testSame()
}

func testWalk() {
	fmt.Println("===== Test Walk =====")
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(buildTree(1), c1)
	go Walk(buildTree(2), c2)
	fmt.Println("Walk tree 1:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c1)
	}
	fmt.Println("Walk tree 2:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c2)
	}
}

func testSame() {
	fmt.Println("===== Test Same =====")
	t1 := buildTree(1)
	t2 := buildTree(1)
	c1 := make(chan int)
	c2 := make(chan int)
	go Preorder(t1, c1)
	go Preorder(t2, c2)
	fmt.Println("Preorder tree 1:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c1)
	}
	fmt.Println("Preorder tree 2:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println(Same(t1, t2))
	t3 := buildTree(1)
	t4 := buildTree(2)
	go Preorder(t3, c1)
	go Preorder(t4, c2)
	fmt.Println("Preorder tree 3:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c1)
	}
	fmt.Println("Preorder tree 4:")
	for i := 0; i < treeSize; i++ {
		fmt.Println(<-c2)
	}
	fmt.Println(Same(t3, t4))
}
