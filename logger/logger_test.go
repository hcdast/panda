/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:35:45
 * @Description:
 */
package logger_test

import (
	"testing"

	"example-hauth/panda/logger"
)

func TestNewLogger(t *testing.T) {
	logger.Info("hello world abcd")
	logger.Info("my name is huang zhan wei")
	conf := logger.NewConfig()
	conf.SetName("newLogName.log")
	lg := logger.NewLogger(conf)
	lg.Error("hello world this is new logger")

	conf2 := logger.NewConfig("log.conf")
	ll := logger.NewLogger(conf2)
	ll.Info("hello world")
}
