/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:34:13
 * @Description:
 */
package crypto_test

import (
	"fmt"
	"testing"

	"example-hauth/panda/crypto"
)

func TestSha1(t *testing.T) {
	fmt.Println(crypto.Sha1("hello world", "abc"))
	newSha1 := crypto.NewSHA1("hzwy23")
	fmt.Println(newSha1.Sha1("zhanwei", "wei"))
	fmt.Println(newSha1.Sha1("hello world", "abc"))
}

func TestNewSHA1(t *testing.T) {
	fmt.Println(crypto.Sha1("hello world", "abc"))

	newSha1 := crypto.NewSHA1("huang")
	fmt.Println(newSha1.Sha1("zhanwei", "wei"))
	fmt.Println(newSha1.Sha1("hello world", "abc"))
}
