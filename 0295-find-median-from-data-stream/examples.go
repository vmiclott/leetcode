package findmedianfromdatastream

import "fmt"

func Examples() {
	fmt.Printf(`Input:
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian())
`)
	fmt.Println("Output:")
	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	fmt.Println(medianFinder.FindMedian())
	medianFinder.AddNum(3)
	fmt.Println(medianFinder.FindMedian())
}
