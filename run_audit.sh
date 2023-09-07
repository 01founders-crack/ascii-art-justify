#!/bin/bash

# Test cases
echo "Running: go run . --align=right left standard"
go run . --align=right left standard
echo

echo "Running: go run . --align=left right standard"
go run . --align=left right standard
echo

echo "Running: go run . --align=center hello shadow"
go run . --align=center hello shadow
echo

echo "Running: go run . --align=justify \"1 Two 4\" shadow"
go run . --align=justify "1 Two 4" shadow
echo

echo "Running: go run . --align=right 23/32 standard"
go run . --align=right 23/32 standard
echo

echo "Running: go run . --align=right ABCabc123 thinkertoy"
go run . --align=right ABCabc123 thinkertoy
echo

echo "Running: go run . --align=center \"#$%&\\\"\" thinkertoy"
go run . --align=center "#$%&\"" thinkertoy
echo

echo "Running: go run . --align=left '23Hello World!' standard"
go run . --align=left '23Hello World!' standard
echo

echo "Running: go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy"
go run . --align=justify 'HELLO there HOW are YOU?!' thinkertoy
echo

echo "Running: go run . --align=right \"a -> A b -> B c -> C\" shadow"
go run . --align=right "a -> A b -> B c -> C" shadow
echo

echo "Running: go run . --align=right abcd shadow"
go run . --align=right abcd shadow
echo

echo "Running: go run . --align=center ola standard"
go run . --align=center ola standard
echo
