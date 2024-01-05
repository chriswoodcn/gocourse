package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

func main() {
	// 设置输出到标准错误流（stderr）
	logger := log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

	// 定义不同类型的日志消息及对应的颜色
	colors := map[string]color.Attribute{
		"debug":   color.Faint,    // 调试信息
		"info":    color.Faint,    // 正常信息
		"warning": color.FgYellow, // 警告信息
		"error":   color.FgRed,    // 错误信息
	}

	// 自定义日志函数
	LogFunc := func(level string, message interface{}) {
		// 根据日志级别选择合适的颜色并格式化输出
		switch level {
		case "debug", "info":
			color.Set(color.Faint, color.Bold)
			defer color.Unset() // Use it in your function
			logger.Println(message)
		default:
			color.Set(colors[level])
			logger.Println(message)
			color.Unset() // Don't forget to unset
		}
	}

	// 测试日志输出
	LogFunc("debug", "This is a debug message.")
	LogFunc("info", "This is an info message.")
	LogFunc("warning", "This is a warning message.")
	LogFunc("error", "This is an error message.")

	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Or just add them to New()
	d := color.New(color.FgCyan, color.Bold)
	d.Printf("This prints bold cyan %s\n", "too!.")

	// Mix up foreground and background colors, create new mixes!
	red := color.New(color.FgRed)

	boldRed := red.Add(color.Bold)
	boldRed.Println("This will print text in bold red.")

	whiteBackground := red.Add(color.BgWhite)
	whiteBackground.Println("Red text with white background.")
}
