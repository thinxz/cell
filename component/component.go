package component

import (
	"../evt"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// 器件接口定义
// ---------- ---------
type IComponent interface {
	Name() string                    // 获取器件名称, 唯一标识符
	GetStitch(no int) (Stitch, bool) // 获取器件针脚
	Event(event, source string, sourceNo int,
		target string, targetNo int) // 器件发布事件
	Calculate(event evt.Event) // 触发对应器件对象计算
	Describe() string          // 器件相关信息描述
}

// 器件定义
// ---------- ---------
type Component struct {
	name string            // 器件名称
	sts  map[int]Stitch    // 针脚定义
	man  *evt.EventManager // 器件 - 事件管理器
}

// 初始化器件数据
// ---------- ----------
func (c *Component) Init(name string, man *evt.EventManager) {
	c.sts = make(map[int]Stitch)
	c.name = name
	c.man = man
}

// 获取器件名称, 唯一标识符
// ---------- ----------
func (c *Component) Name() string {
	return c.name
}

// 获取器件针脚
// ---------- ----------
func (c *Component) GetStitch(no int) (Stitch, bool) {
	s, ok := c.sts[no]
	return s, ok
}

// 初始化器件针脚, 及其关联器件针脚
// ---------- ----------
func (c *Component) AddStitch(no int, target IComponent, targetNo int) {
	s, ok := c.sts[no]
	if !ok {
		s = Stitch{
			no:       no,
			Relation: []*Relation{},
			Signal:   Signal{},
		}
	}

	s.AddRelation(&Relation{
		No:        targetNo,
		Component: target,
	})

	c.sts[no] = s
}

// 器件内部计算
// ---------- ----------
// 器件接收到事件, 触发对应器件对象计算
// 根据不同器件, 有不同的计算规则
// 计算完毕后根据属性发布相应事件
// no 接收器件, 对应针脚
// sourceName 发布事件的器件名称
// ---------- ----------
func (c *Component) Calculate(event evt.Event) {
	fmt.Println(fmt.Sprintf("IComponent [%s] -> not calculate ing .......... ......... ", c.name))
}

// 器件发布事件
// ---------- ----------
// event 事件类型定义
// sourceName 发布事件人
// no    事件触发的针脚号
// ---------- ----------
func (c *Component) Event(event, source string, sourceNo int, target string, targetNo int) {
	fmt.Println(fmt.Sprintf("接收事件, Component [%s:%d] Event [%s] ", c.name, targetNo, event))

	// 发布事件到事件管理器 -> [事件管理器根据器件唯一标识符查询器件, 并触发相联器件计算]
	c.man.Put(evt.Event{
		EventType: event,
		Source:    source,   // 发布事件器件
		SourceNo:  sourceNo, // 发布器件针脚
		Target:    target,   // 接收事件器件
		TargetNo:  targetNo, // 接收器件针脚
	})
}

// 器件描述信息, 返回JSON格式描述
// ---------- ----------
func (c *Component) Describe() string {
	var buff bytes.Buffer
	for _, v := range c.sts {
		var buffer bytes.Buffer
		// 针脚编号
		buffer.WriteString(fmt.Sprintf("\"%d\":[", v.GetNo()))
		for i := 0; i < len(v.Relation); i++ {
			// 器件名称:器件针脚编号,
			buffer.WriteString("\"")
			buffer.WriteString(v.Relation[i].Component.Name())
			buffer.WriteString("\"")
			buffer.WriteString(":")
			buffer.WriteString(strconv.Itoa(v.Relation[i].No))
			buffer.WriteString(",")
		}
		buffer.WriteString("]")
		// 关联器件名称, 关联器件编号
		buff.WriteString(fmt.Sprintf("%s],", strings.TrimRight(buffer.String(), ",]")))
	}
	return fmt.Sprintf("{%s}", strings.TrimRight(buff.String(), ","))
}
