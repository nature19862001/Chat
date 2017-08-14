package main

const (
	ERR_NONE = iota
	//common error
	ERR_NAME_NOT_VALID
	ERR_NAME_MAX_LEN

	//redis error
	ERR_REDIS = iota + 1000

	//user error
	ERR_USER = iota + 1100
	ERR_USER_NOT_EXIST

	//user error
	ERR_FRIEND = iota + 1200
	ERR_FRIEND_GROUP_NOT_EXIST
	ERR_FRIEND_GROUP_EXIST
	ERR_FRIEND_GROUP_MAX_COUNT
	ERR_FRIEND_GROUP_USER_NOT_EMPTY
	ERR_FRIEND_ADD_NEED_REQ
	ERR_FRIEND_IN_BLACKLIST
	ERR_FRIEND_MAX
	ERR_FRIEND_EXIST
	ERR_FRIEND_NOT_EXIST
)

const (
	VERIFY_TYPE_ALLOW_ALL = iota
	VERIFY_TYPE_NEED_AGGRE
	VERIFY_TYPE_REFUSE_ALL
	VERIFY_TYPE_ERR
)

const NAME_MAX_LEN = 64
const GROUP_MAX_COUNT = 128
