package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Symbol struct {
	Code       string `json:"code"`
	DepthLevel int    `json:"depth_level"`
}

type Data struct {
	SymbolList []Symbol `json:"symbol_list"`
}

type Request struct {
	CmdID int    `json:"cmd_id"`
	SeqID int    `json:"seq_id"`
	Trace string `json:"trace"`
	Data  Data   `json:"data"`
}

/*
Special Note:
GitHub: https://github.com/alltick/realtime-forex-crypto-stock-tick-finance-websocket-api
Token Application: https://alltick.co
Replace "testtoken" in the URL below with your own token
API addresses for forex, cryptocurrencies, and precious metals:
wss://quote.tradeswitcher.com/quote-b-ws-api
Stock API address:
wss://quote.tradeswitcher.com/quote-stock-b-ws-api
*/
const baseURL = "wss://quote.tradeswitcher.com/quote-b-ws-api?token="

func HttpGoldPrice(token string) (string, error) {
	// 构造请求体
	requestBody := struct {
		Data struct {
			Code              string `json:"code"`
			KlineType         string `json:"kline_type"`
			KlineTimestampEnd string `json:"kline_timestamp_end"`
			QueryKlineNum     string `json:"query_kline_num"`
			AdjustType        string `json:"adjust_type"`
		} `json:"data"`
		Trace string `json:"trace"`
	}{
		Data: struct {
			Code              string `json:"code"`
			KlineType         string `json:"kline_type"`
			KlineTimestampEnd string `json:"kline_timestamp_end"`
			QueryKlineNum     string `json:"query_kline_num"`
			AdjustType        string `json:"adjust_type"`
		}{
			Code:              "GOLD",
			KlineType:         "1",
			KlineTimestampEnd: "0",
			QueryKlineNum:     "1",
			AdjustType:        "0",
		},
		Trace: "6906e29e-8827-44d4-983d-34b8b0549b3b",
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("marshal request body failed: %w", err)
	}

	url := fmt.Sprintf("https://quote.tradeswitcher.com/quote-b-api/kline?token=%s&query=%s",
		token, url.QueryEscape(string(jsonData)))

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response failed: %w", err)
	}

	return string(body), nil
}

func WebsocketGoldPrice(token string) {
	url := baseURL + token
	log.Println("Connecting to server at", url)

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	// Send heartbeat every 10 seconds
	go func() {
		for range time.NewTicker(10 * time.Second).C {
			req := Request{
				CmdID: 22000,
				SeqID: 123,
				Trace: "3380a7a-3e1f-c3a5-5ee3-9e5be0ec8c241692805462",
				Data:  Data{},
			}
			messageBytes, err := json.Marshal(req)
			if err != nil {
				log.Println("json.Marshal error:", err)
				return
			}
			log.Println("heartbeat req data:", string(messageBytes))

			err = c.WriteMessage(websocket.TextMessage, messageBytes)
			if err != nil {
				log.Println("write:", err)
			}
		}
	}()

	req := Request{
		CmdID: 22002,
		SeqID: 123,
		Trace: uuid.New().String(),
		Data: Data{SymbolList: []Symbol{
			{"GOLD", 1}, // 黄金
			// {"AAPL.US", 5},// 美股
			// {"700.HK", 5},// 港股
			// {"USDJPY", 5},// 美元日元
		}},
	}
	messageBytes, err := json.Marshal(req)
	if err != nil {
		log.Println("json.Marshal error:", err)
		return
	}
	log.Println("req data:", string(messageBytes))

	err = c.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		log.Println("write:", err)
	}

	req.CmdID = 22004
	messageBytes, err = json.Marshal(req)
	if err != nil {
		log.Println("json.Marshal error:", err)
		return
	}
	log.Println("req data:", string(messageBytes))

	err = c.WriteMessage(websocket.TextMessage, messageBytes)
	if err != nil {
		log.Println("write:", err)
	}

	rece_count := 0
	for {
		_, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		} else {
			log.Println("Received message:", string(message))
		}

		rece_count++
		if rece_count%10000 == 0 {
			log.Println("count:", rece_count, " Received message:", string(message))
		}
	}

}
