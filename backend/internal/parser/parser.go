package parser

import (
	"Delivery/backend/internal/parser/worker_pool"
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/models"
	"Delivery/backend/internal/repositories/requests"
	services2 "Delivery/backend/internal/services"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Parser struct {
}

func (p Parser) ParseHTTP(SuppliersLink string) (supplier []requests.SupplierRequest) {
	client := http.DefaultClient

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, SuppliersLink, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	suppliersMap := make(map[string][]requests.SupplierRequest, 0)

	err = json.NewDecoder(res.Body).Decode(&suppliersMap)
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()

	suppliers := make([]requests.SupplierRequest, 0)

	suppliers = suppliersMap["suppliers"]

	for i := range suppliers {
		ctx, cancel = context.WithTimeout(context.Background(), time.Second)
		req, err = http.NewRequestWithContext(ctx, http.MethodGet,
			"http://foodapi.true-tech.php.nixdev.co/suppliers/"+strconv.Itoa(int(suppliers[i].Id))+"/menu",
			nil)
		if err != nil {
			log.Fatal(err)
		}
		res, err = client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		menuMap := make(map[string][]models.Product, 0)
		json.NewDecoder(res.Body).Decode(&menuMap)
		suppliers[i].Menu = menuMap["menu"]
		res.Body.Close()
		cancel()
	}
	return suppliers
}

func (p Parser) UpdateHTTP(suppliers []requests.SupplierRequest, conn *sql.DB) {
	client := http.DefaultClient
	store := repositories.NewStore(conn)
	for {
		time.Sleep(time.Minute)
		for i, sup := range suppliers {
			for j, prod := range sup.Menu {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				req, err := http.NewRequestWithContext(ctx, http.MethodGet,
					"http://foodapi.true-tech.php.nixdev.co/suppliers/"+
						strconv.Itoa(int(sup.Id))+"/menu/"+strconv.Itoa(int(prod.Id)),
					nil)
				res, err := client.Do(req)
				if err != nil {
					log.Println(err)
				}
				var p models.Product
				err = json.NewDecoder(res.Body).Decode(&p)
				if err != nil {
					log.Println(err)
				}
				_, err = store.ProductRepo.UpdatePrice(&p)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(p.Name, prod.Price, "-", p.Price)
				suppliers[i].Menu[j].Price = p.Price
				res.Body.Close()
				cancel()
			}
		}
	}
}

func (p Parser) ParseJson(jsonFile string) (supplier requests.SupplierRequest) {
	supplierData := []byte{}

	file, err := os.Open(jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		chunk := make([]byte, 64)
		n, err := file.Read(chunk)
		if err == io.EOF { //
			break //
		}
		if err != nil {
			panic(err)
		}

		supplierData = append(supplierData, chunk[:n]...)
	}
	err = json.Unmarshal(supplierData, &supplier)
	if err != nil {
		panic(err)
	}

	//PrintSupplier(supplier)
	return supplier
}

func (p Parser) Parse(conn *sql.DB) {
	poolConst := func() worker_pool.Worker {
		return worker{
			DB: conn,
		}
	}

	pool := worker_pool.NewPool(4, poolConst)
	go pool.Run()

	supplier := p.ParseHTTP("http://foodapi.true-tech.php.nixdev.co/suppliers")

	for i := range supplier {
		pool.DataSource <- supplier[i]
	}

	p.UpdateHTTP(supplier, conn)
	pool.Stop()
}

type worker struct {
	DB *sql.DB
}

func (c worker) Do(data interface{}, _ int) {
	sup, ok := data.(requests.SupplierRequest)
	if !ok {
		return
	}
	store := repositories.NewStore(c.DB)
	s := services2.NewSupplierService(store)
	_, err := s.Insert(&sup)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c worker) Stop() {
	return
}
