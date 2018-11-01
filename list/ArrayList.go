package list

const default_capacity = 20

type ArrayList struct {
	elementData []interface{}
	size        int
}

//指定值初始化
func (list *ArrayList) initWithCapacity(initialCapacity int) *ArrayList {
	if initialCapacity > 0 {
		list.elementData = make([]interface{}, initialCapacity)
	} else {
		panic("initialCapacity<=0!")
	}

	return list
}

//默认初始化
func (list *ArrayList) init() *ArrayList {
	list.elementData = make([]interface{}, default_capacity)
	return list
}

func (list *ArrayList) grow(minCapacity int) {
	oldCapacity := len(list.elementData)
	//fmt.Println(oldCapacity)
	newCapacity := oldCapacity + (oldCapacity >> 1)
	// fmt.Println(newCapacity)
	if newCapacity-minCapacity < 0 {
		newCapacity = minCapacity
	}

	newElementData := make([]interface{}, newCapacity)
	//copy old data to new array.
	for i := 0; i < len(list.elementData); i++ {
		newElementData[i] = list.elementData[i]
	}

	list.elementData = newElementData

}

//检查下标是否越界
func (list *ArrayList) rangeCheck(index int) {
	if index >= list.size {
		panic("下标越界异常！")
	}
}

//默认初始化
func New() *ArrayList {
	return new(ArrayList).init()
}

//指定值初始化
func NewList(cap int) *ArrayList {
	return new(ArrayList).initWithCapacity(cap)
}

//判断是否为空
func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

//判断是否为基础数据类型
func (list *ArrayList) IsBasicType(o interface{}) bool {
	switch o.(type) {
	case uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64, bool, string, uintptr, int, uint:
		return true
	}

	return false
}

//判断元素是否存在
func (list *ArrayList) IndexOf(o interface{}) int {
	if o == nil {
		for i := 0; i < list.size; i++ {
			if list.elementData[i] == nil {
				return i
			}
		}
	} else {
		//无法判断结构体是否相等
		if list.IsBasicType(o) {
			for i := 0; i < list.size; i++ {
				if o == list.elementData[i] {
					return i
				}
			}
		}

	}

	return -1
}

//判断是否包含该元素
func (list *ArrayList) Contains(o interface{}) bool {
	return list.IndexOf(o) >= 0
}

//从末尾查看是否存在该元素
func (list *ArrayList) LastIndexOf(o interface{}) int {
	if o == nil {
		for i := list.size - 1; i >= 0; i-- {
			if list.elementData[i] == nil {
				return i
			}
		}
	} else {
		//无法判断结构体是否相等
		if list.IsBasicType(o) {
			for i := list.size - 1; i >= 0; i-- {
				if o == list.elementData[i] {
					return i
				}
			}
		}
	}

	return -1
}

//获取指定元素
func (list *ArrayList) Get(index int) interface{} {
	list.rangeCheck(index)
	return list.elementData[index]
}

//设置指定位置处元素的值
func (list *ArrayList) Set(index int, element interface{}) interface{} {
	list.rangeCheck(index)
	oldValue := list.elementData[index]
	list.elementData[index] = element
	return oldValue
}

//添加元素
func (list *ArrayList) Add(element interface{}) bool {
	if list.size-len(list.elementData) > -2 {
		list.grow(list.size + 1)
	}
	list.elementData[list.size] = element
	list.size++
	return true
}

//删除指定元素
func (list *ArrayList) Remove(element interface{}) bool {

	var delIndex int = -1

	for i := 0; i < list.size; i++ {
		if element == list.elementData[i] {
			delIndex = i
		}
	}

	if delIndex != -1 {
		for i := delIndex + 1; i < list.size; i++ {
			list.elementData[i-1] = list.elementData[i]
		}

		list.size--
	}

	return true
}

//按照索引删除指定元素
func (list *ArrayList) RemoveByIndex(index int) bool {

	list.rangeCheck(index)
	for i := index; i < list.size; i++ {
		list.elementData[i] = list.elementData[i+1]
	}
	list.size--
	return true
}

//获取链表元素个数
func (list *ArrayList) Size() int {
	return list.size
}

func(list *ArrayList)Length()int{
	return len(list.elementData)
}

//清空元素
func (list *ArrayList) Clear() {
	for i := 0; i < list.size; i++ {
		list.elementData[i] = nil
	}

	list.size = 0
}
