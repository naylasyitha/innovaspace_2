package usecase

type CourseUsecaseItf interface {}

type CourseUsecase struct {}

func NewCourseUsecase() CourseUsecaseItf {
    return &CourseUsecase{}
}
