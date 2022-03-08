# Goland For Begginers
#### Created on 2007, Open sourced in 2009

### Usage
* Microservices
* Web Applications
* Database Services

### Advantages
* Simple Syntax
* Fast build time, strat up and run
* Requires fewer resources
* Consistent across different OS

### Data Types
1. Strings
2. Integers
3. Booleans
4. Arrays & Slices
5. Maps
6. Structs

### Variables and Constants
* Formatted Output
* User Input
* Pointers
* Scope Rules
* Loops
* If-else & Switch
* Functions
* Packages
* Goroutines

## Install Go 
1. By link `https://go.dev/doc/install` or manually
2. Manually installation
1.1 Downloading the Go tarball (08.02.2022)
```bash
  wget -c https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
```
1.2 Adjusting the Path Variable
```bash
  export PATH=$PATH:/usr/local/go/bin
```
1.3 Verifying the Go Installation
```bash
  go version
```
3. From console
```bash
  sudo snap install go --classic
```
## Create a new module
```bash
  go mod init <module path>
```
```bash
  go mod init booking-app
```
## In order to run application 
```bash
  go run main.go
```
