package component

import (
	"../evt"
	"fmt"
)

// 线路, 点定义 -> 计算单元
// ---------- ----------
// 所有事件信息都是先发布到点 -> 线路点进行计算
// 然后将事件发布到器件 -> 器件计算后, 再将事件传递到点
// ---------- ----------
type Point struct {
	Component
}

func (p *Point) Transmission(event evt.Event) {
	//fmt.Println(fmt.Sprintf("电源 [%s] -> recalculate ing .......... ......... ", c.Name()))
	fmt.Println(fmt.Sprintf("%s calculating ...", p.Name()))
	fmt.Println()
	p.calculate(event)
	fmt.Println()
	fmt.Println(fmt.Sprintf("%s finish ", p.Name()))
}

func (p *Point) calculate(event evt.Event) {
	e := ElectricCharge{}
	e.ByJson(event.Data)

	// 器件之间计算 [并将电能传输给对应器件]
	if event.EventType == "TransEC" {
		if c, ok := p.GetRelation(p.name); ok {
			for i := 0; i < len(c); i++ {
				if !(c[i].name == event.Source && c[i].no == event.SourceNo) {
					// 正反馈
					p.Event("TransEC", p.name, 0, c[i].name, c[i].no, e)
				}
			}
		}
	}
}
