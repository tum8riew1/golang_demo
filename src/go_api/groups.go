package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Groups struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Guard_name string `json:"guard_name,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Updated_at string `json:"updated_at,omitempty"`
	Status     int    `json:"status,omitempty"`
}

type GroupsNameCreate struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type GroupsNameUpdate struct {
	ID     int    `json:"id" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (h *Handler) groupsHandler(c *gin.Context) {

	results, err := h.db.Query("SELECT id,name,guard_name,created_at,status FROM " + DB_DATABASE + ".roles")
	var groupsDataList []Groups
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var group Groups
		err := results.Scan(&group.ID, &group.Name, &group.Guard_name, &group.Created_at, &group.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		groupsDataList = append(groupsDataList, group)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        groupsDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) groupsDropdownlistHandler(c *gin.Context) {

	results, err := h.db.Query("SELECT id,name FROM " + DB_DATABASE + ".roles WHERE status='1'")
	var groupsDataList []Groups
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var group Groups
		err := results.Scan(&group.ID, &group.Name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		groupsDataList = append(groupsDataList, group)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        groupsDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) groupsAddHandler(c *gin.Context) {

	var groups GroupsNameCreate

	if err := c.ShouldBindJSON(&groups); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	statement, _ := h.db.Prepare(`INSERT INTO ` + DB_DATABASE + `.roles (name, guard_name, status)VALUES (?,?,?)`)
	_, err := statement.Exec(groups.Name, "web", groups.Status)
	log.Println(groups)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Insert" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	// ID         int    `json:"id"`
	// Name       string `json:"name"`
	// Guard_name string `json:"guard_name,omitempty"`
	// Created_at string `json:"created_at,omitempty"`
	// Updated_at string `json:"updated_at,omitempty"`
	// Status     int    `json:"status"`

	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}

func (h *Handler) groupsEditHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	data := &Groups{}
	results, err := h.db.Query("SELECT id,name,status FROM "+DB_DATABASE+".roles where id=?", paramID)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&data.ID, &data.Name, &data.Status)
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

func (h *Handler) groupsUpdateHandler(c *gin.Context) {

	var groups GroupsNameUpdate

	if err := c.ShouldBindJSON(&groups); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}
	statement, _ := h.db.Prepare(`UPDATE ` + DB_DATABASE + `.roles SET name=?,status=? WHERE id=?`)
	_, err := statement.Exec(groups.Name, groups.Status, groups.ID)
	log.Println(groups)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Update" + err.Error(),
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

func (h *Handler) groupsDeleteHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	statement, _ := h.db.Prepare(`DELETE FROM ` + DB_DATABASE + `.roles WHERE id=?`)
	_, err := statement.Exec(paramID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Delete" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          paramID,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}
