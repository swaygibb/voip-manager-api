# VoIP Manager API

## Overview

The VoIP Manager API is a RESTful service built with Go and the Gin framework. It provides functionality to manage call logs and voicemails, including creating, updating, deleting, and searching records.

## Features

- Manage call logs with CRUD operations
- Manage voicemails with CRUD operations
- Search call logs and voicemails using query parameters
- Database initialization and migration using SQLite
- Health check endpoint for verifying API status

## Project Structure

- `main.go`: Entry point of the application
- `db/`: Database setup and configuration
  - `database.go`: Initializes the SQLite database and performs migrations
- `helpers/`: Utility functions
  - `callog_helper.go`: Helper for building search queries for call logs
- `models/`: Data models for the application
  - `calllog.go`: Defines the CallLog model
  - `voicemail.go`: Defines the Voicemail model
- `routes/`: API route handlers
  - `calllogs.go`: Handlers for call log endpoints
  - `voicemails.go`: Handlers for voicemail endpoints

## Endpoints

### Ping

- GET /ping
  - Returns a simple response to verify the API is running

### Call Logs

- GET /call-logs
  - Retrieves all call logs
- POST /call-logs
  - Creates a new call log
- PUT /call-logs/:id
  - Updates an existing call log by ID
- DELETE /call-logs/:id
  - Deletes a call log by ID
- GET /call-logs/:id
  - Retrieves a call log by ID
- GET /call-logs/search
  - Searches call logs based on query parameters

### Voicemails

- GET /voicemails
  - Retrieves all voicemails
- POST /voicemails
  - Creates a new voicemail
- PUT /voicemails/:id
  - Updates an existing voicemail by ID
- DELETE /voicemails/:id
  - Deletes a voicemail by ID
- GET /voicemails/:id
  - Retrieves a voicemail by ID
- GET /voicemails/search
  - Searches voicemails based on query parameters

## Setup

1. Clone the repository
2. Install dependencies using go mod tidy
3. Run the application with go run main.go
4. The API will be available at http://localhost:8080

## Database

The application uses SQLite for data storage. The database file is named voip_manager.db and is automatically created and migrated when the application starts.
