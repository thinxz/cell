// 电源 [1 号针脚 正极, 2 号针脚 负极]
// 计算规则, 负极 0 , 正极 电压值
package component

import (
	"../evt"
	"bytes"
	"fmt"
	"time"
)

// 电源 - 定义
// ----------
type D struct {
	Component
	v  float64         // 电压, 恒压电源
	ec *ElectricCharge // 电荷量 (库伦)
	t  float64         // 时间速率 * 时间单位
}

// 初始化书信值
// ----------
// v 电压值 (单位 V)
// ----------
func (c *D) InitProperty(v float64) {
	c.v = v
	c.ec = NewElectricCharge(int64(c.v))
}

// 电源 - 初始化计算
// 加电自检
// ----------
func (c *D) InitCalculate() {
	fmt.Println(fmt.Sprintf("%s 加电自检 ...", c.Name()))
	fmt.Println()
	// 初始化计算
	c.calculate(evt.Event{})
	// 电动源, 自动发布事件
	go c.powerSupply()
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
// 电荷从正极 发出, 从负极收回
// ----------
func (c *D) calculate(event evt.Event) {
	// 处理数据
	if event.EventType == "" {
		// 启动初始化计算
		c.t = 10
		c.ec.T = c.t
	} else if event.EventType == "TransEC" {
		// 电能传输回负极
	}
}

func (c *D) Describe() string {
	var buff bytes.Buffer
	buff.WriteString(fmt.Sprintf("{\"名称\":\"%s\",\"电压值\":\"%.1fV\"} ", c.Name(), c.v))
	buff.WriteString(c.Component.Describe())
	return buff.String()
}

// 电源为自动发送事件源, 开启协程自动处理
// ---------- ----------
func (c *D) powerSupply() {
	for {
		// 暂停 间隔
		time.Sleep(time.Duration(c.t * 1000000000))
		// 传输电能
		if name, no, ok := c.GetStitchPoint(c.name, 1); ok {
			// 传输电能事件, 从电源正极, 开始传播 (器件事件只能传播到点)
			c.Event("TransEC", c.name, 1, name, no, *c.ec)
		}
	}
}
