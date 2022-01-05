// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq          ComparisonOperator = "EQ"
	ComparisonOperatorNe          ComparisonOperator = "NE"
	ComparisonOperatorLe          ComparisonOperator = "LE"
	ComparisonOperatorLt          ComparisonOperator = "LT"
	ComparisonOperatorGe          ComparisonOperator = "GE"
	ComparisonOperatorGt          ComparisonOperator = "GT"
	ComparisonOperatorContains    ComparisonOperator = "CONTAINS"
	ComparisonOperatorNotContains ComparisonOperator = "NOT_CONTAINS"
	ComparisonOperatorBeginsWith  ComparisonOperator = "BEGINS_WITH"
	ComparisonOperatorIn          ComparisonOperator = "IN"
	ComparisonOperatorBetween     ComparisonOperator = "BETWEEN"
)

// Values returns all known values for ComparisonOperator. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ComparisonOperator) Values() []ComparisonOperator {
	return []ComparisonOperator{
		"EQ",
		"NE",
		"LE",
		"LT",
		"GE",
		"GT",
		"CONTAINS",
		"NOT_CONTAINS",
		"BEGINS_WITH",
		"IN",
		"BETWEEN",
	}
}

type DataLakeResourceType string

// Enum values for DataLakeResourceType
const (
	DataLakeResourceTypeCatalog             DataLakeResourceType = "CATALOG"
	DataLakeResourceTypeDatabase            DataLakeResourceType = "DATABASE"
	DataLakeResourceTypeTable               DataLakeResourceType = "TABLE"
	DataLakeResourceTypeDataLocation        DataLakeResourceType = "DATA_LOCATION"
	DataLakeResourceTypeLfTag               DataLakeResourceType = "LF_TAG"
	DataLakeResourceTypeLfTagPolicy         DataLakeResourceType = "LF_TAG_POLICY"
	DataLakeResourceTypeLfTagPolicyDatabase DataLakeResourceType = "LF_TAG_POLICY_DATABASE"
	DataLakeResourceTypeLfTagPolicyTable    DataLakeResourceType = "LF_TAG_POLICY_TABLE"
)

// Values returns all known values for DataLakeResourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DataLakeResourceType) Values() []DataLakeResourceType {
	return []DataLakeResourceType{
		"CATALOG",
		"DATABASE",
		"TABLE",
		"DATA_LOCATION",
		"LF_TAG",
		"LF_TAG_POLICY",
		"LF_TAG_POLICY_DATABASE",
		"LF_TAG_POLICY_TABLE",
	}
}

type FieldNameString string

// Enum values for FieldNameString
const (
	FieldNameStringResourceArn  FieldNameString = "RESOURCE_ARN"
	FieldNameStringRoleArn      FieldNameString = "ROLE_ARN"
	FieldNameStringLastModified FieldNameString = "LAST_MODIFIED"
)

// Values returns all known values for FieldNameString. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FieldNameString) Values() []FieldNameString {
	return []FieldNameString{
		"RESOURCE_ARN",
		"ROLE_ARN",
		"LAST_MODIFIED",
	}
}

type OptimizerType string

// Enum values for OptimizerType
const (
	OptimizerTypeCompaction        OptimizerType = "COMPACTION"
	OptimizerTypeGarbageCollection OptimizerType = "GARBAGE_COLLECTION"
	OptimizerTypeGeneric           OptimizerType = "ALL"
)

// Values returns all known values for OptimizerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (OptimizerType) Values() []OptimizerType {
	return []OptimizerType{
		"COMPACTION",
		"GARBAGE_COLLECTION",
		"ALL",
	}
}

type Permission string

// Enum values for Permission
const (
	PermissionAll                Permission = "ALL"
	PermissionSelect             Permission = "SELECT"
	PermissionAlter              Permission = "ALTER"
	PermissionDrop               Permission = "DROP"
	PermissionDelete             Permission = "DELETE"
	PermissionInsert             Permission = "INSERT"
	PermissionDescribe           Permission = "DESCRIBE"
	PermissionCreateDatabase     Permission = "CREATE_DATABASE"
	PermissionCreateTable        Permission = "CREATE_TABLE"
	PermissionDataLocationAccess Permission = "DATA_LOCATION_ACCESS"
	PermissionCreateTag          Permission = "CREATE_TAG"
	PermissionAlterTag           Permission = "ALTER_TAG"
	PermissionDeleteTag          Permission = "DELETE_TAG"
	PermissionDescribeTag        Permission = "DESCRIBE_TAG"
	PermissionAssociateTag       Permission = "ASSOCIATE_TAG"
)

