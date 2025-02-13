package service

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type GoldResponse struct {
	Ret   int    `json:"ret"`
	Msg   string `json:"msg"`
	Trace string `json:"trace"`
	Data  struct {
		Code      string `json:"code"`
		KlineType int    `json:"kline_type"`
		KlineList []struct {
			Timestamp  string `json:"timestamp"`
			OpenPrice  string `json:"open_price"`
			ClosePrice string `json:"close_price"`
			HighPrice  string `json:"high_price"`
			LowPrice   string `json:"low_price"`
			Volume     string `json:"volume"`
			Turnover   string `json:"turnover"`
		} `json:"kline_list"`
	} `json:"data"`
}

func FormatGoldPrice(jsonStr string) (string, error) {
	var goldResp GoldResponse
	if err := json.Unmarshal([]byte(jsonStr), &goldResp); err != nil {
		return "", fmt.Errorf("unmarshal response failed: %w", err)
	}

	// 转换时间戳为可读时间
	timestamp, _ := strconv.ParseInt(goldResp.Data.KlineList[0].Timestamp, 10, 64)
	timeStr := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")

	// 格式化输出
	content := fmt.Sprintf(
		"商品代码: %s\n"+
			"K线类型: %d(具体含义需参考对应 API 文档)\n"+
			"K 线详情:\n"+
			"时间: %s\n"+
			"开盘价: %s\n"+
			"收盘价: %s\n"+
			"最高价: %s\n"+
			"最低价: %s\n"+
			"成交量: %s\n"+
			"成交额: %s",
		goldResp.Data.Code,
		goldResp.Data.KlineType,
		timeStr,
		goldResp.Data.KlineList[0].OpenPrice,
		goldResp.Data.KlineList[0].ClosePrice,
		goldResp.Data.KlineList[0].HighPrice,
		goldResp.Data.KlineList[0].LowPrice,
		goldResp.Data.KlineList[0].Volume,
		goldResp.Data.KlineList[0].Turnover,
	)

	return content, nil
}
