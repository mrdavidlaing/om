package api //not in api_tests because we are intentionally testing some functionality internal to the package

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("ResponsePropertyCollection", func() {
	responsePropertyCollectionValueJSON := `[
			{
				"guid": {
					"type": "uuid",
					"configurable": false,
					"credential": false,
					"value": "28bab1d3-4a4b-48d5-8dac-one",
					"optional": false
				},
				"some_property": {
					"type": "boolean",
					"configurable": true,
					"credential": false,
					"value": "true",
					"optional": false
				}
			},
			{
				"guid": {
					"type": "uuid",
					"configurable": false,
					"credential": false,
					"value": "28bab1d3-4a4b-48d5-8dac-two",
					"optional": false
				},
				"name": {
					"type": "string",
					"configurable": true,
					"credential": false,
					"value": "the_name",
					"optional": false
				},
				"yet_another_property": {
					"type": "boolean",
					"configurable": true,
					"credential": false,
					"value": false,
					"optional": false
				}
			},
			{
				"guid": {
					"type": "uuid",
					"configurable": false,
					"credential": false,
					"value": "28bab1d3-4a4b-48d5-8dac-three",
					"optional": false
				},
				"key": {
					"type": "string",
					"configurable": true,
					"credential": false,
					"value": "the_key",
					"optional": false
				},
				"some_additional_property": {
					"type": "boolean",
					"configurable": true,
					"credential": false,
					"value": false,
					"optional": false
				}
			},
			{
				"guid": {
					"type": "uuid",
					"configurable": false,
					"credential": false,
					"value": "28bab1d3-4a4b-48d5-8dac-four",
					"optional": false
				},
				"sqlServerName": {
					"type": "string",
					"configurable": true,
					"credential": false,
					"value": "the_sqlserver_name",
					"optional": false
				},
				"some_additional_property": {
					"type": "boolean",
					"configurable": true,
					"credential": false,
					"value": false,
					"optional": false
				}
			},
			{
				"guid": {
					"type": "uuid",
					"configurable": false,
					"credential": false,
					"value": "28bab1d3-4a4b-48d5-8dac-five",
					"optional": false
				},
				"name": {
					"type": "string",
					"configurable": true,
					"credential": false,
					"value": "the_name",
					"optional": false
				},
				"Filename": {
					"type": "string",
					"configurable": true,
					"credential": false,
					"value": "important_data.tgz",
					"optional": false
				}
			}
		]`
	var rawCollection interface{}
	err := yaml.Unmarshal([]byte(responsePropertyCollectionValueJSON), &rawCollection) //use yaml.Unmarshal to simulate Api.GetStagedProductProperties()
	if err != nil {
		panic(fmt.Errorf("Failed to parse json: %w", err))
	}
	collection, err := parseResponsePropertyCollection(rawCollection)
	if err != nil {
		panic(fmt.Errorf("Failed to parseResponsePropertyCollection from %v\n%w", rawCollection, err))
	}
	When("parseResponsePropertyCollection", func() {
		It("parses all the elements in the collection", func() {
			Expect(len(collection)).To(Equal(len(rawCollection.([]interface{}))))
		})
	})
	When("extracting field values", func() {
		It("correctly extracts guids", func() {
			Expect(collection[2].getFieldValue("guid")).To(Equal("28bab1d3-4a4b-48d5-8dac-three"))
		})
		It("correctly extracts strings", func() {
			Expect(collection[1].getFieldValue("name")).To(Equal("the_name"))
		})
	})
	When("finding the logical key field", func() {
		It("fails to find a logical key when there isn't one", func() {
			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeFalse())
			Expect(key).To(BeEmpty())
		})
		It("finds a 'name' logical key", func() {
			key, ok := collection[1].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("name"))
		})
		It("finds a 'key' logical key", func() {
			key, ok := collection[2].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("key"))
		})
		It("finds a logical key ending in 'name'", func() {
			key, ok := collection[3].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("sqlServerName"))
		})
		It("picks 'name' as the logical key when there is a 'name' field AND a field that ends in 'name' (eg: Filename)", func() {
			key, ok := collection[4].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("name"))
		})
	})
})
