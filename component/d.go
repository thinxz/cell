// 电源 [1 号针脚 正极, 2 号针脚 负极]
// 计算规则, 负极 0 , 正极 电压值
package component

import (
	"../evt"
	"bytes"
	"fmt"
)

// 电源 - 定义
// ----------
type D struct {
	Component
	v float32 // 电压, 恒压电源
}

// 初始化书信值
// ----------
// v 电压值 (单位 V)
// ----------
func (c *D) InitProperty(v float32) {
	c.v = v
}

// 电源 - 初始化计算
// 加电自检
// ----------
func (c *D) InitCalculate() {
	fmt.Println(fmt.Sprintf("%s 加电自检 ...", c.Name()))
	fmt.Println()
	c.calculate(evt.Event{})
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", c.Name()))
}

func (c *D) Transmission(event evt.Event) {
	//fmt.Println(fmt.Sprintf("电源 [%s] -> recalculate ing .......... ......... ", c.Name()))
	fmt.Println(fmt.Sprintf("%s calculating ...", c.Name()))
	fmt.Println()
	c.calculate(event)
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", c.Name()))
}

// 电源 - 计算
// ----------
func (c *D) calculate(event evt.Event) {
	// 计算根据器件属性值, 及计算规则, 计算针脚数据

	// 判断修改各个针脚数据, 并判断该针脚是否需要传输信息
	if s1, ok := c.Stitch(1); ok {
		// 一号针脚 - 正极 -> 设置针脚数据
		s1.Signal.V.Value = c.v
		s1.Signal.V.Direction = evt.Forward
		s1.Signal.I.Direction = evt.Forward

		// 判断是否需要传输信号
		for name, v := range s1.Relation {
			// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
			if t, ok := c.GetComponentStitch(v.Name(), v.No()); ok && (v.Name() == name || name == "") {
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
		// 二号针脚 - 负极 -> 设置针脚数据
		s2.Signal.V.Value = 0
		s2.Signal.V.Direction = evt.Reverse
		// s2.Signal.I.Value
		s2.Signal.I.Direction = evt.Reverse

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

func (c *D) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电压值\":\"%.1fV\"} ", c.Name(), c.v))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
