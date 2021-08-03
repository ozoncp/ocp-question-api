package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-question-api/internal/models"

	"github.com/ozoncp/ocp-question-api/internal/flusher"
	"github.com/ozoncp/ocp-question-api/internal/mocks"
)

var (
	mockCtrl *gomock.Controller
	mockRepo *mocks.MockRepo
)

var _ = Describe("Flusher", func() {
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("invalid arguments", func() {
		When("invalid chunk size is specified", func() {
			It("returns an error", func() {
				f, err := flusher.NewFlusher(0, mockRepo)

				Expect(f).Should(BeNil())
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Context("with correct arguments", func() {
		It("returns a Flusher instance", func() {
			f, err := flusher.NewFlusher(1, mockRepo)
			Expect(f).ShouldNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})

var _ = Describe("Flusher with repository", func() {
	var (
		entities []models.Question
		f        flusher.Flusher
		anyError error
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(mockCtrl)

		anyError = errors.New("error")
		entities = []models.Question{
			{Id: 1, UserId: 1, Text: "Question #1"},
			{Id: 2, UserId: 1, Text: "Question #2"},
			{Id: 3, UserId: 1, Text: "Question #3"},
		}
		f, _ = flusher.NewFlusher(2, mockRepo)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("with exception", func() {
		It("Exception on first chunk", func() {
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[0:2])).Return(anyError)

			result, _ := f.Flush(entities)
			Expect(result).Should(BeEquivalentTo(entities[0:]))
		})

		It("Exception on second chunk", func() {
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[0:2])).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[2:])).Return(anyError)

			result, _ := f.Flush(entities)
			Expect(result).Should(BeEquivalentTo(entities[2:]))
		})
	})

	Context("without exception", func() {
		It("Valid chunks", func() {
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[0:2])).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[2:])).Times(1)

			result, _ := f.Flush(entities)
			Expect(result).Should(BeEmpty())
		})

		It("Valid empty chunks", func() {
			mockRepo.EXPECT().AddEntities([]models.Question{}).Times(0)

			result, _ := f.Flush([]models.Question{})
			Expect(result).Should(BeEmpty())
		})

		It("Valid one incomplete chunk", func() {
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[0:1])).Times(1)

			result, _ := f.Flush(entities[0:1])
			Expect(result).Should(BeEmpty())
		})

		It("Valid one complete chunk", func() {
			mockRepo.EXPECT().AddEntities(gomock.Eq(entities[0:2])).Times(1)

			result, _ := f.Flush(entities[0:2])
			Expect(result).Should(BeEmpty())
		})

		It("Valid with chunk size is 1", func() {
			mockRepo.EXPECT().AddEntities(gomock.Any()).Times(3)

			f, _ = flusher.NewFlusher(1, mockRepo)
			result, _ := f.Flush(entities)
			Expect(result).Should(BeEmpty())
		})
	})
})
