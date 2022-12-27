// Code generated by ent, DO NOT EDIT.

package masterpassword

const (
	// Label holds the string label denoting the masterpassword type in the database.
	Label = "master_password"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// Table holds the table name of the masterpassword in the database.
	Table = "master_passwords"
)

// Columns holds all SQL columns for masterpassword fields.
var Columns = []string{
	FieldID,
	FieldPasswordHash,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
)
