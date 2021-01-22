package lib

func BinarySearch(key int, arr []int) int {
	lo := 0
	hi := len(arr) - 1
	for lo <= hi {
		// 被查的键要么不存在要么必然存在 arr[lo...hi]之间
		// key 5   arr [1,2,3,4,5,6,7,8]
		// 1 mid = 0(lo) + 7 / 2
		// 2 mid = 4(lo) + 3 / 2
		mid := lo + (hi-lo)/2
		if key < arr[mid] {
			hi = mid - 1
		} else if key > arr[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
