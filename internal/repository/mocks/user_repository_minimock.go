package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/kom1ssar/grpc-auth/internal/repository.UserRepository -o user_repository_minimock.go -n UserRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/kom1ssar/grpc-auth/internal/model"
)

// UserRepositoryMock implements repository.UserRepository
type UserRepositoryMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcCreate          func(ctx context.Context, user *model.User) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, user *model.User)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mUserRepositoryMockCreate

	funcGetByEmail          func(ctx context.Context, email string) (up1 *model.User, err error)
	inspectFuncGetByEmail   func(ctx context.Context, email string)
	afterGetByEmailCounter  uint64
	beforeGetByEmailCounter uint64
	GetByEmailMock          mUserRepositoryMockGetByEmail

	funcGetById          func(ctx context.Context, id int64) (up1 *model.User, err error)
	inspectFuncGetById   func(ctx context.Context, id int64)
	afterGetByIdCounter  uint64
	beforeGetByIdCounter uint64
	GetByIdMock          mUserRepositoryMockGetById

	funcGetByName          func(ctx context.Context, name string) (up1 *model.User, err error)
	inspectFuncGetByName   func(ctx context.Context, name string)
	afterGetByNameCounter  uint64
	beforeGetByNameCounter uint64
	GetByNameMock          mUserRepositoryMockGetByName
}

// NewUserRepositoryMock returns a mock for repository.UserRepository
func NewUserRepositoryMock(t minimock.Tester) *UserRepositoryMock {
	m := &UserRepositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*UserRepositoryMockCreateParams{}

	m.GetByEmailMock = mUserRepositoryMockGetByEmail{mock: m}
	m.GetByEmailMock.callArgs = []*UserRepositoryMockGetByEmailParams{}

	m.GetByIdMock = mUserRepositoryMockGetById{mock: m}
	m.GetByIdMock.callArgs = []*UserRepositoryMockGetByIdParams{}

	m.GetByNameMock = mUserRepositoryMockGetByName{mock: m}
	m.GetByNameMock.callArgs = []*UserRepositoryMockGetByNameParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mUserRepositoryMockCreate struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockCreateExpectation
	expectations       []*UserRepositoryMockCreateExpectation

	callArgs []*UserRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// UserRepositoryMockCreateExpectation specifies expectation struct of the UserRepository.Create
type UserRepositoryMockCreateExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockCreateParams
	results *UserRepositoryMockCreateResults
	Counter uint64
}

// UserRepositoryMockCreateParams contains parameters of the UserRepository.Create
type UserRepositoryMockCreateParams struct {
	ctx  context.Context
	user *model.User
}

// UserRepositoryMockCreateResults contains results of the UserRepository.Create
type UserRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Expect(ctx context.Context, user *model.User) *mUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &UserRepositoryMockCreateParams{ctx, user}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Inspect(f func(ctx context.Context, user *model.User)) *mUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by UserRepository.Create
func (mmCreate *mUserRepositoryMockCreate) Return(i1 int64, err error) *UserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &UserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &UserRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the UserRepository.Create method
func (mmCreate *mUserRepositoryMockCreate) Set(f func(ctx context.Context, user *model.User) (i1 int64, err error)) *UserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the UserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the UserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the UserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mUserRepositoryMockCreate) When(ctx context.Context, user *model.User) *UserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("UserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &UserRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &UserRepositoryMockCreateParams{ctx, user},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.Create return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockCreateExpectation) Then(i1 int64, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.UserRepository
func (mmCreate *UserRepositoryMock) Create(ctx context.Context, user *model.User) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, user)
	}

	mm_params := UserRepositoryMockCreateParams{ctx, user}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, &mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := UserRepositoryMockCreateParams{ctx, user}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("UserRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the UserRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, user)
	}
	mmCreate.t.Fatalf("Unexpected call to UserRepositoryMock.Create. %v %v", ctx, user)
	return
}

// CreateAfterCounter returns a count of finished UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of UserRepositoryMock.Create invocations
func (mmCreate *UserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mUserRepositoryMockCreate) Calls() []*UserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*UserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.Create")
	}
}

type mUserRepositoryMockGetByEmail struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetByEmailExpectation
	expectations       []*UserRepositoryMockGetByEmailExpectation

	callArgs []*UserRepositoryMockGetByEmailParams
	mutex    sync.RWMutex
}

// UserRepositoryMockGetByEmailExpectation specifies expectation struct of the UserRepository.GetByEmail
type UserRepositoryMockGetByEmailExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockGetByEmailParams
	results *UserRepositoryMockGetByEmailResults
	Counter uint64
}

