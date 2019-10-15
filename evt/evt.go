package evt

import "../linked"

// 事件定义
type Event struct {
	EventType string // 事件类型
	Source    string // 发布器件名称
	SourceNo  int    // 发布器件针脚
	Target    string // 接收器件名称
	TargetNo  int    // 接收器件针脚
	Data      string // 事件数据
}

// 事件管理器 定义
type EventManager struct {
	// 全局事件表
	events linked.List
}

// 事件管理器 初始化
func (e *EventManager) Init() {
	lt := &linked.List{}
	lt.Init()
	e.events = *lt
}

// 发布事件
func (e *EventManager) Put(event Event) {
	e.events.Put(event)
}

// 处理事件
func (e *EventManager) Push() (Event, bool) {
	c, ok := e.events.Push().(Event)
	if ok {
		return c, true
	} else {
		return Event{}, false
	}
}
