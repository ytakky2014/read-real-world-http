package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("name", "Micheal Jackson")

	fw, err := w.CreateFormFile("thumbnail", "photo.png")
	if err != nil {
		panic(err)
	}

	rf, err := os.Open("photo.png")
	if err != nil {
		panic(err)
	}

	defer rf.Close()
	io.Copy(fw, rf)
	w.Close()

	resp, err := http.Post("http://localhost:18888", w.FormDataContentType(), &buf)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)

}
