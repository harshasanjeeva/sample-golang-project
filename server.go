package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct{
	Id 		string  `json:"Id"`
	Title   string 	`json:"Title"`
	Desc    string 	`json:"desc"`
	Content string	`json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Endpoint hit: homepage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key  := vars["id"]

	for _, article := range Articles{
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request){
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article

	err := json.Unmarshal(reqBody, &article)
	if err !=  nil {
		fmt.Printf("There was an error decoding the json. err = %s", err)
        return
	}
	fmt.Println(string(reqBody))
	fmt.Println(article.Title)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars  :=  mux.Vars(r)
	id := vars["id"]

	for index, article := range  Articles {
		if article.Id  ==  id {
			Articles = append(Articles[:index],Articles[index+1:]...)
		}
	}

}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	err := http.ListenAndServe(":8080",myRouter)
	if err != nil {
		panic(err)
	}
}


func main(){
	Articles =  []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description2", Content: "Article Content 2"},
	}
	handleRequests()

}