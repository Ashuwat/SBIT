package numbers

import "math"

type BigNumber struct {
	LargeNum []int
}

func (b *BigNumber) convertSpliceToInt() int {
	var wholeNum int
	for i := len(b.LargeNum); i > 0; i-- {
		wholeNum = wholeNum + (b.LargeNum[i-1] * int(math.Pow(10, float64(i-1))))
	}
	return wholeNum
}

func (b *BigNumber) IsEven() bool {
	return b.LargeNum[len(b.LargeNum)-1]%2 == 0
}

func (b *BigNumber) IsPrime() bool {
	var wholeNum = b.convertSpliceToInt()

	if wholeNum%2 == 1 {
		return false
	}

	for i := 2; i*i <= wholeNum; i++ {
		if wholeNum/i == 0 {
			return false
		}
	}

	return true
}

func (b *BigNumber) Factors() []int {
	var wholeNum = b.convertSpliceToInt()

	if b.IsPrime() {
		return []int{1, wholeNum}
	}

	var someSlice []int
	for i := 2; i*i <= wholeNum; i++ {
		if wholeNum/i == 0 {
			var tempSlice = []int{i, wholeNum / i}
			someSlice = append(someSlice, tempSlice...)
		}
	}
	return someSlice
}
