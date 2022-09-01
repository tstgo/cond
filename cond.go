/*
 * @Author: leoking
 * @Date: 2022-09-01 15:55:43
 * @LastEditTime: 2022-09-01 20:02:21
 * @LastEditors: leoking
 * @Description:
 */

package cond

/*
*
  - @description:
    判断是否所有的元素都满足p
  - @param {Predicate} p
  - @param {...V} args
  - @return {*}
*/
func All[T any](p Predicate[T], args ...T) bool {
	for _, arg := range args {
		if !p(arg) {
			return false
		}
	}
	return true
}

/*
*
  - @description:
    判断是否至少有一个元素满足p
  - @param {Predicate} p
  - @param {...V} args
  - @return {*}
*/
func Any[T any](p Predicate[T], args ...T) bool {
	for _, v := range args {
		if p(v) {
			return true
		}
	}
	return false
}

/*
*
  - @description:
    判断是不是0
  - @param {any} v
  - @return {*}
*/
func IsZero[T Number](v T) bool {
	return v == 0
}