// UserRepositoryMockGetByEmailParams contains parameters of the UserRepository.GetByEmail
type UserRepositoryMockGetByEmailParams struct {
	ctx   context.Context
	email string
}

// UserRepositoryMockGetByEmailResults contains results of the UserRepository.GetByEmail
type UserRepositoryMockGetByEmailResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for UserRepository.GetByEmail
func (mmGetByEmail *mUserRepositoryMockGetByEmail) Expect(ctx context.Context, email string) *mUserRepositoryMockGetByEmail {
	if mmGetByEmail.mock.funcGetByEmail != nil {
		mmGetByEmail.mock.t.Fatalf("UserRepositoryMock.GetByEmail mock is already set by Set")
	}

	if mmGetByEmail.defaultExpectation == nil {
		mmGetByEmail.defaultExpectation = &UserRepositoryMockGetByEmailExpectation{}
	}

	mmGetByEmail.defaultExpectation.params = &UserRepositoryMockGetByEmailParams{ctx, email}
	for _, e := range mmGetByEmail.expectations {
		if minimock.Equal(e.params, mmGetByEmail.defaultExpectation.params) {
			mmGetByEmail.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetByEmail.defaultExpectation.params)
		}
	}

	return mmGetByEmail
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.GetByEmail
func (mmGetByEmail *mUserRepositoryMockGetByEmail) Inspect(f func(ctx context.Context, email string)) *mUserRepositoryMockGetByEmail {
	if mmGetByEmail.mock.inspectFuncGetByEmail != nil {
		mmGetByEmail.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.GetByEmail")
	}

	mmGetByEmail.mock.inspectFuncGetByEmail = f

	return mmGetByEmail
}

// Return sets up results that will be returned by UserRepository.GetByEmail
func (mmGetByEmail *mUserRepositoryMockGetByEmail) Return(up1 *model.User, err error) *UserRepositoryMock {
	if mmGetByEmail.mock.funcGetByEmail != nil {
		mmGetByEmail.mock.t.Fatalf("UserRepositoryMock.GetByEmail mock is already set by Set")
	}

	if mmGetByEmail.defaultExpectation == nil {
		mmGetByEmail.defaultExpectation = &UserRepositoryMockGetByEmailExpectation{mock: mmGetByEmail.mock}
	}
	mmGetByEmail.defaultExpectation.results = &UserRepositoryMockGetByEmailResults{up1, err}
	return mmGetByEmail.mock
}

// Set uses given function f to mock the UserRepository.GetByEmail method
func (mmGetByEmail *mUserRepositoryMockGetByEmail) Set(f func(ctx context.Context, email string) (up1 *model.User, err error)) *UserRepositoryMock {
	if mmGetByEmail.defaultExpectation != nil {
		mmGetByEmail.mock.t.Fatalf("Default expectation is already set for the UserRepository.GetByEmail method")
	}

	if len(mmGetByEmail.expectations) > 0 {
		mmGetByEmail.mock.t.Fatalf("Some expectations are already set for the UserRepository.GetByEmail method")
	}

	mmGetByEmail.mock.funcGetByEmail = f
	return mmGetByEmail.mock
}

