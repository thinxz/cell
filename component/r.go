// 电阻
package component

import (
	"../evt"
	"bytes"
	"fmt"
	"math"
)

// 电阻
// ----------
type R struct {
	Component
	r float64 // 电阻, 电阻恒定
	i float64
	v float64
}

// 初始化书信值
// ----------
// r 电阻值 (单位, 欧姆)
// ----------
func (c *R) InitProperty(r float64) {
	c.r = r
}

func (c *R) Transmission(event evt.Event) {
	//fmt.Println(fmt.Sprintf("电阻 [%s] -> recalculate ing .......... ......... ", c.Name()))
	fmt.Println(fmt.Sprintf("%s calculating ...", c.Name()))
	fmt.Println()
	c.calculate(event)
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", c.Name()))
}

// 电阻 - 计算
// ----------
func (c *R) calculate(event evt.Event) {
	s := c.initStitch(1, 2)
	one := s[0]
	two := s[1]

	// 01 遍历针脚, 根据传输信号, 变化设置值, 根据源器件事件, 设置目遍器件对应目标改变值
	c.setStitch(event.Source, one)
	c.setStitch(event.Source, two)

	// 02 计算根据器件属性值, 及计算规则, 计算针脚数据
	if one != nil && one.Signal != nil && two != nil && two.Signal != nil {
		c.v = one.Signal.V.Value - two.Signal.V.Value
		if c.v < 0 {
			// one 低电位, two 高电位 | 电压方向 two -> one
			one.Signal.V.Direction = Forward
			one.Signal.I.Direction = Forward
			two.Signal.V.Direction = Reverse
			two.Signal.I.Direction = Reverse
			//
			c.v = math.Abs(c.v)
		} else {
			one.Signal.V.Direction = Reverse
			one.Signal.I.Direction = Reverse
			two.Signal.V.Direction = Forward
			two.Signal.I.Direction = Forward
		}
		c.i = math.Abs(c.v) / c.r
		two.Signal.I.Value = c.i
	}

	// 03 遍历针脚, 判断是否需要传输信号
	c.transmissionStitch(one)
	c.transmissionStitch(two)
}

// 获取需要的针脚
func (c *R) initStitch(no ...int) []*Stitch {
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

func (c *R) setStitch(source string, s *Stitch) {
	if relation, ok := s.Relation[source]; ok {
		// 查询关联器件
		if st, ok := c.GetComponentStitch(relation.Name(), relation.No()); ok {
			// 查询关联器件针脚
			if !s.Signal.V.Equal(st.Signal.V) {
				// 电压不相等
				s.Signal.V.Value = st.Signal.V.Value
			}

			if !s.Signal.I.Equal(st.Signal.I) {
				// 电流不相等
				s.Signal.I.Value = st.Signal.I.Value
			}
		}
	}
}

func (c *R) transmissionStitch(s *Stitch) {
	// 判断是否需要传输信号
	for name, v := range s.Relation {
		// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
		if t, ok := c.GetComponentStitch(v.Name(), v.No()); ok && v.Name() == name {
			ev := false
			if !s.Signal.V.Equal(t.Signal.V) {
				// 电压
				ev = true
			}
			if !s.Signal.I.Equal(t.Signal.I) {
				// 电流
				ev = true
			}
			// 发布事件
			if ev {
				c.Event("DataChange", c.Name(), s.no, name, v.No())
			}
		}
	}
}

func (c *R) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电阻值\":\"%.1f欧\"} ", c.Name(), c.r))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
