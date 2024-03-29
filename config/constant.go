package config

import "gopkg.in/go-playground/validator.v9"

const LOCAL_USERID_UINT = "user_id_uint"
const LOCAL_MANAGERNAME_STRING = "user_name_string"
const LOCAL_USERID_INT64 = "user_id_int64"
const LOCAL_TOKEN = "token"

const MESSAGE_FAIL = -1
const MESSAGE_SUCCESS = 0
const TOKEN_FAIL = -2

// 检测结构体
var Validate = validator.New()
