#!/bin/bash

# Loop through characters a to z
for char in {a..z}; do
    # Echo which character we're testing
    echo "Testing character: $char"
    # Echo the character to the program
    echo "$char" | go run exam2-p14-5.go
    # Add a small delay to avoid overwhelming the server
    sleep 0.5
done 