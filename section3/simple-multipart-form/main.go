package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	w.WriteField("name", "Micheal Jackson")

	p := make(textproto.MIMEHeader)
	p.Set("Content-Type", "image/png")
	p.Set("Content-Disposition", `form-data; name="thumbnail"; filename="photo.png"`)

	fw, err := w.CreatePart(p)
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
