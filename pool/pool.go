package pool

type pool[T any] struct {
	objects   []T
	nextIndex int
}

func MakePool[T any](count int) *pool[T] {
	return &pool[T]{
		objects:   make([]T, count),
		nextIndex: 0,
	}
}

func (p *pool[T]) Pop() *T {
	p.nextIndex++
	return &p.objects[p.nextIndex-1]
}
