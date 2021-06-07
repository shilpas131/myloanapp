# myloanapp
A golang service for loan management

Build the code: 
   go build ./...

Run the tests:
   go test ./...

Run the app: 
   go run main.go

There 3 APIs for this service, Following are the details to work with them :
  1) Initiate Loan :
    
    url: localhost:4200/initiateLoan 
    
    json request body:{
      "amount" : "300000",
      "start-date" : "2020-06-01T08:28:06.801064-04:00", //Date needs to be in this format
      "interest-rate" : "7"
    }
    
    response (if success): {"ID":82,"amount":"500000","start-date":"2020-06-01T08:28:06.801064-04:00","interest-rate":"7","PaymentTracker":null}
    
  2) Add Payment :
    
    url: localhost:4200/addPayment
    
    json request body: {
       "amount" : "80000",
       "date" : "2020-06-05T08:28:06.801064-04:00" //Date needs to be in this format
    }
   
    response (if success): [{"amount":"80000","date":"2020-06-05T08:28:06.801064-04:00"}]
    
  3) Get Balance : 
      
    url: localhost:4200/getBalance?balanceDate=2020-06-11 (date needs to be in this format)
      
    response (if success): The balance is: 235835.48
