// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// KeysColumns holds the columns for the "keys" table.
	KeysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeString},
		{Name: "public_key", Type: field.TypeString},
		{Name: "private_key", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// KeysTable holds the schema information for the "keys" table.
	KeysTable = &schema.Table{
		Name:       "keys",
		Columns:    KeysColumns,
		PrimaryKey: []*schema.Column{KeysColumns[0]},
	}
	// PasswordItemsColumns holds the columns for the "password_items" table.
	PasswordItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "site_name", Type: field.TypeString, Nullable: true},
		{Name: "site_url", Type: field.TypeString, Nullable: true},
		{Name: "username", Type: field.TypeString, Nullable: true},
		{Name: "username_type", Type: field.TypeString, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "tags", Type: field.TypeJSON, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// PasswordItemsTable holds the schema information for the "password_items" table.
	PasswordItemsTable = &schema.Table{
		Name:       "password_items",
		Columns:    PasswordItemsColumns,
		PrimaryKey: []*schema.Column{PasswordItemsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		KeysTable,
		PasswordItemsTable,
	}
)

func init() {
}
