package component

import (
	"../evt"
	"bytes"
	"fmt"
	"strings"
)

// 器件接口定义
// ---------- ---------
type IComponent interface {
	Transmission(event evt.Event)                          // 事件传输到该器件, 触发器发生计算
	Stitch(no int) (*Stitch, bool)                         // 获取器件对应针脚
	GetComponentStitch(name string, no int) (Stitch, bool) // 获取对应器件, 对应针脚数据
	Name() string                                          // 获取器件名称, 唯一标识符
	Describe() string                                      // 器件相关信息描述
}

//  事件传输到该器件, 触发器发生计算 -> 器件内部计算
// ---------- ----------
// 器件接收到事件, 触发对应器件对象计算
// 根据不同器件, 有不同的计算规则
// 计算完毕后根据属性发布相应事件
// no 接收器件, 对应针脚
// sourceName 发布事件的器件名称
// ---------- ----------
func (c *Component) Transmission(event evt.Event) {
	fmt.Println(fmt.Sprintf("IComponent [%s] -> not calculate ing .......... ......... ", c.name))
}

// 根据器件名称, 获取对应器件值
// ---------- ----------
// 值获取, 不能影响获取的其他器件
// ---------- ----------
func (c *Component) getComponent(name string) (IComponent, bool) {
	if v, ok := c.man.GetComponent(name); ok {
		return v, ok
	}
	return nil, false
}

// 获取自身器件针脚数据
// ---------- ----------
func (c *Component) Stitch(no int) (*Stitch, bool) {
	s, ok := c.sts[no]
	return s, ok
}

// 获取对应器件, 对应针脚数据
// 不能对其他器件数据直接进行修改
// ---------- ----------
// name 获取的器件名称
// no   器件对应的针脚号
// ---------- ----------
func (c *Component) GetComponentStitch(name string, no int) (Stitch, bool) {
	if t, ok := c.getComponent(name); ok {
		if s, ok := t.Stitch(no); ok {
			return *s, ok
		}
	}
	return Stitch{}, false
}

// 器件发布事件
// ---------- ----------
// event    事件类型
// source   发布器件名称
// sourceNo 发布器件针脚
// target   接收器件名称
// targetNo 接收器件针脚
// ---------- ----------
func (c *Component) Event(event, source string, sourceNo int, target string, targetNo int) {
	c.man.PutEvent(event, source, sourceNo, target, targetNo)
}

// 获取器件名称, 唯一标识符
// ---------- ----------
func (c *Component) Name() string {
	return c.name
}

// 器件描述信息, 返回JSON格式描述
// ---------- ----------
func (c *Component) Describe() string {
	var buff bytes.Buffer
	for _, v := range c.sts {
		var buffer bytes.Buffer
		// 针脚编号
		buffer.WriteString(fmt.Sprintf("\"%d\":[", v.GetNo()))
		//for i := 0; i < len(v.Relation); i++ {
		//	// 器件名称:器件针脚编号,
		//	buffer.WriteString("\"")
		//	buffer.WriteString(v.Relation[i].Component.Name())
		//	buffer.WriteString("\"")
		//	buffer.WriteString(":")
		//	buffer.WriteString(strconv.Itoa(v.Relation[i].No))
		//	buffer.WriteString(",")
		//}
		buffer.WriteString("]")
		// 关联器件名称, 关联器件编号
		buff.WriteString(fmt.Sprintf("%s],", strings.TrimRight(buffer.String(), ",]")))
	}
	return fmt.Sprintf("{%s}", strings.TrimRight(buff.String(), ","))
}

/********** ********** ********** ********** ********** ********** ********** ********** **********/

// 器件定义
// ---------- ---------
type Component struct {
	name string          // 器件名称
	sts  map[int]*Stitch // 针脚定义
	man  *EnManager      // 器件管理器
}

// 初始化器件数据
// ---------- ----------
func (c *Component) Init(name string, man *EnManager) {
	c.sts = make(map[int]*Stitch)
	c.name = name
	c.man = man
}

// 初始化器件针脚, 及其关联器件针脚
// ---------- ----------
func (c *Component) AddStitch(no int, target IComponent, targetNo int) {
	s, ok := c.sts[no]
	if !ok {
		s = &Stitch{
			no:       no,
			Relation: make(map[string]*Relation),
			Signal:   evt.Signal{},
		}
	}

	s.AddRelation(&Relation{
		name: target.Name(),
		no:   targetNo,
	})

	c.sts[no] = s
}

/********** ********** ********** ********** ********** ********** ********** ********** **********/

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

//func (c *Component) Real(child interface{}) {
//	ref := reflect.ValueOf(child)
//	method := ref.MethodByName("Calculate")
//	if method.IsValid() {
//		// 执行方法
//		v := method.Call(make([]reflect.Value, 0))
//		fmt.Print(v)
//	} else {
//		// 错误处理
//	}
//}

//// 01 传递针脚数据
//// 02 触发关联器件事件 [传递触发的针脚号]
//// ---------- ---------- ----------
//func (c *Component) Transmission() {
//	// 计算完毕, 信息传递
//	fmt.Println(fmt.Sprintf("信息传递 ing ... -> Component [%s] ", c.name))
//	fmt.Println()
//	// 遍历针脚数据, 并传递针脚计算值
//	for _, v := range c.sts {
//		if v.Transmission {
//			for i := 0; i < len(v.Relation); i++ {
//				// 器件 -> 关联其他器件针脚
//				no := v.Relation[i].No
//				// 01 传递针脚数据 [该器件计算的信号, 传递给关联的器件] [数据内联不用传输]
//				//err := v.Relation[i].Component.Write(no, v.Signal)
//				//if err != nil {
//				//	fmt.Println(fmt.Sprintf("%s:%d -> %s", v.Relation[i].Component.Name(), no, err.Error()))
//				//	//
//				//}
//
//				// 02 触发关联器件事件 [传递触发的针脚号]
//				v.Relation[i].Component.event("DataChange", no)
//			}
//		}
//	}
//}
