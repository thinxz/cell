// 电阻
package component

import (
	"../evt"
	"bytes"
	"fmt"
)

// 电阻
// ----------
type R struct {
	Component
	r float32 // 电阻, 电阻恒定
}

// 初始化书信值
// ----------
// r 电阻值 (单位, 欧姆)
// ----------
func (c *R) InitProperty(r float32) {
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
	// 01 遍历针脚, 根据传输信号, 变化设置值
	if s1, ok := c.Stitch(1); ok {
		// 根据源器件事件, 设置目遍器件对应目标改变值
		if relation, ok := s1.Relation[event.Source]; ok {
			// 查询关联器件
			if st, ok := c.GetComponentStitch(relation.Name(), relation.No()); ok {
				// 查询关联器件针脚
				if !s1.Signal.V.Equal(st.Signal.V) {
					// 电压不相等
					s1.Signal.V.Value = st.Signal.V.Value
					s1.Signal.V.Direction = st.Signal.V.Direction
				}

				if !s1.Signal.I.Equal(st.Signal.I) {
					// 电流不相等
					s1.Signal.I.Value = st.Signal.I.Value
					s1.Signal.I.Direction = st.Signal.I.Direction
				}
			}
		}
	}

	if s2, ok := c.Stitch(2); ok {
		// 根据源器件事件, 设置目遍器件对应目标改变值
		if relation, ok := s2.Relation[event.Source]; ok {
			// 查询关联器件
			if st, ok := c.GetComponentStitch(relation.Name(), relation.No()); ok {
				// 查询关联器件针脚
				if !s2.Signal.V.Equal(st.Signal.V) {
					// 电压不相等
					s2.Signal.V.Value = st.Signal.V.Value
					s2.Signal.V.Direction = st.Signal.V.Direction
				}

				if !s2.Signal.I.Equal(st.Signal.I) {
					// 电流不相等
					s2.Signal.I.Value = st.Signal.I.Value
					s2.Signal.I.Direction = st.Signal.I.Direction
				}
			}
		}
	}

	// 02 计算根据器件属性值, 及计算规则, 计算针脚数据

	// 03 遍历针脚, 判断是否需要传输信号
	if s1, ok := c.Stitch(1); ok {
		// 判断是否需要传输信号
		for name, v := range s1.Relation {
			// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
			if t, ok := c.GetComponentStitch(v.Name(), v.No()); ok && v.Name() == name {
				ev := false
				if !s1.Signal.V.Equal(t.Signal.V) {
					// 电压
					ev = true
				}
				if !s1.Signal.I.Equal(t.Signal.I) {
					// 电流
					ev = true
				}

				// 发布事件
				if ev {
					c.Event("DataChange", c.Name(), 1, name, v.No())
				}
			}
		}
	}

	if s2, ok := c.Stitch(2); ok {
		// 判断是否需要传输信号
		for name, v := range s2.Relation {
			// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
			if t, ok := c.GetComponentStitch(v.Name(), v.No()); ok && v.Name() == name {
				ev := false
				if !s2.Signal.V.Equal(t.Signal.V) {
					// 电压
					ev = true
				}
				if !s2.Signal.I.Equal(t.Signal.I) {
					// 电流
					ev = true
				}

				// 发布事件
				if ev {
					c.Event("DataChange", c.Name(), 2, name, v.No())
				}
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
