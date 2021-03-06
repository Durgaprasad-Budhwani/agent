package sdk

import "github.com/pinpt/integration-sdk/work"

// Type aliases for our exported datamodel types to create a stable version
// which Integrations depend on instead of directly depending on a specific
// version of the integration-sdk directly

// WorkIssue is a issue
type WorkIssue = work.Issue

// WorkIssueComment is a issue comment
type WorkIssueComment = work.IssueComment

// WorkIssueStatus is a issue status
type WorkIssueStatus = work.IssueStatus

// WorkIssuePriority is a issue priority
type WorkIssuePriority = work.IssuePriority

// WorkIssueType is a issue type
type WorkIssueType = work.IssueType

// WorkProject is a project
type WorkProject = work.Project

// WorkUser is a user in the work system
type WorkUser = work.User

// WorkIssueAttachments is the work issue attachments
type WorkIssueAttachments = work.IssueAttachments

// WorkIssueChangeLog is the issue changelog
type WorkIssueChangeLog = work.IssueChangeLog

// WorkIssueChangeLogCreatedDate is the issue change log created date
type WorkIssueChangeLogCreatedDate = work.IssueChangeLogCreatedDate

// WorkIssueChangeLogField is the issue change log field enum
type WorkIssueChangeLogField = work.IssueChangeLogField

// WorkIssueChangeLogFieldAssigneeRefID is the enumeration value for assignee_ref_id
const WorkIssueChangeLogFieldAssigneeRefID WorkIssueChangeLogField = work.IssueChangeLogFieldAssigneeRefID

// WorkIssueChangeLogFieldDueDate is the enumeration value for due_date
const WorkIssueChangeLogFieldDueDate WorkIssueChangeLogField = work.IssueChangeLogFieldDueDate

// WorkIssueChangeLogFieldEpicID is the enumeration value for epic_id
const WorkIssueChangeLogFieldEpicID WorkIssueChangeLogField = work.IssueChangeLogFieldEpicID

// WorkIssueChangeLogFieldIdentifier is the enumeration value for identifier
const WorkIssueChangeLogFieldIdentifier WorkIssueChangeLogField = work.IssueChangeLogFieldIdentifier

// WorkIssueChangeLogFieldParentID is the enumeration value for parent_id
const WorkIssueChangeLogFieldParentID WorkIssueChangeLogField = work.IssueChangeLogFieldParentID

// WorkIssueChangeLogFieldPriority is the enumeration value for priority
const WorkIssueChangeLogFieldPriority WorkIssueChangeLogField = work.IssueChangeLogFieldPriority

// WorkIssueChangeLogFieldProjectID is the enumeration value for project_id
const WorkIssueChangeLogFieldProjectID WorkIssueChangeLogField = work.IssueChangeLogFieldProjectID

// WorkIssueChangeLogFieldReporterRefID is the enumeration value for reporter_ref_id
const WorkIssueChangeLogFieldReporterRefID WorkIssueChangeLogField = work.IssueChangeLogFieldReporterRefID

// WorkIssueChangeLogFieldResolution is the enumeration value for resolution
const WorkIssueChangeLogFieldResolution WorkIssueChangeLogField = work.IssueChangeLogFieldResolution

// WorkIssueChangeLogFieldSprintIds is the enumeration value for sprint_ids
const WorkIssueChangeLogFieldSprintIds WorkIssueChangeLogField = work.IssueChangeLogFieldSprintIds

// WorkIssueChangeLogFieldStatus is the enumeration value for status
const WorkIssueChangeLogFieldStatus WorkIssueChangeLogField = work.IssueChangeLogFieldStatus

// WorkIssueChangeLogFieldTags is the enumeration value for tags
const WorkIssueChangeLogFieldTags WorkIssueChangeLogField = work.IssueChangeLogFieldTags

// WorkIssueChangeLogFieldTitle is the enumeration value for title
const WorkIssueChangeLogFieldTitle WorkIssueChangeLogField = work.IssueChangeLogFieldTitle

// WorkIssueChangeLogFieldType is the enumeration value for type
const WorkIssueChangeLogFieldType WorkIssueChangeLogField = work.IssueChangeLogFieldType

// WorkIssueCreatedDate is the issue created date
type WorkIssueCreatedDate = work.IssueCreatedDate

// WorkIssueDueDate is the issue due date
type WorkIssueDueDate = work.IssueDueDate

// WorkIssueLinkedIssues is the issue linked issues
type WorkIssueLinkedIssues = work.IssueLinkedIssues

// WorkIssueLinkedIssuesLinkType is the linked isuse link type enum
type WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkType

// WorkIssueLinkedIssuesLinkTypeBlocks is the enumeration value for blocks
const WorkIssueLinkedIssuesLinkTypeBlocks WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkTypeBlocks

// WorkIssueLinkedIssuesLinkTypeClones is the enumeration value for clones
const WorkIssueLinkedIssuesLinkTypeClones WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkTypeClones

