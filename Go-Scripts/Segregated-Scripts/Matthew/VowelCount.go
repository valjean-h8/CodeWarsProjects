/*
Vowel Count
https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
*/
package kata
import (
  "strings" // for func Count
)
func GetCount(str string) (count int) {
  var x_value string = "aeiou" // all needed values 
  for i := 0; i < len(x_value); i++ {
   count += strings.Count(str, x_value[i:i+1]) // set our count value
 }
  return count
}