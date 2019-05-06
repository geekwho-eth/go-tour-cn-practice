package main

import (
	"fmt"
	"strings"
)

// List[T] 表示一个泛型单链表节点，可以保存任何类型 T 的值
type List[T any] struct {
	val  T        // 节点的值
	next *List[T] // 指向下一个节点的指针
}

// Append 向链表末尾追加一个元素
func (l *List[T]) Append(value T) {
	current := l              // 从当前节点开始遍历
	for current.next != nil { // 找到最后一个节点
		current = current.next
	}
	current.next = &List[T]{val: value} // 在末尾添加新节点
}

// Prepend 在链表头部插入一个元素，返回新的头结点
func Prepend[T any](head *List[T], value T) *List[T] {
	newNode := &List[T]{val: value, next: head} // 创建新节点，其 next 指向原来的头
	return newNode                              // 返回新的头结点
}

// Len 返回链表的长度
func (l *List[T]) Len() int {
	count := 1 // 初始为 1（当前节点）
	for current := l.next; current != nil; current = current.next {
		count++ // 遍历链表，累加节点数
	}
	return count // 返回节点总数
}

// ToSlice 将链表转换为切片
func (l *List[T]) ToSlice() []T {
	var result []T // 初始化一个泛型切片
	current := l   // 从当前节点开始遍历
	for current != nil {
		result = append(result, current.val) // 追加每个节点的值到切片中
		current = current.next               // 移动到下一个节点
	}
	return result // 返回转换后的切片
}

// String 打印链表的内容，实现 fmt.Stringer 接口
func (l *List[T]) String() string {
	var sb strings.Builder // 使用字符串构建器提高效率
	sb.WriteString("[")    // 链表输出起始符
	current := l
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.val)) // 写入当前节点值
		if current.next != nil {
			sb.WriteString(" -> ") // 如果不是最后一个，加分隔符
		}
		current = current.next // 移动到下一个节点
	}
	sb.WriteString("]") // 链表输出结束符
	return sb.String()  // 返回构建好的字符串
}

func main() {
	// 创建一个初始链表节点，值为 1
	list := &List[int]{val: 1}

	// 在链表末尾添加 2 和 3
	list.Append(2)
	list.Append(3)

	// 在链表头部插入 0，list 重新指向新的头节点
	list = Prepend(list, 0)

	// 打印链表内容
	fmt.Println("链表内容：", list)

	// 打印链表长度
	fmt.Println("链表长度：", list.Len())

	// 打印链表转换为切片的形式
	fmt.Println("切片形式：", list.ToSlice())
}
