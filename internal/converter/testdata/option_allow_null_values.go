package testdata

const OptionAllowNullValues = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/OptionAllowNullValues",
    "definitions": {
        "OptionAllowNullValues": {
            "properties": {
                "name2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "string"
                        }
                    ]
                },
                "id2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                },
                "complete2": {
                    "oneOf": [
                        {
                            "type": "null"
                        },
                        {
                            "type": "boolean"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "oneOf": [
                {
                    "type": "null"
                },
                {
                    "type": "object"
                }
            ]
        }
    }
}`

const OptionAllowNullValuesFail = `{"name2": 12345}`

const OptionAllowNullValuesPass = `{"name2": null}`
