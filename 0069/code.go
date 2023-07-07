func mySqrt(x int) int {
    sqrt := 0

	for sqrt*sqrt < x {
		sqrt++
	}

	if sqrt*sqrt > x {
		sqrt--
	}

    return sqrt
}