package data

import (
	"book-microservice/src/model"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

var Books = []model.Book{
	{
		Name:        "Building Microservices: Designing Fine-Grained Systems 2nd Edition",
		Author:      "Sam Newman",
		Description: "As organizations shift from monolithic applications to smaller, self-contained microservices, distributed systems have become more fine-grained. But developing these new systems brings its own host of problems. This expanded second edition takes a holistic view of topics that you need to consider when building, managing, and scaling microservices architectures.",
		Image:       toBase64("src/data/images/1.jpg"),
	},
	{
		Name:        "Building Microservices with Go: Develop seamless, efficient, and robust microservices with Go",
		Author:      "Nic Jackson",
		Description: "Your one-stop guide to the common patterns and practices, showing you how to apply these using the Go programming language",
		Image:       toBase64("src/data/images/2.jpg"),
	},
}

// only jpg format
func toBase64(imageRelativePath string) string {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(pwd + "/" + imageRelativePath)

	bytes, err := ioutil.ReadFile(pwd + "/" + imageRelativePath)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	base64Encoding := "data:image/jpeg;base64,"
	base64Encoding += base64.StdEncoding.EncodeToString(bytes)

	return base64Encoding
}
