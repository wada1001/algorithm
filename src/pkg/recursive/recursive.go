package recursive

// count 1 in supplied number
func CountOneFor(number int) int {
	ret := 0
	for i := number; i > 0; i = i/10 {
		if i % 10 == 1 {
			ret++
		}
	}
	return ret
}

// count 1 in supplied number
func CountOneRecursive(number int) int {
	if number == 0 {
		return 0
	}

	if number % 10 != 1 {
		return 0 + CountOneRecursive(number / 10)
	}

	return 1 + CountOneRecursive(number / 10)
}

func Gcd(nums []int) int {
	if len(nums) == 2 { 
		return GcdIn(nums[0], nums[1])
	}

	return GcdIn(nums[0], Gcd(nums[1:]))
}

func GcdIn(a, b int) int {
	for i:= a; i > 0; i-- {
		if a % i != 0 || b % i != 0 {
			continue
		}
		return i
	}
	return 0
}

