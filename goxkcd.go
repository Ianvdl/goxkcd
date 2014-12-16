package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func get(url string) string {
	response, errget := http.Get(url)
	if errget != nil {
		return "Could not retrieve resource: " + errget.Error()
	}

	result, errread := ioutil.ReadAll(response.Body)
	if errread != nil {
		return "Could not parse content: " + errread.Error()
	}

	body := string(result)
	return body
}

func parse_page(body string) (image, title string) {
	//Get image line
	body = body[strings.Index(body, "<img id=\"comic\""):]
	body = body[:strings.Index(body, "/>")+2]
	//Get image URL
	image = body[strings.Index(body, "src=\"")+5:]
	image = image[:strings.Index(image, "\"")]
	//Get title
	title = body[strings.Index(body, "alt=\"")+5:]
	title = title[:strings.Index(title, "\"")]

	return image, title
}

func main() {
	body := get("http://m.xkcd.com")
	image, title := parse_page(body)

	fmt.Println("Image URL: ", image)
	fmt.Println("Comic Title: ", title)

	imagefile, _ := http.Get(image)
	imagebytes, _ := ioutil.ReadAll(imagefile.Body)

	ioutil.WriteFile(title+".png", imagebytes, 0777)
}
