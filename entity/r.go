// 电阻
package entity

import (
	"../component"
	"fmt"
)

// 电阻
// ----------
type R struct {
	component.Component
}

// 电阻 - 计算
// ----------
func (r *R) calculate() {
	fmt.Println(fmt.Sprintf("%s calculating ...", r.Name()))
	fmt.Println()
	// 计算根据器件属性值, 及计算规则, 计算针脚数据

	// 计算完毕信息传递
	r.Component.Transmission()

	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", r.Name()))
}
