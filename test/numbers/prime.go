package numbers

type NumberProps interface {
	IsPrime() bool
	Factors() []int
	IsEven() bool
}

type Number struct {
	SomeInt int
}

func (n *Number) IsPrime() bool {
	if n.SomeInt%2 == 1 {
		return false
	}

	for i := 2; i*i <= n.SomeInt; i++ {
		if n.SomeInt/i == 0 {
			return false
		}
	}

	return true
}

func (n *Number) Factors() []int {
	if n.IsPrime() {
		return []int{1, n.SomeInt}
	}

	var someSlice []int
	for i := 2; i*i <= n.SomeInt; i++ {
		if n.SomeInt/i == 0 {
			var tempSlice = []int{i, n.SomeInt / i}
			someSlice = append(someSlice, tempSlice...)
		}
	}
	return someSlice
}

func (n *Number) IsEven() bool {
	return n.SomeInt%2 == 0
}
