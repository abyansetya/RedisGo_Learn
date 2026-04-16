package booking

import "sync"

type ConcurrentStore struct {
	mu       sync.RWMutex
	bookings map[string]Booking
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (s *ConcurrentStore) Book(b Booking) (Booking, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	//seat is taken
	if _, exists := s.bookings[b.SeatID]; exists {
		return Booking{}, ErrSeatAlreadyBooked
	}
	s.bookings[b.SeatID] = b
	return b, nil
}

func (s *ConcurrentStore) ListBookings(movieID string) []Booking {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []Booking
	for _, b := range s.bookings {
		if b.MovieID == movieID {
			result = append(result, b)
		}
	}
	return result
}
