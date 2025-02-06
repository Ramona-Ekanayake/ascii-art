#!/bin/bash

# Path to your ASCII art program
ASCII_PROGRAM="./ascii-art"  # Change this to the actual executable

# Define test cases
test_cases=(
    "hello"
    "HELLO"
    "HeLlo HuMaN"
    "1Hello 2There"
    $'Hello\nThere'
    $'Hello\n\nThere'
    "{Hello & There #}"
    "hello There 1 to 2!"
    "MaD3IrA&LiSboN"
    "1a\"#FdwHywR&/()="
    "{|}~"
    "[\\]^_ 'a"
    "RGB"
    ":;<=>?@"
    "\!\" #$%&'()*+,-./"
    "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    "abcdefghijklmnopqrstuvwxyz"
    "aBcDeFgHi"
    "test123 45"
    "XyZ@!#"
    "ab 12 ^% ABC"
)

# Run tests
for test in "${test_cases[@]}"; do
    echo "Running test: '$test'"
    output=$($ASCII_PROGRAM "$test")
    echo -e "Output:\n$output"
    echo "-------------------------------------"
done
