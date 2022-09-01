/*
 * @Author: leoking
 * @Date: 2022-09-01 20:00:55
 * @LastEditTime: 2022-09-01 20:00:55
 * @LastEditors: leoking
 * @Description:
 */
package cond

type Predicate[T any] func(T) bool

type Number interface {
	uint8 | uint16 | uint | uint32 | uint64 | int8 | int16 | int | int32 | int64 | float32 | float64
}
