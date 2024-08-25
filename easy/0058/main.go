func lengthOfLastWord(s string) int {
    input := strings.TrimSpace(s)
    words := strings.Split(input, " ")

    return len(words[len(words)-1])
}