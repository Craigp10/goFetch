package utils

func Average(array []float64) float64 {
	var sum float64
	var average float64

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	average = sum / float64(len(array))
	return average
}
