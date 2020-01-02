package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-grpc-boilerplate/myDB"
	"golang-grpc-boilerplate/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewCalculatorServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseFloat(ctx.Param("a"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		b, err := strconv.ParseFloat(ctx.Param("b"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		req := &proto.Request{A: float64(a), B:float64(b)}
		if response, err := client.Add(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"ans": fmt.Sprint(response.Ans)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})

		}	})
	g.GET("/sub/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseFloat(ctx.Param("a"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		b, err := strconv.ParseFloat(ctx.Param("b"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		req := &proto.Request{A: float64(a), B:float64(b)}
		if response, err := client.Subtract(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"ans": fmt.Sprint(response.Ans)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		}
	})
	g.GET("/mul/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseFloat(ctx.Param("a"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		b, err := strconv.ParseFloat(ctx.Param("b"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		req := &proto.Request{A: float64(a), B:float64(b)}
		if response, err := client.Multiply(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"ans": fmt.Sprint(response.Ans)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})

		}
	})
	g.GET("/div/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseFloat(ctx.Param("a"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		b, err := strconv.ParseFloat(ctx.Param("b"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		req := &proto.Request{A: float64(a), B:float64(b)}
		if response, err := client.Divide(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"ans": fmt.Sprint(response.Ans)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})

		}
	})
	g.GET("/pow/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseFloat(ctx.Param("a"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		b, err := strconv.ParseFloat(ctx.Param("b"), 10)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		req := &proto.Request{A: float64(a), B:float64(b)}
		if response, err := client.Power(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"ans": fmt.Sprint(response.Ans)})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})

		}
	})
	g.GET("/entries", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"items": myDB.GetAllEntry()})

	})
	g.GET("/entries/:id", func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err!= nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		}
		ctx.JSON(http.StatusOK, gin.H{"item": myDB.GetEntry(id)})

	})
	g.POST("/entries", func(ctx *gin.Context) {

		var temp Request
		err := ctx.Bind(&temp)
		if err != nil{
			log.Panic(err.Error())
		}
		if err := myDB.AddEntry(temp.Name, int(temp.Age)); err!=nil {
			log.Panic(err.Error())
		}
	})
	if err := g.Run(":8080"); err!=nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}