package service

import (
	"sort"
	"time"
)

const N = 20

var dataFromNakaiToRoppongi = [][]int{
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

var dataFromRoppongiToNakai = [][]int{
	{0, 2, 9, 18, 28, 37, 49},
	{5, 8, 20, 31, 44, 56},
	{6, 3, 11, 18, 25, 31, 37, 45, 52, 56},
	{7, 1, 5, 9, 14, 18, 23, 27, 30, 33, 36, 39, 42, 45, 48, 51, 55, 58},
	{8, 2, 6, 10, 14, 18, 22, 25, 29, 32, 36, 39, 43, 46, 50, 53, 57},
	{9, 0, 4, 7, 11, 15, 20, 25, 29, 34, 39, 45, 49, 54, 59},
	{10, 4, 9, 14, 19, 24, 29, 34, 39, 44, 50, 56},
	{11, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{12, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{13, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{14, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{15, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{16, 2, 8, 14, 20, 26, 32, 38, 44, 50, 56},
	{17, 2, 8, 14, 19, 25, 30, 35, 40, 46, 51, 56},
	{18, 1, 6, 11, 16, 20, 25, 29, 34, 38, 43, 47, 52, 56},
	{19, 1, 6, 11, 16, 21, 26, 31, 36, 41, 46, 51, 56},
	{20, 1, 6, 12, 18, 25, 31, 37, 43, 50, 56},
	{21, 3, 10, 17, 24, 31, 38, 46, 53},
	{22, 1, 9, 16, 23, 30, 36, 44, 51, 58},
	{23, 5, 12, 19, 26, 33, 40, 47, 54},
}

func GetTime() (int, int, int, int, int, int) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	nowHour := time.Now().In(loc).Hour()
	nowMinute := time.Now().In(loc).Minute()

	goWorkRes1, goWorkRes2 := calculate(nowHour, nowMinute, dataFromNakaiToRoppongi)
	backHomeRes1, backHomeRes2 := calculate(nowHour, nowMinute, dataFromRoppongiToNakai)

	return nowHour, nowMinute, goWorkRes1, goWorkRes2, backHomeRes1, backHomeRes2
}

func calculate(nowHour, nowMinute int, timetable [][]int) (int, int) {
	m := make(map[int][]int, N)
	for i := 0; i < N; i++ {
		inputSlice := timetable[i]
		m[inputSlice[0]] = inputSlice[1:]
	}

	firstTrainHour := timetable[1][0]
	firstTrainMinute := timetable[1][1]
	secondTrainMinute := timetable[1][2]

	var res1, res2 int
	v, ok := m[nowHour]
	if !ok {
		res1 = (firstTrainHour-nowHour)*60 + (firstTrainMinute - nowMinute)
		res2 = res1 + (secondTrainMinute - firstTrainMinute)
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
