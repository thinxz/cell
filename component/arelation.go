// 器件针脚连接关系管理器
package component

type Relation struct {
	name string // 器件名称
	no   int    // 针脚号
}

//func (r Relation) Equals(other Relation) bool {
//	if r.name == other.name && r.no == other.no {
//		return true
//	}
//	return false
//}

// 器件关联关系管理器
type RelationManger struct {
	// 线路点名称:器件名称+针脚号
	points map[string][]Relation
}

func (m *RelationManger) Init() {
	m.points = make(map[string][]Relation)
}

func (m *RelationManger) AddRelation(name string, no int, point *Point) {
	// 01 添加针脚和点之间关系
	if c, ok := m.points[point.name]; ok {
		// 线路存在
		add := true
		for i := 0; i < len(c); i++ {
			if c[i].name == name && c[i].no == no {
				add = false
			}
		}
		if add {
			// 添加关系
			m.points[point.name] = append(c, Relation{name, no})
		}
	} else {
		// 新增关联关系 -  针脚从 1 计数, 0 为线路值
		m.points[point.name] = []Relation{{name, no}}
	}
}

// 获取 获取器件关联的所有点
func (m *RelationManger) GetStitchPoint(name string, no int) (string, int, bool) {
	// 获取器件关联的所有点
	for p, r := range m.points {
		for i := 0; i < len(r); i++ {
			if r[i].name == name && r[i].no == no {
				// 器件:针脚关联关系
				return p, 0, true
			}
		}
	}
	return "", 0, false
}

func (m *RelationManger) GetRelation(name string) ([]Relation, bool) {
	// 获取器件关联的所有点
	relation, ok := m.points[name]
	return relation, ok
}
