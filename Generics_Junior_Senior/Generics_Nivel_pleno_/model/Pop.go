package model


func (p *Stack[T] ) Pop() (T, error) {
	if p.isEmpty() {
		var zero T
		return zero, nil
	}
	ultimo := p.elemento[len(p.elemento)-1]

	p.elemento = p.elemento[:len(p.elemento)-1]
	return ultimo, nil
}