package main

import "fmt"

// Node class
type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
}

// NodeBetweenValues method of LinkedList
// Traverses the list to find out the node that has a property lying between
// the firstProperty and secondProperty

func (linkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty &&
				node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

// AddToHead method of LinkedList
func (linkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
		linkedList.headNode.previousNode = node
	}
	linkedList.headNode = node
}

// NodeWithValue returns Node given parameter property
func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

// AddAfter
func (linkedList *LinkedList) AddAfter(nodeProperty int, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = linkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

// LastNode returns the node at the end of the list. The list is traversed
// to check whether nextNode is nil from nextNode of headNode
func (linkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddToEnd
func (linkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (linkedList *LinkedList) IterateList() {
	var node *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	linkedList := LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToEnd(3)
	linkedList.AddAfter(3, 7)
	fmt.Println("HeadNode:", linkedList.headNode.property)
	fmt.Println("List:")
	linkedList.IterateList()
	fmt.Println("Node between 1 and 7:")
	node := linkedList.NodeBetweenValues(1, 7)
	fmt.Println(node.property)
}
