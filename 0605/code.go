//06:24
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; n > 0 && len(flowerbed)-i >= n; {
		if flowerbed[i] == 1 {
			i++
		} else if (i == 0 || flowerbed[i-1] == 0 ) && (i+1 >= len(flowerbed) || flowerbed[i+1] == 0) {
			i++
			n--
		}
		i++
	}

	return n == 0
}