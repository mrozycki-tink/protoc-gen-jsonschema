package testdata

const MessageKind11 = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "$ref": "#/definitions/MessageKind11",
    "definitions": {
        "MessageKind11": {
            "properties": {
                "name": {
                    "type": "string"
                },
                "ones": {
                    "items": {
                        "$ref": "#/definitions/samples.MessageKind1"
                    },
                    "type": "array"
                },
                "kind2": {
                    "$ref": "#/definitions/samples.MessageKind2",
                    "additionalProperties": true
                },
                "kind3": {
                    "$ref": "#/definitions/samples.MessageKind3",
                    "additionalProperties": true
                },
                "kind4": {
                    "$ref": "#/definitions/samples.MessageKind4",
                    "additionalProperties": true
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.MessageKind1": {
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
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.MessageKind2": {
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
                "isa": {
                    "type": "boolean"
                },
                "hasa": {
                    "type": "boolean"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.MessageKind3": {
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
                "someProp": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        },
        "samples.MessageKind4": {
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
                "special": {
                    "type": "string"
                }
            },
            "additionalProperties": true,
            "type": "object"
        }
    }
}`