// Values returns all known values for Permission. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Permission) Values() []Permission {
	return []Permission{
		"ALL",
		"SELECT",
		"ALTER",
		"DROP",
		"DELETE",
		"INSERT",
		"DESCRIBE",
		"CREATE_DATABASE",
		"CREATE_TABLE",
		"DATA_LOCATION_ACCESS",
		"CREATE_TAG",
		"ALTER_TAG",
		"DELETE_TAG",
		"DESCRIBE_TAG",
		"ASSOCIATE_TAG",
	}
}

type PermissionType string

// Enum values for PermissionType
const (
	PermissionTypeColumnPermission     PermissionType = "COLUMN_PERMISSION"
	PermissionTypeCellFilterPermission PermissionType = "CELL_FILTER_PERMISSION"
)

// Values returns all known values for PermissionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PermissionType) Values() []PermissionType {
	return []PermissionType{
		"COLUMN_PERMISSION",
		"CELL_FILTER_PERMISSION",
	}
}

type QueryStateString string

// Enum values for QueryStateString
const (
	QueryStateStringPending            QueryStateString = "PENDING"
	QueryStateStringWorkunitsAvailable QueryStateString = "WORKUNITS_AVAILABLE"
	QueryStateStringError              QueryStateString = "ERROR"
	QueryStateStringFinished           QueryStateString = "FINISHED"
	QueryStateStringExpired            QueryStateString = "EXPIRED"
)

// Values returns all known values for QueryStateString. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (QueryStateString) Values() []QueryStateString {
	return []QueryStateString{
		"PENDING",
		"WORKUNITS_AVAILABLE",
		"ERROR",
		"FINISHED",
		"EXPIRED",
	}
}

type ResourceShareType string

// Enum values for ResourceShareType
const (
	ResourceShareTypeForeign ResourceShareType = "FOREIGN"
	ResourceShareTypeAll     ResourceShareType = "ALL"
)

// Values returns all known values for ResourceShareType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ResourceShareType) Values() []ResourceShareType {
	return []ResourceShareType{
		"FOREIGN",
		"ALL",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeDatabase ResourceType = "DATABASE"
	ResourceTypeTable    ResourceType = "TABLE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"DATABASE",
		"TABLE",
	}
}

type TransactionStatus string

// Enum values for TransactionStatus
const (
	TransactionStatusActive           TransactionStatus = "ACTIVE"
	TransactionStatusCommitted        TransactionStatus = "COMMITTED"
	TransactionStatusAborted          TransactionStatus = "ABORTED"
	TransactionStatusCommitInProgress TransactionStatus = "COMMIT_IN_PROGRESS"
)

// Values returns all known values for TransactionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransactionStatus) Values() []TransactionStatus {
	return []TransactionStatus{
		"ACTIVE",
		"COMMITTED",
		"ABORTED",
		"COMMIT_IN_PROGRESS",
	}
}

type TransactionStatusFilter string

// Enum values for TransactionStatusFilter
const (
	TransactionStatusFilterAll       TransactionStatusFilter = "ALL"
	TransactionStatusFilterCompleted TransactionStatusFilter = "COMPLETED"
	TransactionStatusFilterActive    TransactionStatusFilter = "ACTIVE"
	TransactionStatusFilterCommitted TransactionStatusFilter = "COMMITTED"
	TransactionStatusFilterAborted   TransactionStatusFilter = "ABORTED"
)

// Values returns all known values for TransactionStatusFilter. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransactionStatusFilter) Values() []TransactionStatusFilter {
	return []TransactionStatusFilter{
		"ALL",
		"COMPLETED",
		"ACTIVE",
		"COMMITTED",
		"ABORTED",
	}
}

type TransactionType string

// Enum values for TransactionType
const (
	TransactionTypeReadAndWrite TransactionType = "READ_AND_WRITE"
	TransactionTypeReadOnly     TransactionType = "READ_ONLY"
)

// Values returns all known values for TransactionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TransactionType) Values() []TransactionType {
	return []TransactionType{
		"READ_AND_WRITE",
		"READ_ONLY",
	}
}
