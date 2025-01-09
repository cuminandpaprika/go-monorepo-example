# go-monorepo-example

## Curl
```
curl -X POST http://localhost:8000/order.v1alpha1.OrderService/CreateOrder \
     -H "Content-Type: application/json" \
     -d '{"order": {
           "id": "order123",
           "customer": {
             "name": "John Doe",
             "phone": "1234567890"
           },
           "items": [
             {
               "name": "Pizza",
               "quantity": 1,
               "price": 1000
             }
           ],
           "total_price": 1000,
           "status": "pending"
         }}'
```