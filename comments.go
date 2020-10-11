package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("comments.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	search := 0
	for _, eachline := range txtlines {
		SearchWords := []string{"username", "password", "adm", "root", "pass", "administrator", "passwd", "pass", "creds", "credentials", "usr", "admin", "user", "bucket_name", "aws_access_key", "aws_secret_key", "S3_BUCKET", "S3_ACCESS_KEY_ID", "S3_SECRET_ACCESS_KEY", "S3_ENDPOINT", "AWS_ACCESS_KEY_ID", "list_aws_accounts"}
		for i := 0; i < len(SearchWords); i++ {
			lookFor := SearchWords[i]
			contain := strings.Contains(eachline, lookFor)
			if contain == true {
				search++
				fmt.Printf("The following line \"%s\" contains\"%s\"\n", eachline, lookFor)
				fmt.Println("                                                ")

			}
		}
	}
	if search == 0 {
		fmt.Println("Search complete. No results were found.")
	}
}
