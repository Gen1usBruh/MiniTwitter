Mini Twitter
This project is a simplified clone of a Twitter.

Description
Mini Twitter allows users to perform basic social networking actions such as creating posts, 
, like tweets and comments, leave comments, following other users, and viewing timelines.

Prerequisites
- postgres docker image postgres:15.6
- golang latest installed
- environment variables exported from .env

Installation and Setup

Clone the repository:
git clone -b dev https://github.com/Gen1usBruh/MiniTwitter.git

Navigate to the project directory:
cd mini-twitter

Install dependencies:
go mod tidy

Run the application:
go run main.go

Usage
Once the application is running, you can access it at http://localhost:8080.
Use the provided API endpoints to interact with Mini Twitter.
