package kata

func GetCount(str string) (count int) {
	var words = [...]string{"a", "e", "i", "o", "u"} // array finds words
	var start int = 0                                // start index array
	var step int = 1                                 // step index array

	// loop for bulkhead elements
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(words); j++ {
			if str[start:step] == words[j] {
				count += 1
			}
		}

		// increment indexes
		start += 1
		step += 1
	}
	return count
}
