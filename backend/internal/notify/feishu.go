package notify

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FeishuNotifyService struct {
	webhook   string
	appSecret string
}

func NewFeishuNotifyService(webhook string, appSecret string) NotifyService {
	return &FeishuNotifyService{
		webhook:   webhook,
		appSecret: appSecret,
	}
}

type feishuMessage struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`
	Content   struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (s *FeishuNotifyService) genSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

// 发送消息
func (s *FeishuNotifyService) Send(msgType NotifyType, title string, content string) error {
	fmt.Printf("开始发送飞书消息，webhook: %s\n", s.webhook)
	timestamp := time.Now().Unix()
	sign, err := s.genSign(s.appSecret, timestamp)
	if err != nil {
		return fmt.Errorf("generate signature failed: %w", err)
	}

	msg := feishuMessage{
		Timestamp: fmt.Sprintf("%d", timestamp),
		Sign:      sign,
		MsgType:   "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: fmt.Sprintf("[%s]\n%s\n%s", s.getTypeString(msgType), title, content),
		},
	}

	jsonData, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal message failed: %w", err)
	}
	fmt.Printf("发送的消息内容: %s\n", string(jsonData))

	resp, err := http.Post(s.webhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("send message failed: %w", err)
	}
	defer resp.Body.Close()

	// 读取并打印完整响应
	respBody, _ := io.ReadAll(resp.Body)
	fmt.Printf("飞书响应状态码: %d, 响应内容: %s\n", resp.StatusCode, string(respBody))

	if resp.StatusCode != http.StatusOK {
		var errResp struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}
		if err := json.Unmarshal(respBody, &errResp); err != nil {
			return fmt.Errorf("decode error response failed: %w", err)
		}
		return fmt.Errorf("send message failed with code %d: %s", errResp.Code, errResp.Msg)
	}

	return nil
}

func (s *FeishuNotifyService) getTypeString(t NotifyType) string {
	switch t {
	case TypeInfo:
		return "INFO"
	case TypeWarning:
		return "WARNING"
	case TypeError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}
