package api //not in api_tests because we are intentionally testing some functionality internal to the package

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("ResponsePropertyCollection", func() {
	unmarshalJSON := func(json string) interface{} {
		var rawCollection interface{}
		err := yaml.Unmarshal([]byte(json), &rawCollection) //use yaml.Unmarshal to simulate Api.GetStagedProductProperties()
		if err != nil {
			panic(fmt.Errorf("Failed to parse json: %w", err))
		}
		return rawCollection
	}
	When("parseResponsePropertyCollection", func() {
		It("parses all the elements in the collection", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())
			Expect(len(collection)).To(Equal(2))
		})
	})
	When("extracting field values", func() {
		It("correctly extracts guids", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())
			Expect(collection[0].getFieldValue("guid")).To(Equal("28bab1d3-4a4b-48d5-8dac-two"))
		})
		It("correctly extracts strings", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())
			Expect(collection[0].getFieldValue("name")).To(Equal("the_name"))
		})
	})
	When("finding the logical key field", func() {
		It("finds a 'name' logical key", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())

			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("name"))
		})
		It("fails to find a logical key when there isn't one", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())

			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeFalse())
			Expect(key).To(BeEmpty())
		})

		It("finds a 'key' logical key", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())

			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("key"))
		})
		It("finds a logical key ending in 'name'", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
				}
			]`))
			Expect(err).To(BeNil())

			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("sqlServerName"))
		})
		It("picks 'name' as the logical key when there is a 'name' field AND a field that ends in 'name' (eg: Filename)", func() {
			collection, err := parseResponsePropertyCollection(unmarshalJSON(`[
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
			]`))
			Expect(err).To(BeNil())

			key, ok := collection[0].findLogicalKeyField()
			Expect(ok).To(BeTrue())
			Expect(key).To(Equal("name"))
		})
	})
})
