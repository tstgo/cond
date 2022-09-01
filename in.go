/*
 * @Author: leoking
 * @Date: 2022-09-01 15:51:43
 * @LastEditTime: 2022-09-01 16:20:33
 * @LastEditors: leoking
 * @Description:
 */
package cond

/*
*
  - @description:
    判断x是否在数组中
  - @param {any} x
  - @param {...any} arr
  - @return {*}
*/
func InArray[T comparable](x T, args ...T) bool {
	return Any(func(t T) bool {
		return x == t
	}, args...)
}
