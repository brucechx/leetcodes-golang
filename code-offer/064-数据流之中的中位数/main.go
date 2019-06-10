package _64_数据流之中的中位数

type DynamicArray struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

func NewDynamicArray() *DynamicArray{
	return &DynamicArray{maxHeap:new(MaxHeap), minHeap:new(MinHeap)}
}

func (d *DynamicArray)insert(num int){
	// 已经有偶数个数据了（可能没有数据）
	// 数据总数是偶数个时把新数据插入到小堆中
	if (d.minHeap.Len() + d.maxHeap.Len()) %2 == 0{
		// 大堆中有数据，并且插入的元素比大堆中的元素小
		if d.maxHeap.Len() > 0 && num < d.maxHeap.GetTop().(int){
			// 将num加入的大堆中去
			Push(d.maxHeap, num)
			// 删除堆顶元素，大堆中的最大元素
			num = Pop(d.maxHeap).(int)
		}
		// num插入到小堆中，当num小于大堆中的最大值进，
		// num就会变成大堆中的最大值，见上面的if操作
		// 如果num不小于大堆中的最大值，num就是自身
		Push(d.minHeap, num)
	}else { // 数据总数是奇数个时把新数据插入到大堆中
		// 小堆中有数据，并且插入的元素比小堆中的元素大
		if d.minHeap.Len() > 0 && num > d.minHeap.GetTop().(int){
			// 将num加入的小堆中去
			Push(d.minHeap, num)
			// 删除堆顶元素，小堆中的最小元素
			num = Pop(d.minHeap).(int)
		}
		// num插入到大堆中，当num大于小堆中的最小值进，
		// num就会变成小堆中的最小值，见上面的if操作
		// 如果num不大于大堆中的最小值，num就是自身
		Push(d.maxHeap, num)
	}
}

func (d *DynamicArray) getMedian() float32{
	size := d.maxHeap.Len() + d.minHeap.Len()
	if size & 1 == 1{ // 奇数个
		return float32(d.minHeap.GetTop().(int))
	}else {
		return float32(d.maxHeap.GetTop().(int) + d.minHeap.GetTop().(int))/float32(2.0)
	}
}