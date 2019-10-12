package component

// 针脚相联的器件
// ----------
type Relation struct {
	// 相联的器件
	Component Component
	// 相联器件的针脚
	No int
}

// 针脚定义
// ---------- ----------
type Stitch struct {
	// 针脚信息[该器件在该针脚上面的信号]
	Signal Signal
	// 针脚序号
	no int
	// 相联的器件
	Relation []*Relation
}

func (s *Stitch) AddRelation(r *Relation) {
	s.Relation = append(s.Relation, r)
}

// 针脚号 该器件针脚号
// ---------- ----------
func (s *Stitch) GetNo() int {
	return s.no
}
