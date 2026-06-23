func hasDuplicate(nums []int) bool {
   seen := make(map[int]struct{})

   // HashMap 
   // Time Complexity - O(n)
   // Space Complexity - O(1)
   for _, num := range nums {
      if _, exists := seen[num]; exists {
         return true
      }
      seen[num] = struct{}{}
   }
   return false
}
