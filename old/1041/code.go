func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	direction := "north"

	directions := map[string][2]int{
		"north": {0, 1},
		"south": {0, -1},
		"east":  {1, 0},
		"west":  {-1, 0},
	}

	for _, letter := range instructions {
		switch letter {
		case 'G':
			x += directions[direction][0]
			y += directions[direction][1]
		case 'L':
			switch direction {
			case "north":
				direction = "west"
			case "south":
				direction = "east"
			case "east":
				direction = "north"
			case "west":
				direction = "south"
			}
		case 'R':
			switch direction {
			case "north":
				direction = "east"
			case "south":
				direction = "west"
			case "east":
				direction = "south"
			case "west":
				direction = "north"
			}
		}
	}

	return (x == 0 && y == 0) || direction != "north"
} 