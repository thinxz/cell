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
	// 触发计算 -> 根据电荷量计算电流
	e := ElectricCharge{}
	e.ByJson(event.Data)

	// 计算电流 、电压
	c.i = e.GetEC() / e.T
	c.v = c.i * c.r
}

func (c *R) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电阻值\":\"%.1f欧\"} ", c.Name(), c.r))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
