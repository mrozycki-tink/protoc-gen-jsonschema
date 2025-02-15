package testdata

const SecondEnum = `{
    "$schema": "http://json-schema.org/draft-04/schema#",
    "enum": [
        "VALUE_4",
        0,
        "VALUE_5",
        1,
        "VALUE_6",
        2,
        "VALUE_7",
        3
    ],
    "oneOf": [
        {
            "type": "string"
        },
        {
            "type": "integer"
        }
    ]
}`

const SecondEnumFail = `"VALUE_3"`

const SecondEnumPass = `"VALUE_7"`
