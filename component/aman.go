// 器件管理器
package component

import (
	"../evt"
	"fmt"
	"time"
)

var (
	r1 R
	r2 R
	d1 D
)

// 器件管理器 定义
// ---------- ----------
type EnManager struct {
	// 实例存储器, 保存所有器件对象
	obj map[string]IComponent
	// 事件处理器, 管理所有事件
	eventManager evt.EventManager
	// 器件关系管理器
	relationManger RelationManger
}

// 器件管理器初始化
// ----------
func (e *EnManager) Init() {
	// 初始化实例存储器
	e.obj = make(map[string]IComponent)

	// 初始化事件管理器
	e.eventManager = evt.EventManager{}
	e.eventManager.Init()

	// 初始化器件关系管理器
	e.relationManger = RelationManger{}
	e.relationManger.Init()

	// 初始化所有对象
	e.initEntity()
}

// 初始化所有实例对象
// ----------
// eventManager 事件管理器, 管理对象间事件传递
// ----------
func (e *EnManager) initEntity() {
	// 初始化, 所有元器对象 -> 并填充默认值(静态属性)
	r1 = R{}
	r1.Init("R1", 2, e)
	r1.InitProperty(1.0)
	r2 = R{}
	r2.Init("R2", 2, e)
	r2.InitProperty(2.0)
	d1 = D{}
	d1.Init("D1", 2, e)
	d1.InitProperty(5.0)

	// 保存器件对象
	e.addEntity(&r1)
	e.addEntity(&r2)
	e.addEntity(&d1)

	// 初始化点线路
	p1 := &Point{}
	p1.Init("P1", 0, e)
	p2 := &Point{}
	p2.Init("P2", 0, e)
	p3 := &Point{}
	p3.Init("P3", 0, e)
	e.addEntity(p1)
	e.addEntity(p2)
	e.addEntity(p3)

	// 初始化, 针脚连接
	e.relationManger.AddRelation("R1", 1, p1)
	e.relationManger.AddRelation("R1", 2, p2)
	e.relationManger.AddRelation("R2", 1, p2)
	e.relationManger.AddRelation("R2", 2, p3)
	e.relationManger.AddRelation("D1", 1, p1)
	e.relationManger.AddRelation("D1", 2, p3)
}

func (e *EnManager) Run() {
	// 加电自检电源初始化计算
	d1.InitCalculate()

	// 循环处理事件
	fmt.Println()
	for {
		if event, ok := e.push(); ok {
			// 获取器件对象
			if c, ok := e.getEntity(event.Target); ok {
				// 传输信息, 触发该器件重新计算
				c.Transmission(event)
				fmt.Println()
			}
		} else {
			// 没有待处理事件, 暂停执行
			time.Sleep(5 * 1000000000) //等待1秒
			fmt.Println()
		}
	}
}

func (e *EnManager) addEntity(c IComponent) {
	e.obj[c.Name()] = c
}

func (e *EnManager) getEntity(name string) (IComponent, bool) {
	if v, ok := e.obj[name]; ok {
		return v, ok
	}
	return nil, false
}

// 获取器件针脚关联的唯一点
func (e *EnManager) GetStitchPoint(name string, no int) (string, int, bool) {
	return e.relationManger.GetStitchPoint(name, no)
}

func (e *EnManager) GetRelation(point string) ([]Relation, bool) {
	return e.relationManger.GetRelation(point)
}

// 器件发布事件
// ---------- ----------
// event    事件类型
// source   发布器件名称
// sourceNo 发布器件针脚
// target   接收器件名称
// targetNo 接收器件针脚
// ---------- ----------
func (e *EnManager) PutEvent(event, source string, sourceNo int, target string, targetNo int, data string) {
	// 发布事件到事件管理器 -> [事件管理器根据器件唯一标识符查询器件, 并触发相联器件计算]
	e.put(evt.Event{
		EventType: event,
		Source:    source,   // 发布事件器件
		SourceNo:  sourceNo, // 发布器件针脚
		Target:    target,   // 接收事件器件
		TargetNo:  targetNo, // 接收器件针脚
		Data:      data,     // 事件数据
	})
}

// 发布事件
func (e *EnManager) put(event evt.Event) {
	e.eventManager.Put(event)
}

// 处理事件
func (e *EnManager) push() (evt.Event, bool) {
	return e.eventManager.Push()
}
