package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	dlog "github.com/fajarhadifirmansyah/bedbo/log"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func DefaultStructuredLogger() gin.HandlerFunc {
	l := dlog.Get()
	return StructuredLogger(&l)
}

func StructuredLogger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		var reqBodyBytes []byte
		if c.Request.Body != nil {
			reqBodyBytes, _ = ioutil.ReadAll(c.Request.Body)

		}
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(reqBodyBytes))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
			defer func() {
				logEvent.Bytes("req_body", reqBodyBytes).Bytes("resp_body", blw.body.Bytes()).Msg(param.ErrorMessage)
			}()
		} else {
			logEvent = logger.Info()
			defer func() {
				logEvent.Bytes("req_body", reqBodyBytes).Bytes("resp_body", blw.body.Bytes()).Msg(param.ErrorMessage)
			}()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String())

	}
}
