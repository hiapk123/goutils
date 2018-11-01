package set

type HashSet struct {
	hashMap map[interface{}]bool
}

//初始化哈希集合
func (set *HashSet) init() *HashSet{
	set.hashMap = make(map[interface{}]bool)
	return set
}

func New()*HashSet{
	return new(HashSet).init()
}

//判断集合是否为空
func (set *HashSet) IsEmpty() bool {
	return len(set.hashMap) == 0
}

//获取集合中元素个数
func (set *HashSet) Size() int {
	return len(set.hashMap)
}

//向集合中添加元素
func (set *HashSet) Add(o interface{}) {
	set.hashMap[o] = true
}

//从集合中移除元素
func (set *HashSet) Remove(o interface{}) {
	delete(set.hashMap, o)
}

//判断集合中是否包含该元素
func (set *HashSet) Contains(o interface{}) bool {
	if _, key := set.hashMap[o]; key {
		return true
	}
	return false
}

//遍历集合获取所有元素
func (set *HashSet) Iterator() []interface{} {
	var (
		result = make([]interface{}, set.Size())
		index  = 0
	)
	for key := range set.hashMap {
		result[index] = key
		index++
	}
	return result
}

