package main

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db = config.ConnectDB()

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
}

type Product struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Unit  string `json:"unit"`
}

type Transaction struct {
	ID          string       `json:"id"`
	BillDate    string       `json:"billDate"`
	EntryDate   string       `json:"entryDate"`
	FinishDate  string       `json:"finishDate"`
	EmployeeID  string       `json:"employeeId"`
	CustomerID  string       `json:"customerId"`
	BillDetails []BillDetail `json:"billDetails"`
	TotalBill   int          `json:"totalBill"`
}

type BillDetail struct {
	ID           string  `json:"id"`
	BillID       string  `json:"billId"`
	Product      Product `json:"product"`
	ProductPrice int     `json:"productPrice"`
	Qty          int     `json:"qty"`
}

func main() {

	router := gin.Default()

	router.POST("/customers", CreateCustomer)
	router.GET("/customers/:id", GetCustomer)
	router.PUT("/customers/:id", UpdateCustomer)
	router.DELETE("/customers/:id", DeleteCustomer)

	router.POST("/products", CreateProduct)
	router.GET("/products", ListProducts)
	router.GET("/products/:id", GetProduct)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)

	router.POST("/transactions", CreateTransaction)
	router.GET("/transactions/:id_bill", GetTransaction)
	router.GET("/transactions", ListTransactions)

	router.Run(":8080")
}

func CreateCustomer(c *gin.Context) {
	var newCustomer Customer
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO customer (name, phone_number, address) VALUES ($1, $2, $3)",
		newCustomer.Name, newCustomer.PhoneNumber, newCustomer.Address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCustomer)
}

func GetCustomer(c *gin.Context) {
	customerID := c.Param("id")

	row := db.QueryRow("SELECT * FROM customer WHERE id = $1", customerID)
	var customer Customer
	err := row.Scan(&customer.ID, &customer.Name, &customer.PhoneNumber, &customer.Address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func UpdateCustomer(c *gin.Context) {
	customerID := c.Param("id")

	row := db.QueryRow("SELECT * FROM customer WHERE id = $1", customerID)
	var existingCustomer Customer
	err := row.Scan(&existingCustomer.ID, &existingCustomer.Name, &existingCustomer.PhoneNumber, &existingCustomer.Address)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	var updatedCustomer Customer
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE customer SET name = $1, phone_number = $2, address = $3 WHERE id = $4",
		updatedCustomer.Name, updatedCustomer.PhoneNumber, updatedCustomer.Address, customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCustomer)
}

func DeleteCustomer(c *gin.Context) {
	customerID := c.Param("id")

	_, err := db.Exec("DELETE FROM customer WHERE id = $1", customerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}

func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO product (name, price, unit) VALUES ($1, $2, $3)",
		newProduct.Name, newProduct.Price, newProduct.Unit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func ListProducts(c *gin.Context) {
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Unit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	productID := c.Param("id")

	row := db.QueryRow("SELECT * FROM product WHERE id = $1", productID)
	var product Product
	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Unit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	row := db.QueryRow("SELECT * FROM product WHERE id = $1", productID)
	var existingProduct Product
	err := row.Scan(&existingProduct.ID, &existingProduct.Name, &existingProduct.Price, &existingProduct.Unit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var updatedProduct Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE product SET name = $1, price = $2, unit = $3 WHERE id = $4",
		updatedProduct.Name, updatedProduct.Price, updatedProduct.Unit, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	_, err := db.Exec("DELETE FROM product WHERE id = $1", productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func CreateTransaction(c *gin.Context) {
	var newTransaction Transaction
	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec(`
		INSERT INTO transaction (bill_date, entry_date, finish_date, employee_id, customer_id, total_bill)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		newTransaction.BillDate, newTransaction.EntryDate, newTransaction.FinishDate,
		newTransaction.EmployeeID, newTransaction.CustomerID, newTransaction.TotalBill)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, billDetail := range newTransaction.BillDetails {
		_, err := db.Exec(`
			INSERT INTO bill_detail (bill_id, product_id, product_price, qty)
			VALUES ($1, $2, $3, $4)`,
			transactionID, billDetail.Product.ID, billDetail.ProductPrice, billDetail.Qty)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func GetTransaction(c *gin.Context) {
	transactionID := c.Param("id_bill")

	row := db.QueryRow(`
		SELECT * FROM transaction WHERE id = $1`,
		transactionID)

	var transaction Transaction
	err := row.Scan(&transaction.ID, &transaction.BillDate, &transaction.EntryDate,
		&transaction.FinishDate, &transaction.EmployeeID, &transaction.CustomerID,
		&transaction.TotalBill)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	rows, err := db.Query(`
		SELECT bd.id, bd.bill_id, bd.product_id, bd.product_price, bd.qty,
		p.id, p.name, p.price, p.unit
		FROM bill_detail bd
		JOIN product p ON bd.product_id = p.id
		WHERE bd.bill_id = $1`,
		transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var billDetails []BillDetail
	for rows.Next() {
		var billDetail BillDetail
		err := rows.Scan(&billDetail.ID, &billDetail.BillID, &billDetail.Product.ID,
			&billDetail.ProductPrice, &billDetail.Qty, &billDetail.Product.ID,
			&billDetail.Product.Name, &billDetail.Product.Price, &billDetail.Product.Unit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		billDetails = append(billDetails, billDetail)
	}

	transaction.BillDetails = billDetails

	c.JSON(http.StatusOK, transaction)
}

func ListTransactions(c *gin.Context) {

	rows, err := db.Query(`
		SELECT id, bill_date, entry_date, finish_date, employee_id, customer_id, total_bill
		FROM transaction`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(&transaction.ID, &transaction.BillDate, &transaction.EntryDate,
			&transaction.FinishDate, &transaction.EmployeeID, &transaction.CustomerID,
			&transaction.TotalBill)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		transactions = append(transactions, transaction)
	}

	c.JSON(http.StatusOK, transactions)
}
