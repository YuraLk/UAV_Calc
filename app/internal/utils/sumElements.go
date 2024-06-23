package utils

func SumElements(arr []float32) float32 {
	var sum float32
	for _, v := range arr {
		sum += v
	}
	return sum
}
