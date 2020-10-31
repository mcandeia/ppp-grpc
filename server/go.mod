module mcandeia/ppp/impl

go 1.15

replace mcandeia/grpc/ppp => ../ppp

require (
	google.golang.org/grpc v1.33.1
	mcandeia/grpc/ppp v0.0.0-00010101000000-000000000000
)
