package entity

func ListAllTables(prefix string) []any {

	list := make([]any, 0)

	// list = append(list, new(Session))

	list = append(list, new(Permission))
	list = append(list, new(Role))
	list = append(list, new(User))

	if len(prefix) > 0 {
		theTableNamePrefix = prefix
	}
	return list
}

////////////////////////////////////////////////////////////////////////////////
// table names

var theTableNamePrefix = "rbac_wa_"

// func (Session) TableName() string {
// 	return theTableNamePrefix + "sessions"
// }

func (User) TableName() string {
	return theTableNamePrefix + "users"
}

func (Permission) TableName() string {
	return theTableNamePrefix + "permissions"
}

func (Role) TableName() string {
	return theTableNamePrefix + "roles"
}

////////////////////////////////////////////////////////////////////////////////
// EOF
