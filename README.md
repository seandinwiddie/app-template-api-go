# app-template-api-go

This project is an API for an application template, written in Go.

## Description

This is an API for an application template. It serves as the backend for fetching and providing data to the frontend application.

This project provides a flexible backend for an application template, allowing easy content updates through the JSON file while maintaining a structured API for frontend consumption.

## Main Components

### API Server (main.go)
- Uses Gorilla Mux for routing
- Implements CORS for cross-origin requests
- Reads initial state from a JSON file
- Provides endpoints for:
  - Homepage
  - Status check
  - Fetching all data
  - Dynamic endpoints for each key in the initial state

### Initial State Data (data/initialState.json)
- Contains structured data for the application template, including:
  - BDD (Behavior-Driven Development) tests
  - Brand information
  - Portfolio features
  - Application procedures
  - Theme toggle settings
  - Navigation data

## Project Structure
- `main.go`: Contains the main API logic
- `data/`: Stores the initial state data

## Dependencies
- Gorilla Mux for routing
- RS/CORS for handling cross-origin requests

## Key Features
- Dynamic endpoint generation based on the initial state structure
- Centralized data management through a JSON file
- Support for multiple themes
- BDD test scenarios included in the data

## Running the Project
The server can be started using `go run main.go`

## Data Flow
The project follows a structure where data is stored in a JSON file, loaded into the API, and then served through various endpoints. This allows for easy updates to the portfolio content without changing the API code.