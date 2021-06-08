/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:36:04
 * @Description:
 */
package route_test

import (
	"example-hauth/panda/route"
	"fmt"
	"net/http"
	"testing"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world", w, r)
}

func Index2(w http.ResponseWriter, r *http.Request, ps route.Params) {
	fmt.Println(ps, w, r)
}

func TestGET(t *testing.T) {
	mux := route.DefaultRouter()
	route.Handler("GET", "/", Index)
	route.GET("/:httprouter", Index2)
	http.ListenAndServe(":8080", mux)
}

func TestNewRouter(t *testing.T) {
	mux := route.NewRouter()
	mux.HandlerFunc("GET", "/hello", Index)
	mux.GET("/", Index2)
	http.ListenAndServe(":8090", mux)
}
