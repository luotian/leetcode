package important

import "math"

//29. 两数相除
//给你两个整数，被除数 dividend 和除数 divisor。将两数相除，要求 不使用 乘法、除法和取余运算。
//整数除法应该向零截断，也就是截去（truncate）其小数部分。例如，8.345 将被截断为 8 ，-2.7335 将被截断至 -2 。
//返回被除数 dividend 除以除数 divisor 得到的 商 。
//注意：假设我们的环境只能存储 32 位 有符号整数，其数值范围是 [−231,  231 − 1] 。
//本题中，如果商 严格大于 231 − 1 ，则返回 231 − 1 ；如果商 严格小于 -231 ，则返回 -231 。
//
//提示：
//-231 <= dividend, divisor <= 231 - 1
//divisor != 0

// 解答技巧：先解决边界溢出问题，再用指数步进进行效率优化
func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == math.MinInt32 {
		if dividend == math.MinInt32 {
			return 1
		}

		return 0
	}

	if dividend == math.MinInt32 {
		if divisor == -1 {
			return math.MaxInt32
		}

		if divisor == 1 {
			return math.MinInt32
		}
	}

	var result int
	var positive bool
	signFlag := (dividend ^ divisor) >> 31
	if signFlag == 0 {
		positive = true
	}

	if dividend < 0 {
		dividend = -dividend
	}

	if divisor < 0 {
		divisor = -divisor
	}

	for dividend >= divisor {
		var step int = 1
		var tmpDivisor int = divisor
		for dividend >= tmpDivisor {
			result += step
			dividend -= tmpDivisor
			step <<= 1
			tmpDivisor <<= 1
		}
	}

	if positive {
		return result
	}
	return -result
}
