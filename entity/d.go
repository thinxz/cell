// 电源 [1 号针脚 正极, 2 号针脚 负极]
// 计算规则, 负极 0 , 正极 电压值
package entity

import (
	"../evt"
	"bytes"
	"fmt"
)
import "../component"

// 电源 - 定义
// ----------
type D struct {
	component.Component
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

func (c *D) Calculate(event evt.Event) {
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
	if s1, ok := c.GetStitch(1); ok {
		// 一号针脚 - 正极 -> 设置针脚数据
		s1.Signal.V = c.v

		// 判断是否需要传输信号
		for i := 0; i < len(s1.Relation); i++ {
			// 器件 -> 关联其他器件针脚
			ts1 := s1.Signal
			if ts2, ok := s1.Relation[i].Component.GetStitch(s1.Relation[i].No); ok {
				if ts1.V != ts2.Signal.V {
					// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
					c.Event("DataChange", c.Name(), 1, s1.Relation[i].Component.Name(), s1.Relation[i].No)
				}
			} else {
				// 关联器件对应针脚已损坏, 或不存在 -> 删除

			}
		}
	}

	if s2, ok := c.GetStitch(1); ok {
		// 二号针脚 - 负极 -> 设置针脚数据
		s2.Signal.V = 0
		// 判断是否需要传输信号
		for i := 0; i < len(s2.Relation); i++ {
			// 器件 -> 关联其他器件针脚
			ts1 := s2.Signal
			if ts2, ok := s2.Relation[i].Component.GetStitch(s2.Relation[i].No); ok {
				if ts1.V != ts2.Signal.V {
					// 关联器件对应针脚信号不匹配配 -> 发布信号改变事件到对应的器件
					c.Event("DataChange", c.Name(), 2, s2.Relation[i].Component.Name(), s2.Relation[i].No)
				}
			} else {
				// 关联器件对应针脚已损坏, 或不存在 -> 删除

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
