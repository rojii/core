package smux

import (
	"bytes"
	"net"
	"testing"
	"time"
)

type buffer struct {
	bytes.Buffer
}

func (b *buffer) Close() error {
	b.Buffer.Reset()
	return nil
}

func (b *buffer) LocalAddr() net.Addr                { return nil }
func (b *buffer) RemoteAddr() net.Addr               { return nil }
func (b *buffer) SetDeadline(t time.Time) error      { return nil }
func (b *buffer) SetReadDeadline(t time.Time) error  { return nil }
func (b *buffer) SetWriteDeadline(t time.Time) error { return nil }

func TestConfig(t *testing.T) {
	VerifyConfig(DefaultConfig())

	config := DefaultConfig()
	config.KeepAliveInterval = 0
	err := VerifyConfig(config)
	if err == nil {
		t.Fatal(err)
	}

	config = DefaultConfig()
	config.KeepAliveInterval = 10
	config.KeepAliveTimeout = 5
	err = VerifyConfig(config)
	if err == nil {
		t.Fatal(err)
	}

	config = DefaultConfig()
	config.MaxFrameSize = 0
	err = VerifyConfig(config)
	if err == nil {
		t.Fatal(err)
	}

	config = DefaultConfig()
	config.MaxFrameSize = 65536
	err = VerifyConfig(config)
	if err == nil {
		t.Fatal(err)
	}

	config = DefaultConfig()
	config.MaxReceiveBuffer = 0
	err = VerifyConfig(config)
	if err == nil {
		t.Fatal(err)
	}

	var bts buffer
	if _, err := Server(&bts, config); err == nil {
		t.Fatal("server started with wrong config")
	}

	if _, err := Client(&bts, config); err == nil {
		t.Fatal("client started with wrong config")
	}
}
