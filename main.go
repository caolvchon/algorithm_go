package main

import (
	_ "container/list"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
	"time"
)

func main1() {
	//res:=CombinationSum([]int{2,3,5},8)
	//fmt.Println(res)
	//res:=combinationSum2([]int{10,1,2,7,6,1,5},8)
	//fmt.Println(res)

	//a:="我是"
	//for i,j:=range a{
	//	fmt.Println(i,j)
	//	if unicode.IsSpace(j){
	//		fmt.Println(rune(j))
	//	}
	//}

	//const placeOfInterest = `⌘`
	//
	//fmt.Printf("plain string: ")
	//fmt.Printf("%s", placeOfInterest)
	//fmt.Printf("\n")
	//
	//fmt.Printf("quoted string: ")
	//fmt.Printf("%+q", placeOfInterest)
	//fmt.Printf("\n")
	//
	//fmt.Printf("hex bytes: ")
	//for i := 0; i < len(placeOfInterest); i++ {
	//	fmt.Printf("%x ", placeOfInterest[i])
	//}
	//fmt.Printf("\n")

	//s := "我是中国人"
	//for index, runeValue := range s {
	//	fmt.Printf("%#U 起始于字位置%d\n", runeValue, index)
	//}
	//fmt.Printf("%#U",'我')
	//fmt.Println()
	//fmt.Println(utf8.RuneCountInString("我是abc"))
	//fmt.Println(len("我是abc"))

	//fmt.Println(utf8.DecodeRuneInString("abc"))
	//
	//
	//res:=replaceSpaceS("We are happy")
	//fmt.Println(res)
	//
	//ret:= replaceSpaceM("We are happy")
	//fmt.Println(ret)
	//
	//a:=[]int{1,2,3,4}
	//b:=[]int{7,8,9,0}
	//fmt.Println(copy(a[1:3],b[1:3]))
	//fmt.Printf("%v %v",a,b)
	//fmt.Println()
	//c:=[]byte("abcd")
	//d:=[]byte("1234")
	//copy(c[2:3],d[0:1])
	//fmt.Printf("%s %s",c,d)

	//s := "We are happy"
	//fmt.Printf("%+v", s[0:1]+"a")
	//fmt.Println(reflect.TypeOf(s[2]))
	//res := replaceSpace("We are happy")
	//fmt.Println(res)
	//
	//fmt.Println(math.MaxInt32)
	//fmt.Println(math.MinInt32)
	//fmt.Println(math.MinInt32 / 10)
	//fmt.Println(7/10, 5/10, 3/10)
	//
	//ss := "abc"
	//fmt.Println(reflect.TypeOf('a'))
	//fmt.Println(ss[0] - 'b')
	//fmt.Println(ss[0] >= 'b')
	//
	//fmt.Println(strToInt("words and 987"))
	//fmt.Println(strToInt("  -0012a42"))

	//s := "fafafad"
	//fmt.Println(reflect.TypeOf(s[1:3]))
	//a := []int{1, 2}
	//fmt.Println(a[0:0])
	//
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	//s := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//s = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	////s = [][]int{{2, 5}, {8, 4}, {0, -1}}
	//ret := SpiralOrder(s)
	//fmt.Println(ret)

	c := false
	fmt.Println(!c)
	a := make(map[int]int)
	fmt.Println(a[1])

	m := make([][]byte, 10)
	for i, _ := range m {
		m[i] = make([]byte, 10)
	}
	fmt.Println(m)
	fmt.Println(m[3][3])

	aa := []A{
		{1}, {2}, {3},
	}
	sort.Slice(aa, func(i, j int) bool {
		return aa[i].a < aa[j].a
	})
	fmt.Println(aa)

	fmt.Println(math.Pow(10, 10))

	ccc := []int{1}
	for i := 1; i < len(ccc); i++ {
		fmt.Println("azzz", ccc[i])
	}

}

