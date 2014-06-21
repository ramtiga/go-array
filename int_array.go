package array

import (
	"fmt"
	"sort"
	"strings"
)

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}

func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p IntArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// length of the array
func (p *IntArray) Length() int {
	return len(*p)
}

// length of the array
func (p *IntArray) Size() int {
	return len(*p)
}

// add to the existing array
func (p *IntArray) Push(value ...int) {
	tmp := *p
	for _, v := range value {
		tmp = append(tmp, v)
	}
	*p = tmp
}

// remove last value of the array
func (p *IntArray) Pop() int {
	tmp := *p
	last := tmp[len(tmp)-1 : len(tmp)]
	tmp = tmp[0 : len(tmp)-1]

	*p = tmp
	return last[0]
}

// get first value of the array
func (p IntArray) First() int {
	return p[0]
}

// get last value of the array
func (p IntArray) Last() int {
	last := p[len(p)-1 : len(p)]
	return last[0]
}

// remove first value of the array
func (p *IntArray) Shift() {
	tmp := *p
	if len(tmp) < 1 {
		return
	}
	*p = tmp[1:len(tmp)]
}

// add value the first index to the existing array
func (p *IntArray) Unshift(arry IntArray) {
	tmp := *p
	arry.Concat(tmp)
	*p = arry
}

// add new array to the existing array
func (p *IntArray) Concat(arry IntArray) {
	tmp := *p
	tmp = append(tmp, arry...)
	*p = tmp
}

// initialize of the array
func (p *IntArray) Clear() {
	var tmp IntArray
	*p = tmp
}

// remove del_target from array
func (p *IntArray) Delete(del_value int) {
	tmp := *p
	var new_array IntArray
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != del_value {
			new_array = append(new_array, tmp[i])
		}
	}
	*p = new_array
}

// remove del_index of index from array
func (p *IntArray) Delete_at(del_index int) {
	tmp := *p
	var new_array IntArray
	for i := 0; i < len(tmp); i++ {
		if i != del_index {
			new_array = append(new_array, tmp[i])
		}
	}
	*p = new_array
}

// get index of value from array
func (p IntArray) Index(n int) int {
	for i, v := range p {
		if v == n {
			return i
		}
	}
	return -1
}

// get a new reverse array
func (p *IntArray) Reverse() {
	tmp := *p
	var new_array IntArray
	length := len(tmp)

	for i := length - 1; i >= 0; i-- {
		new_array = append(new_array, tmp[i])
	}
	*p = new_array
}

// get a new string from array
func (p *IntArray) Join(sep string) string {
	tmp := *p
	var str_ary []string

	for _, v := range tmp {
		str_ary = append(str_ary, fmt.Sprint(v))
	}
	str := strings.Join(str_ary, sep)
	return str
}

// get a new unique array from array
func (p *IntArray) Uniq() IntArray {
	tmp := *p
	var new_array IntArray

	for _, v := range tmp {
		if new_array.Index(v) == -1 {
			new_array = append(new_array, v)
		}
	}
	return new_array
}

// get a new sort array from array
func (p *IntArray) Sort() {
	tmp := *p
	// var tmp2 IntArray
	sort.Sort(tmp)
	*p = tmp
}
