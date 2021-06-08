/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:37:18
 * @Description:
 */
package panda_test

import (
	"example-hauth/panda"
	"fmt"
	"testing"
)

func TestGetKey(t *testing.T) {
	s := panda.JoinKey("hello world", "my name is", "c huang zhan wei")
	fmt.Println(s)
	fmt.Println(panda.GetKey(s, 4))
	fmt.Println([]byte(s))
	fmt.Println(byte(0x1f))
}
