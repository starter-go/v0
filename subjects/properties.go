package subjects

type PropertyName string

const (
	PNameExample PropertyName = "example"

	PNameAvatar        PropertyName = "avatar"
	PNameNickName      PropertyName = "nickname"
	PNameUserName      PropertyName = "username"
	PNameRoles         PropertyName = "roles"
	PNameAuthenticated PropertyName = "authenticated"

	PNameUserID   PropertyName = "user_id"
	PNameUserUUID PropertyName = "user_uuid"

	PNameSessionID   PropertyName = "session_id"
	PNameSessionUUID PropertyName = "session_uuid"

	PNameNotBefore   PropertyName = "not_before"
	PNameNotAfter    PropertyName = "not_after"
	PNameCreatedAt   PropertyName = "created_at"
	PNameUpdatedAt   PropertyName = "updated_at"
	PNameDeletedAt   PropertyName = "deleted_at"
	PNameMaxTokenAge PropertyName = "max_token_age"

	PNameMobile   PropertyName = "mobile"
	PNameEmail    PropertyName = "email"
	PNameLanguage PropertyName = "language"
)
