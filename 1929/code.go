func getConcatenation(nums []int) []int {
copyOfNums := nums

for _, num := range nums {
copyOfNums = append(copyOfNums, num)
}

return copyOfNums
}