# FIT3036 - Backend API

## Overview

This repo contains the backend API for my road surface area calculator.

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
