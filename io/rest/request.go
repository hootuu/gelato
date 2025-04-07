package rest

import (
	"encoding/json"
	"github.com/hootuu/gelato/crtpto/ed25519x"
	"github.com/hootuu/gelato/crtpto/hexx"
	"github.com/hootuu/gelato/crtpto/rand"
	"github.com/hootuu/gelato/errors"
	"github.com/hootuu/gelato/idx"
	"github.com/hootuu/gelato/io/serializer"
	"go.uber.org/zap"
	"time"
)

type Empty struct {
}

func NewEmpty() *Empty {
	return &Empty{}
}

type Request[T any] struct {
	ID        string `json:"id"`
	GuardID   string `json:"guard_id"`
	Timestamp int64  `json:"timestamp"`
	Nonce     int64  `json:"nonce"`
	Signature string `json:"signature"`
	Data      *T     `json:"data,omitempty"`
}

func NewRequest[T any](guardID string, data *T) *Request[T] {
	randInt64, _ := rand.Int64()
	return &Request[T]{
		Data:      data,
		ID:        idx.New(),
		GuardID:   guardID,
		Timestamp: time.Now().UnixMilli(),
		Nonce:     randInt64,
	}
}

func (req *Request[T]) Marshal() ([]byte, *errors.Error) {
	byteArr, err := json.Marshal(req)
	if err != nil {
		return nil, errors.System("101", err)
	}
	return byteArr, nil
}

func (req *Request[T]) JSON() string {
	byteArr, err := json.Marshal(req)
	if err != nil {
		return "{}"
	}
	return string(byteArr)
}

func UnmarshalRequest[T any](data []byte) (*Request[T], *errors.Error) {
	var req Request[T]
	if err := json.Unmarshal(data, &req); err != nil {
		return nil, errors.E("104", "unmarshal request failed: %w", err)
	}
	return &req, nil
}

func (req *Request[T]) Sign(priKey []byte) *errors.Error {
	bytes, err := req.Serialize()
	if err != nil {
		return err
	}
	sig, err := ed25519x.Sign(priKey, bytes)
	if err != nil {
		return err
	}
	req.Signature = hexx.Encode(sig)
	return nil
}

func (req *Request[T]) Verify(pubKey []byte) *errors.Error {
	bytes, err := req.Serialize()
	if err != nil {
		return err
	}
	bytesSign, err := hexx.Decode(req.Signature)
	if err != nil {
		return err
	}
	valid, err := ed25519x.Verify(pubKey, bytes, bytesSign)
	if err != nil {
		return err
	}
	if !valid {
		return errors.E("909", "invalid signature")
	}
	return nil
}

func (req *Request[T]) Serialize() ([]byte, *errors.Error) {
	serializeStr, err := serializer.Serialize(map[string]interface{}{
		"id":        req.ID,
		"guard_id":  req.GuardID,
		"data":      req.Data,
		"timestamp": req.Timestamp,
		"nonce":     req.Nonce,
	})
	if err != nil {
		gLogger.Error("req.Serialize err", zap.Error(err))
		return nil, errors.E("11001", "request serialize failed: %s", err.Error())
	}
	return []byte(serializeStr), nil
}
