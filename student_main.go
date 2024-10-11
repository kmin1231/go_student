package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"	// web framework 'Gin'
)

type Student struct {
	Id int			`json:"id,omitempty"`
	Name string		`json:"name" binding:"required"`
	Age int			`json:"age" binding:"gte=1,lte=150"`
	Score int		`json:"score,omitempty" binding:"gte=0,lte=100"`
	Grade string 	`json:"grade,omitempty"`
}

var students map[int]Student	// map to save student information
var lastId int					// variable to track the last student

func SetupHandlers(g *gin.Engine) {	// routing handler
	g.GET("/students", GetStudentsHandler)
	g.GET("/students/:id", GetStudentHandler)
	g.POST("/students", PostStudentHandler)
	g.DELETE("/students/:id", DeleteStudentHandler)
	g.PUT("/students/:id", PutStudentHandler)
	g.PATCH("/students/:id", PatchStudentHandler)	// age, score, grade
	g.GET("/students/search", SearchStudentHandler)   // name or, grade
	g.POST("/students/:id/grade", CalculateGradeHandler)
	
	students = make(map[int]Student)
	// lastId = 0

	// dummy data
	students[1] = Student{1, "aaa", 16, 87, "NA"}
	students[2] = Student{2, "bbb", 18, 98, "NA"}
	lastId = 2
}


type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentsHandler(c *gin.Context) {
	list := make(Students, 0)

	for _, student := range students {
		list = append(list, student)
	}

	sort.Sort(list)   // sorted student list

	var result string
	for _, student := range list {
		result += fmt.Sprintf(
			"ID: %d, Name: %s, Age: %d, Score: %d, Grade: %s\n",
			student.Id, student.Name, student.Age, student.Score, student.Grade,
		)
	}

	// c.JSON(http.StatusOK, list)   // return in JSON format
	c.String(http.StatusOK, result)
}

func GetStudentHandler(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	student, ok := students[id]

	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	result := fmt.Sprintf(
		"ID: %d\nName: %s\nAge: %d\nScore: %d\nGrade: %s\n",
		student.Id, student.Name, student.Age, student.Score, student.Grade,
	)
		
	// c.JSON(http.StatusOK, list)   // return in JSON format
	c.String(http.StatusOK, result)
}

func PostStudentHandler(c *gin.Context) {
	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	student.Grade = calculateGrade(student.Score)

	lastId++
	student.Id = lastId
	students[lastId] = student
	c.String(http.StatusCreated, fmt.Sprintf("Success to add id: %d\n", lastId))
}

func PutStudentHandler(c *gin.Context) {
	idstr := c.Param("id")

	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var student Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if _, exists := students[id]; !exists {
		c.JSON(http.StatusNotFound, "Student not found")
		return
	}

	student.Id = id
	student.Grade = calculateGrade(student.Score)
	students[id] = student
	c.String(http.StatusCreated, fmt.Sprintf("Success to update id: %d\n", id))
}

type UpdateStudent struct {
	Age   *int   `json:"age" binding:"omitempty,gte=1,lte=150"`
	Score *int   `json:"score" binding:"omitempty,gte=0,lte=100"`
	Grade *string `json:"grade"`
}

func PatchStudentHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var update UpdateStudent
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	student, exists := students[id]
	if !exists {
		c.JSON(http.StatusNotFound, "Student not found")
		return
	}

	if update.Age != nil {
		student.Age = *update.Age
	}
	if update.Score != nil {
		student.Score = *update.Score
	}
	if update.Grade != nil {
		student.Grade = *update.Grade
	}

	students[id] = student

	result := fmt.Sprintf(
		"ID: %d\nName: %s\nAge: %d\nScore: %d\nGrade: %s\n",
		student.Id, student.Name, student.Age, student.Score, student.Grade,
	)
		
	// c.JSON(http.StatusOK, list)   // return in JSON format
	c.String(http.StatusOK, result)
}

func DeleteStudentHandler(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	delete(students, id)
	c.String(http.StatusOK, fmt.Sprintf("Success to delete id: %d\n", id))
}

func calculateGrade(score int) string {
	switch {
		case (score >= 90):
			return "A"
		case (score >= 80):
			return "B"
		case (score >= 70):
			return "C"
		case (score >= 60):
			return "D"
		default:
			return "F"
	}
}

func CalculateGradeHandler(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	student, ok := students[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	grade := calculateGrade(student.Score)
	student.Grade = grade
	students[id] = student

	result := fmt.Sprintf(
		"ID: %d\nName: %s\nAge: %d\nScore: %d\nGrade: %s\n",
		student.Id, student.Name, student.Age, student.Score, grade, // 학생의 정보 및 학점
	)

	c.String(http.StatusOK, result)
}


func SearchStudentHandler(c *gin.Context) {
	nameQuery := c.Query("name")
	gradeQuery := c.Query("grade")
	ageQuery := c.Query("age")

    var age int
    var err error
    if ageQuery != "" {
        age, err = strconv.Atoi(ageQuery)
        if err != nil {
            c.JSON(http.StatusBadRequest, "Invalid age value")
            return
        }
    }

	list := make(Students, 0)
	for _, student := range students {

		if nameQuery != "" && !strings.Contains(student.Name, nameQuery) {
			continue
		}

		if gradeQuery != "" && student.Grade != gradeQuery {
			continue
		}


        if ageQuery != "" && student.Age != age {
            continue
        }

		list = append(list, student)
	}

	sort.Sort(list)

	var result string
	for _, student := range list {
		result += fmt.Sprintf(
			"ID: %d, Name: %s, Age: %d, Score: %d, Grade: %s\n",
			student.Id, student.Name, student.Age, student.Score, student.Grade,
		)
	}

	c.String(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	SetupHandlers(r)
	r.Run(":3000")
}