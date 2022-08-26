package blut

// meta 存储了数据库的表定义, 索引的定义等信息.
type meta struct {
	Id   uint64 `json:"id"`   // 元数据id
	Size int    `json:"size"` // 表容量

	Ns *Namespace
}

func newMeta(id uint64, ns *Namespace) *meta {
	return &meta{
		Id:   id,
		Size: 0,
		Ns:   ns,
	}
}

func (meta) key() []byte {

}
