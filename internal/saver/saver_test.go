package saver_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-question-api/internal/mocks"
	"github.com/ozoncp/ocp-question-api/internal/models"
	"github.com/ozoncp/ocp-question-api/internal/saver"
	"sync"

	"time"
)

var (
	mockCtrl    *gomock.Controller
	mockFlusher *mocks.MockFlusher
)

var _ = Describe("Saver", func() {
	var (
		ctx      context.Context
		s        saver.Saver
		entities []models.Question
	)

	BeforeEach(func() {
		ctx = context.Background()
		mockCtrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockCtrl)

		s = saver.NewSaver(3, mockFlusher, 200*time.Millisecond)
		entities = []models.Question{
			{Id: 1, UserId: 1, Text: "Question #1"},
			{Id: 2, UserId: 1, Text: "Question #2"},
			{Id: 3, UserId: 1, Text: "Question #3"},
			{Id: 4, UserId: 1, Text: "Question #4"},
			{Id: 5, UserId: 1, Text: "Question #5"},
			{Id: 6, UserId: 1, Text: "Question #6"},
			{Id: 7, UserId: 1, Text: "Question #7"},
			{Id: 8, UserId: 1, Text: "Question #8"},
			{Id: 9, UserId: 1, Text: "Question #9"},
			{Id: 10, UserId: 1, Text: "Question #10"},
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("save in repository", func() {
		BeforeEach(func() {
			s.Init()
		})

		AfterEach(func() {
			s.Close()
		})

		It("all saved", func() {
			mockFlusher.EXPECT().Flush(ctx, gomock.Any()).Return([]models.Question{}, nil).MinTimes(1)

			for _, entity := range entities {
				s.Save(entity)
			}
		})

		It("all saved from difference goroutine", func() {
			mockFlusher.EXPECT().Flush(ctx, gomock.Any()).Return([]models.Question{}, nil).MinTimes(0)

			wg := sync.WaitGroup{}

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func() {
					for _, entity := range entities {
						s.Save(entity)
					}

					defer wg.Done()
				}()
			}

			wg.Wait()
		})
	})
})
