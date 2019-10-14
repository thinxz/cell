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
	v float64 // 电压, 恒压电源
}

// 初始化书信值
// ----------
// v 电压值 (单位 V)
// ----------
func (c *D) InitProperty(v float64) {
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
	// 01 定义针脚并获取该器件的针脚信息
	s := c.CalculateInitStitch(1, 2)
	one := s[0]
	two := s[1]

	// 02 根据事件, 设置针脚信息
	c.CalculateSetStitch(event.Source, one)
	c.CalculateSetStitch(event.Source, two)

	// 03 根据器件属性值及计算规则, 计算针脚数据, 并设置自身针脚值

	// 04 传输针脚信号改变事件
	c.CalculatePowerTransmissionStitch(one)
	c.CalculatePowerTransmissionStitch(two)
}

func (c *D) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电压值\":\"%.1fV\"} ", c.Name(), c.v))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
