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

	h := abcHandler{}
	h.Initialize()

	e.GET("/memberships", h.GetAllmembership)
	e.POST("/memberships", h.Savemembership)
	e.GET("/memberships/:id", h.Getmembership)
	e.PUT("/memberships/:id", h.Updatemembership)
	e.DELETE("/memberships/:id", h.Deletemembership)

	e.GET("/bookstores", h.GetAllbook)
	e.POST("/bookstores", h.Savebook)
	e.GET("/bookstores/:bookid", h.Getbook)
	e.PUT("/bookstores/:bookid", h.Updatebook)
	e.DELETE("/bookstores/:bookid", h.Deletebook)

	e.GET("/Orders", h.GetAllOrder)
	e.POST("/Orders", h.SaveOrder)
	e.GET("/Orders/:Orderid", h.GetOrder)
	e.PUT("/Orders/:Orderid", h.UpdateOrder)
	e.DELETE("/Orders/:Orderid", h.DeleteOrder)

	e.Logger.Fatal(e.Start(":1323"))
}

type abcHandler struct {
	DB *gorm.DB
}

type membership struct {
	MemId   uint   `gorm:"primary_key" json:"MemberId"`
	MemName string `json:"Membername"`
	MemTel  string `json:"MemberTelephone"`
	Email   string `json:"email"`
}

type bookstore struct {
	BookID    uint   `gorm:"primary_key" json:"BookID"`
	BookName  string `json:"Bookname"`
	Quatity   uint   `json:"Quatity"`
	BookPrice uint   `json:"BookPrice"`
	Status    string `json:"Status"`
}

type Order struct {
	OrderID  uint `gorm:"primary_key" json:"OrderID"`
	BookOrID uint `json:"BookID"`
	MemOrID  uint `json:"MemberId"`
}

func (h *abcHandler) Initialize() {
	db, err := gorm.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&membership{})
	db.AutoMigrate(&bookstore{})
	db.AutoMigrate(&Order{})

	h.DB = db
}

func (h *abcHandler) GetAllmembership(c echo.Context) error {
	memberships := []membership{}

	h.DB.Find(&memberships)

	return c.JSON(http.StatusOK, memberships)
}

func (h *abcHandler) Getmembership(c echo.Context) error {
	id := c.Param("id")
	membership := membership{}

	if err := h.DB.Find(&membership, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, membership)
}

func (h *abcHandler) Savemembership(c echo.Context) error {
	membership := membership{}

	if err := c.Bind(&membership); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&membership).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, membership)
}

func (h *abcHandler) Updatemembership(c echo.Context) error {
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

func (h *abcHandler) Deletemembership(c echo.Context) error {
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

//============================================================================

func (h *abcHandler) GetAllbook(c echo.Context) error {
	bookstores := []bookstore{}

	h.DB.Find(&bookstores)

	return c.JSON(http.StatusOK, bookstores)
}

func (h *abcHandler) Getbook(c echo.Context) error {
	id := c.Param("bookid")
	bookstore := bookstore{}

	if err := h.DB.Find(&bookstore, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, bookstore)
}

func (h *abcHandler) Savebook(c echo.Context) error {
	bookstore := bookstore{}

	if err := c.Bind(&bookstore); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&bookstore).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, bookstore)
}

func (h *abcHandler) Updatebook(c echo.Context) error {
	id := c.Param("bookid")
	bookstore := bookstore{}

	if err := h.DB.Find(&bookstore, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&bookstore); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&bookstore).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, bookstore)
}

func (h *abcHandler) Deletebook(c echo.Context) error {
	id := c.Param("bookid")
	bookstore := bookstore{}

	if err := h.DB.Find(&bookstore, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Delete(&bookstore).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

//================================================================================
func (h *abcHandler) GetAllOrder(c echo.Context) error {
	Orders := []Order{}

	h.DB.Find(&Orders)

	return c.JSON(http.StatusOK, Orders)
}

func (h *abcHandler) GetOrder(c echo.Context) error {
	id := c.Param("Orderid")
	Order := Order{}

	if err := h.DB.Find(&Order, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, Order)
}

func (h *abcHandler) SaveOrder(c echo.Context) error {
	Order := Order{}

	if err := c.Bind(&Order); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&Order).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, Order)
}

func (h *abcHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("Orderid")
	Order := Order{}

	if err := h.DB.Find(&Order, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := c.Bind(&Order); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := h.DB.Save(&Order).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, Order)
}

func (h *abcHandler) DeleteOrder(c echo.Context) error {
	id := c.Param("Orderid")
	Order := Order{}

	if err := h.DB.Find(&Order, id).Error; err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	if err := h.DB.Delete(&Order).Error; err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
