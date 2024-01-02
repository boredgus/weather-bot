## Overview

This is an implementation of telegram bot to get weather forecast for specified location at specified time. Firstly user creates subscription for weather forecast by providing location coordinates (latitude and longitude) and time when weather forecast should be sent by bot. There can be multiple subscriptions. Each subscription will be sent separately. User can update location/time of any subscription or delete whole subscription. Bot checks for existed subscriptions and sends one weather forecast per subscription every day.

## How it works

Main entrance of application is `cmd/weather-bot/main.go`. It contains bot initialization, bot updates handler, and ticker to send weather foreact for existed subscriptions. All related business logic is contained in `internal/` firectory.

### Run project locally

Basic commands to run project locally are listed in _Makefile_ in the root.

If you are familiar with [make](https://www.gnu.org/software/make/) tool, run `make start` in terminal.If you already run the project locally and want to rebuild bot server, run `make restart`.