func main2() {
	//for i := 0; i < 10; i++ {
	//	go tt()
	//}
	//time.Sleep(5)

	//channel

	/*e := entry{
		ready: make(chan struct{}),
		res:   0,
	}

	go func() {
		if e.ready == nil {
			e.ready = make(chan struct{})
		}
		time.Sleep(time.Second * 3)
		close(e.ready)

	}()

	go func() {
		fmt.Println("func1 ")
		res, ok := <-e.ready
		fmt.Println("func1 res=", res, "ok=", ok)
	}()
	go func() {
		fmt.Println("func2 ")
		res, ok := <-e.ready
		fmt.Println("func2 res=", res, "ok=", ok)
	}()
	time.Sleep(time.Second * 5)
	*/

	//管道pipeline
	/*naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(naturals, squares)
	printer(squares)*/

	//fmt.Println(time.Now())
	//t := time.Tick(time.Second)
	//
	//for i := range t {
	//	fmt.Println(time.Now())
	//	fmt.Println(i)
	//}

	//lengthOfLongestSubstring("pwwkew")

	//fmt.Println(1 << 1)

	//defer func() {
	//	fmt.Println("main defer end")
	//	if err := recover(); err != nil {
	//		fmt.Println("main recover", err)
	//	}
	//}()
	//go panicTest()
	//time.Sleep(time.Second)
	//fmt.Println("main end")

	//&^ 是什么
	//fmt.Println(1 &^ 1)
	//fmt.Println(1 &^ 0)
	//fmt.Println(0 &^ 1)
	//fmt.Println(0 &^ 0)

	a := []int{1, 2, 3, 4, 5}
	s1 := a[:3]
	s2 := a[:3:3]
	s1 = append(s1, 100)
	s2 = append(s2, 101)
	fmt.Println(a, s1, s2)

}

func main() {
	//x := []int{25, 27, 1, 10, 8, 35, 17, 5, 4, 16}
	//fmt.Println(minimumRemoval(x))
	//arr := strings.Split("1.", ".")
	//
	//fmt.Printf("%#v", arr)
	//arr2 := strings.Split("1", ".")
	//
	//fmt.Println(arr2)

	x := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	a := x[3:4]
	b := x[0:5]

	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
	a = append(a, 1111)
	a[0] = 22222
	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)

	fmt.Println(mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift)

	fmt.Println(0 &^ 0)
	fmt.Println(0 &^ 1)
	fmt.Println(1 &^ 0)
	fmt.Println(1 &^ 1)
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

//func isCorrent(s string) bool {
//	s = strings.Trim(s, " ")
//	if len(s) == 0 {
//		return false
//	}
//	arr := strings.Split(s, "e")
//
//	if len(arr) > 2 {
//		return false
//	}
//	if len(arr) == 1 {
//		return is
//	}
//	if len(arr) == 2 {
//
//	}
//}

func isNumber() {

}

func isInteger() {

}

func isFloat() {

}

func minimumRemoval(beans []int) int64 {
	//暴力法吧
	m := make(map[int]int)
	sort.Ints(beans)
	for _, j := range beans {
		if v, ok := m[j]; ok {
			m[j] = v + 1
		} else {
			m[j] = 1
		}
	}
	minBeans := 0

	tfBeans := 0
	packageNum := len(beans)
	for i := 0; i < len(beans); {
		x := beans[i]
		if i == 0 {
			//初始化
			for i := 0; i < len(beans); i++ {
				if beans[i] > x {
					tfBeans += (beans[i] - x)
				}
			}
			i += (m[x])
			continue
		}
		lastX := beans[i-1]

		biggerPkgNum := packageNum - i
		addNum := lastX*m[lastX] - (x-lastX)*biggerPkgNum
		tfBeans += addNum
		if addNum < 0 {
			minBeans = tfBeans
		}
		i += (m[x])
	}
	return int64(minBeans)

}

func panicTest() {
	defer func() {
		fmt.Println("goroutine defer end")
		//if err := recover(); err != nil {
		//	fmt.Println("recover panic", err)
		//}
	}()
	fmt.Println("goroutine start")
	panic("xxx")
	fmt.Println("goroutine end")
}

