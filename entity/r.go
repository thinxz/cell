// 电阻
package entity

import (
	"../component"
	"bytes"
	"fmt"
)

// 电阻
// ----------
type R struct {
	component.Component
	r float32 // 电阻, 电阻恒定
}

// 初始化书信值
// ----------
// r 电阻值 (单位, 欧姆)
// ----------
func (c *R) InitProperty(r float32) {
	c.r = r
}

func (c *R) Calculate() {
	//fmt.Println(fmt.Sprintf("电阻 [%s] -> recalculate ing .......... ......... ", c.Name()))
	fmt.Println(fmt.Sprintf("%s calculating ...", c.Name()))
	fmt.Println()
	c.calculate()
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", c.Name()))
}

// 电阻 - 计算
// ----------
func (c *R) calculate() {
	// 计算根据器件属性值, 及计算规则, 计算针脚数据

}

func (c *R) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电阻值\":\"%.1f欧\"} ", c.Name(), c.r))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}
