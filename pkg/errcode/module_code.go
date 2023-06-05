package errcode

var (
	IDFailed = NewError(20001, "Primary key serialization error")

	UserFailed = NewError(30001, "User query error")

	DaoFailed = NewError(40001, "Dao query error")

	FirebaseSendFailed = NewError(50001, "Firebase send failed")

	MsgSaveFailed = NewError(60001, "Message save failed")

	MsgSendSaveFailed = NewError(70001, "Message send save failed")

	OrganFailed = NewError(80001, "Organ does not exist")
)
