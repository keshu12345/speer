// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	model "github.com/keshu12345/notes/model"
	mock "github.com/stretchr/testify/mock"
)

// NotesService is an autogenerated mock type for the NotesService type
type NotesService struct {
	mock.Mock
}

// GetAllNotes provides a mock function with given fields:
func (_m *NotesService) GetAllNotes() ([]model.Note, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllNotes")
	}

	var r0 []model.Note
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.Note, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.Note); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Note)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNoteByID provides a mock function with given fields: noteId
func (_m *NotesService) GetNoteByID(noteId uint64) (model.Note, error) {
	ret := _m.Called(noteId)

	if len(ret) == 0 {
		panic("no return value specified for GetNoteByID")
	}

	var r0 model.Note
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (model.Note, error)); ok {
		return rf(noteId)
	}
	if rf, ok := ret.Get(0).(func(uint64) model.Note); ok {
		r0 = rf(noteId)
	} else {
		r0 = ret.Get(0).(model.Note)
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(noteId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNotesService creates a new instance of NotesService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotesService(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotesService {
	mock := &NotesService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
