package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username,omitempty"`
	Password   string `json:"password,omitempty"`
	Email      string `json:"email,omitempty"`
	Role_id    int    `json:"role_id,omitempty"`
	Role_name  string `json:"role_name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Created_by string `json:"created_by,omitempty"`
	Updated_by string `json:"updated_by,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Status     int    `json:"status"`
}

type UsersNameCreate struct {
	Name       string `json:"name" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Role_id    int    `json:"role_id" binding:"required"`
	Created_by int    `json:"created_by" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

type UsersNameUpdate struct {
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Password   string `json:"password"`
	Role_id    int    `json:"role_id" binding:"required"`
	Updated_by int    `json:"updated_by" binding:"required"`
	Status     string `json:"status" binding:"required"`
}

func (h *Handler) usersHandler(c *gin.Context) {
	sql := "SELECT users.id,users.name,users.email,roles.name AS role_name,users.created_at,users.status FROM " + DB_DATABASE + ".users,roles WHERE users.role_id = roles.id"
	results, err := h.db.Query(sql)
	var usersDataList []Users
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var user Users
		err := results.Scan(&user.ID, &user.Name, &user.Email, &user.Role_name, &user.Created_at, &user.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		usersDataList = append(usersDataList, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        usersDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) usersAddHandler(c *gin.Context) {

	var users UsersNameCreate

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	pwd := []byte(users.Password)
	hash := hashAndSalt(pwd)

	statement, _ := h.db.Prepare(`INSERT INTO ` + DB_DATABASE + `.users (username, password,name,email,role_id,created_by, status)VALUES (?,?,?,?,?,?,?)`)
	_, err := statement.Exec(users.Username, hash, users.Name, users.Email, users.Role_id, users.Created_by, users.Status)
	log.Println(users)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Insert" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}

func (h *Handler) usersEditHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	data := &Users{}
	results, err := h.db.Query("SELECT id,username,name,email,role_id,status FROM "+DB_DATABASE+".users where id=?", paramID)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&data.ID, &data.Username, &data.Name, &data.Email, &data.Role_id, &data.Status)
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

	c.JSON(http.StatusOK, gin.H{
		//"id":          paramID,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
		"data":        data,
	})

}

func (h *Handler) usersUpdateHandler(c *gin.Context) {

	var users UsersNameUpdate

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	log.Println(users.Password, reflect.TypeOf(users.Password), len(users.Password))

	if len(users.Password) <= 0 {
		log.Println("Case True")
		statement, _ := h.db.Prepare(`UPDATE ` + DB_DATABASE + `.users SET name=?,role_id=?,updated_by=?,status=? WHERE id=?`)
		_, err := statement.Exec(users.Name, users.Role_id, users.Updated_by, users.Status, users.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case Update" + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}

	} else {
		log.Println("Case False")
		pwd := []byte(users.Password)
		hash := hashAndSalt(pwd)
		statement, _ := h.db.Prepare(`UPDATE ` + DB_DATABASE + `.users SET name=?,password=?,role_id=?,updated_by=?,status=? WHERE id=?`)
		_, err := statement.Exec(users.Name, hash, users.Role_id, users.Updated_by, users.Status, users.ID)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case Update" + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
	}
	log.Println(users)
	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}

func (h *Handler) usersDeleteHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	statement, _ := h.db.Prepare(`DELETE FROM ` + DB_DATABASE + `.users WHERE id=?`)
	_, err := statement.Exec(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Delete" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		//"id":          paramID,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}
