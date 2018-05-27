# FIT3036 - Backend API

[![Maintainability](https://api.codeclimate.com/v1/badges/e93c83eb780cf2a1b94a/maintainability)](https://codeclimate.com/github/dylanpinn/FIT3036-backend/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e93c83eb780cf2a1b94a/test_coverage)](https://codeclimate.com/github/dylanpinn/FIT3036-backend/test_coverage)

## Overview

This repo contains the backend API for my road surface area calculator. It is
split up into 3 packages.

* Area
  * Shared functionality.
  * Calculate Distance.
  * Calculate Area.
  * Calculate Surface Area of Roads within Area.
* Local
  * Local server runs on port 8080.
  * Accessed via http://localhost:8080
    * `POST /area`
      * Calculate area of a rectangle.
      * Post body:
        * `north: float64, south: float64, east: float64, west: float64`
    * `POST /road`
      * Calculate surface area of roads within a rectangle.
      * Post body:
        * `north: float64, south: float64, east: float64, west: float64`
* Remote
  * `area`
    * AWS Lambda implementation of the local function.
  * `road`
    * AWS Lambda implementation of the local function.

## Requirements

### Local

* `go` Go programming language >= 1.10 (https://golang.org/)
* `dep` Go dependency management tool (https://github.com/golang/dep)

### Deployment

All local dependencies including:

* `make` GNU Make (https://www.gnu.org/software/make/)
* `node` NodeJS >= 6 (https://nodejs.org/en/)
* `yarn` Yarn package manager (https://yarnpkg.com/en/)
* `aws` AWS CLI (https://aws.amazon.com/cli/)
  * API keys need to be configured.

## Install

* `dep ensure` installs the required dependencies.
* `yarn install` installs the required dependencies for deploying to production.

## Deploy

* `make deploy`
  * This makes the remote binaries and deploys them to AWS using the
    `serverless` framework (http://serverless.com/)

## Test

* `go test ./...` this runs all of the tests in the current directory and all
  subdirectories.