// When sets expectation for the UserRepository.GetByEmail which will trigger the result defined by the following
// Then helper
func (mmGetByEmail *mUserRepositoryMockGetByEmail) When(ctx context.Context, email string) *UserRepositoryMockGetByEmailExpectation {
	if mmGetByEmail.mock.funcGetByEmail != nil {
		mmGetByEmail.mock.t.Fatalf("UserRepositoryMock.GetByEmail mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetByEmailExpectation{
		mock:   mmGetByEmail.mock,
		params: &UserRepositoryMockGetByEmailParams{ctx, email},
	}
	mmGetByEmail.expectations = append(mmGetByEmail.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.GetByEmail return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetByEmailExpectation) Then(up1 *model.User, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetByEmailResults{up1, err}
	return e.mock
}

// GetByEmail implements repository.UserRepository
func (mmGetByEmail *UserRepositoryMock) GetByEmail(ctx context.Context, email string) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGetByEmail.beforeGetByEmailCounter, 1)
	defer mm_atomic.AddUint64(&mmGetByEmail.afterGetByEmailCounter, 1)

	if mmGetByEmail.inspectFuncGetByEmail != nil {
		mmGetByEmail.inspectFuncGetByEmail(ctx, email)
	}

	mm_params := UserRepositoryMockGetByEmailParams{ctx, email}

	// Record call args
	mmGetByEmail.GetByEmailMock.mutex.Lock()
	mmGetByEmail.GetByEmailMock.callArgs = append(mmGetByEmail.GetByEmailMock.callArgs, &mm_params)
	mmGetByEmail.GetByEmailMock.mutex.Unlock()

	for _, e := range mmGetByEmail.GetByEmailMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGetByEmail.GetByEmailMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetByEmail.GetByEmailMock.defaultExpectation.Counter, 1)
		mm_want := mmGetByEmail.GetByEmailMock.defaultExpectation.params
		mm_got := UserRepositoryMockGetByEmailParams{ctx, email}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetByEmail.t.Errorf("UserRepositoryMock.GetByEmail got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetByEmail.GetByEmailMock.defaultExpectation.results
		if mm_results == nil {
			mmGetByEmail.t.Fatal("No results are set for the UserRepositoryMock.GetByEmail")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGetByEmail.funcGetByEmail != nil {
		return mmGetByEmail.funcGetByEmail(ctx, email)
	}
	mmGetByEmail.t.Fatalf("Unexpected call to UserRepositoryMock.GetByEmail. %v %v", ctx, email)
	return
}

// GetByEmailAfterCounter returns a count of finished UserRepositoryMock.GetByEmail invocations
func (mmGetByEmail *UserRepositoryMock) GetByEmailAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByEmail.afterGetByEmailCounter)
}

// GetByEmailBeforeCounter returns a count of UserRepositoryMock.GetByEmail invocations
func (mmGetByEmail *UserRepositoryMock) GetByEmailBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByEmail.beforeGetByEmailCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.GetByEmail.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetByEmail *mUserRepositoryMockGetByEmail) Calls() []*UserRepositoryMockGetByEmailParams {
	mmGetByEmail.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetByEmailParams, len(mmGetByEmail.callArgs))
	copy(argCopy, mmGetByEmail.callArgs)

	mmGetByEmail.mutex.RUnlock()

	return argCopy
}

// MinimockGetByEmailDone returns true if the count of the GetByEmail invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetByEmailDone() bool {
	for _, e := range m.GetByEmailMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByEmailMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByEmailCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByEmail != nil && mm_atomic.LoadUint64(&m.afterGetByEmailCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetByEmailInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetByEmailInspect() {
	for _, e := range m.GetByEmailMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.GetByEmail with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByEmailMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByEmailCounter) < 1 {
		if m.GetByEmailMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.GetByEmail")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.GetByEmail with params: %#v", *m.GetByEmailMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByEmail != nil && mm_atomic.LoadUint64(&m.afterGetByEmailCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.GetByEmail")
	}
}

type mUserRepositoryMockGetById struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetByIdExpectation
	expectations       []*UserRepositoryMockGetByIdExpectation

	callArgs []*UserRepositoryMockGetByIdParams
	mutex    sync.RWMutex
}

// UserRepositoryMockGetByIdExpectation specifies expectation struct of the UserRepository.GetById
type UserRepositoryMockGetByIdExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockGetByIdParams
	results *UserRepositoryMockGetByIdResults
	Counter uint64
}

// UserRepositoryMockGetByIdParams contains parameters of the UserRepository.GetById
type UserRepositoryMockGetByIdParams struct {
	ctx context.Context
	id  int64
}

// UserRepositoryMockGetByIdResults contains results of the UserRepository.GetById
type UserRepositoryMockGetByIdResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for UserRepository.GetById
func (mmGetById *mUserRepositoryMockGetById) Expect(ctx context.Context, id int64) *mUserRepositoryMockGetById {
	if mmGetById.mock.funcGetById != nil {
		mmGetById.mock.t.Fatalf("UserRepositoryMock.GetById mock is already set by Set")
	}

	if mmGetById.defaultExpectation == nil {
		mmGetById.defaultExpectation = &UserRepositoryMockGetByIdExpectation{}
	}

	mmGetById.defaultExpectation.params = &UserRepositoryMockGetByIdParams{ctx, id}
	for _, e := range mmGetById.expectations {
		if minimock.Equal(e.params, mmGetById.defaultExpectation.params) {
			mmGetById.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetById.defaultExpectation.params)
		}
	}

	return mmGetById
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.GetById
func (mmGetById *mUserRepositoryMockGetById) Inspect(f func(ctx context.Context, id int64)) *mUserRepositoryMockGetById {
	if mmGetById.mock.inspectFuncGetById != nil {
		mmGetById.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.GetById")
	}

	mmGetById.mock.inspectFuncGetById = f

	return mmGetById
}

