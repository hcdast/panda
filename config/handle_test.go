/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:33:58
 * @Description:
 */
package config_test

import (
	"fmt"
	"testing"

	"example-hauth/panda/config"
)

func TestLoad(t *testing.T) {
	c, err := config.Load("./testData/data.txt")
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.INI)
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.YAML)
	fmt.Println(c, err)
	c, err = config.Load("./testData/data.txt", config.JSON)
	fmt.Println(c, err)

	fmt.Println(c.Set("myyfwe", "demfdsdfo"))
	fmt.Println(c.Get("abc"))
	fmt.Println(c.Get("hello"))
}
