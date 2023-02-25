package main

import "time"

type Credentials struct {
	id        string
	userId    string
	title     string
	login     string
	password  string
	createdAt time.Time
	updatedAt time.Time
	meta      string
}

type Text struct {
	id        string
	userId    string
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
	meta      string
}

type BinaryData struct {
	id        string
	userId    string
	title     string
	content   string
	createdAt time.Time
	updatedAt time.Time
	meta      string
}

type BankCard struct {
	id         string
	userId     string
	title      string
	cardHolder string
	cardNumber string
	cardExpire string
	cardCVV    string
	createdAt  time.Time
	updatedAt  time.Time
	meta       string
}

type MemoryStorage struct {
	credentials Credentials
	text        Text
	binaryData  BinaryData
	bankCard    BankCard
}
