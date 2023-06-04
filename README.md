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

