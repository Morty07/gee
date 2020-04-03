package main

import (
	"fmt"
	"gee/gee"
)

// 1.
// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/hello", helloHandler) //使用默认的处理器 ServeMux
// 	// http.Handle() 可以自己传入自己实现的处理器
// 	log.Fatal(http.ListenAndServe(":8000", nil))
// }

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	for k, v := range r.Header {
// 		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 	}
// }

// 2.
// type Engine struct{}

// func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
// 	case "/hello":
// 		for k, v := range r.Header {
// 			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
// 		}
// 	default:
// 		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL.Path)
// 	}
// }

// // 此时我们自己实现了处理逻辑，在以上的注释代码中，我们只能用http.HandleFunc实现路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。
// // 但是实现了Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口，在这里我们可以自定义路由映射的规则，也可以添加一些处理逻辑，比如日志，异常处理等。
// func main() {
// 	engine := new(Engine) // 我们自己实现的处理器
// 	http.ListenAndServe(":8000", engine)
// }

// 3.
func main() {
	g := gee.New()
	g.GET("/", func(c *gee.Context) {
		fmt.Fprintf(c.W, "URL.Path = %q\n", c.Path)
	})
	g.GET("/hello", func(c *gee.Context) {
		for k, v := range c.R.Header {
			fmt.Fprintf(c.W, "Header[%q] = %q\n", k, v)
		}
	})
	g.Run(":8000")
}
