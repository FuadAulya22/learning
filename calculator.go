package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func BilanganHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	log.Println("test")
	bilsatu := r.FormValue("BilanganSatu")
	operasi := r.FormValue("Kabataku")
	bildua := r.FormValue("BilanganDua")
	log.Println("test", bilsatu, bildua, operasi)
	if len(bilsatu) > 0 && len(operasi) > 0 && len(bildua) > 0 {
		log.Println("test2")
		i1, err := strconv.ParseFloat(bilsatu, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bilangan Satu yang Anda input salah!\r\n"))
			return
		}

		i2, err := strconv.ParseFloat(bildua, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bilangan Dua yang Anda input salah!\r\n"))
			return
		}

		if operasi == "tambah" {
			hasil := i1 + i2
			hasilStr := fmt.Sprintf("Total: %9.2f\r\n", hasil)
			w.Write([]byte(hasilStr))
		} else if operasi == "kurang" {
			hasil := i1 - i2
			hasilStr := fmt.Sprintf("Total: %9.2f\r\n", hasil)
			w.Write([]byte(hasilStr))
		} else if operasi == "bagi" {
			hasil := i1 / i2
			hasilStr := fmt.Sprintf("Total: %9.2f\r\n", hasil)
			w.Write([]byte(hasilStr))
		} else if operasi == "kali" {
			hasil := i1 * i2
			hasilStr := fmt.Sprintf("Total: %9.2f\r\n", hasil)
			w.Write([]byte(hasilStr))
		}
	} else {
		w.Write([]byte("Paramater yang Anda masukkan salah"))
	}
}

func main() {
	router := httprouter.New()

	router.GET("/calculator", BilanganHandler)

	http.ListenAndServe(":9001", router)
}
