// Named return parameter에 리턴값을 할당
package main

func sum(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
