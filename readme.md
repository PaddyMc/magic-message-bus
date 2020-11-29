# magicmessagebus

sample application to use a message bus for module communication in the cosmos-sdk

code in this POC removes the need to pass keepers into other keeper's constructors

# how to test

- startport serve

in another terminal

- cd /cmd/magic-message-bus
- go run main.go tx magicmessagebus create-poll bus asd --from user1 --keyring-backend test --chain-id magicmessagebus
- go run main.go query bank balances cosmos1xm82mkw2jkwkdgq3r0cu8f92r9t2emm8m8xpuw
