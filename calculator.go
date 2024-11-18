package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	// 创建应用
	myApp := app.New()
	myWindow := myApp.NewWindow("计算器")

	// 初始化变量
	var input string
	output := widget.NewLabel("")

	// 创建按钮
	buttons := []string{
		"7", "8", "9", "/", "C",
		"4", "5", "6", "*", "←",
		"1", "2", "3", "-", "=",
		"0", ".", "+",
	}

	// 输入框
	entry := widget.NewEntry()
	entry.Disable()

	// 按钮事件处理
	grid := container.NewGridWithColumns(5)
	for _, button := range buttons {
		btn := button
		grid.Add(widget.NewButton(btn, func() {
			switch btn {
			case "C": // 清空输入
				input = ""
				output.SetText("")
				entry.SetText("")
			case "←": // 删除最后一个字符
				if len(input) > 0 {
					input = input[:len(input)-1]
					entry.SetText(input)
				}
			case "=": // 计算结果
				result, err := evaluate(input)
				if err != nil {
					output.SetText("错误: " + err.Error())
				} else {
					output.SetText("结果: " + strconv.FormatFloat(result, 'f', -1, 64))
				}
			default: // 更新输入
				input += btn
				entry.SetText(input)
			}
		}))
	}

	// 布局
	content := container.NewVBox(
		widget.NewLabel("简单计算器"),
		entry,
		grid,
		output,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 400))
	myWindow.ShowAndRun()
}

// 计算逻辑
func evaluate(expression string) (float64, error) {
	// 简单实现一个解析器 (仅支持基础运算)
	var stack []float64
	var operator byte
	num := 0.0
	var err error

	for i := 0; i < len(expression); i++ {
		ch := expression[i]
		switch {
		case ch >= '0' && ch <= '9' || ch == '.':
			start := i
			for i+1 < len(expression) && (expression[i+1] >= '0' && expression[i+1] <= '9' || expression[i+1] == '.') {
				i++
			}
			num, err = strconv.ParseFloat(expression[start:i+1], 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		case ch == '+' || ch == '-' || ch == '*' || ch == '/':
			operator = ch
		default:
			return 0, nil
		}
	}

	if len(stack) < 2 {
		return 0, nil
	}
	// 此处简单逻辑，需扩展支持括号。
	return stack[0], nil
}

