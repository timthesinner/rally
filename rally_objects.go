//Package rally by TimTheSinner (this file was auto generated)
package rally

/**
 * Copyright (c) 2016 TimTheSinner All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//Base minimal Rally object type
type Base struct {
	APIMajor      string `json:"_rallyAPIMajor,omitempty"`
	APIMinor      string `json:"_rallyAPIMinor,omitempty"`
	URI           string `json:"_ref"`
	RefObjectName string `json:"_refObjectName,omitempty"`
	RefObjectUUID string `json:"_refObjectUUID,omitempty"`
	RefType       string `json:"_type,omitempty"`
}

//QueryResult standard query result
type QueryResult struct {
	Errors           []interface{}
	PageSize         float64
	Results          []map[string]interface{}
	StartIndex       float64
	TotalResultCount float64
	Warnings         []interface{}
	APIMajor         string `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string `json:"_rallyAPIMinor,omitempty"`
}

//AllowedQueryOperator auto-generated
type AllowedQueryOperator struct {
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	OperatorName     string                 `json:"OperatorName,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Attachment auto-generated
type Attachment struct {
	Artifact         map[string]interface{} `json:"Artifact,omitempty"`
	Content          *AttachmentContent     `json:"Content,omitempty"`
	ContentType      string                 `json:"ContentType,omitempty"`
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Size             float64                `json:"Size,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	User             *User                  `json:"User,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//AttachmentContent auto-generated
type AttachmentContent struct {
	Content          string                 `json:"Content,omitempty"`
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Errors           []interface{}          `json:"Errors,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Warnings         []interface{}          `json:"Warnings,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//AttributeDefinition auto-generated
type AttributeDefinition struct {
	AllowedQueryOperators *AllowedQueryOperator  `json:"AllowedQueryOperators,omitempty"`
	AllowedValues         map[string]interface{} `json:"AllowedValues,omitempty"`
	AttributeType         string                 `json:"AttributeType,omitempty"`
	CreationDate          string                 `json:"CreationDate,omitempty"`
	ElementName           string                 `json:"ElementName,omitempty"`
	MaxFractionalDigits   float64                `json:"MaxFractionalDigits,omitempty"`
	MaxLength             float64                `json:"MaxLength,omitempty"`
	Name                  string                 `json:"Name,omitempty"`
	Note                  string                 `json:"Note,omitempty"`
	ObjectID              float64                `json:"ObjectID,omitempty"`
	ObjectUUID            string                 `json:"ObjectUUID,omitempty"`
	RealAttributeType     string                 `json:"RealAttributeType,omitempty"`
	Subscription          *Subscription          `json:"Subscription,omitempty"`
	Type                  string                 `json:"Type,omitempty"`
	TypeDefinition        *TypeDefinition        `json:"TypeDefinition,omitempty"`
	VersionID             string                 `json:"VersionId,omitempty"`
	CreatedAt             string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion         string                 `json:"_objectVersion,omitempty"`
	APIMajor              string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor              string                 `json:"_rallyAPIMinor,omitempty"`
	URI                   string                 `json:"_ref,omitempty"`
	RefObjectName         string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID         string                 `json:"_refObjectUUID,omitempty"`
	RefType               string                 `json:"_type,omitempty"`
	CustomAttributes      map[string]interface{} `json:"CustomAttributes"`
}

//BuildDefinition auto-generated
type BuildDefinition struct {
	Builds           map[string]interface{} `json:"Builds,omitempty"`
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	LastStatus       string                 `json:"LastStatus,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Project          *Project               `json:"Project,omitempty"`
	Projects         *Project               `json:"Projects,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//HierarchicalRequirement auto-generated
type HierarchicalRequirement struct {
	AcceptedDate         string                   `json:"AcceptedDate,omitempty"`
	Attachments          *Attachment              `json:"Attachments,omitempty"`
	Changesets           map[string]interface{}   `json:"Changesets,omitempty"`
	Children             *HierarchicalRequirement `json:"Children,omitempty"`
	CreationDate         string                   `json:"CreationDate,omitempty"`
	DefectStatus         string                   `json:"DefectStatus,omitempty"`
	Defects              map[string]interface{}   `json:"Defects,omitempty"`
	Description          string                   `json:"Description,omitempty"`
	DirectChildrenCount  float64                  `json:"DirectChildrenCount,omitempty"`
	Discussion           map[string]interface{}   `json:"Discussion,omitempty"`
	DragAndDropRank      string                   `json:"DragAndDropRank,omitempty"`
	Feature              map[string]interface{}   `json:"Feature,omitempty"`
	FormattedID          string                   `json:"FormattedID,omitempty"`
	InProgressDate       string                   `json:"InProgressDate,omitempty"`
	Iteration            *Iteration               `json:"Iteration,omitempty"`
	LastUpdateDate       string                   `json:"LastUpdateDate,omitempty"`
	Milestones           map[string]interface{}   `json:"Milestones,omitempty"`
	Name                 string                   `json:"Name,omitempty"`
	Notes                string                   `json:"Notes,omitempty"`
	ObjectID             float64                  `json:"ObjectID,omitempty"`
	ObjectUUID           string                   `json:"ObjectUUID,omitempty"`
	Owner                *User                    `json:"Owner,omitempty"`
	PassingTestCaseCount float64                  `json:"PassingTestCaseCount,omitempty"`
	PlanEstimate         float64                  `json:"PlanEstimate,omitempty"`
	PortfolioItem        map[string]interface{}   `json:"PortfolioItem,omitempty"`
	Predecessors         *HierarchicalRequirement `json:"Predecessors,omitempty"`
	Project              *Project                 `json:"Project,omitempty"`
	Release              *Release                 `json:"Release,omitempty"`
	RevisionHistory      *RevisionHistory         `json:"RevisionHistory,omitempty"`
	ScheduleState        string                   `json:"ScheduleState,omitempty"`
	ScheduleStatePrefix  string                   `json:"ScheduleStatePrefix,omitempty"`
	Subscription         *Subscription            `json:"Subscription,omitempty"`
	Successors           *HierarchicalRequirement `json:"Successors,omitempty"`
	Tags                 *Tag                     `json:"Tags,omitempty"`
	TaskActualTotal      float64                  `json:"TaskActualTotal,omitempty"`
	TaskEstimateTotal    float64                  `json:"TaskEstimateTotal,omitempty"`
	TaskRemainingTotal   float64                  `json:"TaskRemainingTotal,omitempty"`
	TaskStatus           string                   `json:"TaskStatus,omitempty"`
	Tasks                *Task                    `json:"Tasks,omitempty"`
	TestCaseCount        float64                  `json:"TestCaseCount,omitempty"`
	TestCaseStatus       string                   `json:"TestCaseStatus,omitempty"`
	TestCases            map[string]interface{}   `json:"TestCases,omitempty"`
	VersionID            string                   `json:"VersionId,omitempty"`
	Workspace            *Workspace               `json:"Workspace,omitempty"`
	CreatedAt            string                   `json:"_CreatedAt,omitempty"`
	ObjectVersion        string                   `json:"_objectVersion,omitempty"`
	APIMajor             string                   `json:"_rallyAPIMajor,omitempty"`
	APIMinor             string                   `json:"_rallyAPIMinor,omitempty"`
	URI                  string                   `json:"_ref,omitempty"`
	RefObjectName        string                   `json:"_refObjectName,omitempty"`
	RefObjectUUID        string                   `json:"_refObjectUUID,omitempty"`
	RefType              string                   `json:"_type,omitempty"`
	CustomAttributes     map[string]interface{}   `json:"CustomAttributes"`
}

//Initiative auto-generated
type Initiative struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount,omitempty"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal,omitempty"`
	Attachments                        *Attachment            `json:"Attachments,omitempty"`
	Changesets                         map[string]interface{} `json:"Changesets,omitempty"`
	Children                           map[string]interface{} `json:"Children,omitempty"`
	Collaborators                      *User                  `json:"Collaborators,omitempty"`
	CreationDate                       string                 `json:"CreationDate,omitempty"`
	Description                        string                 `json:"Description,omitempty"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount,omitempty"`
	Discussion                         map[string]interface{} `json:"Discussion,omitempty"`
	DragAndDropRank                    string                 `json:"DragAndDropRank,omitempty"`
	Errors                             []interface{}          `json:"Errors,omitempty"`
	FormattedID                        string                 `json:"FormattedID,omitempty"`
	InvestmentCategory                 string                 `json:"InvestmentCategory,omitempty"`
	JobSize                            float64                `json:"JobSize,omitempty"`
	LastUpdateDate                     string                 `json:"LastUpdateDate,omitempty"`
	LeafStoryCount                     float64                `json:"LeafStoryCount,omitempty"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal,omitempty"`
	Milestones                         map[string]interface{} `json:"Milestones,omitempty"`
	Name                               string                 `json:"Name,omitempty"`
	Notes                              string                 `json:"Notes,omitempty"`
	ObjectID                           float64                `json:"ObjectID,omitempty"`
	ObjectUUID                         string                 `json:"ObjectUUID,omitempty"`
	Owner                              *User                  `json:"Owner,omitempty"`
	Parent                             map[string]interface{} `json:"Parent,omitempty"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount,omitempty"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate,omitempty"`
	PortfolioItemType                  *TypeDefinition        `json:"PortfolioItemType,omitempty"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName,omitempty"`
	Predecessors                       map[string]interface{} `json:"Predecessors,omitempty"`
	Project                            *Project               `json:"Project,omitempty"`
	RROEValue                          float64                `json:"RROEValue,omitempty"`
	RefinedEstimate                    float64                `json:"RefinedEstimate,omitempty"`
	RevisionHistory                    *RevisionHistory       `json:"RevisionHistory,omitempty"`
	RiskScore                          float64                `json:"RiskScore,omitempty"`
	StateChangedDate                   string                 `json:"StateChangedDate,omitempty"`
	Subscription                       *Subscription          `json:"Subscription,omitempty"`
	Successors                         map[string]interface{} `json:"Successors,omitempty"`
	Tags                               *Tag                   `json:"Tags,omitempty"`
	TimeCriticality                    float64                `json:"TimeCriticality,omitempty"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount,omitempty"`
	UserBusinessValue                  float64                `json:"UserBusinessValue,omitempty"`
	ValueScore                         float64                `json:"ValueScore,omitempty"`
	VersionID                          string                 `json:"VersionId,omitempty"`
	WSJFScore                          float64                `json:"WSJFScore,omitempty"`
	Warnings                           []interface{}          `json:"Warnings,omitempty"`
	Workspace                          *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                          string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion                      string                 `json:"_objectVersion,omitempty"`
	APIMajor                           string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                           string                 `json:"_rallyAPIMinor,omitempty"`
	URI                                string                 `json:"_ref,omitempty"`
	RefObjectName                      string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID                      string                 `json:"_refObjectUUID,omitempty"`
	RefType                            string                 `json:"_type,omitempty"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//Iteration auto-generated
type Iteration struct {
	CreationDate            string                 `json:"CreationDate,omitempty"`
	EndDate                 string                 `json:"EndDate,omitempty"`
	Name                    string                 `json:"Name,omitempty"`
	Notes                   string                 `json:"Notes,omitempty"`
	ObjectID                float64                `json:"ObjectID,omitempty"`
	ObjectUUID              string                 `json:"ObjectUUID,omitempty"`
	Project                 *Project               `json:"Project,omitempty"`
	RevisionHistory         *RevisionHistory       `json:"RevisionHistory,omitempty"`
	StartDate               string                 `json:"StartDate,omitempty"`
	State                   string                 `json:"State,omitempty"`
	Subscription            *Subscription          `json:"Subscription,omitempty"`
	Theme                   string                 `json:"Theme,omitempty"`
	UserIterationCapacities map[string]interface{} `json:"UserIterationCapacities,omitempty"`
	VersionID               string                 `json:"VersionId,omitempty"`
	Workspace               *Workspace             `json:"Workspace,omitempty"`
	CreatedAt               string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion           string                 `json:"_objectVersion,omitempty"`
	APIMajor                string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                string                 `json:"_rallyAPIMinor,omitempty"`
	URI                     string                 `json:"_ref,omitempty"`
	RefObjectName           string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID           string                 `json:"_refObjectUUID,omitempty"`
	RefType                 string                 `json:"_type,omitempty"`
	CustomAttributes        map[string]interface{} `json:"CustomAttributes"`
}

//PortfolioItemFeature auto-generated
type PortfolioItemFeature struct {
	AcceptedLeafStoryCount             float64                  `json:"AcceptedLeafStoryCount,omitempty"`
	AcceptedLeafStoryPlanEstimateTotal float64                  `json:"AcceptedLeafStoryPlanEstimateTotal,omitempty"`
	Attachments                        *Attachment              `json:"Attachments,omitempty"`
	Changesets                         map[string]interface{}   `json:"Changesets,omitempty"`
	Collaborators                      *User                    `json:"Collaborators,omitempty"`
	CreationDate                       string                   `json:"CreationDate,omitempty"`
	Description                        string                   `json:"Description,omitempty"`
	DirectChildrenCount                float64                  `json:"DirectChildrenCount,omitempty"`
	Discussion                         map[string]interface{}   `json:"Discussion,omitempty"`
	DisplayColor                       string                   `json:"DisplayColor,omitempty"`
	DragAndDropRank                    string                   `json:"DragAndDropRank,omitempty"`
	FormattedID                        string                   `json:"FormattedID,omitempty"`
	InvestmentCategory                 string                   `json:"InvestmentCategory,omitempty"`
	JobSize                            float64                  `json:"JobSize,omitempty"`
	LastUpdateDate                     string                   `json:"LastUpdateDate,omitempty"`
	LateChildCount                     float64                  `json:"LateChildCount,omitempty"`
	LeafStoryCount                     float64                  `json:"LeafStoryCount,omitempty"`
	LeafStoryPlanEstimateTotal         float64                  `json:"LeafStoryPlanEstimateTotal,omitempty"`
	Milestones                         map[string]interface{}   `json:"Milestones,omitempty"`
	Name                               string                   `json:"Name,omitempty"`
	Notes                              string                   `json:"Notes,omitempty"`
	ObjectID                           float64                  `json:"ObjectID,omitempty"`
	ObjectUUID                         string                   `json:"ObjectUUID,omitempty"`
	Owner                              *User                    `json:"Owner,omitempty"`
	Parent                             map[string]interface{}   `json:"Parent,omitempty"`
	PercentDoneByStoryCount            float64                  `json:"PercentDoneByStoryCount,omitempty"`
	PercentDoneByStoryPlanEstimate     float64                  `json:"PercentDoneByStoryPlanEstimate,omitempty"`
	PortfolioItemType                  *TypeDefinition          `json:"PortfolioItemType,omitempty"`
	PortfolioItemTypeName              string                   `json:"PortfolioItemTypeName,omitempty"`
	Predecessors                       map[string]interface{}   `json:"Predecessors,omitempty"`
	Project                            *Project                 `json:"Project,omitempty"`
	RROEValue                          float64                  `json:"RROEValue,omitempty"`
	RefinedEstimate                    float64                  `json:"RefinedEstimate,omitempty"`
	RevisionHistory                    *RevisionHistory         `json:"RevisionHistory,omitempty"`
	RiskScore                          float64                  `json:"RiskScore,omitempty"`
	State                              *State                   `json:"State,omitempty"`
	StateChangedDate                   string                   `json:"StateChangedDate,omitempty"`
	Subscription                       *Subscription            `json:"Subscription,omitempty"`
	Successors                         map[string]interface{}   `json:"Successors,omitempty"`
	Tags                               *Tag                     `json:"Tags,omitempty"`
	TimeCriticality                    float64                  `json:"TimeCriticality,omitempty"`
	UnEstimatedLeafStoryCount          float64                  `json:"UnEstimatedLeafStoryCount,omitempty"`
	UserBusinessValue                  float64                  `json:"UserBusinessValue,omitempty"`
	UserStories                        *HierarchicalRequirement `json:"UserStories,omitempty"`
	ValueScore                         float64                  `json:"ValueScore,omitempty"`
	VersionID                          string                   `json:"VersionId,omitempty"`
	WSJFScore                          float64                  `json:"WSJFScore,omitempty"`
	Workspace                          *Workspace               `json:"Workspace,omitempty"`
	CreatedAt                          string                   `json:"_CreatedAt,omitempty"`
	ObjectVersion                      string                   `json:"_objectVersion,omitempty"`
	APIMajor                           string                   `json:"_rallyAPIMajor,omitempty"`
	APIMinor                           string                   `json:"_rallyAPIMinor,omitempty"`
	URI                                string                   `json:"_ref,omitempty"`
	RefObjectName                      string                   `json:"_refObjectName,omitempty"`
	RefObjectUUID                      string                   `json:"_refObjectUUID,omitempty"`
	RefType                            string                   `json:"_type,omitempty"`
	CustomAttributes                   map[string]interface{}   `json:"CustomAttributes"`
}

//PortfolioItemInitiative auto-generated
type PortfolioItemInitiative struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount,omitempty"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal,omitempty"`
	Attachments                        *Attachment            `json:"Attachments,omitempty"`
	Changesets                         map[string]interface{} `json:"Changesets,omitempty"`
	Children                           map[string]interface{} `json:"Children,omitempty"`
	Collaborators                      *User                  `json:"Collaborators,omitempty"`
	CreationDate                       string                 `json:"CreationDate,omitempty"`
	Description                        string                 `json:"Description,omitempty"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount,omitempty"`
	Discussion                         map[string]interface{} `json:"Discussion,omitempty"`
	DragAndDropRank                    string                 `json:"DragAndDropRank,omitempty"`
	FormattedID                        string                 `json:"FormattedID,omitempty"`
	InvestmentCategory                 string                 `json:"InvestmentCategory,omitempty"`
	JobSize                            float64                `json:"JobSize,omitempty"`
	LastUpdateDate                     string                 `json:"LastUpdateDate,omitempty"`
	LeafStoryCount                     float64                `json:"LeafStoryCount,omitempty"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal,omitempty"`
	Milestones                         map[string]interface{} `json:"Milestones,omitempty"`
	Name                               string                 `json:"Name,omitempty"`
	Notes                              string                 `json:"Notes,omitempty"`
	ObjectID                           float64                `json:"ObjectID,omitempty"`
	ObjectUUID                         string                 `json:"ObjectUUID,omitempty"`
	Owner                              *User                  `json:"Owner,omitempty"`
	Parent                             map[string]interface{} `json:"Parent,omitempty"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount,omitempty"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate,omitempty"`
	PortfolioItemType                  *TypeDefinition        `json:"PortfolioItemType,omitempty"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName,omitempty"`
	Predecessors                       map[string]interface{} `json:"Predecessors,omitempty"`
	Project                            *Project               `json:"Project,omitempty"`
	RROEValue                          float64                `json:"RROEValue,omitempty"`
	RefinedEstimate                    float64                `json:"RefinedEstimate,omitempty"`
	RevisionHistory                    *RevisionHistory       `json:"RevisionHistory,omitempty"`
	RiskScore                          float64                `json:"RiskScore,omitempty"`
	StateChangedDate                   string                 `json:"StateChangedDate,omitempty"`
	Subscription                       *Subscription          `json:"Subscription,omitempty"`
	Successors                         map[string]interface{} `json:"Successors,omitempty"`
	Tags                               *Tag                   `json:"Tags,omitempty"`
	TimeCriticality                    float64                `json:"TimeCriticality,omitempty"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount,omitempty"`
	UserBusinessValue                  float64                `json:"UserBusinessValue,omitempty"`
	ValueScore                         float64                `json:"ValueScore,omitempty"`
	VersionID                          string                 `json:"VersionId,omitempty"`
	WSJFScore                          float64                `json:"WSJFScore,omitempty"`
	Workspace                          *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                          string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion                      string                 `json:"_objectVersion,omitempty"`
	APIMajor                           string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                           string                 `json:"_rallyAPIMinor,omitempty"`
	URI                                string                 `json:"_ref,omitempty"`
	RefObjectName                      string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID                      string                 `json:"_refObjectUUID,omitempty"`
	RefType                            string                 `json:"_type,omitempty"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//PortfolioItemTheme auto-generated
type PortfolioItemTheme struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount,omitempty"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal,omitempty"`
	Attachments                        *Attachment            `json:"Attachments,omitempty"`
	Changesets                         map[string]interface{} `json:"Changesets,omitempty"`
	Children                           map[string]interface{} `json:"Children,omitempty"`
	Collaborators                      *User                  `json:"Collaborators,omitempty"`
	CreationDate                       string                 `json:"CreationDate,omitempty"`
	Description                        string                 `json:"Description,omitempty"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount,omitempty"`
	Discussion                         map[string]interface{} `json:"Discussion,omitempty"`
	DragAndDropRank                    string                 `json:"DragAndDropRank,omitempty"`
	FormattedID                        string                 `json:"FormattedID,omitempty"`
	InvestmentCategory                 string                 `json:"InvestmentCategory,omitempty"`
	JobSize                            float64                `json:"JobSize,omitempty"`
	LastUpdateDate                     string                 `json:"LastUpdateDate,omitempty"`
	LeafStoryCount                     float64                `json:"LeafStoryCount,omitempty"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal,omitempty"`
	Milestones                         map[string]interface{} `json:"Milestones,omitempty"`
	Name                               string                 `json:"Name,omitempty"`
	Notes                              string                 `json:"Notes,omitempty"`
	ObjectID                           float64                `json:"ObjectID,omitempty"`
	ObjectUUID                         string                 `json:"ObjectUUID,omitempty"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount,omitempty"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate,omitempty"`
	PortfolioItemType                  *TypeDefinition        `json:"PortfolioItemType,omitempty"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName,omitempty"`
	Predecessors                       map[string]interface{} `json:"Predecessors,omitempty"`
	Project                            *Project               `json:"Project,omitempty"`
	RROEValue                          float64                `json:"RROEValue,omitempty"`
	RefinedEstimate                    float64                `json:"RefinedEstimate,omitempty"`
	RevisionHistory                    *RevisionHistory       `json:"RevisionHistory,omitempty"`
	RiskScore                          float64                `json:"RiskScore,omitempty"`
	StateChangedDate                   string                 `json:"StateChangedDate,omitempty"`
	Subscription                       *Subscription          `json:"Subscription,omitempty"`
	Successors                         map[string]interface{} `json:"Successors,omitempty"`
	Tags                               *Tag                   `json:"Tags,omitempty"`
	TimeCriticality                    float64                `json:"TimeCriticality,omitempty"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount,omitempty"`
	UserBusinessValue                  float64                `json:"UserBusinessValue,omitempty"`
	ValueScore                         float64                `json:"ValueScore,omitempty"`
	VersionID                          string                 `json:"VersionId,omitempty"`
	WSJFScore                          float64                `json:"WSJFScore,omitempty"`
	Workspace                          *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                          string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion                      string                 `json:"_objectVersion,omitempty"`
	APIMajor                           string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                           string                 `json:"_rallyAPIMinor,omitempty"`
	URI                                string                 `json:"_ref,omitempty"`
	RefObjectName                      string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID                      string                 `json:"_refObjectUUID,omitempty"`
	RefType                            string                 `json:"_type,omitempty"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//PreliminaryEstimate auto-generated
type PreliminaryEstimate struct {
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	Errors           []interface{}          `json:"Errors,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	RevisionHistory  *RevisionHistory       `json:"RevisionHistory,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	Value            float64                `json:"Value,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Warnings         []interface{}          `json:"Warnings,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Project auto-generated
type Project struct {
	BuildDefinitions *BuildDefinition       `json:"BuildDefinitions,omitempty"`
	Children         *Project               `json:"Children,omitempty"`
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	Editors          *User                  `json:"Editors,omitempty"`
	Iterations       *Iteration             `json:"Iterations,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	Notes            string                 `json:"Notes,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Owner            *User                  `json:"Owner,omitempty"`
	Releases         *Release               `json:"Releases,omitempty"`
	RevisionHistory  *RevisionHistory       `json:"RevisionHistory,omitempty"`
	SchemaVersion    string                 `json:"SchemaVersion,omitempty"`
	State            string                 `json:"State,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	TeamMembers      *User                  `json:"TeamMembers,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//ProjectPermission auto-generated
type ProjectPermission struct {
	CustomObjectID   string                 `json:"CustomObjectID,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Project          *Project               `json:"Project,omitempty"`
	Role             string                 `json:"Role,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	User             *User                  `json:"User,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Release auto-generated
type Release struct {
	CreationDate                 string                 `json:"CreationDate,omitempty"`
	GrossEstimateConversionRatio float64                `json:"GrossEstimateConversionRatio,omitempty"`
	Name                         string                 `json:"Name,omitempty"`
	Notes                        string                 `json:"Notes,omitempty"`
	ObjectID                     float64                `json:"ObjectID,omitempty"`
	ObjectUUID                   string                 `json:"ObjectUUID,omitempty"`
	Project                      *Project               `json:"Project,omitempty"`
	ReleaseDate                  string                 `json:"ReleaseDate,omitempty"`
	ReleaseStartDate             string                 `json:"ReleaseStartDate,omitempty"`
	RevisionHistory              *RevisionHistory       `json:"RevisionHistory,omitempty"`
	State                        string                 `json:"State,omitempty"`
	Subscription                 *Subscription          `json:"Subscription,omitempty"`
	Theme                        string                 `json:"Theme,omitempty"`
	VersionID                    string                 `json:"VersionId,omitempty"`
	Workspace                    *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                    string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion                string                 `json:"_objectVersion,omitempty"`
	APIMajor                     string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                     string                 `json:"_rallyAPIMinor,omitempty"`
	URI                          string                 `json:"_ref,omitempty"`
	RefObjectName                string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID                string                 `json:"_refObjectUUID,omitempty"`
	RefType                      string                 `json:"_type,omitempty"`
	CustomAttributes             map[string]interface{} `json:"CustomAttributes"`
}

//Revision auto-generated
type Revision struct {
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Description      string                 `json:"Description,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	RevisionHistory  *RevisionHistory       `json:"RevisionHistory,omitempty"`
	RevisionNumber   float64                `json:"RevisionNumber,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	User             *User                  `json:"User,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//RevisionHistory auto-generated
type RevisionHistory struct {
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Errors           []interface{}          `json:"Errors,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Revisions        *Revision              `json:"Revisions,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Warnings         []interface{}          `json:"Warnings,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//State auto-generated
type State struct {
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Errors           []interface{}          `json:"Errors,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	OrderIndex       float64                `json:"OrderIndex,omitempty"`
	RevisionHistory  *RevisionHistory       `json:"RevisionHistory,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	TypeDef          *TypeDefinition        `json:"TypeDef,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	WIPLimit         float64                `json:"WIPLimit,omitempty"`
	Warnings         []interface{}          `json:"Warnings,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Subscription auto-generated
type Subscription struct {
	CreationDate            string                 `json:"CreationDate,omitempty"`
	Errors                  []interface{}          `json:"Errors,omitempty"`
	MaximumCustomUserFields float64                `json:"MaximumCustomUserFields,omitempty"`
	MaximumProjects         float64                `json:"MaximumProjects,omitempty"`
	Modules                 string                 `json:"Modules,omitempty"`
	Name                    string                 `json:"Name,omitempty"`
	ObjectID                float64                `json:"ObjectID,omitempty"`
	ObjectUUID              string                 `json:"ObjectUUID,omitempty"`
	PasswordExpirationDays  float64                `json:"PasswordExpirationDays,omitempty"`
	PreviousPasswordCount   float64                `json:"PreviousPasswordCount,omitempty"`
	RevisionHistory         *RevisionHistory       `json:"RevisionHistory,omitempty"`
	StoryHierarchyType      string                 `json:"StoryHierarchyType,omitempty"`
	SubscriptionID          float64                `json:"SubscriptionID,omitempty"`
	SubscriptionType        string                 `json:"SubscriptionType,omitempty"`
	VersionID               string                 `json:"VersionId,omitempty"`
	Warnings                []interface{}          `json:"Warnings,omitempty"`
	Workspaces              *Workspace             `json:"Workspaces,omitempty"`
	CreatedAt               string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion           string                 `json:"_objectVersion,omitempty"`
	APIMajor                string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                string                 `json:"_rallyAPIMinor,omitempty"`
	URI                     string                 `json:"_ref,omitempty"`
	RefObjectName           string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID           string                 `json:"_refObjectUUID,omitempty"`
	CustomAttributes        map[string]interface{} `json:"CustomAttributes"`
}

//Tag auto-generated
type Tag struct {
	CreationDate     string                 `json:"CreationDate,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Task auto-generated
type Task struct {
	Attachments      *Attachment              `json:"Attachments,omitempty"`
	Changesets       map[string]interface{}   `json:"Changesets,omitempty"`
	CreationDate     string                   `json:"CreationDate,omitempty"`
	Description      string                   `json:"Description,omitempty"`
	Discussion       map[string]interface{}   `json:"Discussion,omitempty"`
	DragAndDropRank  string                   `json:"DragAndDropRank,omitempty"`
	Estimate         float64                  `json:"Estimate,omitempty"`
	FormattedID      string                   `json:"FormattedID,omitempty"`
	Iteration        *Iteration               `json:"Iteration,omitempty"`
	LastUpdateDate   string                   `json:"LastUpdateDate,omitempty"`
	Milestones       map[string]interface{}   `json:"Milestones,omitempty"`
	Name             string                   `json:"Name,omitempty"`
	Notes            string                   `json:"Notes,omitempty"`
	ObjectID         float64                  `json:"ObjectID,omitempty"`
	ObjectUUID       string                   `json:"ObjectUUID,omitempty"`
	Owner            *User                    `json:"Owner,omitempty"`
	Project          *Project                 `json:"Project,omitempty"`
	Release          *Release                 `json:"Release,omitempty"`
	RevisionHistory  *RevisionHistory         `json:"RevisionHistory,omitempty"`
	State            string                   `json:"State,omitempty"`
	Subscription     *Subscription            `json:"Subscription,omitempty"`
	Tags             *Tag                     `json:"Tags,omitempty"`
	TaskIndex        float64                  `json:"TaskIndex,omitempty"`
	TimeSpent        float64                  `json:"TimeSpent,omitempty"`
	ToDo             float64                  `json:"ToDo,omitempty"`
	VersionID        string                   `json:"VersionId,omitempty"`
	WorkProduct      *HierarchicalRequirement `json:"WorkProduct,omitempty"`
	Workspace        *Workspace               `json:"Workspace,omitempty"`
	CreatedAt        string                   `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                   `json:"_objectVersion,omitempty"`
	APIMajor         string                   `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                   `json:"_rallyAPIMinor,omitempty"`
	URI              string                   `json:"_ref,omitempty"`
	RefObjectName    string                   `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                   `json:"_refObjectUUID,omitempty"`
	RefType          string                   `json:"_type,omitempty"`
	CustomAttributes map[string]interface{}   `json:"CustomAttributes"`
}

//Theme auto-generated
type Theme struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount,omitempty"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal,omitempty"`
	Attachments                        *Attachment            `json:"Attachments,omitempty"`
	Changesets                         map[string]interface{} `json:"Changesets,omitempty"`
	Children                           map[string]interface{} `json:"Children,omitempty"`
	Collaborators                      *User                  `json:"Collaborators,omitempty"`
	CreationDate                       string                 `json:"CreationDate,omitempty"`
	Description                        string                 `json:"Description,omitempty"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount,omitempty"`
	Discussion                         map[string]interface{} `json:"Discussion,omitempty"`
	DragAndDropRank                    string                 `json:"DragAndDropRank,omitempty"`
	Errors                             []interface{}          `json:"Errors,omitempty"`
	FormattedID                        string                 `json:"FormattedID,omitempty"`
	InvestmentCategory                 string                 `json:"InvestmentCategory,omitempty"`
	JobSize                            float64                `json:"JobSize,omitempty"`
	LastUpdateDate                     string                 `json:"LastUpdateDate,omitempty"`
	LeafStoryCount                     float64                `json:"LeafStoryCount,omitempty"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal,omitempty"`
	Milestones                         map[string]interface{} `json:"Milestones,omitempty"`
	Name                               string                 `json:"Name,omitempty"`
	Notes                              string                 `json:"Notes,omitempty"`
	ObjectID                           float64                `json:"ObjectID,omitempty"`
	ObjectUUID                         string                 `json:"ObjectUUID,omitempty"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount,omitempty"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate,omitempty"`
	PortfolioItemType                  *TypeDefinition        `json:"PortfolioItemType,omitempty"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName,omitempty"`
	Predecessors                       map[string]interface{} `json:"Predecessors,omitempty"`
	Project                            *Project               `json:"Project,omitempty"`
	RROEValue                          float64                `json:"RROEValue,omitempty"`
	RefinedEstimate                    float64                `json:"RefinedEstimate,omitempty"`
	RevisionHistory                    *RevisionHistory       `json:"RevisionHistory,omitempty"`
	RiskScore                          float64                `json:"RiskScore,omitempty"`
	StateChangedDate                   string                 `json:"StateChangedDate,omitempty"`
	Subscription                       *Subscription          `json:"Subscription,omitempty"`
	Successors                         map[string]interface{} `json:"Successors,omitempty"`
	Tags                               *Tag                   `json:"Tags,omitempty"`
	TimeCriticality                    float64                `json:"TimeCriticality,omitempty"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount,omitempty"`
	UserBusinessValue                  float64                `json:"UserBusinessValue,omitempty"`
	ValueScore                         float64                `json:"ValueScore,omitempty"`
	VersionID                          string                 `json:"VersionId,omitempty"`
	WSJFScore                          float64                `json:"WSJFScore,omitempty"`
	Warnings                           []interface{}          `json:"Warnings,omitempty"`
	Workspace                          *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                          string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion                      string                 `json:"_objectVersion,omitempty"`
	APIMajor                           string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                           string                 `json:"_rallyAPIMinor,omitempty"`
	URI                                string                 `json:"_ref,omitempty"`
	RefObjectName                      string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID                      string                 `json:"_refObjectUUID,omitempty"`
	RefType                            string                 `json:"_type,omitempty"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//TypeDefinition auto-generated
type TypeDefinition struct {
	Attributes       *AttributeDefinition   `json:"Attributes,omitempty"`
	CreationDate     string                 `json:"CreationDate,omitempty"`
	DisplayName      string                 `json:"DisplayName,omitempty"`
	ElementName      string                 `json:"ElementName,omitempty"`
	IDPrefix         string                 `json:"IDPrefix,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	Note             string                 `json:"Note,omitempty"`
	ObjectID         float64                `json:"ObjectID,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Ordinal          float64                `json:"Ordinal,omitempty"`
	Parent           *TypeDefinition        `json:"Parent,omitempty"`
	RevisionHistory  *RevisionHistory       `json:"RevisionHistory,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	TypePath         string                 `json:"TypePath,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//User auto-generated
type User struct {
	CostCenter             string                 `json:"CostCenter,omitempty"`
	CreationDate           string                 `json:"CreationDate,omitempty"`
	Department             string                 `json:"Department,omitempty"`
	DisplayName            string                 `json:"DisplayName,omitempty"`
	EmailAddress           string                 `json:"EmailAddress,omitempty"`
	Errors                 []interface{}          `json:"Errors,omitempty"`
	FirstName              string                 `json:"FirstName,omitempty"`
	LandingPage            string                 `json:"LandingPage,omitempty"`
	LastLoginDate          string                 `json:"LastLoginDate,omitempty"`
	LastName               string                 `json:"LastName,omitempty"`
	LastPasswordUpdateDate string                 `json:"LastPasswordUpdateDate,omitempty"`
	LastSystemTimeZoneName string                 `json:"LastSystemTimeZoneName,omitempty"`
	ObjectID               float64                `json:"ObjectID,omitempty"`
	ObjectUUID             string                 `json:"ObjectUUID,omitempty"`
	OfficeLocation         string                 `json:"OfficeLocation,omitempty"`
	RevisionHistory        *RevisionHistory       `json:"RevisionHistory,omitempty"`
	Role                   string                 `json:"Role,omitempty"`
	Subscription           *Subscription          `json:"Subscription,omitempty"`
	SubscriptionID         float64                `json:"SubscriptionID,omitempty"`
	SubscriptionPermission string                 `json:"SubscriptionPermission,omitempty"`
	TeamMemberships        *Project               `json:"TeamMemberships,omitempty"`
	UserName               string                 `json:"UserName,omitempty"`
	UserPermissions        map[string]interface{} `json:"UserPermissions,omitempty"`
	UserProfile            *UserProfile           `json:"UserProfile,omitempty"`
	VersionID              string                 `json:"VersionId,omitempty"`
	Warnings               []interface{}          `json:"Warnings,omitempty"`
	WorkspacePermission    string                 `json:"WorkspacePermission,omitempty"`
	CreatedAt              string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion          string                 `json:"_objectVersion,omitempty"`
	APIMajor               string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor               string                 `json:"_rallyAPIMinor,omitempty"`
	URI                    string                 `json:"_ref,omitempty"`
	RefObjectName          string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID          string                 `json:"_refObjectUUID,omitempty"`
	RefType                string                 `json:"_type,omitempty"`
	CustomAttributes       map[string]interface{} `json:"CustomAttributes"`
}

//UserProfile auto-generated
type UserProfile struct {
	CreationDate          string                 `json:"CreationDate,omitempty"`
	DateFormat            string                 `json:"DateFormat,omitempty"`
	DateTimeFormat        string                 `json:"DateTimeFormat,omitempty"`
	Errors                []interface{}          `json:"Errors,omitempty"`
	ObjectID              float64                `json:"ObjectID,omitempty"`
	ObjectUUID            string                 `json:"ObjectUUID,omitempty"`
	SessionTimeoutSeconds float64                `json:"SessionTimeoutSeconds,omitempty"`
	Subscription          *Subscription          `json:"Subscription,omitempty"`
	TimeZone              string                 `json:"TimeZone,omitempty"`
	VersionID             string                 `json:"VersionId,omitempty"`
	Warnings              []interface{}          `json:"Warnings,omitempty"`
	CreatedAt             string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion         string                 `json:"_objectVersion,omitempty"`
	APIMajor              string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor              string                 `json:"_rallyAPIMinor,omitempty"`
	URI                   string                 `json:"_ref,omitempty"`
	RefObjectUUID         string                 `json:"_refObjectUUID,omitempty"`
	RefType               string                 `json:"_type,omitempty"`
	CustomAttributes      map[string]interface{} `json:"CustomAttributes"`
}

//Workspace auto-generated
type Workspace struct {
	Children               *Project                `json:"Children,omitempty"`
	CreationDate           string                  `json:"CreationDate,omitempty"`
	Description            string                  `json:"Description,omitempty"`
	Errors                 []interface{}           `json:"Errors,omitempty"`
	Name                   string                  `json:"Name,omitempty"`
	Notes                  string                  `json:"Notes,omitempty"`
	ObjectID               float64                 `json:"ObjectID,omitempty"`
	ObjectUUID             string                  `json:"ObjectUUID,omitempty"`
	Owner                  *User                   `json:"Owner,omitempty"`
	Projects               *Project                `json:"Projects,omitempty"`
	RevisionHistory        *RevisionHistory        `json:"RevisionHistory,omitempty"`
	SchemaVersion          string                  `json:"SchemaVersion,omitempty"`
	State                  string                  `json:"State,omitempty"`
	Style                  string                  `json:"Style,omitempty"`
	Subscription           *Subscription           `json:"Subscription,omitempty"`
	TypeDefinitions        *TypeDefinition         `json:"TypeDefinitions,omitempty"`
	VersionID              string                  `json:"VersionId,omitempty"`
	Warnings               []interface{}           `json:"Warnings,omitempty"`
	WorkspaceConfiguration *WorkspaceConfiguration `json:"WorkspaceConfiguration,omitempty"`
	CreatedAt              string                  `json:"_CreatedAt,omitempty"`
	ObjectVersion          string                  `json:"_objectVersion,omitempty"`
	APIMajor               string                  `json:"_rallyAPIMajor,omitempty"`
	APIMinor               string                  `json:"_rallyAPIMinor,omitempty"`
	URI                    string                  `json:"_ref,omitempty"`
	RefObjectName          string                  `json:"_refObjectName,omitempty"`
	RefObjectUUID          string                  `json:"_refObjectUUID,omitempty"`
	RefType                string                  `json:"_type,omitempty"`
	CustomAttributes       map[string]interface{}  `json:"CustomAttributes"`
}

//WorkspaceConfiguration auto-generated
type WorkspaceConfiguration struct {
	CreationDate              string                 `json:"CreationDate,omitempty"`
	DateFormat                string                 `json:"DateFormat,omitempty"`
	DateTimeFormat            string                 `json:"DateTimeFormat,omitempty"`
	Errors                    []interface{}          `json:"Errors,omitempty"`
	IterationEstimateUnitName string                 `json:"IterationEstimateUnitName,omitempty"`
	ObjectID                  float64                `json:"ObjectID,omitempty"`
	ObjectUUID                string                 `json:"ObjectUUID,omitempty"`
	ReleaseEstimateUnitName   string                 `json:"ReleaseEstimateUnitName,omitempty"`
	Subscription              *Subscription          `json:"Subscription,omitempty"`
	TaskUnitName              string                 `json:"TaskUnitName,omitempty"`
	TimeZone                  string                 `json:"TimeZone,omitempty"`
	VersionID                 string                 `json:"VersionId,omitempty"`
	Warnings                  []interface{}          `json:"Warnings,omitempty"`
	WorkDays                  string                 `json:"WorkDays,omitempty"`
	Workspace                 *Workspace             `json:"Workspace,omitempty"`
	CreatedAt                 string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion             string                 `json:"_objectVersion,omitempty"`
	APIMajor                  string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor                  string                 `json:"_rallyAPIMinor,omitempty"`
	URI                       string                 `json:"_ref,omitempty"`
	RefObjectUUID             string                 `json:"_refObjectUUID,omitempty"`
	RefType                   string                 `json:"_type,omitempty"`
	CustomAttributes          map[string]interface{} `json:"CustomAttributes"`
}

//WorkspacePermission auto-generated
type WorkspacePermission struct {
	CustomObjectID   string                 `json:"CustomObjectID,omitempty"`
	Name             string                 `json:"Name,omitempty"`
	ObjectUUID       string                 `json:"ObjectUUID,omitempty"`
	Role             string                 `json:"Role,omitempty"`
	Subscription     *Subscription          `json:"Subscription,omitempty"`
	User             *User                  `json:"User,omitempty"`
	VersionID        string                 `json:"VersionId,omitempty"`
	Workspace        *Workspace             `json:"Workspace,omitempty"`
	CreatedAt        string                 `json:"_CreatedAt,omitempty"`
	ObjectVersion    string                 `json:"_objectVersion,omitempty"`
	APIMajor         string                 `json:"_rallyAPIMajor,omitempty"`
	APIMinor         string                 `json:"_rallyAPIMinor,omitempty"`
	URI              string                 `json:"_ref,omitempty"`
	RefObjectName    string                 `json:"_refObjectName,omitempty"`
	RefObjectUUID    string                 `json:"_refObjectUUID,omitempty"`
	RefType          string                 `json:"_type,omitempty"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}
