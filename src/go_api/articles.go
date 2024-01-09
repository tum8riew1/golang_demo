package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Articles struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Description       string `json:"description,omitempty"`
	Short_description string `json:"short_description,omitempty"`
	Category_name     string `json:"category_name,omitempty"`
	Category_id       int    `json:"category_id,omitempty"`
	Created_by        string `json:"created_by,omitempty"`
	Updated_by        string `json:"updated_by,omitempty"`
	Created_at        string `json:"created_at,omitempty"`
	Updated_at        string `json:"updated_at,omitempty"`
	Status            int    `json:"status"`
}

type ArticlesCreate struct {
	Title             string `json:"title" binding:"required"`
	Description       string `json:"description"`
	Short_description string `json:"short_description"`
	Category_id       int    `json:"category_id" binding:"required"`
	Created_by        int    `json:"created_by" binding:"required"`
	Status            string `json:"status" binding:"required"`
}

type ArticlesUpdate struct {
	ID                int    `json:"id" binding:"required"`
	Title             string `json:"title" binding:"required"`
	Description       string `json:"description"`
	Short_description string `json:"short_description"`
	Category_id       int    `json:"category_id" binding:"required"`
	Updated_by        int    `json:"updated_by" binding:"required"`
	Status            string `json:"status" binding:"required"`
}

func (h *Handler) articlesHandler(c *gin.Context) {
	sql := "SELECT article.id,article.title,article_category.title AS category_name,article.created_at,article.status FROM " + DB_DATABASE + ".article,article_category WHERE article.category_id = article_category.id"
	results, err := h.db.Query(sql)
	var articlesDataList []Articles
	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	for results.Next() {
		var article Articles
		err := results.Scan(&article.ID, &article.Title, &article.Category_name, &article.Created_at, &article.Status)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message":     "Case2 " + err.Error(),
				"status_code": http.StatusInternalServerError,
			})
			return
		}
		articlesDataList = append(articlesDataList, article)
	}
	c.JSON(http.StatusOK, gin.H{
		"message":     http.StatusText(http.StatusOK),
		"data":        articlesDataList,
		"status_code": http.StatusOK,
	})

}

func (h *Handler) articlesAddHandler(c *gin.Context) {

	var articles ArticlesCreate

	if err := c.ShouldBindJSON(&articles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	statement, _ := h.db.Prepare(`INSERT INTO ` + DB_DATABASE + `.article (title, description,short_description,category_id,created_by, status)VALUES (?,?,?,?,?,?)`)
	_, err := statement.Exec(articles.Title, articles.Description, articles.Short_description, articles.Category_id, articles.Created_by, articles.Status)
	log.Println(articles)
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

func (h *Handler) articlesEditHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	data := &Articles{}
	results, err := h.db.Query("SELECT id,title, description,short_description,category_id,status FROM "+DB_DATABASE+".article where id=?", paramID)

	if err != nil {
		//og.Println("Err", err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case1 " + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	if results.Next() {
		err = results.Scan(&data.ID, &data.Title, &data.Description, &data.Short_description, &data.Category_id, &data.Status)
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

func (h *Handler) articlesUpdateHandler(c *gin.Context) {

	var articles ArticlesUpdate

	if err := c.ShouldBindJSON(&articles); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message":     http.StatusText(http.StatusBadRequest),
			"status_code": http.StatusBadRequest,
		})
		return
	}

	statement, _ := h.db.Prepare(`UPDATE ` + DB_DATABASE + `.article SET title=?,description=?,short_description=?,category_id=?,updated_by=?,status=? WHERE id=?`)
	_, err := statement.Exec(articles.Title, articles.Description, articles.Short_description, articles.Category_id, articles.Updated_by, articles.Status, articles.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message":     "Case Update" + err.Error(),
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	log.Println(articles)
	c.JSON(http.StatusOK, gin.H{
		//"id":          user.Username,
		"message":     http.StatusText(http.StatusOK),
		"status_code": http.StatusOK,
	})

}

func (h *Handler) articlesDeleteHandler(c *gin.Context) {

	paramID := c.Param("id")

	// if paramID == nil || len(paramID) <= 0 {
	// 	return http.StatusBadRequest, errors.New(http.StatusText(http.StatusBadRequest))
	// }

	statement, _ := h.db.Prepare(`DELETE FROM ` + DB_DATABASE + `.article WHERE id=?`)
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
