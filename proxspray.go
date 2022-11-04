package main

//import packages
import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

//set variable
var (
	err      error
	response *http.Response
	body     []byte
)

func main() {

	//Get user input
	if len(os.Args) <= 1 {
		fmt.Printf("USAGE : %s <URL LIST FILE> \n", os.Args[0])
		os.Exit(0)
	}

	domain := os.Args[1]

	//create the file to write out the report to
	fileHandle, _ := os.Create("report.html")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	//open file with the URLS
	f, err := os.Open(domain)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//write out the top HTML part of the report
	fmt.Fprintln(writer, "<!DOCTYPE html>")
	fmt.Fprintln(writer, "<html>")
	fmt.Fprintln(writer, "<style>")
	fmt.Fprintln(writer, "table, th, td {")
	fmt.Fprintln(writer, "  border:1px solid black;")
	fmt.Fprintln(writer, "}")
	fmt.Fprintln(writer, "</style>")
	fmt.Fprintln(writer, "<body>")
	fmt.Fprintln(writer, "<table style=\"width:100%\">")
	fmt.Fprintln(writer, "<tr>")
	fmt.Fprintln(writer, "<th>Proxy Target</th>")
	fmt.Fprintln(writer, "<th>Response Message</th>")
	fmt.Fprintln(writer, "<th>URL Attempt</th>")
	fmt.Fprintln(writer, "<th>Status</th>")
	fmt.Fprintln(writer, "<th>Length</th>")
	fmt.Fprintln(writer, "</tr>")

	//read the file line by line
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, err := net.DialTimeout("tcp", scanner.Text()+":80", 2*time.Second)
		fmt.Println("Trying:" + scanner.Text() + ":80")
		if err == nil {

			proxyStr := "http://" + scanner.Text() + ":80"
			proxyURL, err := url.Parse(proxyStr)
			if err != nil {
				log.Println(err)
			}

			//adding the proxy settings to the Transport object
			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}

			//adding the Transport object to the http Client
			client := &http.Client{
				Transport: transport,
			}

			var target = "google.com"

			//Loop through the array
			//for _, hackDirsandFiles := range dirsAndFiles {
			URLresult := "https://" + target
			urlStr := URLresult
			url, err := url.Parse(urlStr)
			if err := recover(); err != nil {
				log.Println(err)
			}

			//generating the HTTP GET request
			request, err := http.NewRequest("GET", url.String(), nil)
			if err := recover(); err != nil {
				log.Println(err)
			}

			request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:76.0) Gecko/20100101 Firefox/76.0")

			//Uncomment the below line to add a two second delay.
			//time.Sleep(2 * time.Second)
			//calling the URL
			response, err := client.Do(request)
			if err != nil {
				fmt.Println("Attempting proxy through: " + proxyStr)
				log.Println(err)
				log.Println("*************************************************************")
				log.Println("                                                             ")
				//fmt.Fprintln(writer, "<tr>")
				//fmt.Fprintln(writer, "<td>"+proxyStr+"</td>")
				//fmt.Fprintln(writer, "<td>"+err.Error()+"</td>")
				continue
			}

			//start writing out the report
			fmt.Println("Testing", URLresult)
			fmt.Println("HTTP Response Status:", response.StatusCode, http.StatusText(response.StatusCode))
			if response.StatusCode >= 200 && response.StatusCode <= 299 {
				fmt.Println("HTTP Status is in the 2xx range")
				fmt.Println(response.ContentLength)
				length := strconv.Itoa(int(response.ContentLength))
				fmt.Fprintln(writer, "<tr>")
				fmt.Fprintln(writer, "<td>"+proxyStr+"</td>")
				fmt.Fprintln(writer, "<td>"+err.Error()+"</td>")
				fmt.Fprintln(writer, "<td><b><a href=\""+URLresult+"\"target=\"_blank\">"+URLresult+"</a></b></td>")
				fmt.Println("-------------------------------")

				//Convert the status code integer to string so it can be printed out
				status := strconv.Itoa(response.StatusCode)

				//Write out more of the report.
				fmt.Fprintln(writer, "<td><b>"+status+"</b></td>")
				fmt.Fprintln(writer, "<td><b>"+length+"</b></td>")
				fmt.Println("                               ")
				fmt.Fprintln(writer, "</tr>")

			} else {
				fmt.Fprintln(writer, "<tr>")
				fmt.Fprintln(writer, "<td>"+URLresult+"</td>")
				fmt.Println("-------------------------------")
				length := strconv.Itoa(int(response.ContentLength))
				status := strconv.Itoa(response.StatusCode)
				fmt.Fprintln(writer, "<td>"+status+"</td>")
				fmt.Fprintln(writer, "<td><b>"+length+"</b></td>")
				fmt.Println("                               ")
				fmt.Fprintln(writer, "</tr>")
			}

			//}
			writer.Flush()
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(writer, "</table>")
	fmt.Fprintln(writer, "</body>")
	fmt.Fprintln(writer, "</html>")

} //end main