// WorkIssueLinkedIssuesLinkTypeDuplicates is the enumeration value for duplicates
const WorkIssueLinkedIssuesLinkTypeDuplicates WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkTypeDuplicates

// WorkIssueLinkedIssuesLinkTypeCauses is the enumeration value for causes
const WorkIssueLinkedIssuesLinkTypeCauses WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkTypeCauses

// WorkIssueLinkedIssuesLinkTypeRelates is the enumeration value for relates
const WorkIssueLinkedIssuesLinkTypeRelates WorkIssueLinkedIssuesLinkType = work.IssueLinkedIssuesLinkTypeRelates

// WorkIssuePlannedEndDate is the issue planned end date
type WorkIssuePlannedEndDate = work.IssuePlannedEndDate

// WorkIssuePlannedStartDate is the issue planned start date
type WorkIssuePlannedStartDate = work.IssuePlannedStartDate

// WorkIssueUpdatedDate is the issue updated date
type WorkIssueUpdatedDate = work.IssueUpdatedDate

// WorkIssueCommentCreatedDate is the issue comment created date
type WorkIssueCommentCreatedDate = work.IssueCommentCreatedDate

// WorkIssueCommentUpdatedDate is the issue comment updated date
type WorkIssueCommentUpdatedDate = work.IssueCommentUpdatedDate

// WorkIssueTypeMappedType is the issue type mapped type enum
type WorkIssueTypeMappedType = work.IssueTypeMappedType

// WorkIssueTypeMappedTypeUnknown is the enumeration value for unknown
const WorkIssueTypeMappedTypeUnknown WorkIssueTypeMappedType = work.IssueTypeMappedTypeUnknown

// WorkIssueTypeMappedTypeFeature is the enumeration value for feature
const WorkIssueTypeMappedTypeFeature WorkIssueTypeMappedType = work.IssueTypeMappedTypeFeature

// WorkIssueTypeMappedTypeBug is the enumeration value for bug
const WorkIssueTypeMappedTypeBug WorkIssueTypeMappedType = work.IssueTypeMappedTypeBug

// WorkIssueTypeMappedTypeEnhancement is the enumeration value for enhancement
const WorkIssueTypeMappedTypeEnhancement WorkIssueTypeMappedType = work.IssueTypeMappedTypeEnhancement

// WorkIssueTypeMappedTypeEpic is the enumeration value for epic
const WorkIssueTypeMappedTypeEpic WorkIssueTypeMappedType = work.IssueTypeMappedTypeEpic

// WorkIssueTypeMappedTypeStory is the enumeration value for story
const WorkIssueTypeMappedTypeStory WorkIssueTypeMappedType = work.IssueTypeMappedTypeStory

// WorkIssueTypeMappedTypeTask is the enumeration value for task
const WorkIssueTypeMappedTypeTask WorkIssueTypeMappedType = work.IssueTypeMappedTypeTask

// WorkIssueTypeMappedTypeSubtask is the enumeration value for subtask
const WorkIssueTypeMappedTypeSubtask WorkIssueTypeMappedType = work.IssueTypeMappedTypeSubtask

// WorkConfig is the work config model
type WorkConfig = work.Config

// WorkConfigStatuses is the work config status type
type WorkConfigStatuses = work.ConfigStatuses

// WorkProjectVisibility is the visibility of the project
type WorkProjectVisibility = work.ProjectVisibility

// WorkProjectVisibilityPrivate is the enumeration value for private
const WorkProjectVisibilityPrivate = work.ProjectVisibilityPrivate

// WorkProjectVisibilityPublic is the enumeration value for public
const WorkProjectVisibilityPublic = work.ProjectVisibilityPublic

// WorkProjectAffiliation is the project affiliation
type WorkProjectAffiliation = work.ProjectAffiliation

// WorkProjectAffiliationOrganization is the enumeration value for organization
const WorkProjectAffiliationOrganization = work.ProjectAffiliationOrganization

// WorkProjectAffiliationUser is the enumeration value for user
const WorkProjectAffiliationUser = work.ProjectAffiliationUser

// WorkProjectAffiliationThirdparty is the enumeration value for thirdparty
const WorkProjectAffiliationThirdparty = work.ProjectAffiliationThirdparty

// WorkProjectIssueResolutions is a work.ProjectIssueResolutions
type WorkProjectIssueResolutions = work.ProjectIssueResolutions

// WorkProjectIssueTypes is a work.ProjectIssueTypes
type WorkProjectIssueTypes = work.ProjectIssueTypes

// WorkIssueTransitions is the issue transitions
type WorkIssueTransitions = work.IssueTransitions

// WorkProjectCapability is the project capability
type WorkProjectCapability = work.ProjectCapability

// WorkProjectCapabilityIssueMutationFields represents the object structure for issue_mutation_fields
type WorkProjectCapabilityIssueMutationFields = work.ProjectCapabilityIssueMutationFields

