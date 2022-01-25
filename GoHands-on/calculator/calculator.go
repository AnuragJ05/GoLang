//10.Custom Packages
package calculator

func Calculate(num1 int, num2 int, op string) int {
	if op == "+" {
		return num1 + num2
	} else if op == "-" {
		return num1 - num2
	} else if op == "*" {
		return num1 * num2
	} else if op == "/" {
		return num1 / num2
	} else if op == "%" {
		return num1 % num2
	}
	return 0
}
