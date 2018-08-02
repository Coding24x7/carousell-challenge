#!/bin/bash

set +x
set +e

gitpath=github.com/Coding24x7/carousell-challenge
dirpath=${GOPATH}/src/${gitpath}

# work from webapp
cd ${dirpath}

# goa generate app
# goagen does not work with vendor, see https://github.com/goadesign/goa/issues/923

if [ -d ${dirpath}/vendor ]; then
	mv ${dirpath}/vendor ${dirpath}/.vendor
fi

design=${gitpath}/design
if [ "$#" -ne 1 ]; then
	if [ -d ${dirpath}/goa_temp ]; then
		rm -rf ${dirpath}/goa_temp
	fi
	mkdir ${dirpath}/goa_temp

	goagen bootstrap -o ${dirpath}/goa_temp -d ${design}
    goagen app -d ${design}
    goagen swagger -d ${design}
else
    goagen $1 -d ${design}
fi

if [ -d ${dirpath}/.vendor ]; then
	mv ${dirpath}/.vendor ${dirpath}/vendor
fi
