package Array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func Constructor(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

func GetArrayFromArr(arr []interface{}) *Array {
	a := &Array{
		data: make([]interface{}, len(arr)),
		size: len(arr),
	}
	for i := 0; i < len(arr); i++ {
		a.data[i] = arr[i]
	}

	return a
}

// 获取数组的容量
func (this *Array) GetCapacity() int {
	return len(this.data)
}

// 获得数组中的元素个数
func (this *Array) GetSize() int {
	return this.size
}

// 返回数组是否为空
func (this *Array) IsEmpty() bool {
	return this.size == 0
}

// 在第 index 个位置插入一个新元素 e
func (this *Array) Add(index int, e interface{}) {
	if index < 0 || index > this.GetCapacity() {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if this.size == len(this.data) {
		this.resize(2 * this.size)
	}

	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = e
	this.size++
}

// 向所有元素后添加一个新元素
func (this *Array) AddLast(e interface{}) {
	this.Add(this.size, e)
}

// 向所有元素前添加一个新元素
func (this *Array) AddFirst(e interface{}) {
	this.Add(0, e)
}

// 获取 index 索引位置的元素
func (this *Array) Get(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

// 修改 index 索引位置的元素
func (this *Array) Set(index int, e interface{}) {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}

// 查找数组中是否有元素 e
func (this *Array) Contains(e interface{}) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (this *Array) Find(e interface{}) int {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (this *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (this *Array) Remove(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Remove failed,Index is illegal.")
	}

	e := this.data[index]
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
	this.data[this.size] = nil //loitering object != memory leak

	// 考虑边界条件，避免长度为 1 时，resize 为 0
	if this.size == len(this.data)/4 && len(this.data)/2 != 0 {
		this.resize(len(this.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (this *Array) RemoveFirst() interface{} {
	return this.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (this *Array) RemoveLast() interface{} {
	return this.Remove(this.size - 1)
}

// 从数组中删除一个元素 e
func (this *Array) RemoveElement(e interface{}) bool {
	index := this.Find(e)
	if index == -1 {
		return false
	}

	this.Remove(index)
	return true
}

// 从数组中删除所有元素 e
func (this *Array) RemoveAllElement(e interface{}) bool {
	if this.Find(e) == -1 {
		return false
	}

	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			this.Remove(i)
		}
	}
	return true
}

// 为数组扩容
func (this *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[i]
	}

	this.data = newData
}

func (a *Array) Swap(i int, j int) {
	if i < 0 || i >= a.size || j < 0 || j >= a.size {
		panic("Index is illegal.")
	}
	a.data[i], a.data[j] = a.data[j], a.data[i]
}

// 重写 Array 的 string 方法
func (this *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", this.size, len(this.data)))
	buffer.WriteString("[")
	for i := 0; i < this.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(this.data[i]))
		if i != (this.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
