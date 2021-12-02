package redis

import (
	"bytes"
	"context"
	"io"

	"github.com/CloudyKit/jet/v6"
	"github.com/syncfuture/go/serr"
	"github.com/syncfuture/go/sredis"
	"github.com/syncfuture/go/u"
)

type RedisTemplateLoader struct {
	redisBase
}

func NewRedisTemplateLoader(redisConfig *sredis.RedisConfig) jet.Loader {
	r := new(RedisTemplateLoader)
	r.redisClient = sredis.NewClient(redisConfig)
	return r
}

func (x *RedisTemplateLoader) Exists(templatePath string) bool {
	ex := x.redisClient.HExists(context.Background(), "xxx:TEMPLATE", templatePath)
	r, err := ex.Result()
	if u.LogError(err) {
		return false
	}

	return r
}

func (x *RedisTemplateLoader) Open(templatePath string) (io.ReadCloser, error) {
	s := x.redisClient.HGet(context.Background(), "xxx:TEMPLATE", templatePath)
	b, err := s.Bytes()
	if err != nil {
		return nil, serr.WithStack(err)
	}
	return io.NopCloser(bytes.NewReader(b)), nil
}
