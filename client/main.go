package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang-grpc-boilerplate/proto"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Name string //`json:"Name"`
	Type string //`json:"Type"`
	Quantity int64 //`json:"Quantity"`
	Price float64
	Tax float64
	Total_Price float64

}

func (i Item) String() string {
	return fmt.Sprintf("Name: %v \n Type: %v \n Quantity: %v \n Price: %v \n", i.Name, i.Type, i.Quantity, i.Price)
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
	g.GET("/database-connection", func(ctx *gin.Context) {
		db, err := sql.Open("mysql","root:nuclei@123@tcp(127.0.0.1:3306)/Nuclei" )
		if err != nil {
			log.Panic(err.Error())
		}
		log.Println( "DB Connection Successful")
		defer db.Close()

		res, err := db.Query("Select * FROM Items")
		if err != nil {
			log.Panic(err.Error())
		}

		var items []Item
		for res.Next() {
			var i Item
			err = res.Scan(&i.Name, &i.Type, &i.Quantity, &i.Price)
			if err != nil {
				log.Panic(err.Error())
			}
			items = append(items, i)
			log.Print(i)
		}
		ctx.JSON(http.StatusOK, gin.H{"items": fmt.Sprint(items)})

	})
	if err := g.Run(":8080"); err!=nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}