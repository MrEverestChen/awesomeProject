package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

func main() {
	// 创建带有选项的 Chrome 浏览器上下文
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),          // 设为 false 以显示浏览器窗口（调试用）
		chromedp.Flag("disable-gpu", true),        // 适用于某些系统
		chromedp.Flag("enable-automation", false), // 让网站认为是正常用户
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"),
	)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var pageContent string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://onlineservices.immigration.govt.nz"),
		chromedp.Sleep(5*time.Second), // 等待页面加载
		chromedp.OuterHTML("html", &pageContent),
	)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Page Content:", pageContent)
}
