package selectactivities

// finishTime has to be ordered 
func RecursiveActivitySelection(startTime, finishTime []int, numberOfActivities int) []int {
	i := 0

	activities := make([]int, 0, numberOfActivities)
	// always select the first activity
	activities = append(activities, i)

	for j := 1; j < numberOfActivities; j++ {
		if startTime[j] > finishTime[i] {
			activities = append(activities, j)
			i = j
		}
	}

	return activities
}
