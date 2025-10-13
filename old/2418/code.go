//14:54
func sortPeople(names []string, heights []int) []string {
	type Person struct {
		Name string
		Height int
	}
	var people []Person
	for i, name := range names {
		people = append(people, Person{Name: name, Height: heights[i]})
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})
	var sortedNames []string
	for _, person := range people {
		sortedNames = append(sortedNames, person.Name)
	}

	return sortedNames
}

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string)
	result := make([]string, len(names))

	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}

	sort.Sort(sort.Reverse(sort.IntSlice(heights)))
	for i := 0; i < len(names); i++ {
		result[i] = m[heights[i]]
	}

	return result
}
