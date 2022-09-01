/*
 * @Author: leoking
 * @Date: 2022-09-01 15:51:43
 * @LastEditTime: 2022-09-01 16:21:54
 * @LastEditors: leoking
 * @Description:
 */
package cond

import "testing"

func TestInArray(t *testing.T) {
	t.Log(InArray(10, 1, 2, 3, 4, 5))
	t.Log(InArray("", "x", "y", ""))
}
