package service

import (
	"sort"
	"time"
)

const N = 20

var data = [][]int{
	{0, 0, 7, 14, 22, 30, 38, 49, 56},
	{5, 10, 28, 41, 50},
	{6, 3, 15, 22, 30, 37, 44, 50, 56},
	{7, 4, 11, 16, 20, 25, 29, 34, 38, 43, 47, 50, 53, 56, 59},
	{8, 2, 5, 8, 11, 15, 18, 22, 26, 30, 34, 38, 42, 45, 49, 52, 56, 59},
	{9, 3, 7, 10, 14, 17, 21, 24, 28, 32, 36, 41, 45, 50, 54, 59},
	{10, 5, 10, 14, 19, 25, 29, 34, 39, 44, 49, 54, 59},
	{11, 5, 11, 17, 22, 28, 33, 39, 45, 51, 57},
	{12, 3, 9, 15, 21, 27, 33, 39, 45, 51, 57},
	{13, 3, 9, 15, 21, 27, 33, 39, 45, 51, 57},
	{14, 3, 9, 15, 21, 27, 33, 39, 45, 51, 57},
	{15, 3, 9, 15, 21, 27, 33, 39, 45, 51, 57},
	{16, 3, 9, 15, 21, 27, 33, 39, 45, 51, 57},
	{17, 3, 9, 15, 21, 27, 33, 39, 46, 51, 56},
	{18, 1, 6, 12, 17, 22, 27, 32, 36, 41, 46, 51, 54, 59},
	{19, 3, 8, 12, 17, 22, 27, 32, 37, 43, 48, 53, 57},
	{20, 2, 7, 12, 17, 22, 28, 33, 40, 46, 52, 58},
	{21, 5, 11, 17, 25, 32, 39, 45, 51, 59},
	{22, 6, 14, 22, 30, 37, 44, 50, 58},
	{23, 6, 13, 18, 25, 32, 39, 46, 54},
}

func GetTime() (int, int) {
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
		res1 = (4-nowHour)*60 + (60 + 10 - nowMinute)
		res2 = res1 + (28 - 10)
		return res1, res2
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
	return res1, res2
}