func pathSum(root *TreeNode, target int) [][]int {
	res := [][]int{}
	path := []int{}
	var f func(*TreeNode, int)
	f = func(node *TreeNode, target int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Right == nil && node.Left == nil {
			if node.Val == target {
				path = append(path, node.Val)
				r := make([]int, len(path))
				copy(r, path)
				res = append(res, r)
			}
		}
		f(node.Left, target-node.Val)
		f(node.Right, target-node.Val)

	}

	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNums2(n int) int {
	res := 0
	var f func(n int) int
	f = func(n int) int {
		_ = n > 0 && f(n-1) > 0
		res += n
		return 1
	}
	f(n)
	return res

}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		cha := s[i]
		if len(m) == 0 {
			m[cha] = i
			maxLen = 1
			continue
		}
		if index, ok := m[cha]; ok {
			length := len(m)
			fmt.Println(length, m, i, index)
			for j := i - length; j <= index; j++ {
				delete(m, s[j])
			}
			m[cha] = i
		} else {
			m[cha] = i
			if len(m) > maxLen {
				fmt.Println("-", i, m)
				maxLen = len(m)
			}
		}

	}
	return maxLen
}

// 排序，直到第k个数
func partition(arr []int, i, j int) int {
	com, left, right := i, i, j
	for left < right {
		for left < right && arr[right] >= arr[com] {
			right--
		}
		for left < right && arr[left] <= arr[com] {
			left++
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	arr[com], arr[left] = arr[left], arr[com]
	return left
}

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func square(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(read <-chan int) {
	for x := range read {
		fmt.Println(x)
	}
}

type entry struct {
	ready chan struct{}
	res   int
}

var once sync.Once

func tt() {
	once.Do(func() {
		fmt.Println(111)
		time.Sleep(time.Millisecond)
		fmt.Println(333)
	})
	fmt.Println(222)
}

type A struct {
	a int
}

func replaceSpaceS(s string) string {
	res := ""
	for _, j := range s {
		if j != ' ' {
			res += string(j)
		}
	}
	return res

}

func replaceSpace(s string) string {
	blankNum := 0
	for _, j := range s {
		if j == ' ' { //todo
			blankNum++
		}
	}

	t := make([]byte, len(s)+blankNum*2)
	newS := []byte("%20")
	tNum := 0
	for k := 0; k < len(s); k++ {
		if s[k] == byte(' ') {
			copy(t[tNum:], newS)
			tNum += 3
		} else {
			t[tNum] = s[k]
			tNum += 1
		}
	}
	return string(t)

}

func strToInt(str string) int {
	maxInt32 := 1<<31 - 1
	minInt32 := -1 << 31

	symbol := 1

	comp := maxInt32 / 10

	i := 0
	for i < len(str) && (str[i] == ' ') {
		i++
	}
	if i > len(str)-1 {
		return 0
	}

	if str[i] == '-' {
		symbol = -1
		i++
	} else if str[i] == '+' {
		symbol = 1
		i++
	} else if str[i] >= '0' && str[i] <= '9' {
		symbol = 1
	} else {
		return 0
	}
	if i > len(str)-1 {
		return 0
	}

	ret := 0

	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			x := int(str[i] - '0')
			if ret < comp || (ret == comp && x <= 7) {
				ret = ret*10 + x
			} else {
				if symbol == 1 {
					return maxInt32
				} else {
					return minInt32
				}
			}
		} else {
			return ret * symbol
		}
	}
	return ret * symbol
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {

	var arr []int
	now := head
	for now.Val != 0 {
		arr = append(arr, now.Val)
	}

	//1 2
	for i := 0; i < len(arr)/2; i++ {
		arr[i] = arr[len(arr)-1]
	}
	return arr
}

func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	pre, post := len(s), len(s)
	ret := ""
	for i := len(s) - 1; i >= -1; i-- {
		if i == -1 || s[i] == ' ' {
			if pre <= len(s)-1 && s[pre] != ' ' {
				ret = ret + s[pre:post]
				if i != -1 {
					ret = ret + " "
				}

			}
			pre, post = i, i
		} else {
			pre = i
		}
	}
	return ret

}
