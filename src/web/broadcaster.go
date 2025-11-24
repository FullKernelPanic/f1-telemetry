package web

type Broadcaster interface {
	Broadcast(b []byte)
}
