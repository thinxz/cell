package component

// 获取需要的针脚
func (c *Component) CalculateInitStitch(no ...int) []*Stitch {
	s := [SMaxNum]*Stitch{}
	for i := 0; i < len(no); i++ {
		if ts, ok := c.Stitch(no[i]); ok {
			s[i] = ts
		} else {
			// 错误->查询自身针脚不存在
		}
	}
	return s[0:len(no)]
}

// 获取自身在对应器件中, 对应针脚与自身针脚的关系
func (c *Component) GetComponentStitchRelation(name string, no int) (Relation, bool) {
	if s, ok := c.GetComponentStitch(name, no); ok {
		if v, ok := s.relation[c.name]; ok {
			return *v, ok
		}
	}
	return Relation{
		signal: &Signal{
			V{},
			I{},
		},
	}, false
}

func (c *Component) CalculateSetStitch(source string, s *Stitch) {
	if ownerRelation, ok := s.relation[source]; ok {
		// 查询关联器件
		if other, ok := c.GetComponentStitchRelation(ownerRelation.nameRelation, ownerRelation.noRelation); ok {
			if !ownerRelation.signal.V.Equal(other.signal.V) {
				// 电压不相等
				ownerRelation.signal.V.Value = other.signal.V.Value
			}

			if !ownerRelation.signal.I.Equal(other.signal.I) {
				// 电流不相等
				ownerRelation.signal.I.Value = other.signal.I.Value
			}
		}
	}
}

func (c *Component) CalculateTransmissionStitch(s *Stitch) {
	// 判断是否需要传输信号
	for _, ownerRelation := range s.relation {
		// 查询关联器件
		if otherRelation, ok := c.GetComponentStitchRelation(ownerRelation.nameRelation, ownerRelation.noRelation); ok {
			// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
			ev := false
			if !ownerRelation.signal.V.Equal(otherRelation.signal.V) {
				// 电压
				ev = true
			}

			if !ownerRelation.signal.I.Equal(otherRelation.signal.I) {
				// 电流
				ev = true
			}
			// 发布事件
			if ev {
				c.Event("DataChange", c.Name(), s.no, ownerRelation.nameRelation, ownerRelation.noRelation)
			}
		}
	}
}

func (c *Component) CalculatePowerTransmissionStitch(s *Stitch) {
	// 判断是否需要传输信号
	for _, ownerRelation := range s.relation {
		// 查询关联器件
		if otherRelation, ok := c.GetComponentStitchRelation(ownerRelation.nameRelation, ownerRelation.noRelation); ok {
			// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
			ev := false
			if !ownerRelation.signal.V.PowerEqual(otherRelation.signal.V) {
				// 电压
				ev = true
			}

			if !ownerRelation.signal.I.PowerEqual(otherRelation.signal.I) {
				// 电流
				ev = true
			}
			// 发布事件
			if ev {
				c.Event("DataChange", c.Name(), s.no, ownerRelation.nameRelation, ownerRelation.noRelation)
			}
		}
	}
}
