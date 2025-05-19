/*
Vowel Count
https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
*/

package kata

func GetCount(str string) (count int) {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true} // Creates a map (dictionary) where the keys are vowels (a, e, i, o, u), and the values are true.
	for _, char := range str {                                                     // Cycle runs on every character (char) in the string str
		if vowels[char] {
			count++
		}
	}
	return
}
