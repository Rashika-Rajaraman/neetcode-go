func groupAnagrams(strs []string) [][]string {
     res := make(map[[26]int][]string)

     // Time Complexity = O(m * n)
     // Space Complexity = O(m * n)
     for _, str := range strs {
        var count [26]int
        for _, ch := range str {
            count[ch - 'a'] = count[ch - 'a'] + 1
        }
        res[count] = append(res[count], str)
     }

     var result [][]string
     for _, group := range res {
        result = append(result, group)
     }
     return result
}
