package component

import (
	"../evt"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

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

// 初始化器件针脚, 及其关联器件针脚
// ---------- ----------
func (c *Component) AddStitch(no int, target Component, targetNo int) {
	s, ok := c.sts[no]
	if !ok {
		s = Stitch{
			no:       no,
			Relation: [] *Relation{},
			Signal:   Signal{},
		}
	}

	s.AddRelation(&Relation{
		No:        targetNo,
		Component: target,
	})

	c.sts[no] = s
}

//// 信息传入接口 [被其他器件调用]
//// ---------- ----------
//// no      事件触发的针脚号
//// data    传入该器件的信息内容
//// ---------- ----------
//func (c *Component) Write(no int, signal Signal) error {
//	_, ok := c.sts[no]
//	if !ok {
//		return err.NewErr("针脚号错误, 已关闭或不存在")
//	}
//	return nil
//}

// 管理器接收到事件 -> 触发对应器件对象计算
// ---------- ----------
func (c *Component) Calculate() {
	fmt.Println(fmt.Sprintf("Component [%s] -> calculate ing .......... ......... ", c.name))
}

// 01 传递针脚数据
// 02 触发关联器件事件 [传递触发的针脚号]
// ---------- ---------- ----------
func (c *Component) Transmission() {
	// 计算完毕, 信息传递
	fmt.Println(fmt.Sprintf("信息传递 ing ... -> Component [%s] ", c.name))
	fmt.Println()
	// 遍历针脚数据, 并传递针脚计算值
	for _, v := range c.sts {
		for i := 0; i < len(v.Relation); i++ {
			// 器件 -> 关联其他器件针脚
			no := v.Relation[i].No
			// 01 传递针脚数据 [该器件计算的信号, 传递给关联的器件] [数据内联不用传输]
			//err := v.Relation[i].Component.Write(no, v.Signal)
			//if err != nil {
			//	fmt.Println(fmt.Sprintf("%s:%d -> %s", v.Relation[i].Component.Name(), no, err.Error()))
			//	//
			//}

			// 02 触发关联器件事件 [传递触发的针脚号]
			v.Relation[i].Component.event("DataChange", no)
		}
	}
}

// 器件描述信息, 返回JSON格式描述
// ---------- ----------
// event 事件类型定义
// no    事件触发的针脚号
// ---------- ----------
func (c *Component) event(event string, no int) {
	fmt.Println(fmt.Sprintf("接收事件, Component [%s:%d] Event [%s] ", c.name, no, event))

	// 发布事件到事件管理器 -> [事件管理器根据器件唯一标识符查询器件, 并触发相联器件计算]
	c.man.Put(evt.Event{
		EventType: event,
		Name:      c.name, // 接收事件人
		No:        no,
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
