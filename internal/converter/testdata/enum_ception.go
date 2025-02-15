package testdata

const EnumCeption = `{
    "$schema": "http://json-schema.org/draft-06/schema#",
    "$ref": "#/definitions/Enumception",
    "definitions": {
        "Enumception": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "failureMode": {
                    "enum": [
                        "RECURSION_ERROR",
                        0,
                        "SYNTAX_ERROR",
                        1
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ],
                    "description": "FailureModes enum"
                },
                "payload": {
                    "$ref": "#/definitions/samples.PayloadMessage",
                    "additionalProperties": true
                },
                "payloads": {
                    "items": {
                        "$ref": "#/definitions/samples.PayloadMessage"
                    },
                    "type": "array"
                },
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
        },
        "samples.PayloadMessage": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "rating": {
                    "type": "number"
                },
                "complete": {
                    "type": "boolean"
                },
                "topology": {
                    "enum": [
                        "FLAT",
                        0,
                        "NESTED_OBJECT",
                        1,
                        "NESTED_MESSAGE",
                        2,
                        "ARRAY_OF_TYPE",
                        3,
                        "ARRAY_OF_OBJECT",
                        4,
                        "ARRAY_OF_MESSAGE",
                        5
                    ],
                    "oneOf": [
                        {
                            "type": "string"
                        },
                        {
                            "type": "integer"
                        }
                    ]
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`

const EnumCeptionFail = `{"payloads": [ {"topology": "MAP"} ]}`

const EnumCeptionPass = `{"payloads": [ {"topology": "ARRAY_OF_MESSAGE"} ]}`
