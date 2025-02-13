package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type BalanceResponse struct {
	Result struct {
		Value uint64 `json:"value"`
	} `json:"result"`
}

func TestGetBalance(t *testing.T) {
	// RPC端点
	rpcEndpoint := "https://solana-api.projectserum.com"
	// rpcEndpoint := "https://api.mainnet-beta.solana.com"
	// 使用测试网而不是主网
	// rpcEndpoint := "https://api.testnet.solana.com"
	// 其他公共节点
	// https://solana-api.projectserum.com
	// https://api.mainnet.rpcpool.com

	// 构造获取余额请求
	requestBody, err := json.Marshal(RPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "getBalance",
		Params:  []interface{}{""},
	})
	if err != nil {
		t.Fatalf("Failed to marshal request: %v", err)
	}

	// 创建带超时的客户端
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// 发送HTTP POST请求
	resp, err := client.Post(rpcEndpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 解析响应
	var balanceResp BalanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&balanceResp); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	fmt.Printf("Balance: %d lamports\n", balanceResp.Result.Value)
}
