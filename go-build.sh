#!/bin/bash

filesFolder=files
filesToReadFolder=filesToRead
logsFolder=logs
param=$1

create_folders(){
    #create files folder
    if [ -d $filesFolder ]
    then
        echo “La capeta ya existe.”
    else
        mkdir $filesFolder
    if [ $? -eq 0 ]
    then
        echo “se ha creado con éxito $filesFolder”
    else
        echo “Ups! Algo ha fallado al crear ”
    fi
    fi

    #create filesToRead folder
    if [ -d $filesToReadFolder ]
    then
        echo “La capeta ya existe.”
    else
        mkdir $filesToReadFolder
    if [ $? -eq 0 ]
    then
        echo “se ha creado con éxito $filesToReadFolder”
    else
        echo “Ups! Algo ha fallado al crear ”
    fi
    fi

        #create filesToRead folder
    if [ -d $logsFolder ]
    then
        echo “La capeta ya existe.”
    else
        mkdir $logsFolder
    if [ $? -eq 0 ]
    then
        echo “se ha creado con éxito $logsFolder”
    else
        echo “Ups! Algo ha fallado al crear ”
    fi
    fi
}

create_files(){
    touch ./filesToRead/dnis.txt
    touch ./filesToRead/names.txt
    touch ./filesToRead/phoneNumbers.txt
    touch ./filesToRead/shops.txt
}

build_project(){
    cd ./methods/converts
    go build
    cd ../csv
    go build
    cd ../dni
    go build
    cd ../file
    go build
    cd ../json
    go build
    cd ../log
    go build
    cd ../sql
    go build
    cd ..
    go build
    cd ../cmd/main
    go build
    rm main
    cd ../create-shops
    go build
    rm create-shops
    cd ../..

}

run_project(){
    go run ./cmd/main/main.go
}

run_insert_query(){
    go run ./cmd/create-shops/main.go
}

clear_project(){
    rm ./files/*
    rm ./logs/*
}

clear_all_project(){
    rm ./files/*
    rm ./logs/*
    rm ./filesToRead/*

    rmdir ./files
    rmdir ./filesToRead
    rmdir ./logs
}

run_test() {
    go test -timeout 30s -run ^TestComprobeDniAndNie$ drivers-create/methods/dni
    go test -timeout 30s -run ^TestGenerateJson$ drivers-create/methods/json
    go test -timeout 30s -run ^TestSql$ drivers-create/methods/sql
    go test -timeout 30s -run ^TestUsersToPasswords$ drivers-create/methods/converts
    go test -timeout 30s -run ^TestConvertAllDnisToUsers$ drivers-create/methods/converts

    rmdir methods/*/coverage
}

case $param in
    "i" | "init" | "-i")
        create_folders
        create_files
            ;;
    "b" | "build" | "-b")
        build_project
        echo "Project compiled"
        echo "/ __| | | |/ __/ __/ _ \/ __/ __|"
        echo "\__ \ |_| | (_| (_|  __/\__ \__ \\"
        echo "|___/\__,_|\___\___\___||___/___/"
                                   
        ;;
    "r" | "run" | "-r")
        run_project
        ;;
    "c" | "clear" | "-c")
        clear_project
        ;;
    "ca" | "clear-project" | "--clear-project")
        clear_all_project
        ;;
    "cc" | "clear-cache" | "--clear-cache")
        go clean -testcache
        ;; 
    "q" | "query" | "-q")
        run_insert_query
        ;;
    "t" | "test" | "-t")
        go clean -testcache
        run_test
        ;;
    "h" | "??" | "help" | "-q" | "--help")
        echo "Help:"
        echo "i: Inicialice the project, when download the project is the fist choice to use"
        echo "b: build the project after execute"
        echo "r: run the project to create the files"
        echo "c: to clear files folder"
        echo "ca: to clear all project"
        echo "cc: to clear the cache"
        echo "q: run the project to create insert sql tables"
        echo "t: run the test"
        echo "h: help"
            ;;
esac