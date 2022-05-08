package exercise

// 长度为n的数组中,其中元素小于等于n切大于0,相同元素最多出现两次
//找出数组中出现2次的元素
//要求:时间复杂度O(n),空间复杂度O(1)

/*
解题思路:数组元素小于等于n条件，可以联想到原地hash操作
即将元素放到对应下标位置,当元素和下标不同时,即此数组出现多次
由于数组下标为0到n-1,因此比较是nums[i]和nums[nums[i]-1]
若nums[i]不等于nums[nums[i]-1],则交换
示例:3,4,2,3,1中找出现两次的元素

原始数据:3,4,2,3,1
第一趟(i=0):nums[0]=3,nums[3-1]=2,交换后 2,4,3,3,1
           nums[0]=2,nums[2-1]=4,交换后 4,2,3,3,1
		   nums[0]=4,nums[4-1]=1,交换后 1,2,3,3,4
		   nums[0]=1,nums[1-1]=1, 第一趟结束
第二趟(i=1):nums[1]=2,nums[2-1]=2, 第二趟结束
第三趟(i=2):nums[2]=3,nums[3-1]=3, 第三趟结束
第四趟(i=3):nums[3]=3,nums[3-1]=3, 第四趟结束
第五趟(i=4):nums[4]=4,nums[4-1]=3, 交换后 1,2,3,4,3
			nums[4]=3,nums[3-1]=3,第五躺结束
最终得到数组:1,2,3,4,3
最后遍历数组,nums[i]-1!=i的即为重复元素
*/

func FindDuplicates(nums []int) []int {
	var res []int
	for index := range nums {
		for {
			if nums[index] == nums[nums[index]-1] {
				break
			}
			nums[index], nums[nums[index]-1] = nums[nums[index]-1], nums[index]
		}
	}

	for index := range nums {
		if nums[index]-1 != index {
			res = append(res, nums[index])
		}
	}
	return res
}
