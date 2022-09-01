/*
 * @Author: leoking
 * @Date: 2022-09-01 15:54:57
 * @LastEditTime: 2022-09-01 16:18:51
 * @LastEditors: leoking
 * @Description:
 */
package cond

/*
*
  - @description:
    判断两个参数是否相等
  - @return {*}
*/
func Equal[T comparable](a, b T) bool {
	return a == b
}
