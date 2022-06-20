package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	log2()
}

// 第一種用法
func log1() {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()

	url := "https://www.google.com.tw/"
	sugar.Infow(
		"failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)

	sugar.Infof("failed to fetch URL: %s", url)
}

// 第二種用法
func log2() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "https://www.google.com.tw/"
	logger.Info(
		"https://www.google.com.tw/",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	// 除了Info以外還有Debug/Error/Warn

	/** 基本類型
	* zap.Type (Type 可以為bool/int/uint/float64/complex64/time.Time/time.Duration/error等)
	*
	* 指針類型寫法
	* zap.int(key string, val int) Field => 一般
	* zap.intp(key string, val *int) Field => 指針
	* zap.ints(key string, val []int) Field
	*
	* 特殊類型
	* zap.Any(key string, val interface{}) Field => 任意
	* zap.Binary(key string, val []byte) Field => 2進制
	 */
}
