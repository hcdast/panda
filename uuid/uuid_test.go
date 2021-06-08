/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:36:43
 * @Description:
 */
package uuid_test

import (
	"fmt"
	"testing"

	"example-hauth/panda/uuid"
)

func TestUuid(t *testing.T) {
	fmt.Println(uuid.Random())
	fmt.Println(uuid.Random())
	fmt.Println(uuid.Random())
	fmt.Println(uuid.Random())

	fmt.Println(uuid.UUID())
	fmt.Println(uuid.UUID())
	fmt.Println(uuid.UUID())
}
