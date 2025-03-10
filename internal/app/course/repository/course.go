package repository

import(
    "gorm.io/gorm"
)

type CourseMySQLItf interface {}

type CourseMySQL struct {
    db *gorm.DB
}

func NewCourseMySQL(db *gorm.DB) CourseMySQLItf {
    return &CourseMySQL{db}
}
