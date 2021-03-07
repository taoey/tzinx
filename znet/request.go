package znet

import "github.com/taoey/tzinx/ziface"

type Requeset struct {
	conn ziface.IConnect

	data []byte
}

func (r *Requeset) GetConnection() ziface.IConnect {
	return r.conn
}

func (r *Requeset) GetData() []byte {
	return r.data
}
