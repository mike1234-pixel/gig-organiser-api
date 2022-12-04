# Gig Organiser

Gig Organiser is a job tracking app that uses Docker, a Golang REST API, and a Postgres database to help users manage their job search and application process.

## Prerequisites

To run Gig Organiser, you will need the following software installed on your machine:

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- [Go](https://golang.org/)

## Running the application

To run Gig Organiser, follow these steps:

1. Clone this repository to your local machine.

   ```
   git clone https://github.com/mike1234-pixel/gig-organiser-api.git
   ```

2. Change into the project directory.

   ```
   cd gig-organiser-api
   ```

3. Build and run the Docker containers using Docker Compose.

   ```
   docker-compose up
   ```

4. Once the containers are running, you can access Gig Organiser at http://localhost:3000/.

## Stopping the application

To stop Gig Organiser, press `CTRL+C` in the terminal window.
