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
        echo “se ha creaco con éxito $filesFolder”
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
        echo “se ha creaco con éxito $filesToReadFolder”
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
        echo “se ha creaco con éxito $logsFolder”
    else
        echo “Ups! Algo ha fallado al crear ”
    fi
    fi
}

create_files(){
    touch ./filesToRead/dnis.txt
    touch ./filesToRead/names.txt
    touch ./filesToRead/phoneNumbers.txt
    touch ./filesToRead/shopCode.txt
}

build_project(){
    cd ./methods/dni
    go build
    cd ../converts
    go build
    cd ../file
    go build
    cd ../json
    go build
    cd ../sql
    go build
    cd ../csv
    go build
    cd ../log
    go build
    cd ..
    go build
    cd ../cmd/main
    go build
    rm main
    cd ../..
}

run_project(){
    go run ./cmd/main/main.go
}

case $param in
    "i")
        create_folders
        create_files
            ;;
    "b")
        build_project
        echo "Project compiled"
        ;;
    "r")
        run_project
        ;;
    "h")
        echo "Help:"
            ;;
esac