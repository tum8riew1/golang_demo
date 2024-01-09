package main

import (
	"database/sql"
	"internal/validate"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const DB_CONNECTION = "mysql"
const DB_HOST = "192.168.1.2"
const DB_PORT = "3306"
const DB_USERNAME = "xxxxxxx"
const DB_PASSWORD = "xxxxxxx"
const DB_DATABASE = "xxxxxxx"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Role_id  string `json:"role_id"`
	Status   int    `json:"status"`
}

type JSONTime time.Time

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserGetByUsername struct {
	Username string `json:"username" binding:"required"`
}

type Handler struct {
	db *sql.DB
}

func newHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	db, err := sql.Open(DB_CONNECTION, DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_DATABASE+"?parseTime=true")
	defer db.Close()
	if err != nil {
		//log.Println("connect database fail !")
		panic("connect database fail !")
	} else {
		log.Println("connect database success !")
	}
	handler := newHandler(db)

	router := gin.Default()
	router.Use(corsMiddleware())
	router.POST("/api/v1/backend/login", handler.loginHandler)
	//router.GET("/api/v1/hashpassword", handler.hashpasswordHandler)
	//router.GET("/api/v1/checkpassword", handler.checkpasswordHandler)
	v1_backend := router.Group("/api/v1/backend", authorizationMiddleware)
	{
		v1_backend.POST("/dashboard", handler.dashboardHandler)
		v1_backend.POST("/users/get-by-username", handler.getUserByUsernameHandler)
		//Groups
		v1_backend.POST("/groups", handler.groupsHandler)
		v1_backend.POST("/groups-dropdownlist", handler.groupsDropdownlistHandler)
		v1_backend.POST("/groups/create", handler.groupsAddHandler)
		v1_backend.GET("/groups/edit/:id", handler.groupsEditHandler)
		v1_backend.POST("/groups/update", handler.groupsUpdateHandler)
		v1_backend.GET("/groups/delete/:id", handler.groupsDeleteHandler)
		//Users
		v1_backend.POST("/users", handler.usersHandler)
		v1_backend.POST("/users/create", handler.usersAddHandler)
		v1_backend.GET("/users/edit/:id", handler.usersEditHandler)
		v1_backend.POST("/users/update", handler.usersUpdateHandler)
		v1_backend.GET("/users/delete/:id", handler.usersDeleteHandler)
		//Categorys
		v1_backend.POST("/categorys", handler.categorysHandler)
		v1_backend.POST("/categorys-dropdownlist", handler.categorysDropdownlisHandler)
		v1_backend.POST("/categorys/create", handler.categorysAddHandler)
		v1_backend.GET("/categorys/edit/:id", handler.categorysEditHandler)
		v1_backend.POST("/categorys/update", handler.categorysUpdateHandler)
		v1_backend.GET("/categorys/delete/:id", handler.categorysDeleteHandler)
		//Categorys
		v1_backend.POST("/articles", handler.articlesHandler)
		v1_backend.POST("/articles/create", handler.articlesAddHandler)
		v1_backend.GET("/articles/edit/:id", handler.articlesEditHandler)
		v1_backend.POST("/articles/update", handler.articlesUpdateHandler)
		v1_backend.GET("/articles/delete/:id", handler.articlesDeleteHandler)
	}

	s := &http.Server{
		Addr:           ":3001",
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func authorizationMiddleware(c *gin.Context) {

	if v_code, err := validate.ValidateHeader(c); err != nil {
		c.AbortWithStatusJSON(v_code, gin.H{
			"message":     err.Error(),
			"status_code": v_code,
		})
		return
	}

	s := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(s, "Bearer ")

	if err := validate.ValidateToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":     err.Error(),
			"status_code": http.StatusUnauthorized,
		})
		return
	}
}

func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func (h *Handler) hashpasswordHandler(c *gin.Context) {

	pwd := []byte("secret")
	hash := hashAndSalt(pwd)

	c.JSON(http.StatusOK, gin.H{
		"password":    hash,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) checkpasswordHandler(c *gin.Context) {

	password := "secret"
	hash := "$2a$04$MvzbGTn3YwULXO0.pBIZj.68mH9zEZyCj0b1Xd27BUVJlgtCNGy3i"
	match := CheckPasswordHash(password, hash)
	c.JSON(http.StatusOK, gin.H{
		"password":    password,
		"hash":        hash,
		"match":       match,
		"status_code": http.StatusOK,
	})

}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (h *Handler) loginHandler(c *gin.Context) {

	var user_login UserLogin
	//c.ShouldBindJSON(&user_login)
	if err := c.ShouldBindJSON(&user_login); err != nil {
		//.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	prod := &User{}
	results, err := h.db.Query("SELECT id,username,name,email,password FROM "+DB_DATABASE+".users where status =1 AND username =?", user_login.Username)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&prod.ID, &prod.Username, &prod.Name, &prod.Email, &prod.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}

		password := user_login.Password
		hash := prod.Password
		if match := CheckPasswordHash(password, hash); match != true {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":     "Username & Password Invalid !",
				"status_code": http.StatusUnauthorized,
			})
			return
		}

	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case3 ",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		Id:        prod.Email,
	})
	//log.Println("User id --->", string(prod.Email), prod.Email)
	ss, err := token.SignedString([]byte("Echo31"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": http.StatusText(http.StatusOK),
		"token":   ss,
		//"user":        prod.Username,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) dashboardHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        "test",
		"status_code": http.StatusOK,
	})

}

func (h *Handler) getUserByUsernameHandler(c *gin.Context) {

	var user UserGetByUsername

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	data := &User{}
	results, err := h.db.Query("SELECT id,name,email,avatar,role_id,status FROM "+DB_DATABASE+".users where status =1 AND username =?", user.Username)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&data.ID, &data.Name, &data.Email, &data.Avatar, &data.Role_id, &data.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}

	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case3 ",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	//paramID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"data":        data,
		"status_code": http.StatusOK,
	})

}

func StringToDate(date string) *time.Time {
	parsed, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return &parsed
}
