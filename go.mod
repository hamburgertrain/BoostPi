module github.com/hamburgertrain/boostpi

go 1.23.2

require (
	github.com/d2r2/go-i2c v0.0.0-20191123181816-73a8a799d6bc
	github.com/hamburgertrain/elmobd v0.0.0
)

require (
	github.com/d2r2/go-logger v0.0.0-20210606094344-60e9d1233e22 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/tarm/serial v0.0.0-20180830185346-98f6abe2eb07 // indirect
	golang.org/x/sys v0.16.0 // indirect
)

replace github.com/hamburgertrain/elmobd v0.0.0 => ./internal/elmobd
