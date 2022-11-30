package sort

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

//Len()
func (s StuScores) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

//Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func Sort1StuScores1() {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	// 打印未排序的 stus 数据
	fmt.Println("Default:\n\t", stus)
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)

	sort.Sort(sort.Reverse(stus))
	fmt.Println(stus)

}

/*
Default:
	 [{alan 95} {hikerell 91} {acmfly 96} {leao 90}]
IS Sorted?
	 true
Sorted:
	 [{leao 90} {hikerell 91} {alan 95} {acmfly 96}]
[{acmfly 96} {alan 95} {hikerell 91} {leao 90}]
*/
func Sort1Search1() {
	x := 11
	s := []int{3, 6, 8, 11, 45, 50} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool {
		fmt.Println("i:", i)
		return s[i] >= x
	})
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
}

/*
11  在 s 中的位置为： 3
*/

func Sort1GuessingGame1() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
