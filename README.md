# gorest 💤

A test Go REST API

## Resources
1. [GORM](https://gorm.io/)
2. [Chi](https://go-chi.io/#/)
3. [Docker](https://docs.docker.com/)
4. [Git](https://git-scm.com/)

## Setup
1. Install Docker
2. Install Git and clone the repository using the following command:
   ```
   git clone https://github.com/snretuerma/gorest.git
   ```
3. Create an empty directory and name it `db` for database storage
4. Create a `.env` file and copy the content of `.env.example`
   ```
   cp env.example env
   ```
5. Build the Docker container using the following command:
   ```
    docker compose up --build
   ```
6. Access the api through a REST client like [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/)
    ```
    localhost:3200/api/posts
    ```