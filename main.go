package main

import (
	"fmt"

	"github.com/fajarhadifirmansyah/bedbo/api"
	"github.com/fajarhadifirmansyah/bedbo/api/middleware"
	cv "github.com/fajarhadifirmansyah/bedbo/api/validator"
	"github.com/fajarhadifirmansyah/bedbo/config"
	"github.com/fajarhadifirmansyah/bedbo/store"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("BE Test DBO: Fajar Hadifirmansyah")
	decimal.MarshalJSONWithoutQuotes = true

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.DefaultStructuredLogger())
	r.Use(gin.Recovery())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("gender", cv.GenderValidate)
	}

	configDB := &config.DBConfig{}
	db := configDB.PostgreSql()

	v1 := r.Group("api/v1")

	customerRoutes(v1, db)
	productRoutes(v1, db)
	ordersRoutes(v1, db)
	authRoutes(v1, db)
	userRoutes(v1, db)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}

func customerRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/customers")

	store := store.NewCustomerStore(db)
	handler := api.NewCustomerHandler(store)

	r.Use(middleware.Auth())
	r.GET("/", handler.GetCustomersHandler)
	r.POST("/", handler.CreateCustomerHandler)
	r.GET(":customerid", handler.GetCustomerByIDHandler)
	r.DELETE(":customerid", handler.DeleteCustomerHandler)
	r.PUT(":customerid", handler.UpdateCustomerHandler)
}

func productRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/products")

	store := store.NewProductStore(db)
	handler := api.NewProductHandler(store)

	r.Use(middleware.Auth())
	r.GET("/", handler.GetAll)
}

func ordersRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/orders")

	store := store.NewOrderStore(db)
	handler := api.NewOrderHandler(store)

	r.Use(middleware.Auth())
	r.GET("/", handler.GetOrderHandler)
	r.POST("/", handler.CreateOrderHandler)
	r.GET(":orderID", handler.GetByIDHandler)
	r.PUT(":orderID", handler.UpdateOrderHandler)
	r.DELETE(":orderID", handler.DeleteOrderHandler)

}

func authRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/auth")

	aStore := store.NewAuthStore(db)
	uStore := store.NewUserStore(db)
	handler := api.NewAuthHandler(aStore, uStore)

	r.POST("/register", handler.SignUpHandler)
	r.POST("/login", handler.LoginHandler)
	r.POST("/refresh", handler.RefreshAccessToken)
	r.GET("/logout", middleware.DeserializeUser(uStore), handler.RefreshAccessToken)
}

func userRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	r := rg.Group("/users")

	store := store.NewUserStore(db)
	handler := api.NewUserHandler(store)

	r.Use(middleware.DeserializeUser(store))
	r.GET("/me", handler.GetMe)
}
