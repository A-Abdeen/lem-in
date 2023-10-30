#!/bin/bash

cd main
echo "TURN TEST START"
echo ""
echo "========== example00.txt -> Maximum turns: 6"
go run main.go example00.txt
echo ""
echo "========== example01.txt -> Maximum turns: 8"
go run main.go example01.txt
echo ""
echo "========== example02.txt -> Maximum turns: 11"
go run main.go example02.txt
echo ""
echo "========== example03.txt -> Maximum turns: 6"
go run main.go example03.txt
echo ""
echo "========== example04.txt -> Maximum turns: 6"
go run main.go example04.txt
echo ""
echo "========== example05.txt -> Maximum turns: 8"
go run main.go example05.txt
echo ""
echo "========== badexample00.txt -> Return: ERROR message"
go run main.go badexample00.txt
echo ""
echo "========== badexample01.txt -> Return: ERROR message"
go run main.go badexample01.txt
echo ""
echo "TURN TEST END"
echo ""