// Return sets up results that will be returned by UserRepository.GetById
func (mmGetById *mUserRepositoryMockGetById) Return(up1 *model.User, err error) *UserRepositoryMock {
	if mmGetById.mock.funcGetById != nil {
		mmGetById.mock.t.Fatalf("UserRepositoryMock.GetById mock is already set by Set")
	}

	if mmGetById.defaultExpectation == nil {
		mmGetById.defaultExpectation = &UserRepositoryMockGetByIdExpectation{mock: mmGetById.mock}
	}
	mmGetById.defaultExpectation.results = &UserRepositoryMockGetByIdResults{up1, err}
	return mmGetById.mock
}

// Set uses given function f to mock the UserRepository.GetById method
func (mmGetById *mUserRepositoryMockGetById) Set(f func(ctx context.Context, id int64) (up1 *model.User, err error)) *UserRepositoryMock {
	if mmGetById.defaultExpectation != nil {
		mmGetById.mock.t.Fatalf("Default expectation is already set for the UserRepository.GetById method")
	}

	if len(mmGetById.expectations) > 0 {
		mmGetById.mock.t.Fatalf("Some expectations are already set for the UserRepository.GetById method")
	}

	mmGetById.mock.funcGetById = f
	return mmGetById.mock
}

// When sets expectation for the UserRepository.GetById which will trigger the result defined by the following
// Then helper
func (mmGetById *mUserRepositoryMockGetById) When(ctx context.Context, id int64) *UserRepositoryMockGetByIdExpectation {
	if mmGetById.mock.funcGetById != nil {
		mmGetById.mock.t.Fatalf("UserRepositoryMock.GetById mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetByIdExpectation{
		mock:   mmGetById.mock,
		params: &UserRepositoryMockGetByIdParams{ctx, id},
	}
	mmGetById.expectations = append(mmGetById.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.GetById return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetByIdExpectation) Then(up1 *model.User, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetByIdResults{up1, err}
	return e.mock
}

// GetById implements repository.UserRepository
func (mmGetById *UserRepositoryMock) GetById(ctx context.Context, id int64) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGetById.beforeGetByIdCounter, 1)
	defer mm_atomic.AddUint64(&mmGetById.afterGetByIdCounter, 1)

	if mmGetById.inspectFuncGetById != nil {
		mmGetById.inspectFuncGetById(ctx, id)
	}

	mm_params := UserRepositoryMockGetByIdParams{ctx, id}

	// Record call args
	mmGetById.GetByIdMock.mutex.Lock()
	mmGetById.GetByIdMock.callArgs = append(mmGetById.GetByIdMock.callArgs, &mm_params)
	mmGetById.GetByIdMock.mutex.Unlock()

	for _, e := range mmGetById.GetByIdMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGetById.GetByIdMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetById.GetByIdMock.defaultExpectation.Counter, 1)
		mm_want := mmGetById.GetByIdMock.defaultExpectation.params
		mm_got := UserRepositoryMockGetByIdParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetById.t.Errorf("UserRepositoryMock.GetById got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetById.GetByIdMock.defaultExpectation.results
		if mm_results == nil {
			mmGetById.t.Fatal("No results are set for the UserRepositoryMock.GetById")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGetById.funcGetById != nil {
		return mmGetById.funcGetById(ctx, id)
	}
	mmGetById.t.Fatalf("Unexpected call to UserRepositoryMock.GetById. %v %v", ctx, id)
	return
}

// GetByIdAfterCounter returns a count of finished UserRepositoryMock.GetById invocations
func (mmGetById *UserRepositoryMock) GetByIdAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetById.afterGetByIdCounter)
}

// GetByIdBeforeCounter returns a count of UserRepositoryMock.GetById invocations
func (mmGetById *UserRepositoryMock) GetByIdBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetById.beforeGetByIdCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.GetById.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetById *mUserRepositoryMockGetById) Calls() []*UserRepositoryMockGetByIdParams {
	mmGetById.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetByIdParams, len(mmGetById.callArgs))
	copy(argCopy, mmGetById.callArgs)

	mmGetById.mutex.RUnlock()

	return argCopy
}

