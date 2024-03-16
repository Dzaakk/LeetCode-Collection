//07:59
func destCity(paths [][]string) string {
	sources := make(map[string]struct{})
	for i := range paths {
		sources[paths[i][0]] = struct {}{}
	}

	for i := range paths {
		if _, ok := sources[paths[i][1]]; !ok{
			return paths[i][1]
		}
	}
	return ""
}
