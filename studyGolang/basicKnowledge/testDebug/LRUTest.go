package main

import "fmt"


// 手撸双向链表 + 自带的map哈希
type LRUCache struct {
    Capacity int //最大容量
    Size int    // 当前大小
    cache map[int]*Node //哈希标识 node在链表中出现的位置,使查找复杂度为O(1)
    head,tail *Node // 头尾哑结点
}

type Node struct {
    key,value int   //存储实际的kv
    pre,next *Node  // 因为是双向链表，所以有前后节点
}

func NewNode(key int,value int) *Node {
    return &Node{
        key:key,
        value:value,
        pre:nil,
        next:nil,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        Capacity:capacity,
        Size:0,
        cache:map[int]*Node{},
        head:NewNode(0,0),
        tail:NewNode(0,0),
    }
    l.head.next = l.tail
    l.tail.pre = l.head
	fmt.Println("LRU constructor success,capacity = ",capacity,"size = ",0)
    return l
}


func (this *LRUCache) Get(key int) int {
    // 如果key不在cache中，直接返回-1
    if _,ok := this.cache[key];!ok {
		fmt.Println("LRU has not found key = ",key)
        return -1
    }
    // 否则需要更新该节点在链表中的位置,即放在队头
    node := this.cache[key]
    this.MoveToHead(node)
	fmt.Println("[Get] key = ",key,"value = ",node.value)
    return node.value
}


func (this *LRUCache) Put(key int, value int)  {
    // 如果链表中没有该节点
    if _,ok := this.cache[key];!ok {
        node := NewNode(key,value)
        this.cache[key] = node
        this.AddNodeToHead(node)
        this.Size++
        // 检查链表还有无位置
        if this.Size > this.Capacity {
            this.RemoveTail()
            this.Size--
        }
    } else {
        // 如果有节点，则需要更新value并把它移动到队头即可
        node := this.cache[key]
        node.value = value
        this.MoveToHead(node)
    }
	fmt.Println("[Put] key = ",key,"value = ",value," success,now size = ",this.Size)
}

// 把节点移动到队头
func (this *LRUCache) MoveToHead(node *Node) {
    // 把节点删除
    this.RemoveNode(node)
    // 重新添加节点到队头
    this.AddNodeToHead(node)
}


// 删除节点
func (this *LRUCache) RemoveNode(node *Node) {
    // 操作指针即可
    node.pre.next = node.next
    node.next.pre = node.pre
}

// 添加节点到队头
func (this *LRUCache) AddNodeToHead(node *Node) {
    // 注意操作顺序
    node.pre = this.head
    node.next = this.head.next
    this.head.next.pre = node
    this.head.next = node
}

// 删除链表中的节点
func (this *LRUCache) RemoveTail(){
    node := this.tail.pre
    // 1.先操作指针删掉它
    this.RemoveNode(node)
    // 2.释放空间,从cache中删掉它
    delete(this.cache,node.key)
}
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	lrucache := Constructor(5)
    lrucache.Put(1,1)
    lrucache.Get(2)
}