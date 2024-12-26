# Events-Registrations REST API

A RESTful API that allows users to manage events and reservations with ease.
## Features

Users can:

  -  Register: Sign up for an account.
  - Login: Authenticate to access the API.
  -  Create Events: Organize and share events.
  -  Edit Events: Update event details.
  -  Delete Events: Remove events from the system.
  -  Create Reservations: Register for events.
  -  Delete Reservations: Cancel event registrations.

Technology Stack

This project uses:
  -  Golang: The core programming language for building the API.
  -  Gin Framework: For handling HTTP requests and creating a lightweight yet powerful web server.
  -  SQLite: A simple, file-based relational database to store user, event, and reservation data.

Getting Started

   Clone the repository:

    git clone <repository-url>
    cd events-registrations-api

Install dependencies:
Ensure you have Go installed, then run:

    go mod tidy

Run the application:

    go run main.go
 The API will be accessible at http://localhost:8080.

## API Endpoints

### Authentication

    POST /register: Register a new user.
    POST /login: Authenticate and obtain a token.

### Events

    POST /events: Create a new event.
    GET /events: Retrieve all events.
    GET /events/:id: Retrieve details of a specific event.
    PUT /events/:id: Update an event.
    DELETE /events/:id: Delete an event.

### Reservations

    POST /reservations: Create a reservation for an event.
    GET /reservations: Retrieve all reservations.
    DELETE /reservations/:id: Cancel a reservation.
