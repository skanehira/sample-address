package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type server struct {
	db *gorm.DB
	e  *echo.Echo
}

type Address struct {
	ID      int    `json:"id" gorm:"AUTO_INCREMENT"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

type ErrResponse struct {
	Message string `json:"message"`
}

func main() {
	db, err := gorm.Open("sqlite3", "address.db")
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Address{})

	s := server{
		db: db,
		e:  echo.New(),
	}

	s.Start()
}

func (s *server) Start() {
	s.e.Static("/", ".")
	s.e.GET("/address", s.AllAddress)
	s.e.POST("/address", s.RegisterAddress)
	s.e.DELETE("/address/:id", s.DeleteAddress)
	s.e.PUT("/address/:id", s.UpdateAddress)

	log.Println("start server")
	if err := s.e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}

func (s *server) AllAddress(c echo.Context) error {
	address := []Address{}

	if err := s.db.Find(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, address)
}

func (s *server) RegisterAddress(c echo.Context) error {
	var address Address

	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse{Message: err.Error()})
	}

	if err := s.db.Save(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, address)
}

func (s *server) DeleteAddress(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse{Message: err.Error()})
	}
	address := Address{ID: id}

	if err := s.db.Delete(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (s *server) UpdateAddress(c echo.Context) error {
	var address Address
	if err := c.Bind(&address); err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse{Message: err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse{Message: err.Error()})
	}

	address.ID = id

	if err := s.db.Model(&address).Update(&address).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
