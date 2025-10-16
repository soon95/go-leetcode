package no295

import (
	"fmt"
	"testing"
)

func TestMedianFinder(t *testing.T) {

	medianFinder := Constructor()

	medianFinder.AddNum(40)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(12)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(16)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(14)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(35)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(19)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(34)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(35)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(28)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(35)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(26)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(6)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(8)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(2)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(14)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(25)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(25)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(4)
	fmt.Printf("%v\n", medianFinder.FindMedian())

	medianFinder.AddNum(33)
	fmt.Printf("%v\n", medianFinder.FindMedian())

}

//[null,null,40.0,null,26.0,null,16.0,null,15.0,null,16.0,null,17.5,null,19.0,null,26.5,null,28.0,null,31.0,null,28.0,null,27.0,null,26.0,null,22.5,null,19.0,null,22.0,null,25.0,null,22.0,null,25.0,null,22.0,null,19.0,null,18.5,null,19.0,null,18.5,null,19.0,null,18.5,null,19.0,null,21.5,null,19.0,null,18.5,null,18.0,null,18.5,null,19.0,null,19.0,null,19.0,null,18.5,null,19.0,null,19.0,null,19.0,null,19.0,null,19.0]
//[null,null,40.0,null,26.0,null,16.0,null,15.0,null,16.0,null,17.5,null,19.0,null,26.5,null,28.0,null,31.0,null,28.0,null,27.0,null,26.0,null,22.5,null,19.0,null,22.0,null,19.0,null,22.0,null,19.0,null,22.0,null,25.0,null,18.5,null,19.0,null,18.5,null,19.0,null,18.5,null,19.0,null,21.5,null,19.0,null,18.5,null,18.0,null,18.5,null,19.0,null,19.0,null,19.0,null,18.5,null,19.0,null,19.0,null,19.0,null,19.0,null,19.0]
