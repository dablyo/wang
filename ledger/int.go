package ledger

type Provider interface {
	Init(int)
	Open()
	Read() int
	Write(int)
	Close()
}
