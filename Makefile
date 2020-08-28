# Make sure you've copied .env.example to .env and added your API-keys
include .env

build:
	go build -ldflags="-X github.com/jadlers/gonow/location.apiKey=$(PLATSUPPSLAG_KEY) -X github.com/jadlers/gonow/trip.apiKey=$(RESEPLANERARE_KEY)"

install:
	go install -ldflags="-X github.com/jadlers/gonow/location.apiKey=$(PLATSUPPSLAG_KEY) -X github.com/jadlers/gonow/trip.apiKey=$(RESEPLANERARE_KEY)"
