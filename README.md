# Culture-Hub

A collaborative platform where team members can submit ideas to enhance team culture and engagement. CultureHub allows users to propose, upvote, downvote, and track upcoming/ongoing activities, while HR/Admin members can manage approvals and track progress.

## Project URL
[GitHub Repository](https://github.com/tarunngusain08/Culture-Hub)

---

## Table of Contents
- [Features](#features)
- [API Endpoints](#api-endpoints)
- [Technologies](#technologies)
- [Installation](#installation)
- [Usage](#usage)
- [Database Models](#database-models)
- [Contributing](#contributing)
- [License](#license)

---

## Features

### User Roles
1. **Team Members**
   - Can submit new ideas for team-building, culture improvement, and wellness.
   - Vote on submitted ideas.
   - View ongoing activities and their progress.

2. **HR/Admin**
   - Approve or reject submitted ideas.
   - Assign coordinators or leaders to activities.
   - Track and manage the status of activities.
   - Generate reports on completed activities and their impact.

### Core Functionalities
- **Idea Submission**: Users can submit ideas for activities, with optional timeline and impact estimation.
- **Voting**: Users can upvote or downvote ideas. Once an activity is completed, users can vote on its impact on team culture.
- **Activity Updates**: Users and coordinators can post updates on ongoing activities.
- **Notifications**: Team members receive notifications for updates on ideas and activities.

---

## API Endpoints

Below are the primary endpoints available in the CultureHub platform:

### 1. Register/Login API
- **POST** `/api/v1/register`
- **POST** `/api/v1/login`

### 2. Post Idea API
- **POST** `/api/v1/ideas`

### 3. Get Idea API
- **GET** `/api/v1/ideas/{id}`

### 4. Upvote/Downvote API
- **POST** `/api/v1/ideas/{id}/vote?action={upvote|downvote}`

### 5. Update Idea Status API
- **PATCH** `/api/v1/ideas/{id}/status?action={submitted|approved|ongoing|completed}`

### 6. Fetch all ideas
- **GET** `/api/v1/ideas`

### 7. Ongoing/Upcoming Activities
- **GET** `/api/v1/activities`
---

## Technologies

- **Backend**: Golang
- **Frontend**: React (planned)
- **Database**: PostgreSQL
- **Authentication**: JWT
- **Containerization**: Docker
- **Orchestration**: Kubernetes (future enhancement)

---

## Installation

### Prerequisites
- Go 1.16+ installed
- PostgreSQL installed and running
- Docker (optional, for containerized deployment)

### Steps

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/tarunngusain08/Culture-Hub.git
   cd Culture-Hub
   ```

2. **Set up the Database:**
   - Create a PostgreSQL database named `culturehub`.
   - Run the migrations (if using a migration tool) to set up the tables.

3. **Configure Environment Variables:**
   Create a `.env` file and set the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=culturehub
   DB_PASSWORD=culturehub
   DB_NAME=culturehub
   JWT_SECRET=jwt_secret
   ```

4. **Run the Application:**
   ```bash
   go run main.go
   ```

---

## Usage

### Running the Application Locally

- Access the platform through `http://localhost:8080` (default port for the Go server).
- Use tools like Postman or Curl to interact with the API.

### API Interaction Example (Using Curl)

1. **Register a User:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/register \
   -H "Content-Type: application/json" \
   -d '{"username": "john_doe", "email": "john.doe@example.com", "password": "password123"}'
   ```

2. **Post a New Idea:**
   ```bash
   curl -X POST http://localhost:8080/api/v1/ideas \
   -H "Authorization: Bearer JWT_TOKEN" \
   -H "Content-Type: application/json" \
   -d '{"title": "New Team Building Activity", "description": "Fun activity", "tags": ["team-building"]}'
   ```

---

## Database Models

```sql
CREATE TYPE user_role AS ENUM ('TeamMember', 'HR', 'Admin');
CREATE TYPE idea_status AS ENUM ('Submitted', 'Approved', 'InProgress', 'Completed');

CREATE TABLE "User" (
                        id SERIAL PRIMARY KEY,
                        username VARCHAR(50) UNIQUE NOT NULL,
                        email VARCHAR(100) UNIQUE NOT NULL,
                        password_hash VARCHAR(255) NOT NULL,
                        role user_role DEFAULT 'TeamMember',
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        last_login TIMESTAMP
);

CREATE TABLE "Idea" (
                        id SERIAL PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        description TEXT NOT NULL,
                        tags VARCHAR(255),
                        timeline DATE,
                        impact_estimation TEXT,
                        user_id INT REFERENCES "User"(id) ON DELETE CASCADE,
                        status idea_status DEFAULT 'Submitted',
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        vote_count INT DEFAULT 0
);

CREATE OR REPLACE FUNCTION update_timestamp()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON "Idea"
    FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

```

## Flow Diagram
![CultureHub](https://github.com/user-attachments/assets/482de8d9-1577-4d65-ad3b-492f781d14d2)


## Future Enhancements

[CultureHub.md](https://github.com/user-attachments/files/17455456/CultureHub.md)<img width="1085" alt="Screenshot 2024-10-20 at 6 09 42 PM" src="https://github.com/user-attachments/assets/2f3e1c19-64db-4890-8155-e4b7a2f3bd59">

