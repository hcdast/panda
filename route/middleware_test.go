/*
 * @Author: hc
 * @Date: 2021-06-07 11:04:31
 * @LastEditors: hc
 * @LastEditTime: 2021-06-07 11:36:28
 * @Description:
 */
package route_test

import (
	"example-hauth/panda/route"
	"fmt"
	"net/http"
	"testing"
)

type a struct {
}

func (a) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("hello func1")
	next(w, r)
}

type b struct {
}

func (b) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("hello func2")
	next(w, r)
}

func TestNewMiddle(t *testing.T) {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
	})

	handle := route.Wrap(mux)
	mid := route.NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}

func TestNewMiddleware(t *testing.T) {
	mux := route.NewRouter()

	mux.POST("/:name/:bcd", Index2)
	mux.HandlerFunc("POST", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hi HandleFunc")
	})

	handle := route.Wrap(mux)
	mid := route.NewMiddleware(handle, &b{}, &a{})
	http.ListenAndServe(":8080", mid)
}
