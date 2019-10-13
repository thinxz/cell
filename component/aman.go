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
}

// 器件管理器初始化
// ----------
func (e *EnManager) Init() {
	// 初始化实例存储器
	e.obj = make(map[string]IComponent)

	// 初始化事件管理器
	e.eventManager = evt.EventManager{}
	e.eventManager.Init()

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
	r1.Init("R1", e)
	r1.InitProperty(1.0)
	r2 = R{}
	r2.Init("R2", e)
	r2.InitProperty(2.0)
	d1 = D{}
	d1.Init("D1", e)
	d1.InitProperty(5.0)

	// 保存器件对象
	e.addEntity(&r1)
	e.addEntity(&r2)
	e.addEntity(&d1)

	//
	// 初始化, 针脚连接
	r1.AddStitch(1, &r2, 2)
	r1.AddStitch(2, &d1, 2)

	r2.AddStitch(1, &d1, 1)
	r2.AddStitch(2, &r1, 1)

	d1.AddStitch(1, &r2, 1)
	d1.AddStitch(2, &r1, 2)
}

func (e *EnManager) Run() {
	// 加电自检电源初始化计算
	d1.InitCalculate()

	// 循环处理事件
	fmt.Println()
	for {
		if event, ok := e.eventManager.Push(); ok {
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

//
func (e *EnManager) GetComponent(name string) (IComponent, bool) {
	if v, ok := e.getEntity(name); ok {
		return v, ok
	}
	return nil, false
}

// 器件发布事件
// ---------- ----------
// event    事件类型
// source   发布器件名称
// sourceNo 发布器件针脚
// target   接收器件名称
// targetNo 接收器件针脚
// ---------- ----------
func (e *EnManager) PutEvent(event, source string, sourceNo int, target string, targetNo int) {
	fmt.Println(fmt.Sprintf("接收事件, Component [%s:%d] Event [%s] ", source, targetNo, event))

	// 发布事件到事件管理器 -> [事件管理器根据器件唯一标识符查询器件, 并触发相联器件计算]
	e.put(evt.Event{
		EventType: event,
		Source:    source,   // 发布事件器件
		SourceNo:  sourceNo, // 发布器件针脚
		Target:    target,   // 接收事件器件
		TargetNo:  targetNo, // 接收器件针脚
	})
}

// 发布事件
func (e *EnManager) put(event evt.Event) {
	e.eventManager.Put(event)
}

// 处理事件
func (e *EnManager) Push() (evt.Event, bool) {
	return e.eventManager.Push()
}
