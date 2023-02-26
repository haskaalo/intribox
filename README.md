# Intribox (In development)

## What is this?

Intribox is a "self-hostable" media backup service. This service allows users to safely backup pictures and videos to local storage or a S3-compotable service. It also allows them to search, view and share them.

In other words, a end-to-end encrypted alternative to Google Photos

## Features

* See all of your pictures and video online on the web
* Ability to share pictures privately with a link (todo)
* Create albums (in progress)
* Self-hostable with ability to store locally on a drive or a S3-compatible cloud provider
* Search pictures and videos based on location, date, etc... (in progress)

## Getting started (For testing)

Requirements: Docker, Docker-compose and Go

1. Download the repository
2. Run `docker compose -f "docker-compose.dev.yml" up -d --build`
3. Inside the folder client, run `npm install` and `npm build`
4. Run `go build` and `make setupaws` inside the repository
5. Run intribox

## Developers

### Explanation of the stack

The webserver is made in Golang while the front-end is a SAP made using React and TypeScript