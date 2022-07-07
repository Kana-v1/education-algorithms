/*
 * @lc app=leetcode id=1344 lang=golang
 *
 * [1344] Angle Between Hands of a Clock
 */

// @lc code=start
package main

import "math"

func angleClock(hour int, minutes int) float64 {
	if hour == 12 {
		hour = 0
	}

	if minutes == 60 {
		minutes = 0
	}

	var hours float64 = 12
	var degree float64 = 360

	degreeInHour := degree / hours
	degreeInMinute := degree / 60
	hourDegreePerMinute := degreeInHour / 60

	minutesDegree := float64(minutes) * degreeInMinute
	hoursDegree := float64(hour)*degreeInHour + float64(minutes)*hourDegreePerMinute

	angle := math.Abs(minutesDegree - hoursDegree)

	if angle < degree/2 {
		return angle
	}

	return degree - angle
}

// @lc code=end
