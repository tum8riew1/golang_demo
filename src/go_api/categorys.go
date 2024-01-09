package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Categorys struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	Created_by        string `json:"created_by,omitempty"`
	Updated_by        string `json:"updated_by,omitempty"`
	Created_at        string `json:"created_at,omitempty"`
	Updated_at        string `json:"updated_at,omitempty"`
	Status            int    `json:"status"`
}

type CategorysCreate struct {
	Title             string `json:"title" binding:"required"`
	Description       string `json:"description"`
	Short_description string `json:"short_description"`
	Created_by        int    `json:"created_by" binding:"required"`
	Status            string `json:"status" binding:"required"`
}

type CategorysUpdate struct {
	ID                int    `json:"id" binding:"required"`
	Title             string `json:"title" binding:"required"`
	Description       string `json:"description"`
	Short_description string `json:"short_description"`
	Updated_by        int    `json:"updated_by" binding:"required"`
	Status            string `json:"status" binding:"required"`
}

func (h *Handler) categorysHandler(c *gin.Context) {
	sql := "SELECT id,title,short_description,created_at,status FROM " + DB_DATABASE + ".article_category"
	results, err := h.db.Query(sql)
	var categorysDataList []Categorys
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var category Categorys
		err := results.Scan(&category.ID, &category.Title, &category.Short_description, &category.Created_at, &category.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		categorysDataList = append(categorysDataList, category)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        categorysDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) categorysDropdownlisHandler(c *gin.Context) {
	sql := "SELECT id,title,status FROM " + DB_DATABASE + ".article_category WHERE status ='1'"
	results, err := h.db.Query(sql)
	var categorysDataList []Categorys
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var category Categorys
		err := results.Scan(&category.ID, &category.Title, &category.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		categorysDataList = append(categorysDataList, category)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        categorysDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) categorysAddHandler(c *gin.Context) {

	var categorys CategorysCreate

	if err := c.ShouldBindJSON(&categorys); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	statement, _ := h.db.Prepare(`INSERT INTO ` + DB_DATABASE + `.article_category (title, description,short_description,created_by, status)VALUES (?,?,?,?,?)`)
	_, err := statement.Exec(categorys.Title, categorys.Description, categorys.Short_description, categorys.Created_by, categorys.Status)
	log.Println(categorys)
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

func (h *Handler) categorysEditHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	data := &Categorys{}
	results, err := h.db.Query("SELECT id,title, description,short_description,status FROM "+DB_DATABASE+".article_category where id=?", paramID)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&data.ID, &data.Title, &data.Description, &data.Short_description, &data.Status)
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

func (h *Handler) categorysUpdateHandler(c *gin.Context) {

	var categorys CategorysUpdate

	if err := c.ShouldBindJSON(&categorys); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	statement, _ := h.db.Prepare(`UPDATE ` + DB_DATABASE + `.article_category SET title=?,description=?,short_description=?,updated_by=?,status=? WHERE id=?`)
	_, err := statement.Exec(categorys.Title, categorys.Description, categorys.Short_description, categorys.Updated_by, categorys.Status, categorys.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Update" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	log.Println(categorys)
	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}

func (h *Handler) categorysDeleteHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	statement, _ := h.db.Prepare(`DELETE FROM ` + DB_DATABASE + `.article_category WHERE id=?`)
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
