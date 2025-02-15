package testdata

const OptionDisallowAdditionalProperties = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionDisallowAdditionalProperties",
    "definitions": {
        "OptionDisallowAdditionalProperties": {
            "properties": {
                "name2": {
                    "type": "string"
                },
                "timestamp2": {
                    "type": "string"
                },
                "id2": {
                    "type": "integer"
                },
                "rating2": {
                    "type": "number"
                },
                "complete2": {
                    "type": "boolean"
                }
            },
            "additionalProperties": false,
            "type": "object"
        }
    }
}`

const OptionDisallowAdditionalPropertiesFail = `{"something": 12345}`

const OptionDisallowAdditionalPropertiesPass = `{"rating2": 12345}`
