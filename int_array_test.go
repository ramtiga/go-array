package array

import (
    "testing"
)

func TestLength(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    result := d.Length()
    if result != 5 {
        t.Errorf("result=%v", result)
    }
}

func TestSize(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    result := d.Size()
    if result != 5 {
        t.Errorf("result=%v", result)
    }
}

func TestPush(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{5, 3, 1, 2, 6, 7}
    d.Push(7)
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestFirst(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    resutl := d.First()
    
    if resutl != 5 {
        t.Errorf("result=%v", d)
    } 
}

func TestLast(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    resutl := d.Last()
    
    if resutl != 6 {
        t.Errorf("result=%v", d)
    } 
}

func TestShift(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{3, 1, 2, 6}    
    d.Shift()
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestUnshift(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{10, 11, 5, 3, 1, 2, 6}    

    add := IntArray{10, 11}
    d.Unshift(add)
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestConcat(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{5, 3, 1, 2, 6, 10, 11}    

    add := IntArray{10, 11}
    d.Concat(add)
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestClear(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    d.Clear()
    if d.Length() != 0 {
        t.Errorf("result=%v", d)
    } 
}

func TestDelete(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6, 3}    
    chk := IntArray{5, 1, 2, 6}

    d.Delete(3)
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestDelete_at(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{5, 1, 2, 6}

    d.Delete_at(1)
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestIndex(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    

    result := d.Index(1)
    
    if result == 3 {
        t.Errorf("result=%v", result)
    } 
    result = d.Index(7)
    
    if result != -1 {
        t.Errorf("result=%v", result)
    } 
}

func TestReverse(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{6, 2, 1, 3, 5}

    d.Reverse()
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}

func TestJoin(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := "5-3-1-2-6"

    result := d.Join("-")
    
    if result != chk {
        t.Errorf("result=%v", result)
    } 
}

func TestUniq(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 3, 2, 1}
    chk := IntArray{5, 3, 1, 2}

    result := d.Uniq()
    
    if result.Length() != 4 {
        t.Errorf("result.Length=%v", result.Length())
    }
    for i, _ := range result {
       if result[i] != chk[i] {
          t.Errorf("result=%v, index=%v", result, i)
       } 
    }
}

func TestSort(t *testing.T) {
    var d IntArray = IntArray{5, 3, 1, 2, 6}    
    chk := IntArray{1, 2, 3, 5, 6}

    d.Sort()
    
    for i, _ := range d {
       if d[i] != chk[i] {
          t.Errorf("result=%v, index=%v", d, i)
       } 
    }
}



