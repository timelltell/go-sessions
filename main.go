package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	//"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)
func MyMiddleware1(ctx *gin.Context){
	fmt.Println("I am MyMiddleware1")
	ctx.Next()
	//ctx.Abort()
	fmt.Println("I am MyMiddleware1 end")

}
func MyMiddleware2() gin.HandlerFunc{
	return func(ctx *gin.Context){
		fmt.Println("I am MyMiddleware2")
		ctx.Next()
		return
		//ctx.Abort()
		fmt.Println("I am MyMiddleware2 end")
	}
}
func main(){
	router:=gin.Default()
	//store, _ := redis.NewStore(10, "tcp", "localhost: "", []byte("secret"))
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: 0,
	})
	router.Use(MyMiddleware1)
	router.Use(MyMiddleware2())
	router.Use(sessions.Sessions("mysession", store))
	router.GET("/test",func(ctx *gin.Context){
		//设置session
		s:=sessions.Default(ctx)
		s.Set("timellchen","cb")
		s.Save()

		v:=s.Get("timellchen")
		fmt.Println("v: ",v)

		ctx.Writer.WriteString("cookie test")
		fmt.Println("success")
	})
	router.Group("/v1")
	router.Run(":9999")
}
func main_cookie(){
	router:=gin.Default()
	router.GET("/test",func(ctx *gin.Context){
		//设置session
		//ctx.SetCookie("timelltell","my cookie",600,"","",false,true)

		//拿cookie
		cookie, err:=ctx.Cookie("timelltell")
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println("cookie: ",cookie)
		ctx.Writer.WriteString("cookie test")
		fmt.Println("success")
	})
	router.Run(":9999")
}
