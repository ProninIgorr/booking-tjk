#!/bin/bash

go build -o booking-tjk cmd/web/*.go 
./booking-tjk -dbname=booking -dbuser=postgres -cache=false -production=false -dbpass=Pizdechuynya99