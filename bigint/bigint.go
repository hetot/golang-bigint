package bigint

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

type Bigint struct {
	Value string
}

var ErrorBadInput = errors.New("bad input, please input only number")

func check_to_valid(num string) bool {
	if match, err := regexp.MatchString(`^[+-]?[0-9]*$`, num); err == nil {
		if !match {
			return false
		}
	} else {
		return false
	}

	return true
}

func NewInt(num string) (Bigint, error) {
	if !check_to_valid(num) {
		return Bigint{}, ErrorBadInput
	}

	tmp_num := []rune(num)
	if string(tmp_num[0]) == "+" {
		num = string(tmp_num[1:])
	}

	return Bigint{Value: num}, nil
}

func (z *Bigint) Set(num string) error {
	if !check_to_valid(num) {
		return ErrorBadInput
	}

	tmp_num := []rune(num)
	if string(tmp_num[0]) == "+" {
		num = string(tmp_num[1:])
	}

	z.Value = num
	return nil
}

func Add(a, b Bigint) Bigint {
	ans := ""

	max_len := math.Max(float64(len(a.Value)), float64(len(b.Value)))

	if a.Value[0] != '-' {
		if b.Value[0] != '-' {
			i := 1
			var tmp_a, tmp_b, tmp_inc int
			for i <= int(max_len) {
				if len(a.Value)-i > -1 {
					tmp_a, _ = strconv.Atoi(string(a.Value[len(a.Value)-i]))
				} else {
					tmp_a = 0
				}

				if len(b.Value)-i > -1 {
					tmp_b, _ = strconv.Atoi(string(b.Value[len(b.Value)-i]))
				} else {
					tmp_b = 0
				}

				tmp_sum := tmp_inc + tmp_a + tmp_b
				tmp_inc = 0

				if tmp_sum > 9 {
					tmp_inc = 1
					tmp_sum = tmp_sum % 10
				}
				ans = strconv.Itoa(tmp_sum) + ans
				i += 1
			}
			if tmp_inc == 1 {
				ans = "1" + ans
			}
		} else {
			return Sub(a, b.Abs())
		}
	} else {
		if b.Value[0] == '-' {
			ans = "-" + Add(a.Abs(), b.Abs()).Value
		} else {
			return Sub(b, a.Abs())
		}
	}

	return Bigint{Value: ans}
}

func Sub(a, b Bigint) Bigint {
	ans := ""

	if a.Value[0] != '-' {
		if b.Value[0] != '-' {
			var max_num, min_num string
			var neg bool

			if len(a.Value) > len(b.Value) {
				max_num, min_num = a.Value, b.Value
			} else if len(a.Value) < len(b.Value) {
				max_num, min_num = b.Value, a.Value
				neg = true
			} else {
				for i := 0; i < len(a.Value); i++ {
					tmp_a, _ := strconv.Atoi(string(a.Value[i]))
					tmp_b, _ := strconv.Atoi(string(b.Value[i]))
					if tmp_a > tmp_b {
						max_num, min_num = a.Value, b.Value
						break
					} else if tmp_a < tmp_b {
						max_num, min_num = b.Value, a.Value
						neg = true
						break
					}
				}
			}
			if max_num != "" {
				i := 1
				tmp_a, tmp_b, tmp_dec, tmp_sum := 0, 0, 0, 0
				for len(max_num)-i > -1 {
					tmp_a, _ = strconv.Atoi(string(max_num[len(max_num)-i]))

					if len(min_num)-i > -1 {
						tmp_b, _ = strconv.Atoi(string(min_num[len(min_num)-i]))
						tmp_b = (-1) * tmp_b
					} else {
						tmp_b = 0
					}

					tmp_sum = tmp_a + tmp_b + tmp_dec
					tmp_dec = 0

					if tmp_sum < 0 {
						tmp_sum = 10 + tmp_sum
						tmp_dec = -1
					}

					ans = strconv.Itoa(tmp_sum) + ans

					i += 1
				}

				for ans[0] == '0' {
					ans = ans[1:]
				}

				if neg {
					ans = "-" + ans
				}
			} else {
				ans = "0"
			}
		} else {
			return Add(a, b.Abs())
		}
	} else {
		if b.Value[0] == '-' {
			return Add(b.Abs(), a)
		} else {
			ans = "-" + Add(a.Abs(), b).Value
		}
	}
	return Bigint{Value: ans}
}

func Multiply(a, b Bigint) Bigint {
	neg := false
	ans := ""

	if a.Value != "0" && b.Value != "0" {
		if (a.Value[0] == '-' && b.Value[0] != '-') || (a.Value[0] != '-' && b.Value[0] == '-') {
			neg = true
		}

		new_a := a.Abs()
		tmp_a := new_a.Value
		new_b := b.Abs()
		new_b.Set(Sub(b, Bigint{Value: "1"}).Value)

		for new_b.Value != "0" {
			new_a = Add(new_a, Bigint{Value: tmp_a})
			new_b.Set(Sub(new_b, Bigint{Value: "1"}).Value)
		}
		ans = new_a.Value

		if neg {
			ans = "-" + ans
		}
	} else {
		ans = "0"
	}

	return Bigint{Value: ans}
}

func Mod(a, b Bigint) Bigint {
	ans := ""

	if b.Value == "0" {
		ans = a.Value
	} else if a.Value[0] != '-' {
		new_a_val := a.Value
		for Sub(Bigint{Value: new_a_val}, b.Abs()).Value[0] != '-' && Sub(Bigint{Value: new_a_val}, b.Abs()).Value != "0" {
			new_a_val = Sub(Bigint{Value: new_a_val}, b.Abs()).Value
		}

		if Sub(Bigint{Value: new_a_val}, b.Abs()).Value == "0" {
			ans = "0"
		} else {
			ans = new_a_val
		}
	} else {
		new_a_val := a.Value
		for Add(Bigint{Value: new_a_val}, b.Abs()).Value[0] == '-' && Sub(Bigint{Value: new_a_val}, b.Abs()).Value != "0" {
			new_a_val = Add(Bigint{Value: new_a_val}, b.Abs()).Value
		}

		if Add(Bigint{Value: new_a_val}, b.Abs()).Value == "0" {
			ans = "0"
		} else {
			ans = Add(Bigint{Value: new_a_val}, b.Abs()).Value
		}
	}

	return Bigint{Value: ans}
}

func (x *Bigint) Abs() Bigint {
	if x.Value[0] == '-' {
		return Bigint{
			Value: x.Value[1:],
		}
	}
	return Bigint{Value: x.Value}
}
