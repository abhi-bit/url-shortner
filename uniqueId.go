package urlshortner

type Generator struct {
	c       chan int64
	counter int64
}

func NewGenerator() *Generator {
	return &Generator{
		c:       make(chan int64, 1),
		counter: 10000,
	}
}

func (this *Generator) Start() {
	go this.Sequencer()
}

func (this *Generator) Sequencer() {
	for {
		this.counter++
		this.c <- this.counter
	}
}

func (this *Generator) GetID() int64 {
	return <-this.c
}
