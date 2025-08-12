# Subscription Service

This is a simple REST service for managing user subscriptions, built with Go. It's a test task for a Junior Golang Developer position. We basically built a full CRUDL API, hooked it up to a Postgres database, and containerized the whole thing with Docker.

### What's Inside?

*   **Go:** The programming language for the API itself.
*   **Gin:** A fast, easy-to-use web framework for Go.
*   **PostgreSQL:** Our database for storing all the subscription data.
*   **Docker & Docker Compose:** For running the app and the database in containers, making setup a breeze.
*   **golang-migrate:** For handling database migrations, so the schema is always up to date.
*   **zap:** For structured, high-performance logging.
*   **Swagger:** For automatically generating API documentation.

### Getting Started

Everything runs in Docker, so you don't need to install Go or Postgres on your machine. Just make sure you have Docker Desktop running.

1.  **Clone the repo** (if you haven't already).
2.  **Open your terminal** and navigate to the project directory.
3.  **Run this command:**

    ```bash
    docker-compose up --build
    ```

This will build the Go application, start the Postgres database, and run the migrations. The API will be available at `http://localhost:8080`.

### Testing the API

You can test the API using `curl` or any API client you like (like Postman or Insomnia). You can also use the interactive Swagger docs.

#### Interactive Docs (Easy Mode)

1.  Open your browser and go to `http://localhost:8080/swagger/index.html`.
2.  You'll see all the available endpoints. You can expand them, fill in parameters, and hit "Execute" to try them out live.

#### Using `curl` (Hacker Mode)

Here’s a quick run-through of the API using `curl`.

**1. Create a subscription:**

```bash
curl -X POST \
  http://localhost:8080/api/subscriptions/ \
  -H 'Content-Type: application/json' \
  -d '{
    "service_name": "Yandex Plus",
    "price": 400,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "07-2025"
  }'
```

This will return a JSON object with the ID of the new subscription. **Copy that ID!**

**2. List all subscriptions:**

```bash
curl http://localhost:8080/api/subscriptions/
```

**3. Get a specific subscription by its ID:**

(Replace `your_id_goes_here` with the ID you copied.)

```bash
curl http://localhost:8080/api/subscriptions/your_id_goes_here
```

**4. Update a subscription:**

(Replace `your_id_goes_here` with the ID you copied.)

```bash
curl -X PUT \
  http://localhost:8080/api/subscriptions/your_id_goes_here \
  -H 'Content-Type: application/json' \
  -d '{
    "service_name": "Yandex Plus Premium",
    "price": 500,
    "user_id": "60601fee-2bf1-4721-ae6f-7636e79a0cba",
    "start_date": "07-2025"
  }'
```

**5. Calculate total cost:**

This calculates the total subscription cost for a specific user for a given month and year.

```bash
curl "http://localhost:8080/api/subscriptions/total?user_id=60601fee-2bf1-4721-ae6f-7636e79a0cba&year=2025&month=7"
```

**6. Delete a subscription:**

(Replace `your_id_goes_here` with the ID you copied.)

```bash
curl -X DELETE http://localhost:8080/api/subscriptions/your_id_goes_here
```