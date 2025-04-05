module kyc

go 1.23.4

replace db => ../db

replace middlewares => ../middlewares

require (
	db v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	middlewares v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
)
