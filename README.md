# Property Slot Assignment

## Introduction

This project contains a simple function that calculates the number of item slots in a webpage that will be assigned to a client.

## Setup and Usage

### Clone the repository
```bash
git clone https://github.com/srsaurav0/PropertySlotAssignment
cd PropertySlotAssignment
```

### Install Dependencies
```bash
go mod init math-test
go mod tidy
```

### Execute the Function
```bash
go run math.go
```

### Test Function
Change the values of this section (***line 10***) to test for different values:
```bash
input := map[string]float64{
	"11": X,
	"12": X,
	"24": X,
}
```
Replace X with a float value.

## Testing

For testing the `math.go` file, run this command:
```bash
go test .
```
Test with coverage
```bash
go test . -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
open coverage.html
```
View total coverage in terminal:
```bash
go tool cover -func=coverage.out | grep total: | awk '{print $3}'
```
Open `coverage.html` in a browser to view detailed coverage.