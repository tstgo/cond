/*
 * @Author: leoking
 * @Date: 2022-09-01 16:03:10
 * @LastEditTime: 2022-09-01 16:16:58
 * @LastEditors: leoking
 * @Description:
 */
package cond

import "testing"

func TestAll(t *testing.T) {
	t.Log(All(func(t int) bool {
		return t > 0
	}, 1, 2, 3, 4))
	t.Log(All(func(s string) bool {
		return len(s) == 0
	}, "", "xx"))
}

func TestAny(t *testing.T) {
	t.Log(Any(func(x int) bool {
		return x > 5
	}, 0, 1))
	t.Log(Any(func(x int) bool {
		return x >= 5
	}, 5, 1))
}
