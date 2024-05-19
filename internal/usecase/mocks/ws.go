// Code generated by MockGen. DO NOT EDIT.
// Source: ws.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	domain "tourism/internal/domain"
	ws "tourism/internal/domain/ws"

	gomock "github.com/golang/mock/gomock"
)

// MockWsRepository is a mock of WsRepository interface.
type MockWsRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWsRepositoryMockRecorder
}

// MockWsRepositoryMockRecorder is the mock recorder for MockWsRepository.
type MockWsRepositoryMockRecorder struct {
	mock *MockWsRepository
}

// NewMockWsRepository creates a new mock instance.
func NewMockWsRepository(ctrl *gomock.Controller) *MockWsRepository {
	mock := &MockWsRepository{ctrl: ctrl}
	mock.recorder = &MockWsRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWsRepository) EXPECT() *MockWsRepositoryMockRecorder {
	return m.recorder
}

// AddClient mocks base method.
func (m *MockWsRepository) AddClient(ctx context.Context, client *ws.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddClient", ctx, client)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddClient indicates an expected call of AddClient.
func (mr *MockWsRepositoryMockRecorder) AddClient(ctx, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddClient", reflect.TypeOf((*MockWsRepository)(nil).AddClient), ctx, client)
}

// CreateRoom mocks base method.
func (m *MockWsRepository) CreateRoom(ctx context.Context, room *ws.Room) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRoom", ctx, room)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRoom indicates an expected call of CreateRoom.
func (mr *MockWsRepositoryMockRecorder) CreateRoom(ctx, room interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRoom", reflect.TypeOf((*MockWsRepository)(nil).CreateRoom), ctx, room)
}

// GetClientsByRoomID mocks base method.
func (m *MockWsRepository) GetClientsByRoomID(ctx context.Context, roomID string) ([]*ws.ClientResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientsByRoomID", ctx, roomID)
	ret0, _ := ret[0].([]*ws.ClientResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClientsByRoomID indicates an expected call of GetClientsByRoomID.
func (mr *MockWsRepositoryMockRecorder) GetClientsByRoomID(ctx, roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientsByRoomID", reflect.TypeOf((*MockWsRepository)(nil).GetClientsByRoomID), ctx, roomID)
}

// GetRoomsByClientID mocks base method.
func (m *MockWsRepository) GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoomsByClientID", ctx, clientID)
	ret0, _ := ret[0].([]*ws.RoomResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoomsByClientID indicates an expected call of GetRoomsByClientID.
func (mr *MockWsRepositoryMockRecorder) GetRoomsByClientID(ctx, clientID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoomsByClientID", reflect.TypeOf((*MockWsRepository)(nil).GetRoomsByClientID), ctx, clientID)
}

// GetUserByEmail mocks base method.
func (m *MockWsRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockWsRepositoryMockRecorder) GetUserByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockWsRepository)(nil).GetUserByEmail), ctx, email)
}

// GetUserByID mocks base method.
func (m *MockWsRepository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*domain.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockWsRepositoryMockRecorder) GetUserByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockWsRepository)(nil).GetUserByID), ctx, id)
}
