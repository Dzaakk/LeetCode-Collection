func numberOfSteps(num int) int {
    out := 0
    for {
        if num ==0 {
            break
        }else if num %2 ==0{
            num /= 2
            out += 1
        } else {
            num -= 1
            out += 1
        }
    }
    return out
}