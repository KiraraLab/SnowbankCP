package datachannel

import "fmt"

type TableSchema struct {
	TableName string
	Columns   []ColumnSchema
}

type ColumnSchema struct {
	ColumnName   string
	DataType     DataType
	IsNullable   bool
	DefaultValue Datum
}

func (ts TableSchema) Validate(data Record) error {
	for _, column := range ts.Columns {
		datum, exists := data[column.ColumnName]

		//Fallback to default value if datum is nil and default value is set
		if column.DefaultValue.Datum != nil && datum.Datum == nil {
			data[column.ColumnName] = column.DefaultValue
		}

		// Check if the column exists in the data and validate its type and nullability
		if !exists {
			if !column.IsNullable {
				return fmt.Errorf("SchemaCheckError: missing required column: %s", column.ColumnName)
			}
		} else if datum.Datum == nil && !column.IsNullable {
			return fmt.Errorf("SchemaCheckError: null value for non-nullable column: %s", column.ColumnName)
		} else if datum.DataType != column.DataType {
			return fmt.Errorf("SchemaCheckError: type mismatch for column %s: expected %s, got %s", column.ColumnName, column.DataType, datum.DataType)
		}
	}
	return nil
}
