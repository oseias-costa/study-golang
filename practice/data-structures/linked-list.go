package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int8
}

type List struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value Person
	Next  *Node
}

func main() {
	list := List{}

	oseias := Person{"Oseias", "Costa", 32}
	outro := Person{"Outro", "Qualquer", 30}
	maria := Person{"maria", "dos santos", 30}
	dani := Person{"dani", "shemer", 30}
	marcos := Person{"marcos", "da costa", 30}
	novo := Person{"novo", "ultimo", 30}

	list.append(oseias)
	list.append(outro)
	list.append(maria)
	list.append(dani)
	list.append(marcos)
	list.append(novo)

	list.display()
	fmt.Println("--------------------------")

	search := list.seach("dani")
	fmt.Println(search)

	fmt.Println("--------------------------")

	list.delete("dani")
	list.display()
}

func (l *List) append(person Person) {
	node := &Node{Value: person}
	if l.Head == nil {
		l.Head = node
	}
	if l.Tail != nil {
		l.Tail.Next = node
	}
	l.Tail = node
}

func (l *List) display() {
	node := l.Head
	for node != nil {
		fmt.Println(node.Value.FirstName)
		node = node.Next
	}
}

func (l *List) seach(name string) (person Person) {
	node := l.Head
	for node != nil {
		if node.Value.FirstName == name {
			person = node.Value
			break
		}
		node = node.Next
	}

	return
}

func (l *List) delete(name string) {
	if l.Head.Value.FirstName == name {
		l.Head = l.Head.Next
		return
	}
	prev := l.Head
	node := l.Head.Next
	for node != nil {
		if node.Value.FirstName == name {
			prev.Next = node.Next
			break
		}
		prev = node
		node = node.Next
	}
	if l.Tail == node {
		l.Tail = prev
	}
}
