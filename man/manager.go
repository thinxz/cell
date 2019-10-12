package man

import (
	"../component"
	"../entity"
	"../evt"
	"fmt"
)

var (
	r1 entity.R
	r2 entity.R
	d1 entity.D
)

type Manager struct {
	// 实例存储器, 保存所有器件对象
	obj map[string]component.Component
	// 事件处理器, 管理所有事件
	eventManager *evt.EventManager
}

func (m *Manager) init() {
	// 初始化实例存储器
	m.obj = make(map[string]component.Component)

	// 初始化事件管理器
	m.eventManager = &evt.EventManager{}
	m.eventManager.Init()

	// 初始化, 所有元器对象 -> 并填充默认值(静态属性)
	r1 = entity.R{}
	r1.Init("R1", m.eventManager)
	r2 = entity.R{}
	r2.Init("R2", m.eventManager)
	d1 = entity.D{}
	d1.Init("D1", m.eventManager)

	// 保存器件对象
	m.addEntity(r1.Component)
	m.addEntity(r2.Component)
	m.addEntity(d1.Component)
}

func (m *Manager) initStitch() {
	// 初始化, 针脚连接
	r1.AddStitch(1, r2.Component, 2)
	r1.AddStitch(2, d1.Component, 2)

	r2.AddStitch(1, d1.Component, 1)
	r2.AddStitch(2, r1.Component, 1)

	d1.AddStitch(1, r2.Component, 1)
	d1.AddStitch(2, r1.Component, 2)
}

func (m *Manager) Start() {
	// 初始化并创建, 所有器件对象
	m.init()
	// 初始化, 所有器件的针脚数据
	m.initStitch()

	fmt.Println()
	// 加电自检电源初始化计算
	d1.InitCalculate()

	// 循环处理事件
	fmt.Println()
	event := m.eventManager.Push()
	e := m.getEntity(event.Name)
	fmt.Println(e.Describe())

	e.Calculate()
}

func (m *Manager) addEntity(c component.Component) {
	m.obj[c.Name()] = c
}

func (m *Manager) getEntity(name string) *component.Component {
	if v, ok := m.obj[name]; ok {
		return &v
	}
	return nil
}
