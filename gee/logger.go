/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-03-02 19:50:59
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-03-02 19:51:53
 * @FilePath: /gee/gee/logger.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package gee

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
