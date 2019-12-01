package handler

import (
	"testing"

	"github.com/impal-lms/lms-backend/services"
	"github.com/labstack/echo/v4"
)

func TestHandler_GetAllClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.GetAllClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.CreateClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetClassroomByID(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.GetClassroomByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.UpdateClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteClassroomByID(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.DeleteClassroomByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetAllStudentClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.GetAllStudentClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetAllStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_CreateStudentClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.CreateStudentClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.CreateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_GetStudentClassroomByID(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.GetStudentClassroomByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.GetStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_UpdateStudentClassroom(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.UpdateStudentClassroom(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.UpdateStudentClassroom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHandler_DeleteStudentClassroomByID(t *testing.T) {
	type fields struct {
		Services services.Services
	}
	type args struct {
		ctx echo.Context
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
			h := &Handler{
				Services: tt.fields.Services,
			}
			if err := h.DeleteStudentClassroomByID(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Handler.DeleteStudentClassroomByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
