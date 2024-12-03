package utils

func MapOccurrences(input []int) map[int]int {
	occurrenceMapRight := make(map[int]int)

	for _, value := range input {
		if _, ok := occurrenceMapRight[value]; ok {
			occurrenceMapRight[value]++
		} else {
			occurrenceMapRight[value] = 1
		}
	}
	//fmt.Println(occurrenceMapRight)
	return occurrenceMapRight
}
