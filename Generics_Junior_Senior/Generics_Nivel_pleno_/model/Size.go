package model

func (s *Stack[T]) Size() int{
	return len(s.elemento)
}