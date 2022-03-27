package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum3(nums))
}

// 注意: 此题的边界是个很麻烦的问题。需要考虑清楚
func threeSum(nums []int) [][]int {
	// 1、先排序
	// 2、转换思路，将 a+b+c=0  看做为 a+b = -c
	// 3、双指针方式，跳过重复值
	rest := make([][]int, 0)
	n := len(nums)
	sort.Ints(nums)
	//fmt.Println(nums)
	for c := 0; c < n; c++ {
		// 上升排序，C>0 就可以结束了s
		if nums[c] > 0 {
			break
		}
		if c > 0 && nums[c] == nums[c-1] {
			continue
		}
		for i, j := c+1, n-1; i < n; {
			if i == j {
				break
			}
			if i > c+1 && nums[i] == nums[i-1] {
				i++
				continue
			}
			if nums[i]+nums[j]+nums[c] < 0 {
				i++
				continue
			} else if nums[i]+nums[j]+nums[c] > 0 {
				j--
			} else {
				//fmt.Println(nums[c], nums[i], nums[j])
				rest = append(rest, []int{nums[c], nums[i], nums[j]})
				i++
			}
		}
	}
	return rest
}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

// 注意: 此题的边界是个很麻烦的问题。需要考虑清楚
func threeSum3(nums []int) [][]int {

	rest := make([][]int, 0)
	n := len(nums)
	// 排序
	sort.Ints(nums)

	// 遍历 第一个数
	for i := 0; i < n-2; i++ {

		// 跳过i重复值
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			lo := i + 1
			hi := n - 1
			sum := 0 - nums[i]

			// 左边指针小于右边指针
			for lo < hi {

				// 相等
				if nums[lo]+nums[hi] == sum {

					// 保存结果
					rest = append(rest, []int{nums[i], nums[lo], nums[hi]})

					// 分别处理掉左右重复的数
					for lo < hi && nums[lo] == nums[lo+1] {
						lo++
					}
					for lo < hi && nums[hi] == nums[hi-1] {
						hi--
					}

					// 因为只有唯一解，故不可能，一边不动，另一边再匹配到不重复的数
					lo++
					hi--
				} else if nums[lo]+nums[hi] < sum {
					// 小于，说明需要更大的值进行匹配
					lo++
				} else {
					// 大于，说明需要更小的值匹配
					hi--
				}
			}
		}
	}
	return rest
}
