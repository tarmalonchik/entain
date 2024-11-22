package logger

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type Client struct {
	logger *logrus.Logger
}

func NewLogger() *Client {
	return &Client{
		logger: logrus.New(),
	}
}

func (c *Client) Log(ctx context.Context, message string) {
	c.logger.WithContext(ctx).WithError(fmt.Errorf("application runner")).Errorf(message)
}

func (c *Client) Infof(format string, args ...any) {
	c.logger.Infof(format, args...)
}

func (c *Client) Errorf(err error, format string, args ...any) {
	c.logger.WithError(err).Errorf(format, args...)
}
