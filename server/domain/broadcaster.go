package domain

type Broadcaster interface {
	Broadcast(b []byte)
}
