/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:37:04
 * @Description:
 */
package panda_test

import (
	"example-hauth/panda"
	"fmt"
	"testing"
)

func TestDateFormat(t *testing.T) {
	fmt.Println(panda.DateFormat("2017-12-06", "YYYY-MM-DD"))
	fmt.Println(panda.DateFormat("2017-12-06 11:03:21", "YYYY-MM-DD HH:MM:SS"))
	fmt.Println(panda.DateFormat("2017-12-06 21:03:21", "YYYY-MM-DD HH24:MM:SS"))
}
