# Web-App-Pen-Testing
These are applications useful during a web application penetration test.<br>
<br>
<h2>comments.go</h2>
Compile using Go: go build comments.go. Use burp to find and export comments (comments.txt) to a directory containing comments.exe. Run "comments' from cmd. It automatically looks for "comments.txt" and dsiplays the output to the terminal.
<br>
<h2>dirspray.go</h2>
Compile using: go build dirspray.go.
<br>
This application takes two arguments: The URL list to attack and the discovery list. The application then enumerates the list of targets with a single discovery list. This is good for Cloud Penetration testing when you have an idea of what you are looking for at scale across multiple targets.
  <br>
  <h2>proxspray.go</h2>
Compile using: go build proxspray.go
This application takes one argument - the list of domains to attack. The application then tries to proxy traffic through those sites. This came about because I noticed certain installations of nginx and Apache allowed proxying through the server due to a misconfiguration.
<br>
<h2>CoWitness</h2>

CoWitness is a powerful web application testing tool that enhances the accuracy and efficiency of your testing efforts. It allows you to mimic an HTTP server and a DNS server, providing complete responses and valuable insights during your testing process.

### Features

- Simulate an HTTP server and a DNS server for comprehensive web application testing
- Capture and log all incoming requests to analyze the requested resource and user agent information
- Identify false positives and distinguish genuine vulnerabilities
- Discover hidden vulnerabilities by examining requested resource paths and file names
- Monitor and analyze HTTP and DNS logs simultaneously with the multitail utility
- Improve the overall effectiveness of your web application testing

### Usage

1. Choose a domain name for your testing environment.
2. Set up a remote server and obtain a public IP address for it.
3. Register your name servers to point to the public IP address.
4. Create glue records to associate the IP address with your remote server.
5. Ensure that ports 80 and 53 are available on the remote server.
6. Compile and run CoWitness on the remote server.
7. Monitor the HTTP log and DNS log simultaneously using multitail in your terminal.

### Community and Contributions

We welcome contributions and feedback from the community. If you encounter any issues or have suggestions for improvements, please open an issue or submit a pull request on our GitHub repository.

### License

CoWitness is released under the [MIT License](LICENSE).

