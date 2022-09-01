/*
 * @Author: leoking
 * @Date: 2022-09-01 16:03:10
 * @LastEditTime: 2022-09-01 19:57:44
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

func TestIsZero(t *testing.T) {
	t.Log(IsZero(0))
	t.Log(IsZero(int32(0)))
	t.Log(IsZero(0.0))
	t.Log(IsZero(float32(0)))
	t.Log(IsZero(uint8(0)))
}
