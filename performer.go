package performer

type Performer interface {
	Perform(HandlerFunc) Performer
}
