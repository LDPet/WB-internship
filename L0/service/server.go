package main

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"service/cache"
	"service/model"
	"strings"
	"sync"
)

const (
	protocolPrefix = "http://"
	host           = "localhost:8099"
	indexURL       = "/orders"
	showURL        = "/orders/"
)

type OrderLink struct {
	Uid  string
	Link template.URL
}

type IndexOrdersView struct {
	SearchUrl string
	Orders    []OrderLink
}

func runServer(ctx context.Context, cacher cache.Cacher[string, []byte]) {
	indexTmpl := template.Must(template.ParseFiles("templates/index.html")) //todo handle error
	showTmpl := template.Must(template.ParseFiles("templates/show.html"))   //todo handle error

	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		okIndex, _ := regexp.Match("\\/orders\\/$", []byte(p))
		okShow, _ := regexp.Match("\\/orders\\/[a-zA-z0-9]+", []byte(p))

		if okIndex && r.Method == http.MethodGet {
			code, err := index(w, cacher, indexTmpl)
			if err != nil {
				w.WriteHeader(code)
				w.Write([]byte(err.Error()))
				log.Println(err)
			}
			return
		} else if okShow && r.Method == http.MethodGet {
			code, err := show(w, r, cacher, showTmpl)
			if err != nil {
				w.WriteHeader(code)
				w.Write([]byte(err.Error()))
				log.Println(err)
			}
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	})

	httpServer := http.Server{
		Addr: host,
	}
	wg := ctx.Value("wg").(*sync.WaitGroup)
	quit := ctx.Value("chan").(chan os.Signal)

	go func() {
		sig := <-quit
		quit <- sig

		wg.Done()
		err := httpServer.Shutdown(ctx)
		if err != nil {
			log.Fatal("Shutdown server error: ", err)
		}
	}()

	wg.Add(1)
	err := httpServer.ListenAndServe()
	if err != nil {
		wg.Done()
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server shutdown")
		} else {
			log.Fatal("Server run error: ", err)
		}
	}
}

func index(w http.ResponseWriter, cacher cache.Cacher[string, []byte], tmpl *template.Template) (int, error) {
	keys := cacher.Keys()

	orders := make([]OrderLink, 0)
	for _, key := range keys {
		orders = append(orders, OrderLink{
			Uid:  key,
			Link: template.URL(protocolPrefix + host + showURL + key),
		})
	}

	view := IndexOrdersView{
		SearchUrl: host + showURL,
		Orders:    orders,
	}

	err := tmpl.Execute(w, view)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func show(w http.ResponseWriter, r *http.Request, cacher cache.Cacher[string, []byte], tmpl *template.Template) (int, error) {
	p := r.URL.Path
	uid := strings.TrimPrefix(p, "/orders/")

	data, ok := cacher.Get(uid)
	if !ok {
		return http.StatusNotFound, errors.New("not found order")
	}
	order := model.Order{}
	err := json.Unmarshal(data, &order)
	if err != nil {
		return http.StatusInternalServerError, errors.New("bad data in cache")
	}

	//w.Write(order)
	err = tmpl.Execute(w, order)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}
