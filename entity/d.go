// 电源
package entity

import "fmt"
import "../component"

// 电源 - 定义
// ----------
type D struct {
	component.Component
}

// 电源 - 初始化计算
// 加电自检
// ----------
func (c *D) InitCalculate() {
	c.calculate()
}

// 电源 - 计算
// ----------
func (c *D) calculate() {
	fmt.Println(fmt.Sprintf("%s calculating ...", c.Name()))
	fmt.Println()
	// 计算根据器件属性值, 及计算规则, 计算针脚数据

	// 计算完毕信息传递
	c.Component.Transmission()
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", c.Name()))
}
