package simplearraysum

func SimpleArraySum(arr []int32) int32 {
	sum := int32(0)

	for _, v := range arr {
		sum += v
	}

	return sum
}
