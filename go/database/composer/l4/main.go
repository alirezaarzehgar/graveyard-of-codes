package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
	Rate  uint   `json:"rate"`
}

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]any{
		"message": "Fuck off",
	})
}

func NewProduct(c echo.Context) error {
	var p Product
	err := json.NewDecoder(c.Request().Body).Decode(&p)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = db.Query("INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func GetProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "invalid product id",
		})
	}

	var p Product
	err = db.QueryRow("SELECT id, name, price, rate FROM products WHERE id = ? LIMIT 1", id).Scan(&p.ID, &p.Name, &p.Price, &p.Rate)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": "message not found",
		})
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func EditProduct(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	p := Product{ID: uint(id)}
	err = json.NewDecoder(c.Request().Body).Decode(&p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	stmt, _ := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	r, err := stmt.Exec(p.Name, p.Price, id)
	stmt.Close()
	rn, _ := r.RowsAffected()
	if rn == 0 {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, p)
}

func DeleteProduct(c echo.Context) error {
	stmt, _ := db.Prepare("DELETE FROM products WHERE id = ?")
	r, err := stmt.Exec(c.Param("id"))
	stmt.Close()
	rn, _ := r.RowsAffected()
	if rn == 0 {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]any{
		"message": "deleted successfully",
	})
}

func SearchProducts(c echo.Context) error {
	limit := c.QueryParam("limit")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	page := c.QueryParam("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	rows, err := db.Query("SELECT id, name, price, rate FROM products LIMIT ? OFFSET ?", limitInt, (pageInt-1)*limitInt)
	if err == sql.ErrNoRows {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}
	if err != nil {
		return err
	}

	var products []Product
	for rows.Next() {
		var p Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Rate)

		products = append(products, p)
	}

	return c.JSON(http.StatusOK, map[string]any{
		"page":     pageInt,
		"limit":    limitInt,
		"products": products,
	})
}

func main() {
	var err error
	db, err = sql.Open("mysql", "user:pass@tcp(db:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	r := echo.New()
	r.GET("/test", Test)
	r.POST("/product", NewProduct)
	r.GET("/product/:id", GetProduct)
	r.PUT("/product/:id", EditProduct)
	r.DELETE("/product/:id", DeleteProduct)
	r.GET("/products", SearchProducts)

	http.ListenAndServe(":8000", r)
}
