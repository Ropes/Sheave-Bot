package history

import (
	"container/heap"
	"testing"
)

var x []string
var a []string

func TestInit(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}
	//fmt.Println(x, a)

	hh := NewHistory(20)
	heap.Init(hh)
	//hh.PrintDump()
	heap.Push(hh, x)
	//hh.Add(x)
	//hh.PrintDump()
	if hh.heap[0][0] != "x" {
		t.Errorf("First element incorrect: %#v", hh.heap[0][0])
	}
	heap.Push(hh, a)
	//hh.Add(a)

	//fmt.Printf("hh: %#v\n", hh)
	if hh == nil {
		t.Errorf("First history stored is: %#v", hh)
	}
	if hh.heap[0][0] != "a" {
		t.Errorf("First element incorrect: %#v", hh.heap[0][0])
	}
}

func TestPop(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}
	//fmt.Println(x, a)

	hh := NewHistory(20)
	heap.Init(hh)
	heap.Push(hh, x)
	heap.Push(hh, x)
	heap.Push(hh, a)
	/*
		hh.Add(x)
		hh.Add(a)
	*/

	lasti := hh.lastIndex()
	if lasti != 2 {
		hh.PrintDump()
		t.Errorf("Last index not being calculated correctly %#v", lasti)
	}

	poped := hh.Pop().([]string)

	//fmt.Printf("Popped: %#v", poped)
	if poped[0] != "x" {
		t.Errorf("Wrong element poped from stack: %#v\n", poped)
	}
	if hh.Len() != 2 {
		hh.PrintDump()
		t.Errorf("Too many items in heap? %#v\n", hh)
	}
}

func TestLastIndex(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}
	c := []string{"d", "e", "f"}

	hh := NewHistory(20)
	heap.Init(hh)
	heap.Push(hh, x)

	/*
		hh.Add(x)
		hh.Add(a)
		hh.Add(c)
	*/
	heap.Push(hh, a)
	heap.Push(hh, c)

	lasti := hh.lastIndex()
	if lasti != 2 {
		t.Errorf("LastIndex being calculated wrong: %#v", lasti)
	}
}

func TestRound(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}
	c := []string{"d", "e", "f"}
	g := []string{"j", "k", "l"}

	hh := NewHistory(2)
	heap.Init(hh)
	heap.Push(hh, x)
	heap.Push(hh, a)
	heap.Push(hh, c)
	heap.Push(hh, g)
	/*
		hh.Add(x)
		hh.Add(a)
		hh.Add(c)
		hh.Add(g)
	*/

	if len(hh.heap) > 2 {
		t.Errorf("Heap grew beyond limit? \n%#v\n", hh)
	}
	//fmt.Println(hh) :D
}

func TestInsertion(t *testing.T) {
	c := []string{"d", "e", "f"}
	//g := []string{"j", "k", "l"}

	hh := NewHistory(20)
	heap.Init(hh)
	heap.Push(hh, c)
	heap.Push(hh, x)
	heap.Push(hh, a)
	heap.Push(hh, c)
	heap.Push(hh, a)
	/*
		hh.Add(c)
		hh.Add(x)
	*/

	if hh.Hist(0) == nil {
		t.Errorf("Heap Push not adding to correct end of list")
	}
	//hh.PrintDump()
}

func TestLen(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}
	c := []string{"d", "e", "f"}
	g := []string{"j", "k", "l"}

	hh := NewHistory(20)
	heap.Init(hh)

	if hh.Len() != 0 {
		t.Errorf("Limit should be zero! %#v", hh)
	}

	heap.Push(hh, x)
	heap.Push(hh, a)
	heap.Push(hh, c)
	heap.Push(hh, g)

	if hh.Len() != 4 {
		t.Errorf("Error: hh.Len() not 4! %#v", hh)
	}

	for i := 0; i < 21; i++ {
		heap.Push(hh, g)
	}
	hh.PrintDump()

	if hh.Len() != 20 {
		hh.PrintDump()
		t.Errorf("Error: Length grew beyond it's supposed limit: %#v", hh.Len())
	}
}

func TestEmptyHeap(t *testing.T) {
	x = []string{"x", "y", "z"}
	a = []string{"a", "b", "c"}

	hh := NewHistory(20)
	heap.Init(hh)

	emp := hh.Hist(1)
	if len(emp) != 0 {
		t.Errorf("A panic will be thrown before this, but should be an empty string")
	}

	heap.Push(hh, x)
	heap.Push(hh, a)

	emp = hh.Hist(1)
	if len(emp) == 0 {
		t.Errorf("Valid value wasn't returned")
	}

}
