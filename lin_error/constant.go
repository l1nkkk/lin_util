package lin_error

type ErrorCode int32

const (
	NoError       ErrorCode = 0
	InternalError ErrorCode = 1000
	DBError       ErrorCode = 1100
	NetError      ErrorCode = 1200
)

const (
	NoErrorStr       = "NoError"
	InternalErrorStr = "InternalError"
	DBErrorStr       = "DBError"
	NetErrorStr      = "NetError"

	Undefine = "Undefine"
)

func (e ErrorCode) String() string {
	switch e {
	case NoError:
		return NoErrorStr
	case InternalError:
		return InternalErrorStr
	case DBError:
		return DBErrorStr
	case NetError:
		return NetErrorStr
	default:
		return "undefine"
	}
}

func (e ErrorCode) Int32() int32 {
	return (int32)(e)
}
