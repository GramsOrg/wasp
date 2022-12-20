package pipe

// InfinitePipe provides deserialised sender and receiver: it queues messages
// sent by the sender and returns them to the receiver whenever it is ready,
// without blocking the sender process. Depending on the backing queue, the pipe
// might have other characteristics.
type InfinitePipe[E Hashable] struct {
	input  chan E
	output chan E
	length chan int
	buffer Queue[E]
}

var _ Pipe[Hashable] = &InfinitePipe[Hashable]{}

func NewDefaultInfinitePipe[E Hashable]() Pipe[E] {
	return newInfinitePipe(NewDefaultLimitedPriorityHashQueue[E]())
}

func NewPriorityInfinitePipe[E Hashable](priorityFun func(E) bool) Pipe[E] {
	return newInfinitePipe(NewPriorityLimitedPriorityHashQueue(priorityFun))
}

func NewLimitInfinitePipe[E Hashable](limit int) Pipe[E] {
	return newInfinitePipe(NewLimitLimitedPriorityHashQueue[E](limit))
}

func NewLimitPriorityInfinitePipe[E Hashable](priorityFun func(E) bool, limit int) Pipe[E] {
	return newInfinitePipe(NewLimitPriorityLimitedPriorityHashQueue(priorityFun, limit))
}

func NewHashInfinitePipe[E Hashable]() Pipe[E] {
	return newInfinitePipe(NewHashLimitedPriorityHashQueue[E](true))
}

func NewPriorityHashInfinitePipe[E Hashable](priorityFun func(E) bool) Pipe[E] {
	return newInfinitePipe(NewPriorityHashLimitedPriorityHashQueue(priorityFun, true))
}

func NewLimitHashInfinitePipe[E Hashable](limit int) Pipe[E] {
	return newInfinitePipe(NewLimitHashLimitedPriorityHashQueue[E](limit, true))
}

func NewInfinitePipe[E Hashable](priorityFun func(E) bool, limit int) Pipe[E] {
	return newInfinitePipe(NewLimitedPriorityHashQueue(priorityFun, limit, true))
}

func newInfinitePipe[E Hashable](queue Queue[E]) *InfinitePipe[E] {
	ch := &InfinitePipe[E]{
		input:  make(chan E),
		output: make(chan E),
		length: make(chan int),
		buffer: queue,
	}
	go ch.infiniteBuffer()
	return ch
}

func (ch *InfinitePipe[E]) In() chan<- E {
	return ch.input
}

func (ch *InfinitePipe[E]) Out() <-chan E {
	return ch.output
}

func (ch *InfinitePipe[E]) Len() int {
	return <-ch.length
}

func (ch *InfinitePipe[E]) Close() {
	close(ch.input)
}

func (ch *InfinitePipe[E]) infiniteBuffer() {
	var input, output chan E
	var next E
	var nilE E
	input = ch.input

	for input != nil || output != nil {
		select {
		case elem, open := <-input:
			if open {
				ch.buffer.Add(elem)
			} else {
				input = nil
			}
		case output <- next:
			ch.buffer.Remove()
		case ch.length <- ch.buffer.Length():
		}

		if ch.buffer.Length() > 0 {
			output = ch.output
			next = ch.buffer.Peek()
		} else {
			output = nil
			next = nilE
		}
	}

	close(ch.output)
	close(ch.length)
}
