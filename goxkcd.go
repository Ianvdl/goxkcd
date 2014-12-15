package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	source, _ := http.Get("http://m.xkcd.com")
	response, _ := ioutil.ReadAll(source.Body)
	body := string(response)
	//Get image line
	body = body[strings.Index(body, "<img id=\"comic\""):]
	body = body[:strings.Index(body, "/>")+2]
	//Get image URL
	image := body[strings.Index(body, "src=\"")+5:]
	image = image[:strings.Index(image, "\"")]
	//Get title
	title := body[strings.Index(body, "alt=\"")+5:]
	title = title[:strings.Index(title, "\"")]

	fmt.Println("HTML Source Line: ", body)
	fmt.Println("Image URL: ", image)
	fmt.Println("Comic Title: ", title)

	imagefile, _ := http.Get(image)
	imagebytes, _ := ioutil.ReadAll(imagefile.Body)

	ioutil.WriteFile(title+".png", imagebytes, 0777)
}
