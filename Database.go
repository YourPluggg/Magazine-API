package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

create database productdb;
use productdb;

create table product (
		ID         int auto_increment primary key,
		Definition varchar(30) not null,
		Name       varchar(30) not null,
		Price      float ot null
)
