package dynamo_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/emman27/dynamo"
)

var _ = Describe("Dynamo", func() {
	type FakeModel struct {
		Model
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Describe("Model embedding", func() {
		It("Marshals to JSON without including Model.client and Model.TableName", func() {
			var (
				err error
				mp  map[string]string
				i   int
			)
			m := FakeModel{
				ID:   "12345",
				Name: "This is a fake model",
			}
			s, err := json.Marshal(m)
			Expect(err).To(BeNil())
			Expect(json.Unmarshal(s, &mp)).To(BeNil())
			Expect(mp["id"]).To(Equal("12345"))
			Expect(mp["name"]).To(Equal("This is a fake model"))
			for _ = range mp {
				i++
			}
			Expect(i).To(Equal(2))
		})
	})

	PDescribe("Creating a record", func() {
		Context("With a FakeModel of ID 27 and Name Alpha", func() {
			It("Calls PutItem", func() {})
		})
	})
})
