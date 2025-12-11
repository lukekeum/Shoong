package rtmp

import (
	"errors"
	"strings"

	"github.com/yutopp/go-rtmp"
	rtmpmsg "github.com/yutopp/go-rtmp/message"
)

type Handler struct {
	rtmp.DefaultHandler

	ChannelId string
	StreamKey string
}

func (h *Handler) OnConnect(timestamp uint32, cmd *rtmpmsg.NetConnectionConnect) error {
	key := strings.Trim(cmd.Command.App, "/")

	if len(key) < 6 || !strings.HasPrefix(key, "live/") {
		return errors.New("invalid app name: " + key[0:4])
	}

	h.ChannelId = key[5:]

	return nil
}

func (h *Handler) OnPublish(ctx *rtmp.StreamContext, ts uint32, cmd *rtmpmsg.NetStreamPublish) error {
	id := cmd.PublishingName

	if id == "" {
		return errors.New("invalid stream key: nil")
	}

	err := isOverlap(id)

	if err != nil {
		return err
	}

	err = isValid(id)

	if err != nil {
		return err
	}

	err = setIsStreamTrue(id)

	if err != nil {
		return err
	}

	h.StreamKey = id

	// TODO: 트랜스코드 서버에 grpc 요청하기
	// (strean_key, user_id, ...)

	return nil
}

// TODO: 기능들 추가하기 + 함수 변환 타입 바꾸기
func isOverlap(k string) error {
	return nil
}

func isValid(k string) error {
	return nil
}

func setIsStreamTrue(u string) error {
	return nil
}
