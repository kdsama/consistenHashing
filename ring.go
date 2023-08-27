package main

import (
	"fmt"
	"math/big"
)

type Ringer interface{}

type KVStore map[string]string

type CircularListNode struct {
	Key      string
	Val      *big.Int
	PrevDist *big.Int
	PrevKey  string
	Next     *CircularListNode
}

type RingService struct {
	Head *CircularListNode
	H    Hasher
}

func NewRingService(h Hasher) *RingService {
	return &RingService{
		H: h,
	}
}
func (r *RingService) PlaceNodeHash(key string) {
	tmp := r.Head
	constant := r.H.CalculateHash(key)
	if tmp == nil {
		r.Head = &CircularListNode{
			Key: key,
			Val: constant,
		}
		r.Head.Next = r.Head
		r.Head.PrevDist = big.NewInt(0)
		return
	}

	Node := r.Head
	Diff := r.H.Distance(constant, Node.Val)
	current := Node.Next
	// prev := Node
	for current != nil && current != Node {
		distance := r.H.Distance(constant, current.Val)
		if distance.Sign() > 0 && distance.Cmp(Diff) < 0 {
			Node = current
			Diff = Diff.Set(distance)
		}
		// prev = current
		current = current.Next
	}
	// Need to do the circular bit as well
	// we found a node with minimum distance
	// lets say it is b
	// we are trying to insert c
	// b.Next = c
	newNode := &CircularListNode{
		Key:     key,
		Val:     constant,
		PrevKey: Node.Key,
	}
	fmt.Println(" previous node and New Node ", Node, newNode)
	tt := Node.Next
	Node.Next = newNode
	newNode.PrevDist = r.H.Distance(newNode.Val, Node.Val)
	newNode.Next = tt
}

func (r *RingService) IterateNodes() {
	head := r.Head
	n := head.Next
	i := 0

	for n != nil && n != head {
		n = n.Next
		fmt.Println("rev and key ", n.PrevKey, n.Key)
		i++
	}
}

func (r *RingService) FindMinimumDifference(key string) {
	if r.Head == nil {
		return
	}
	constant := r.H.CalculateHash(key)
	Node := r.Head
	Difference := new(big.Int).Abs(new(big.Int).Sub(constant, Node.Val))
	current := Node.Next

	for current != Node {
		distance := new(big.Int).Abs(new(big.Int).Sub(constant, current.Val))
		if distance.Cmp(Difference) < 0 {
			Node = current
			Difference = Difference.Set(distance)
		}
		current = current.Next
	}
	fmt.Println(Node.Key, Difference, "For the key", key)
}
