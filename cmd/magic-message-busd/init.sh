#!/bin/bash

go run main.go init testing
echo "y" | go run main.go keys add validator
go run main.go add-genesis-account validator 1000000000stake
go run main.go gentx validator --chain-id testing
go run main.go collect-gentxs
#./simd start

