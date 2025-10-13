//23:27
func garbageCollection(garbage []string, travel []int) int {
	collectionTime := 0
	glassTravelTime:= 0
	paperTravelTime:= 0
	metalTravelTime := 0
	currentTime := 0

	for i, garbageType := range garbage {
		if i!= 0 {
			currentTime += travel[i-1]
		}
		collectionTime += len(garbageType)
		for j := 0; j < len(garbageType); j++ {
			if garbageType[j] == 'M' {
				metalTravelTime = currentTime
			} else if  garbageType[j] == 'G' {
				glassTravelTime = currentTime
			} else if  garbageType[j] == 'P' {
				paperTravelTime = currentTime
			}
		}
	}
	result := collectionTime + glassTravelTime + paperTravelTime + metalTravelTime
	return result
}
