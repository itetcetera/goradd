package db

// IndexDescription is used by SQL analysis to extract details about an Index in the database. We can use indexes
// to know how to get to sorted data easily.
type IndexDescription struct {
	// KeyName is the name of the index
	KeyName     string
	// IsUnique indicates whether the index is for a unique index
	IsUnique    bool
	// IsPrimary indicates whether this is the index for the primary key
	IsPrimary   bool
	// ColumnNames are the columns that are part of the index
	ColumnNames []string
}
