#!/bin/bash

BASE_URL="http://server:3000"


echo "(1) Testing GET /students"
curl -X GET "$BASE_URL/students"
sleep 3
echo -e "\n"


echo "(2) Testing GET /students/:id"
curl -X GET "$BASE_URL/students/1"
sleep 3
echo -e "\n"


echo "(3) Testing POST /students"
curl -X POST "$BASE_URL/students" -d '{"name": "Alice", "age": 23, "score": 89}'
curl -X POST "$BASE_URL/students" -d '{"name": "Harry", "age": 21, "score": 77}'
curl -X POST "$BASE_URL/students" -d '{"name": "Peter", "age": 24, "score": 95}'
echo -e "\n"

curl -X GET "$BASE_URL/students"
sleep 3
echo -e "\n"


echo "(4) Testing DELETE /students/:id"
curl -X DELETE "$BASE_URL/students/2"

echo -e "\n"

curl -X GET "$BASE_URL/students"
sleep 3
echo -e "\n"


echo "(5) Testing PUT /students/:id"
echo -e "before 'PUT' request"
curl -X GET "$BASE_URL/students/1"
echo -e "\n"

curl -X PUT "$BASE_URL/students/1" -d '{"name": "Steve", "age": 26, "score": 90}'
echo -e "\n"

echo -e "after 'PUT' request"
curl -X GET "$BASE_URL/students/1"
sleep 3
echo -e "\n"


echo "(6) Testing PATCH /students/:id"
echo -e "before 'PATCH' request"
curl -X GET "$BASE_URL/students/3"
echo -e "\n"

curl -X PATCH "$BASE_URL/students/3" -d '{"score": 91}'
sleep 3
echo -e "\n"


echo "(7) Testing POST /students/:id/grade"
curl -X POST "$BASE_URL/students/1/grade"
curl -X POST "$BASE_URL/students/3/grade"
echo "grade updated for the score modified by 'PUT' & 'PATCH' request"
sleep 3
echo -e "\n"


echo "(8) Testing GET /students/search"
echo "search students with grade 'A'"
curl -X GET "$BASE_URL/students/search?grade=A"
sleep 3
echo -e "\n"