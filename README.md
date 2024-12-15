<h1><code>Go</code>&nbsp;Student Information Management System</h1>

## 📖 Overview
This project is focused on building and deploying a <b>RESTful API</b> for managing student information. It demonstrates how to implement a RESTful API with <code><b>Go</b></code> and the <code><b>Gin</b></code> web framework, and how to <b>containerize</b> the application using <b><code>Docker</code></b> for local development and <b><code>Kubernetes</b></code> in a scalable environment.

## 💫 SERVER

- provides a set of **RESTful APIs** to manage student records
- handles API requests for performing **CRUD operations** (Create, Read, Update, Delete) on student information

## 💫 CLIENT

- a shell script (**`request.sh`**) that sends requests to the server using **cURL**
- simulates interactions with the **server** by sending requests to create, retrieve, update, and delete student records, demonstrating how to **interact** with the API endpoints

## 💫 Containerized Deployment

- both the **server** and the **client** is fully containerized → can be deployed using **`Docker Compose`** for local development and testing, and **`Kubernetes`** for distributed and scalable production environments

<br>

## 📜 `Student` struct
- <b><code>Id</code> (int)</b>: a <b>unique</b> identifier for each student<br>
- <b><code>Name</code> (string)</b>: the name of the student -- required for creating or updating a student record<br>
- <b><code>Age</code> (int)</b>: the age of the student<br>
- <b><code>Score</code> (int)</b>: the score of the student<br>
- <b><code>Grade</code> (string)</b>: the grade of the student -- automatically determined based on <code><b>Score</b></code>

## 📜 API Endpoints

- <code><b>GET</b> /students</code>: retrieve a list of all students<br>
- <code><b>GET</b> /students/:id</code>: retrieve a specific student by ID<br>
- <code><b>POST</b> /students</code>: add a new student -- <b>JSON</b> body required<br>
- <code><b>DELETE</b> /students/:id</code>: delete a student by ID<br>
- <code><b>PUT</b> /students/:id</code>: update an existing student -- <b>JSON</b> body required<br>
- <code><b>PATCH</b> /students/:id</code>: update specific fields (age, score, grade) of an existing student -- <b>JSON</b> body optional<br>
- <code><b>GET</b> /students/search</code>: search for students by name, grade, or age<br>
<code><b>POST</b> /students/:id/grade</code>: calculate and update the grade for a specific student<br>

<br>

## Project Structure

```
.
├── client
│   ├── 🐳 Dockerfile
│   └── request.sh
├── server
│   ├── main.go
│   ├── 🐳 Dockerfile
│   ├── go.mod
│   ├── go.sum
│   ├── student_mux/
│   └── wait-for-it.sh
├── 🐳🐙 docker-compose.yaml
├── ☸️ client-deployment.yaml
├── ☸️ service.yaml
├── ☸️ server-deployment.yaml
└── README.md
```

📌 <code><b>student_mux</b></code>: a simple <b>RESTful API</b> for managing student information using the <code><b>Gorilla/Mux</b></code> web framework

<br>

## 📚 Reference
- 공봉식, 『[<b>Tucker의 Go 언어 프로그래밍</b>](https://product.kyobobook.co.kr/detail/S000213858928)』, 골든래빗(주), 2024.

<br>
<br>
