/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:35:06
 * @Description:
 */
package jwt_test

import (
	"testing"

	"fmt"

	"example-hauth/panda/jwt"
)

func TestToken(t *testing.T) {
	token, err := jwt.GenToken(jwt.NewUserdata().SetUserId("huang").SetOrgunitId("100"))
	fmt.Println(token, err)
	claims, err := jwt.ParseToken(token)
	fmt.Println(claims, err)

	flag := jwt.ValidToken(token)
	fmt.Println("true or false:", flag)

	handle := jwt.NewHandle(nil)
	flag = handle.ValidToken(token)
	fmt.Println("true or false:", flag)
	token, err = handle.GenToken(jwt.NewUserdata().SetUserId("huang").SetOrgunitId("100"))
	fmt.Println(token, err)
	claims, err = jwt.ParseToken(token)
	fmt.Println(claims, err)

	flag = jwt.ValidToken(token)
	fmt.Println("true or false:", flag)
}
