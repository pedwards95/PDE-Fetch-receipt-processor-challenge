# PDE-Fetch-receipt-processor-challenge
Public API application for Fetch Software Engineer Position

To run locally (No Docker):
1. Navigate to root directory (same directory where this readme is)
2. Input into terminal:
```
go build -o ./bin/PDE-Fetch-receipt-processor-challenge ./cmd/PDE-Fetch-receipt-processor-challenge
```
3. Input into terminal:
```
bin/PDE-Fetch-receipt-processor-challenge
```
4. Send requests to localhost:8080


To run locally (With Docker):
1. Navigate to root directory (same directory where this readme is)
2. Input into terminal:
```
docker build -t pde-app .
```
3. Input into terminal:
```
docker run -d --name PDE-App -p 8080:8080 pde-app
```
4. Send requests to localhost:8080