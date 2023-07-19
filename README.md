# drivers-upload
- [Description](#Description)
- [Build, test and run](#Build)
- [How to run the builds](#Run)

## Description

This is the support team's repository which is responsible for creating users for transport, acl and generating the shipping and picking layouts for wms warehouses.

## Build
To execute the project, the first thing to do is to start the project with the script "go-build.sh" and with the parameter "i"

The script have 10 options:

i: Inicialice the project, when download the project is the fist choice

b: Build the project before execute

r: Run to generate driver users, publish to Auth and publish role in ACL

u: Run the project to create users in Auth and publis role in ACL

l: Run to transform the layouts excell to sql file

c: Clear files folder

ca: Clear all project files and folders

cc: Clear the go cache

t: Launch the unit test

h: Help

## Run

The 3 builds are
- Drivers create
- Warehouses users
- Generate Layouts

### Drivers Create

To create drivers are 4 files:

- names
- dnis
- shops
- phoneNumbers

In names put the names to share with you to create users separete to line space

In dnis put the dnis to share with you to create users separete to line space

In shops put the warehouse numbers separate to line (-) to the warehouse name to share with you to create users, with publish two or more warehouse separete to line space

In phoneNumbers put the phoneNumbers to share with you to create users separete to line space

To generate drivers of two or more warehouse, in order, separete with white line to distinct warehouse number

### Warehouses users

To generate the warehouses users yo have know the count of users and modify "userList.json" file with the find parameters

### Genearate Layouts

This build generate the sql file with layout, to build, in the layouts folder you put the xlsx file with layouts, to generate sql files, the xlsx have this names:
- "expedition.xlsx"
- "picking.xlsx"