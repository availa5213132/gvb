package main

import "fmt"

// 优化的冒泡排序，添加了标志变量
func optimizedBubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // 标志变量，用于检测这一轮是否发生交换
		for j := 0; j < n-i-1; j++ {
			// 如果前一个元素比后一个元素大，交换它们
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 交换元素
				swapped = true                      // 发生了交换，设置为 true
			}
		}
		// 如果没有发生交换，说明数组已排序好，提前退出
		if !swapped {
			break
		}
	}
}

func main() {
	// 示例数组
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("排序前:", arr)

	// 调用优化的冒泡排序
	optimizedBubbleSort(arr)

	fmt.Println("排序后:", arr)
}
