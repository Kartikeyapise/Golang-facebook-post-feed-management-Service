package service

import (
	"github.com/kartikeya/sample_app/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func TestFinaAll(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	post := entity.Post{
		Id:    1,
		Title: "A",
		Text:  "B",
	}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.FindAll()

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.Equal(t, 1, result[0].Id)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
	assert.Nil(t, err)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectations
	post := entity.Post{
		Id:    1,
		Title: "A",
		Text:  "B",
	}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	//Mock Assertion: Behavioural
	mockRepo.AssertExpectations(t)

	//data Assertion
	assert.NotNil(t, result.Id)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The Post is Empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	testService := NewPostService(nil)
	samplePost := entity.Post{Id: 1, Title: "", Text: "A"}
	err := testService.Validate(&samplePost)
	assert.NotNil(t, err)
	assert.Equal(t, "The Post Title is Empty", err.Error())
}
