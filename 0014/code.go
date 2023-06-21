func longestCommonPrefix(strs []string) string {
    pref := strs[0]
    for i := 1; i<len(strs); i++ {
        for !strings.HasPrefix(strs[i], pref) {
            pref = pref[:len(pref)-1]
        }
    }
    return pref
}