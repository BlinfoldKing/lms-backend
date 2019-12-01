package repository

import (
	"reflect"
	"testing"

	"github.com/impal-lms/lms-backend/domain"
	"github.com/jinzhu/gorm"
)

func TestGORM_GetAllClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllClassroom()
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		classroom domain.Classroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateClassroom(tt.args.classroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetClassroomByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Classroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		classroom domain.Classroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateClassroom(tt.args.classroom); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteClassroomByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteClassroomByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_GetAllStudentClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		studentID   int64
		classroomID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetAllStudentClassroom(tt.args.studentID, tt.args.classroomID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetAllStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetAllStudentClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_CreateStudentClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		studentClassroom domain.StudentClassroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.CreateStudentClassroom(tt.args.studentClassroom)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.CreateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.CreateStudentClassroom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_GetStudentClassroomByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.StudentClassroom
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			got, err := g.GetStudentClassroomByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GORM.GetStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GORM.GetStudentClassroomByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGORM_UpdateStudentClassroom(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		studentClassroom domain.StudentClassroom
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.UpdateStudentClassroom(tt.args.studentClassroom); (err != nil) != tt.wantErr {
				t.Errorf("GORM.UpdateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGORM_DeleteStudentClassroomByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GORM{
				DB: tt.fields.DB,
			}
			if err := g.DeleteStudentClassroomByID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("GORM.DeleteStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
