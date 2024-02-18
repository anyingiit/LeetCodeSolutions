package main

// removeElement 函数的目标是原地移除数组nums中所有等于val的元素，并返回移除后数组的新长度。
// 该函数采用双指针技巧来实现原地修改，避免使用额外的数组空间。
//
// 思路概述：
// - 使用两个指针：i和J。i作为外层循环的索引，遍历整个数组；J用于寻找不等于val的元素，以便替换等于val的元素。
// - 遍历数组：对于每个元素，检查它是否等于val。
//   - 如果等于val，使用J指针寻找下一个不等于val的元素。找到后，将该元素移动到当前i指针的位置，并将J位置的元素标记为已处理。
//   - 如果不等于val，表示当前元素是需要保留在数组中的，直接增加最终数组的长度计数。
//
// - 通过这种方式，可以确保所有等于val的元素被移除，同时保留下来的元素填充在数组的前面，后面的元素不做考虑。
// - 函数最终返回新的数组长度，即保留下来的元素数量。
func removeElement(nums []int, val int) int {
	var length, J int
	length = 0 // 最终长度
	J = 0      // i后正在审查的当前元素

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			// 确保J指向i之后的位置
			if J <= i {
				J = i + 1
			}
			// 寻找第一个不等于val的元素
			for ; J < len(nums); J++ {
				if nums[J] != val {
					break
				}
			}

			// 如果J超出数组范围，结束循环
			if J > len(nums)-1 {
				break
			} else {
				// 替换nums[i]为找到的不等于val的元素，增加长度
				nums[i] = nums[J]
				length++

				// 将处理过的元素标记为val，并移动J到下一个位置
				nums[J] = val
				J++
			}

		} else {
			// 当前元素不等于val，直接增加长度
			length++
		}
	}

	return length
}
