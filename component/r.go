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
	// 01 定义针脚并获取该器件的针脚信息
	s := c.CalculateInitStitch(1, 2)
	one := s[0]
	two := s[1]

	// 02 根据事件, 设置针脚信息
	c.CalculateSetStitch(event.Source, one)
	c.CalculateSetStitch(event.Source, two)

	// 03 根据器件属性值及计算规则, 计算针脚数据, 并设置自身针脚值
	if one != nil && two != nil {
		oneRelation, oneOwnerOk := one.GetRelation(event.Source)
		twoRelation, twoOwnerOk := two.GetRelation(event.Source)

		c.v = math.Abs(oneRelation.signal.V.Value - twoRelation.signal.V.Value)
		c.i = c.v / c.r
		if other, ok := c.GetComponentStitchRelation(event.Source, event.SourceNo); ok {
			if oneOwnerOk {
				oneRelation.signal.V.Direction = other.signal.V.Direction * Invert
				oneRelation.signal.I.Direction = other.signal.I.Direction * Invert
				twoRelation.signal.V.Direction = other.signal.V.Direction
				twoRelation.signal.I.Direction = other.signal.I.Direction
			} else if twoOwnerOk {
				twoRelation.signal.V.Direction = other.signal.V.Direction * Invert
				twoRelation.signal.I.Direction = other.signal.V.Direction * Invert
				oneRelation.signal.V.Direction = other.signal.V.Direction
				oneRelation.signal.I.Direction = other.signal.I.Direction
			} else {
				// 错误
			}

			// 计算值
			//oneRelation.signal.V.Value = c.v
			//oneRelation.signal.I.Value = c.i
			//
			//twoRelation.signal.V.Value = c.v
			//twoRelation.signal.I.Value = c.i
		}
	}

	// 04 传输针脚信号改变事件
	c.CalculateTransmissionStitch(one)
	c.CalculateTransmissionStitch(two)
}

func (c *R) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电阻值\":\"%.1f欧\"} ", c.Name(), c.r))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