// MinimockGetByIdDone returns true if the count of the GetById invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetByIdDone() bool {
	for _, e := range m.GetByIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByIdCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetById != nil && mm_atomic.LoadUint64(&m.afterGetByIdCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetByIdInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetByIdInspect() {
	for _, e := range m.GetByIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.GetById with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByIdCounter) < 1 {
		if m.GetByIdMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.GetById")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.GetById with params: %#v", *m.GetByIdMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetById != nil && mm_atomic.LoadUint64(&m.afterGetByIdCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.GetById")
	}
}

type mUserRepositoryMockGetByName struct {
	mock               *UserRepositoryMock
	defaultExpectation *UserRepositoryMockGetByNameExpectation
	expectations       []*UserRepositoryMockGetByNameExpectation

	callArgs []*UserRepositoryMockGetByNameParams
	mutex    sync.RWMutex
}

// UserRepositoryMockGetByNameExpectation specifies expectation struct of the UserRepository.GetByName
type UserRepositoryMockGetByNameExpectation struct {
	mock    *UserRepositoryMock
	params  *UserRepositoryMockGetByNameParams
	results *UserRepositoryMockGetByNameResults
	Counter uint64
}

// UserRepositoryMockGetByNameParams contains parameters of the UserRepository.GetByName
type UserRepositoryMockGetByNameParams struct {
	ctx  context.Context
	name string
}

// UserRepositoryMockGetByNameResults contains results of the UserRepository.GetByName
type UserRepositoryMockGetByNameResults struct {
	up1 *model.User
	err error
}

// Expect sets up expected params for UserRepository.GetByName
func (mmGetByName *mUserRepositoryMockGetByName) Expect(ctx context.Context, name string) *mUserRepositoryMockGetByName {
	if mmGetByName.mock.funcGetByName != nil {
		mmGetByName.mock.t.Fatalf("UserRepositoryMock.GetByName mock is already set by Set")
	}

	if mmGetByName.defaultExpectation == nil {
		mmGetByName.defaultExpectation = &UserRepositoryMockGetByNameExpectation{}
	}

	mmGetByName.defaultExpectation.params = &UserRepositoryMockGetByNameParams{ctx, name}
	for _, e := range mmGetByName.expectations {
		if minimock.Equal(e.params, mmGetByName.defaultExpectation.params) {
			mmGetByName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetByName.defaultExpectation.params)
		}
	}

	return mmGetByName
}

// Inspect accepts an inspector function that has same arguments as the UserRepository.GetByName
func (mmGetByName *mUserRepositoryMockGetByName) Inspect(f func(ctx context.Context, name string)) *mUserRepositoryMockGetByName {
	if mmGetByName.mock.inspectFuncGetByName != nil {
		mmGetByName.mock.t.Fatalf("Inspect function is already set for UserRepositoryMock.GetByName")
	}

	mmGetByName.mock.inspectFuncGetByName = f

	return mmGetByName
}

// Return sets up results that will be returned by UserRepository.GetByName
func (mmGetByName *mUserRepositoryMockGetByName) Return(up1 *model.User, err error) *UserRepositoryMock {
	if mmGetByName.mock.funcGetByName != nil {
		mmGetByName.mock.t.Fatalf("UserRepositoryMock.GetByName mock is already set by Set")
	}

	if mmGetByName.defaultExpectation == nil {
		mmGetByName.defaultExpectation = &UserRepositoryMockGetByNameExpectation{mock: mmGetByName.mock}
	}
	mmGetByName.defaultExpectation.results = &UserRepositoryMockGetByNameResults{up1, err}
	return mmGetByName.mock
}

// Set uses given function f to mock the UserRepository.GetByName method
func (mmGetByName *mUserRepositoryMockGetByName) Set(f func(ctx context.Context, name string) (up1 *model.User, err error)) *UserRepositoryMock {
	if mmGetByName.defaultExpectation != nil {
		mmGetByName.mock.t.Fatalf("Default expectation is already set for the UserRepository.GetByName method")
	}

	if len(mmGetByName.expectations) > 0 {
		mmGetByName.mock.t.Fatalf("Some expectations are already set for the UserRepository.GetByName method")
	}

	mmGetByName.mock.funcGetByName = f
	return mmGetByName.mock
}

// When sets expectation for the UserRepository.GetByName which will trigger the result defined by the following
// Then helper
func (mmGetByName *mUserRepositoryMockGetByName) When(ctx context.Context, name string) *UserRepositoryMockGetByNameExpectation {
	if mmGetByName.mock.funcGetByName != nil {
		mmGetByName.mock.t.Fatalf("UserRepositoryMock.GetByName mock is already set by Set")
	}

	expectation := &UserRepositoryMockGetByNameExpectation{
		mock:   mmGetByName.mock,
		params: &UserRepositoryMockGetByNameParams{ctx, name},
	}
	mmGetByName.expectations = append(mmGetByName.expectations, expectation)
	return expectation
}

