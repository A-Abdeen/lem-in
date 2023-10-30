#!/bin/bash

cd main
echo "TIME TEST START"
echo ""
echo "========== example06.txt -> Maximum time: 1.5 minutes"
go run main.go example06.txt
echo ""
echo "========== example07.txt -> Maximum time: 2.5 minutes"
go run main.go example07.txt
echo ""
echo "TIME TEST END"
echo ""