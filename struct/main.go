package main

import "fmt"

// Department represents a university department.
type Department struct {
    Name string      // Name of the department, e.g., "Computer Science"
    Courses []Course // Slice of courses offered by the department
}

// Course represents a course within a department.
type Course struct {
    Code string      // Unique identifier for the course, e.g., "CS101"
    Name string      // Human-readable name of the course
    Credits int      // Number of credit hours for the course
    Department *Department // Pointer to the department this course belongs to
    Enrollments []Enrollment // Slice of enrollments, linking students to this course
}

// Student represents a student.
type Student struct {
    ID   string       // Student's unique identifier
    Name string       // Student's name
    Enrollments []Enrollment // Slice of student's enrollments in various courses
}

// Enrollment links a student with a course.
type Enrollment struct {
    Student *Student  // Pointer to the student involved in this enrollment
    Course  *Course   // Pointer to the course involved in this enrollment
    Grade   float32   // Grade received by the student in this course
}

// Enroll enrolls a student in a course.
func (c *Course) Enroll(student *Student) {
    enrollment := Enrollment{Student: student, Course: c} // Creates a new Enrollment instance
    c.Enrollments = append(c.Enrollments, enrollment) // Adds the new enrollment to the course's slice
    student.Enrollments = append(student.Enrollments, enrollment) // Adds the new enrollment to the student's slice
}

// SetGrade sets a grade for a given enrollment.
func (e *Enrollment) SetGrade(grade float32) {
    e.Grade = grade // Assigns a grade to the enrollment
}

// GetCourses returns all courses a student is enrolled in.
func (s *Student) GetCourses() []Course {
    var courses []Course
    for _, enrollment := range s.Enrollments {
        courses = append(courses, *enrollment.Course) // Appends each course to the slice
    }
    return courses
}

func main() {
    // Create departments
    compSci := Department{Name: "Computer Science"} // Initializes a Department instance

    // Create courses
    course1 := Course{Code: "CS101", Name: "Introduction to Programming", Credits: 4, Department: &compSci} // Initializes a Course instance
    compSci.Courses = append(compSci.Courses, course1) // Adds the course to the department's course list

    // Create students
    student1 := Student{ID: "S1001", Name: "John Doe"} // Initializes a Student instance
    student2 := Student{ID: "S1002", Name: "Jane Smith"} // Initializes another Student instance

    // Enroll students in courses
    course1.Enroll(&student1) // Enrolls John Doe in CS101
    course1.Enroll(&student2) // Enrolls Jane Smith in CS101

    // Set grades for students
    course1.Enrollments[0].SetGrade(92.3) // Sets grade for John Doe
    course1.Enrollments[1].SetGrade(88.5) // Sets grade for Jane Smith

    // Print student courses and grades
    for _, enrollment := range student1.Enrollments {
        fmt.Printf("%s is enrolled in %s and received a grade of %.2f\n", 
            enrollment.Student.Name, enrollment.Course.Name, enrollment.Grade)
    }
}
