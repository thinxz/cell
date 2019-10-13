package component

import "../evt"

// 针脚相联的器件
// ----------
type Relation struct {
	// 相联的器件
	name string
	// 相联器件的针脚
	no int
}

func (r *Relation) Name() string {
	return r.name
}

func (r *Relation) No() int {
	return r.no
}

// 针脚定义
// ---------- ----------
type Stitch struct {
	// 针脚信息[该器件在该针脚上面的信号]
	Signal evt.Signal
	// 针脚序号
	no int
	// 相联的器件
	Relation map[string]*Relation
}

func (s *Stitch) AddRelation(r *Relation) {
	if _, ok := s.Relation[r.Name()]; !ok {
		s.Relation[r.Name()] = r
	} else {
		// 器件已关联
	}
}

// 针脚号 该器件针脚号
// ---------- ----------
func (s *Stitch) GetNo() int {
	return s.no
}
