package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!")
	})

	h := membershipHandler{}
	h.Initialize()

	e.GET("/memberships", h.GetAllmembership)
	e.POST("/memberships", h.Savemembership)
	e.GET("/memberships/:id", h.Getmembership)
	e.PUT("/memberships/:id", h.Updatemembership)
	e.DELETE("/memberships/:id", h.Deletemembership)

	e.Logger.Fatal(e.Start(":1323"))
}

type membershipHandler struct {
	DB *gorm.DB
}

func (h *membershipHandler) Initialize() {
	db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&membership{})

	h.DB = db
}

type membership struct {
	MemId   uint   `gorm:"primary_key" json:"MemberId"`
	MemName string `json:"Membername"`
	MemTel  string `json:"MemberTelephone"`
	Email   string `json:"email"`
}

func (h *membershipHandler) GetAllmembership(c echo.Context) error {
	memberships := []membership{}

	h.DB.Find(&memberships)

	return c.JSON(http.StatusOK, memberships)
}

func (h *membershipHandler) Getmembership(c echo.Context) error {
	id := c.Param("id")
	membership := membership{}

	if err := h.DB.Find(&membership, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, membership)
}

func (h *membershipHandler) Savemembership(c echo.Context) error {
	membership := membership{}

	if err := c.Bind(&membership); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&membership).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, membership)
}

func (h *membershipHandler) Updatemembership(c echo.Context) error {
	id := c.Param("id")
	membership := membership{}

	if err := h.DB.Find(&membership, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&membership); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&membership).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, membership)
}

func (h *membershipHandler) Deletemembership(c echo.Context) error {
	id := c.Param("id")
	membership := membership{}

	if err := h.DB.Find(&membership, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Delete(&membership).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
