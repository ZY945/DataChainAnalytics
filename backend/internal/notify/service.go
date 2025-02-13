package notify

import (
	"fmt"
	"log"
)

type NotifyType int

const (
	TypeInfo NotifyType = iota
	TypeWarning
	TypeError
)

type NotifyService interface {
	Send(msgType NotifyType, title string, content string) error
	SendCard(title string, content interface{}, templateID string, templateVersionName string) error
}

// 默认通知服务实现
type defaultNotifyService struct {
	logger    *log.Logger
	feishuSvc NotifyService // 改回接口类型
}

func NewDefaultNotifyService(feishuSvc NotifyService) NotifyService {
	return &defaultNotifyService{
		logger:    log.Default(),
		feishuSvc: feishuSvc,
	}
}

func (s *defaultNotifyService) Send(msgType NotifyType, title string, content string) error {
	// 同时发送到日志和飞书
	// msg := fmt.Sprintf("[%s] %s: %s", s.getTypeString(msgType), title, content)
	// s.logger.Println(msg)

	// 调用飞书通知
	if s.feishuSvc != nil {
		if err := s.feishuSvc.Send(msgType, title, content); err != nil {
			return fmt.Errorf("send feishu message failed: %w", err)
		} else {
			s.logger.Println("send feishu message success")
		}
	} else {
		s.logger.Println("feishuSvc is nil")
	}

	return nil
}

func (s *defaultNotifyService) SendCard(title string, content interface{}, templateID string, templateVersionName string) error {
	// 调用飞书通知
	if s.feishuSvc != nil {
		if err := s.feishuSvc.SendCard(title, content, templateID, templateVersionName); err != nil {
			return fmt.Errorf("send feishu card message failed: %w", err)
		} else {
			s.logger.Println("send feishu card message success")
		}
	} else {
		s.logger.Println("feishuSvc is nil")
	}

	return nil
}

func (s *defaultNotifyService) getTypeString(t NotifyType) string {
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
