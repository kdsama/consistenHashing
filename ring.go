package main

import (
	"fmt"
	"log"
	"math/big"
)

type Ringer interface{}

var KVStore = map[*big.Int]string{}

type CircularListNode struct {
	Key  string
	Val  *big.Int
	Prev *CircularListNode
	Next *CircularListNode
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
	constant := r.H.CalculateHash(key)
	if _, ok := KVStore[constant]; ok {
		log.Fatal("HAsh already exists ")
	}
	newNode := &CircularListNode{
		Key: key,
		Val: constant,
	}
	KVStore[constant] = key
	if r.Head == nil {
		r.Head = newNode
		r.Head.Next = r.Head
		r.Head.Prev = r.Head
		return
	}

	current := r.Head
	minDist := r.H.Distance(constant, current.Val)
	closestNode := current

	for current.Next != r.Head {
		distance := r.H.Distance(constant, current.Val)
		if distance.Cmp(minDist) < 0 {
			minDist.Set(distance)
			closestNode = current
		}
		current = current.Next
	}

	newNode.Next = closestNode.Next
	closestNode.Next = newNode
	newNode.Prev = closestNode
	newNode.Next.Prev = newNode
}

func (r *RingService) IterateNodes() {
	head := r.Head
	n := head.Next
	i := 0

	for n != nil && n != head {
		n = n.Next
		fmt.Println("Previous and Current  ", n.Prev, n)
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
