# Go Array Libray

go-array is a library of like ruby's Array.

## Installation

    $ go get github.com/ramtiga/go-array

## Usage

```go
imort (
    "github.com/ramtiga/go-array"
)

arry := array.IntArray{1, 2, 3, 4, 5}

arry.Length()    // returns => 5
arry.Push(6)     // arry => {1, 2, 3, 4, 5, 6}
arry.Pop()       // arry => {1, 2, 3, 4, 5}
arry.First()     // returns => 1
arry.Last()      // returns => 5
arry.Shift()     // arry => {2, 3, 4, 5}

// other methods  
// Unshift, Concat, Clear, Delete, Delete_at, 
// Index, Index, Reverse, Join, Uniq, Sort

```

## TODO
- StringArray

## License
MIT

## Author
ramtiga

