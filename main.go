package main

import (
	"math/rand"
	"time"
)

func main() {
	h := NewHashService()
	r := NewRingService(h)
	a := []byte{}
	for i := 0; i < 26; i = i + 2 {
		rand.Seed(time.Now().UnixMilli())
		time.Sleep(100 * time.Millisecond)
		a = append(a, (byte('a') + uint8(i)))
		r.PlaceNodeHash(string(a[len(a)-1]))
	}

	r.IterateNodes()
	for i, v := range a {
		r.FindMinimumDifference(string(v + uint8(i+1)))
	}

}