// Then sets up UserRepository.GetByName return parameters for the expectation previously defined by the When method
func (e *UserRepositoryMockGetByNameExpectation) Then(up1 *model.User, err error) *UserRepositoryMock {
	e.results = &UserRepositoryMockGetByNameResults{up1, err}
	return e.mock
}

// GetByName implements repository.UserRepository
func (mmGetByName *UserRepositoryMock) GetByName(ctx context.Context, name string) (up1 *model.User, err error) {
	mm_atomic.AddUint64(&mmGetByName.beforeGetByNameCounter, 1)
	defer mm_atomic.AddUint64(&mmGetByName.afterGetByNameCounter, 1)

	if mmGetByName.inspectFuncGetByName != nil {
		mmGetByName.inspectFuncGetByName(ctx, name)
	}

	mm_params := UserRepositoryMockGetByNameParams{ctx, name}

	// Record call args
	mmGetByName.GetByNameMock.mutex.Lock()
	mmGetByName.GetByNameMock.callArgs = append(mmGetByName.GetByNameMock.callArgs, &mm_params)
	mmGetByName.GetByNameMock.mutex.Unlock()

	for _, e := range mmGetByName.GetByNameMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.up1, e.results.err
		}
	}

	if mmGetByName.GetByNameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetByName.GetByNameMock.defaultExpectation.Counter, 1)
		mm_want := mmGetByName.GetByNameMock.defaultExpectation.params
		mm_got := UserRepositoryMockGetByNameParams{ctx, name}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetByName.t.Errorf("UserRepositoryMock.GetByName got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetByName.GetByNameMock.defaultExpectation.results
		if mm_results == nil {
			mmGetByName.t.Fatal("No results are set for the UserRepositoryMock.GetByName")
		}
		return (*mm_results).up1, (*mm_results).err
	}
	if mmGetByName.funcGetByName != nil {
		return mmGetByName.funcGetByName(ctx, name)
	}
	mmGetByName.t.Fatalf("Unexpected call to UserRepositoryMock.GetByName. %v %v", ctx, name)
	return
}

// GetByNameAfterCounter returns a count of finished UserRepositoryMock.GetByName invocations
func (mmGetByName *UserRepositoryMock) GetByNameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByName.afterGetByNameCounter)
}

// GetByNameBeforeCounter returns a count of UserRepositoryMock.GetByName invocations
func (mmGetByName *UserRepositoryMock) GetByNameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetByName.beforeGetByNameCounter)
}

// Calls returns a list of arguments used in each call to UserRepositoryMock.GetByName.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetByName *mUserRepositoryMockGetByName) Calls() []*UserRepositoryMockGetByNameParams {
	mmGetByName.mutex.RLock()

	argCopy := make([]*UserRepositoryMockGetByNameParams, len(mmGetByName.callArgs))
	copy(argCopy, mmGetByName.callArgs)

	mmGetByName.mutex.RUnlock()

	return argCopy
}

// MinimockGetByNameDone returns true if the count of the GetByName invocations corresponds
// the number of defined expectations
func (m *UserRepositoryMock) MinimockGetByNameDone() bool {
	for _, e := range m.GetByNameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByNameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByName != nil && mm_atomic.LoadUint64(&m.afterGetByNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetByNameInspect logs each unmet expectation
func (m *UserRepositoryMock) MinimockGetByNameInspect() {
	for _, e := range m.GetByNameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserRepositoryMock.GetByName with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetByNameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetByNameCounter) < 1 {
		if m.GetByNameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserRepositoryMock.GetByName")
		} else {
			m.t.Errorf("Expected call to UserRepositoryMock.GetByName with params: %#v", *m.GetByNameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetByName != nil && mm_atomic.LoadUint64(&m.afterGetByNameCounter) < 1 {
		m.t.Error("Expected call to UserRepositoryMock.GetByName")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserRepositoryMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockCreateInspect()

			m.MinimockGetByEmailInspect()

			m.MinimockGetByIdInspect()

			m.MinimockGetByNameInspect()
			m.t.FailNow()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *UserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetByEmailDone() &&
		m.MinimockGetByIdDone() &&
		m.MinimockGetByNameDone()
}
