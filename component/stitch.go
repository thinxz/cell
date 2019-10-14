package component

// 针脚相联的器件
// ----------
type Relation struct {
	// 关联针脚之间信息 [该器件在该针脚上面的信号]
	signal *Signal
	// 相联的器件
	nameRelation string
	// 相联器件的针脚
	noRelation int
}

func NewRelation(nameRelation string, noRelation int) *Relation {
	return &Relation{
		signal:       NewSignal(),
		nameRelation: nameRelation,
		noRelation:   noRelation,
	}
}

//func (r *Relation) Name() string {
//	return r.nameRelation
//}
//
//func (r *Relation) No() int {
//	return r.noRelation
//}

// 针脚定义
// ---------- ----------
type Stitch struct {
	// 针脚序号
	no int
	// 相联的器件
	relation map[string]*Relation
}

func NewStitch(no int) *Stitch {
	return &Stitch{
		no:       no,
		relation: make(map[string]*Relation),
	}
}

func (s *Stitch) AddRelation(r *Relation) {
	if _, ok := s.relation[r.nameRelation]; !ok {
		s.relation[r.nameRelation] = r
	} else {
		// 器件已关联
	}
}

// 针脚号 该器件针脚号
// ---------- ----------
func (s *Stitch) GetNo() int {
	return s.no
}

// 获取器件中, 一个针脚中, 关联器件的关系
func (s *Stitch) GetRelation(target string) (*Relation, bool) {
	if relation, ok := s.relation[target]; ok {
		return relation, ok
	}

	// 获取本身的值 -> 没有改变则此接口上所有值都是相同
	if s.relation != nil && len(s.relation) > 0 {
		// 该针脚有关联值[否则为悬空状态]
		for _, relation := range s.relation {
			// 返回获取的第一个 | false 参与计算, 单位该针脚为非改变值
			return relation, false
		}
	}

	// 悬空状态返回空数据
	return NewRelation("", 0), false
}