// WorkProjectCapabilityIssueMutationFieldsType is the enumeration type for type
type WorkProjectCapabilityIssueMutationFieldsType = work.ProjectCapabilityIssueMutationFieldsType

// WorkProjectCapabilityIssueMutationFieldsValues represents the object structure for values
type WorkProjectCapabilityIssueMutationFieldsValues = work.ProjectCapabilityIssueMutationFieldsValues

// WorkProjectCapabilityIssueMutationFieldsTypeString is the enumeration value for string
const WorkProjectCapabilityIssueMutationFieldsTypeString = work.ProjectCapabilityIssueMutationFieldsTypeString

// WorkProjectCapabilityIssueMutationFieldsTypeNumber is the enumeration value for number
const WorkProjectCapabilityIssueMutationFieldsTypeNumber = work.ProjectCapabilityIssueMutationFieldsTypeNumber

// WorkProjectCapabilityIssueMutationFieldsTypeWorkIssueType is the enumeration value for work_issue_type
const WorkProjectCapabilityIssueMutationFieldsTypeWorkIssueType = work.ProjectCapabilityIssueMutationFieldsTypeWorkIssueType

// WorkProjectCapabilityIssueMutationFieldsTypeWorkIssuePriority is the enumeration value for work_issue_priority
const WorkProjectCapabilityIssueMutationFieldsTypeWorkIssuePriority = work.ProjectCapabilityIssueMutationFieldsTypeWorkIssuePriority

// WorkProjectCapabilityIssueMutationFieldsTypeStringArray is the enumeration value for string_array
const WorkProjectCapabilityIssueMutationFieldsTypeStringArray = work.ProjectCapabilityIssueMutationFieldsTypeStringArray

// WorkProjectCapabilityIssueMutationFieldsTypeUser is the enumeration value for user
const WorkProjectCapabilityIssueMutationFieldsTypeUser = work.ProjectCapabilityIssueMutationFieldsTypeUser

// WorkProjectCapabilityIssueMutationFieldsTypeAttachment is the enumeration value for attachment
const WorkProjectCapabilityIssueMutationFieldsTypeAttachment = work.ProjectCapabilityIssueMutationFieldsTypeAttachment

// WorkProjectCapabilityIssueMutationFieldsTypeTextbox is the enumeration value for textbox
const WorkProjectCapabilityIssueMutationFieldsTypeTextbox = work.ProjectCapabilityIssueMutationFieldsTypeTextbox

// WorkProjectCapabilityIssueMutationFieldsTypeEpic is the enumeration value for Epic
const WorkProjectCapabilityIssueMutationFieldsTypeEpic = work.ProjectCapabilityIssueMutationFieldsTypeEpic

// WorkProjectCapabilityIssueMutationFieldsTypeWorkSprint is the enumeration value for Sprint
const WorkProjectCapabilityIssueMutationFieldsTypeWorkSprint = work.ProjectCapabilityIssueMutationFieldsTypeWorkSprint

// WorkProjectCapabilityIssueMutationFieldsTypeWorkIssue is the enumeration value for work_issue
const WorkProjectCapabilityIssueMutationFieldsTypeWorkIssue = work.ProjectCapabilityIssueMutationFieldsTypeWorkIssue

// WorkProjectCapabilityIssueMutationFieldsTypeDate is common date object value
const WorkProjectCapabilityIssueMutationFieldsTypeDate = work.ProjectCapabilityIssueMutationFieldsTypeDate

// NewWorkProjectID will return the work project id
func NewWorkProjectID(customerID string, refID string, refType string) string {
	return work.NewProjectID(customerID, refID, refType)
}

// NewWorkIssueID will return the work issue id
func NewWorkIssueID(customerID string, refID string, refType string) string {
	return work.NewIssueID(customerID, refID, refType)
}

// NewWorkIssueCommentID will return the work issue comment id
func NewWorkIssueCommentID(customerID string, refID string, refType string) string {
	return work.NewIssueCommentID(customerID, refID, refType)
}

// NewWorkIssuePriorityID will return the work issue priority id
func NewWorkIssuePriorityID(customerID string, refType string, refID string) string {
	return work.NewIssuePriorityID(customerID, refType, refID)
}

// NewWorkIssueStatusID will return the work issue status id
func NewWorkIssueStatusID(customerID string, refType string, refID string) string {
	return work.NewIssueStatusID(customerID, refType, refID)
}

// NewWorkIssueTypeID will return the work issue type id
func NewWorkIssueTypeID(customerID string, refType string, refID string) string {
	return work.NewIssueTypeID(customerID, refType, refID)
}

// NewWorkUserID will return the work user id
func NewWorkUserID(customerID string, refID string, refType string) string {
	return work.NewUserID(customerID, refID, refType)
}

// NewWorkConfigID will return the work config id
func NewWorkConfigID(customerID string, refType string, integrationInstanceID string) string {
	return work.NewConfigID(customerID, refType, integrationInstanceID)
}
