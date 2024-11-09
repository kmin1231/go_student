<h2><code>Go</code>&nbsp;Student Information Management System</h2>

### 📖 Overview
This is a simple <b>Command Line Interface (CLI)</b> based program for managing student information, using <code><b>Go</b></code> and the <code><b>Gin</b></code> web framework. This program serves as a practical exercise for implementing <b>RESTful APIs</b>. It facilitates the management of student records through <b>CRUD</b> functionalities, allowing users to add, update, delete, and retrieve student details.

### 📜 `Student` struct
<b><code>Id</code> (int)</b>: a <b>unique</b> identifier for each student<br>
<b><code>Name</code> (string)</b>: the name of the student -- required for creating or updating a student record<br>
<b><code>Age</code> (int)</b>: the age of the student<br>
<b><code>Score</code> (int)</b>: the score of the student<br>
<b><code>Grade</code> (string)</b>: the grade of the student -- automatically determined based on <code><b>Score</b></code>

### 📜 API Endpoints

<code><b>GET</b> /students</code>: retrieve a list of all students<br>
<code><b>GET</b> /students/:id</code>: retrieve a specific student by ID<br>
<code><b>POST</b> /students</code>: add a new student -- <b>JSON</b> body required<br>
<code><b>DELETE</b> /students/:id</code>: delete a student by ID<br>
<code><b>PUT</b> /students/:id</code>: update an existing student -- <b>JSON</b> body required<br>
<code><b>PATCH</b> /students/:id</code>: update specific fields (age, score, grade) of an existing student -- <b>JSON</b> body optional<br>
<code><b>GET</b> /students/search</code>: search for students by name, grade, or age<br>
<code><b>POST</b> /students/:id/grade</code>: calculate and update the grade for a specific student<br>

<br>

📌 <code><b>student_mux</b></code>: a simple <b>RESTful API</b> for managing student information using the <code><b>Gorilla/Mux</b></code> web framework

<br>

## 📚 Reference
- 공봉식, 『[<b>Tucker의 Go 언어 프로그래밍</b>](https://product.kyobobook.co.kr/detail/S000213858928)』, 골든래빗(주), 2024.

<br>
<br>
