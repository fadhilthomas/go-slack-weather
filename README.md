# go-slack-weather
[![go-slack-weather](https://github.com/fadhilthomas/go-slack-weather/actions/workflows/go-slack-weather.yml/badge.svg?branch=main)](https://github.com/fadhilthomas/go-slack-weather/actions/workflows/go-slack-weather.yml)
[![license](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/fadhilthomas/go-slack-weather/issues)

![slack](https://user-images.githubusercontent.com/29804796/129493231-2de98bac-09ac-4686-a97d-487a344dd6a1.png)

Send current weather updates from OpenWeatherMap API to your Slack profile status using GitHub Action

---

## Contents

- [Setup](#setup)
- [Change Update Period](#change-update-period)

---

## Setup
* Fork this repository.
* Set these environment variables in GitHub repository secrets.

| **Variable** | **Value** |
|--|--|
| WEATHER_API | Get OpenWeatherMap API at `https://home.openweathermap.org/users/sign_up`. |
| CITY | Find your city id at `http://bulk.openweathermap.org/sample/`. |
| TIMEZONE | Find your timezone city at `https://www.iana.org/time-zones`. |
| SLACK_TOKEN | Your Slack User OAuth Token. |

## Change Update Period
* Change `cron` value inside `.github/workflows/go-slack-weather.yml`.
* If you need help to set the value, visit `https://crontab.guru`.