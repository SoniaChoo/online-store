package service

import (
	"sort"
	"time"
)

const N = 20

var data = [][]int{
	{0, 4, 15, 31, 49},
	{5, 14, 28, 38, 51, 59},
	{6, 7, 15, 23, 30, 34, 39, 44, 50, 56},
	{7, 2, 8, 14, 19, 24, 29, 34, 38, 42, 46, 50, 53, 56, 59},
	{8, 2, 5, 8, 11, 14, 17, 20, 23, 26, 29, 32, 35, 38, 41, 44, 48, 51, 55, 58},
	{9, 2, 6, 10, 14, 18, 23, 27, 31, 36, 40, 44, 48, 52, 57},
	{10, 2, 6, 9, 14, 20, 26, 32, 38, 44, 50, 56},
	{11, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{12, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{13, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{14, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{15, 2, 8, 14, 20, 26, 32, 38, 44, 49, 53, 57},
	{16, 2, 8, 14, 19, 24, 29, 34, 39, 45, 50, 55},
	{17, 0, 5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55},
	{18, 0, 5, 10, 16, 22, 27, 32, 38, 44, 50, 56},
	{19, 2, 7, 12, 17, 22, 27, 31, 36, 41, 46, 52, 57},
	{20, 2, 7, 12, 16, 20, 26, 31, 36, 42, 46, 50, 54, 59},
	{21, 5, 11, 17, 23, 29, 35, 42, 48, 54},
	{22, 3, 11, 18, 26, 35, 43, 51, 58},
	{23, 8, 17, 27, 35, 44, 54},
}

func GetTime() (int, int, int, int) {
	m := make(map[int][]int, N)
	for i := 0; i < N; i++ {
		inputSlice := data[i]
		m[inputSlice[0]] = inputSlice[1:]
	}

	loc, _ := time.LoadLocation("Asia/Tokyo")
	nowHour := time.Now().In(loc).Hour()
	nowMinute := time.Now().In(loc).Minute()

	var res1, res2 int
	v, ok := m[nowHour]
	if !ok {
		res1 = (4-nowHour)*60 + (60 + 14 - nowMinute)
		res2 = res1 + (28 - 14)
		return res1, res2, nowHour, nowMinute
	}

	index := sort.SearchInts(v, nowMinute)
	nextHourIndex := (nowHour + 1) % 24
	if index == len(v) {
		res1 = m[nextHourIndex][0] - nowMinute + 60
		res2 = m[nextHourIndex][1] - nowMinute + 60
	} else if index == len(v)-1 {
		res1 = v[index] - nowMinute
		res2 = m[nextHourIndex][0] - nowMinute + 60
	} else {
		res1 = v[index] - nowMinute
		res2 = v[index+1] - nowMinute
	}
	return res1, res2, nowHour, nowMinute
}
