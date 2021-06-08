/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:34:44
 * @Description:
 */
package hret_test

import (
	"example-hauth/panda/hret"
	"fmt"
	"testing"
)

func TestHttpPanic(t *testing.T) {
	defer hret.RecvPanic(func() {
		fmt.Println("捕获异常")
	})
	panic("抛出异常")
}
