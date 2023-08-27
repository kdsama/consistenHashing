package main

type Ringer interface {
	AddNode(node string) error
}

type RingService struct {
	h Hasher // hasher service
}

func NewRingService(h Hasher) *RingService {

	return &RingService{
		h: h,
	}
}

func (r *RingService) Initialize(nodes ...[]string) error {

	return nil
}
func (r *RingService) AddNode(node string) error {

	return nil
}

func (r *RingService) GetNodeFromKey(key string) error {

	return nil
}
