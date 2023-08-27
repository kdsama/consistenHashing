package main

type ConHasher interface {
	Initialize(nodes ...string) error
	AddNode(node string) error
	GetNodeFromKey(key string) error
	PlaceHash() error
}

type ConHashService struct {
	h Hasher // hasher service
	r Ringer
}

func NewConHashService(h Hasher) *ConHashService {

	return &ConHashService{
		h: h,
	}
}

func (r *ConHashService) Initialize(nodes ...string) error {
	for _, node := range nodes {
		err := r.AddNode(node)
		if err != nil {
			return err
		}

	}
	return nil
}
func (r *ConHashService) AddNode(node string) error {
	r.h.CalculateHash(node)
	return nil

}

func (r *ConHashService) PlaceHash() error {
	// place the hash on the ring
	return nil
}

func (r *ConHashService) GetNodeFromKey(key string) error {

	return nil
}
