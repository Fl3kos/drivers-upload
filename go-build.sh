#!/bin/bash

filesFolder=files
filesToReadFolder=filesToRead
filesToReadLayouts=filesToRead/layouts
filesSql=files/sql
filesSqlShop=files/sqlShops
filesJson=files/userCouchbase
filesCsv=files/usersAndPasswords
filesNames=files/names
filesAclSql=files/aclSql
filesUsersAcl=files/usersEndPoint
logsFolder=logs
filesUserList=files/userList
filesExpedition=files/expeditionSql
filesSorterMap=files/sorterMap
param=$1

create_folders(){
    #create files folder
    create_folder $filesFolder
    create_folder $filesToReadFolder
    create_folder %filesToReadLayouts
    create_folder $logsFolder

    create_folder $filesSql
    create_folder $filesCsv
    create_folder $filesJson
    create_folder $filesSqlShop
    create_folder $filesNames
    create_folder $filesAclSql
    create_folder $filesUsersAcl
    create_folder $filesUserList
    create_folder $filesExpedition
    create_folder $filesSorterMap

}

create_folder(){
    if [ -d $1 ]
    then
        echo “La capeta ya existe.”
    else
        mkdir $1
    if [ $? -eq 0 ]
    then
        echo “se ha creado con éxito $1
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
    touch ./filesToRead/token.txt
}

create_user_list_file(){
    touch ./filesToRead/userList.json
    echo "{
        \"PKR\": 5,
        \"CRD\": 3,
        \"ADM\": 1
}" >> ./filesToRead/userList.json
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
    cd ../layouts
    go build
    rm layouts
    cd ../acl-users
    go build
    rm acl-users
    cd ../users-shop
    go build
    rm users-shop
    cd ../..

}

run_project(){
    rm ./logs/logs.log
    go run ./cmd/main/main.go
}

run_insert_query(){
    rm ./logs/logs.log
    go run ./cmd/create-shops/main.go
}

run_layouts(){
    rm ./logs/logs.log
    go run ./cmd/layouts/main.go
}

run_users_list(){
    rm ./logs/logs.log
    go run ./cmd/users-shop/main.go
}

run_acl_users(){
    rm ./logs/logs.log
    go run ./cmd/acl-users/main.go

clear_project(){
    rm ./files/*/*
    rm ./logs/*
}

clear_all_project(){
    rm ./files/*
    rm ./files/*/*
    rm ./logs/*
    rm ./filesToRead/*

    rmdir ./files/*
    rmdir ./files
    rmdir ./filesToRead
    rmdir ./logs
}

run_test() {
    rm ./logs/logs_test.log
    go test -timeout 30s -run ^TestComprobeDniAndNie$ drivers-create/methods/dni
    go test -timeout 30s -run ^TestGenerateJson$ drivers-create/methods/json
    go test -timeout 30s -run ^TestSql$ drivers-create/methods/sql
    go test -timeout 30s -run ^TestUsersToPasswords$ drivers-create/methods/userToPassword
    go test -timeout 30s -run ^TestConvertAllDnisToUsers$ drivers-create/methods/dniToUser
}

case $param in
    "i" | "init" | "-i")
        create_folders
        create_files
        create_user_list_file
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
    "u" | "users-list" | "-u")
        run_users_list
        ;;
    "l" | "layouts" | "-l")
        run_layouts
        ;;
    "a" | "acl" | "-a")
        run_acl_users
        ;;
    "t" | "test" | "-t")
        go clean -testcache
        run_test
        ;;
    "p" | "pull" | "-p")
        git pull
        clear_all_project
        create_folders
        create_files
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
        echo "u: run the users list main to create users to insert in wms warehouse"
        echo "l: run the layouts constructor sql file"
        echo "t: run the test"
        echo "p: pull repo, delete and create new folders"
        echo "h: help"
            ;;
esac