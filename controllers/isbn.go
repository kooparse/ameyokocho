package controllers

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/kooparse/ameyokocho/models"
	"github.com/kooparse/ameyokocho/utils"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"time"
)

var client = http.Client{
	Timeout: time.Duration(5 * time.Second),
}

func CreateBookFromIsbn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isbn := params.ByName("isbn")

	uri := "https://www.googleapis.com/books/v1/volumes?q=isbn:" + isbn

	res, err := client.Get(uri)
	defer res.Body.Close()

	utils.ErrorCheck(err)

	body, err := ioutil.ReadAll(res.Body)
	utils.ErrorCheck(err)

	content := string(body)
	accessor := "items.0.volumeInfo."

	title := gjson.Get(content, accessor+"title")
	description := gjson.Get(content, accessor+"description")
	isbn_10 := gjson.Get(content, accessor+"industryIdentifiers.#[type==ISBN_10].identifier")
	isbn_13 := gjson.Get(content, accessor+"industryIdentifiers.#[type==ISBN_13].identifier")
	language := gjson.Get(content, accessor+"language")
	publishedDate := gjson.Get(content, accessor+"publishedDate")

	var authors []string
	for _, v := range gjson.Get(content, accessor+"authors").Array() {
		authors = append(authors, v.Str)
	}

	if isbn_10.Exists() || isbn_13.Exists() {
		book := models.Book{
			Title:         title.Str,
			Description:   description.Str,
			ISBN_10:       isbn_10.Str,
			ISBN_13:       isbn_13.Str,
			Language:      language.Str,
			PublishedDate: publishedDate.Str,
			Authors:       authors,
		}

		models.AddBook(book)

		// Send the formatted response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
