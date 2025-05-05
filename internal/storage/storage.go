package storage

import "sync"

type Stats struct {
	mu    sync.Mutex
	Cats  int
	Memes int
}

func (s *Stats) IncCat() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Cats++
}

func (s *Stats) IncMeme() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Memes++
}

func (s *Stats) GetStats() (int, int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.Cats, s.Memes
}
