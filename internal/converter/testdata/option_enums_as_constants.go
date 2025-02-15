package testdata

const OptionEnumsAsConstants = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/OptionEnumsAsConstants",
    "definitions": {
        "OptionEnumsAsConstants": {
            "properties": {
                "importedEnum": {
                    "enum": [
                        "VALUE_0",
                        0,
                        "VALUE_1",
                        1,
                        "VALUE_2",
                        2,
                        "VALUE_3",
                        3
                    ],
                    "oneOf": [
                        {
                            "description": "Zero",
                            "const": "VALUE_0"
                        },
                        {
                            "description": "Zero",
                            "const": 0
                        },
                        {
                            "description": "One",
                            "const": "VALUE_1"
                        },
                        {
                            "description": "One",
                            "const": 1
                        },
                        {
                            "description": "Two",
                            "const": "VALUE_2"
                        },
                        {
                            "description": "Two",
                            "const": 2
                        },
                        {
                            "description": "Three",
                            "const": "VALUE_3"
                        },
                        {
                            "description": "Three",
                            "const": 3
                        }
                    ],
                    "description": "This is an enum"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const OptionEnumsAsConstantsFail = `{"importedEnum": "VALUE_4"}`

const OptionEnumsAsConstantsPass = `{"importedEnum": "VALUE_3"}`
