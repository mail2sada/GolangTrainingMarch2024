package main

import (
	"fmt"
	"time"
)

type Product struct {
	Pid      int       `json:"pid"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Price    float64   `json:"price"`
	Expiry   time.Time `json:"expiry"`
	Desc     string    `json:"desc"`
}

var ProductDetails []Product = []Product{Product{Pid: 1, Name: "Laptop", Category: "Electronics", Price: 145000.00, Expiry: time.Now().Add(365 * time.Hour * 24)},
	Product{Pid: 2, Name: "PlayStation", Category: "Electronics", Price: 145000.00, Expiry: time.Now().Add(365 * time.Hour * 24)},
}

func GetAllProduct() []Product {
	return ProductDetails
}

func GetProduct(pid int) (Product, error) {
	for _, v := range ProductDetails {
		if v.Pid == pid {
			return v, nil
		}
	}
	p := Product{}
	return p, fmt.Errorf("Not found")
}

func InserProduct(p Product) error {
	p, e := GetProduct(p.Pid)
	if e == nil {
		return fmt.Errorf("Product already exists..")
	}
	ProductDetails = append(ProductDetails, p)
	return nil
}

func DeleteProduct(pID int) error {
	_, e := GetProduct(pID)
	if e != nil {
		return fmt.Errorf("Product does not exist..")
	}
	idx := -1
	for i, v := range ProductDetails {
		if v.Pid == pID {
			idx = i
			break
		}
	}
	if idx == -1 {
		return fmt.Errorf("internal error")
	}
	ProductDetails = append(ProductDetails[:idx-1], ProductDetails[idx+1:]...)
	return nil
}
