package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 主函数
func main() {
	fmt.Println("欢迎使用简单计算器！")
	fmt.Println("请输入数学表达式（如：3 + 4 或 8 / 2），输入 'exit' 退出。")

	scanner := bufio.NewScanner(os.Stdin) // 读取用户输入

	for {
		fmt.Print(">>> ")
		scanner.Scan()
		input := scanner.Text() // 获取用户输入

		// 退出程序
		if strings.ToLower(strings.TrimSpace(input)) == "exit" {
			fmt.Println("感谢使用，再见！")
			break
		}

		// 计算结果
		result, err := calculate(input)
		if err != nil {
			fmt.Println("错误:", err)
		} else {
			fmt.Printf("结果: %v\n", result)
		}
	}
}

// 计算函数
func calculate(expression string) (float64, error) {
	// 去掉多余的空格
	expression = strings.TrimSpace(expression)

	// 支持的运算符
	operators := []string{"+", "-", "*", "/"}

	for _, operator := range operators {
		if strings.Contains(expression, operator) {
			// 按运算符拆分
			parts := strings.Split(expression, operator)
			if len(parts) != 2 {
				return 0, fmt.Errorf("无效的表达式: %s", expression)
			}

			// 将操作数转换为浮点数
			num1, err1 := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
			num2, err2 := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)

			if err1 != nil || err2 != nil {
				return 0, fmt.Errorf("无效的数字")
			}

			// 执行运算
			switch operator {
			case "+":
				return num1 + num2, nil
			case "-":
				return num1 - num2, nil
			case "*":
				return num1 * num2, nil
			case "/":
				if num2 == 0 {
					return 0, fmt.Errorf("除数不能为零")
				}
				return num1 / num2, nil
			}
		}
	}

	return 0, fmt.Errorf("无法识别的表达式: %s", expression)
}
