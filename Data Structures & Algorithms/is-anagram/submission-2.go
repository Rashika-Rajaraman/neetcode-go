func isAnagram(s string, t string) bool {
     if len(s) != len(t) {
         return false
     }
     
     freqS, freqT := make(map[rune]int), make(map[rune]int)

     // HashMap
     // Time Complexity - O(n + m)
     // Space Complexity: O(k), 
     // where k is the number of unique characters. 
     // If the input is restricted to lowercase English letters, it becomes O(1).
     for i, ch := range s {
         freqS[ch]++
         freqT[rune(t[i])]++
     }

      for k, v := range freqS {
         if freqT[k] != v {
            return false
         }
      }
      return true
}
