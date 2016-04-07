#calc-server
RPC calculator server  
Calculates simple arithmetic tasks. Supports addition, subtraction, multiplication and division.

##Requirements
You need Golang 1.5 or newer [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/code.html)  
Then install Godep:  
`go get github.com/tools/godep`

##Installation
Download calc-server:  
`go get github.com/alyoshka/calc-server`  
`cd $GOPATH/src/github.com/alyoshka/calc-server`  
Install  
`godep go install`  

##Usage
`calc-server`  
You might specify host and port to listen  
`calc-server -host localhost:8081`
