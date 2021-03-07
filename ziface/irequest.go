package ziface

type IRequest interface {
	GetConnection() IConnect

	GetData() []byte
}
