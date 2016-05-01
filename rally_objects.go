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
	APIMajor      string `json:"_rallyAPIMajor"`
	APIMinor      string `json:"_rallyAPIMinor"`
	URI           string `json:"_ref"`
	RefObjectName string `json:"_refObjectName"`
	RefObjectUUID string `json:"_refObjectUUID"`
	RefType       string `json:"_type"`
}

//QueryResult standard query result
type QueryResult struct {
	Errors           []interface{}
	PageSize         float64
	Results          []interface{}
	StartIndex       float64
	TotalResultCount float64
	Warnings         []interface{}
	APIMajor         string `json:"_rallyAPIMajor"`
	APIMinor         string `json:"_rallyAPIMinor"`
}

//AllowedQueryOperator auto-generated
type AllowedQueryOperator struct {
	ObjectUUID       string                 `json:"ObjectUUID"`
	OperatorName     string                 `json:"OperatorName"`
	VersionID        string                 `json:"VersionId"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Attachment auto-generated
type Attachment struct {
	Artifact         map[string]interface{} `json:"Artifact"`
	Content          AttachmentContent      `json:"Content"`
	ContentType      string                 `json:"ContentType"`
	CreationDate     string                 `json:"CreationDate"`
	Name             string                 `json:"Name"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Size             float64                `json:"Size"`
	Subscription     Subscription           `json:"Subscription"`
	User             User                   `json:"User"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//AttachmentContent auto-generated
type AttachmentContent struct {
	Content          string                 `json:"Content"`
	CreationDate     string                 `json:"CreationDate"`
	Errors           []interface{}          `json:"Errors"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Subscription     Subscription           `json:"Subscription"`
	VersionID        string                 `json:"VersionId"`
	Warnings         []interface{}          `json:"Warnings"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//AttributeDefinition auto-generated
type AttributeDefinition struct {
	AllowedQueryOperators []AllowedQueryOperator `json:"AllowedQueryOperators"`
	AllowedValues         map[string]interface{} `json:"AllowedValues"`
	AttributeType         string                 `json:"AttributeType"`
	CreationDate          string                 `json:"CreationDate"`
	ElementName           string                 `json:"ElementName"`
	MaxFractionalDigits   float64                `json:"MaxFractionalDigits"`
	MaxLength             float64                `json:"MaxLength"`
	Name                  string                 `json:"Name"`
	Note                  string                 `json:"Note"`
	ObjectID              float64                `json:"ObjectID"`
	ObjectUUID            string                 `json:"ObjectUUID"`
	RealAttributeType     string                 `json:"RealAttributeType"`
	Subscription          Subscription           `json:"Subscription"`
	Type                  string                 `json:"Type"`
	TypeDefinition        TypeDefinition         `json:"TypeDefinition"`
	VersionID             string                 `json:"VersionId"`
	CreatedAt             string                 `json:"_CreatedAt"`
	ObjectVersion         string                 `json:"_objectVersion"`
	APIMajor              string                 `json:"_rallyAPIMajor"`
	APIMinor              string                 `json:"_rallyAPIMinor"`
	URI                   string                 `json:"_ref"`
	RefObjectName         string                 `json:"_refObjectName"`
	RefObjectUUID         string                 `json:"_refObjectUUID"`
	RefType               string                 `json:"_type"`
	CustomAttributes      map[string]interface{} `json:"CustomAttributes"`
}

//BuildDefinition auto-generated
type BuildDefinition struct {
	Builds           map[string]interface{} `json:"Builds"`
	CreationDate     string                 `json:"CreationDate"`
	Description      string                 `json:"Description"`
	LastStatus       string                 `json:"LastStatus"`
	Name             string                 `json:"Name"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Project          Project                `json:"Project"`
	Projects         []Project              `json:"Projects"`
	Subscription     Subscription           `json:"Subscription"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//HierarchicalRequirement auto-generated
type HierarchicalRequirement struct {
	AcceptedDate         string                    `json:"AcceptedDate"`
	Attachments          []Attachment              `json:"Attachments"`
	Changesets           map[string]interface{}    `json:"Changesets"`
	Children             []HierarchicalRequirement `json:"Children"`
	CreationDate         string                    `json:"CreationDate"`
	DefectStatus         string                    `json:"DefectStatus"`
	Defects              map[string]interface{}    `json:"Defects"`
	Description          string                    `json:"Description"`
	DirectChildrenCount  float64                   `json:"DirectChildrenCount"`
	Discussion           map[string]interface{}    `json:"Discussion"`
	DragAndDropRank      string                    `json:"DragAndDropRank"`
	Feature              map[string]interface{}    `json:"Feature"`
	FormattedID          string                    `json:"FormattedID"`
	InProgressDate       string                    `json:"InProgressDate"`
	Iteration            Iteration                 `json:"Iteration"`
	LastUpdateDate       string                    `json:"LastUpdateDate"`
	Milestones           map[string]interface{}    `json:"Milestones"`
	Name                 string                    `json:"Name"`
	Notes                string                    `json:"Notes"`
	ObjectID             float64                   `json:"ObjectID"`
	ObjectUUID           string                    `json:"ObjectUUID"`
	Owner                User                      `json:"Owner"`
	PassingTestCaseCount float64                   `json:"PassingTestCaseCount"`
	PlanEstimate         float64                   `json:"PlanEstimate"`
	PortfolioItem        map[string]interface{}    `json:"PortfolioItem"`
	Predecessors         []HierarchicalRequirement `json:"Predecessors"`
	Project              Project                   `json:"Project"`
	Release              Release                   `json:"Release"`
	RevisionHistory      RevisionHistory           `json:"RevisionHistory"`
	ScheduleState        string                    `json:"ScheduleState"`
	ScheduleStatePrefix  string                    `json:"ScheduleStatePrefix"`
	Subscription         Subscription              `json:"Subscription"`
	Successors           []HierarchicalRequirement `json:"Successors"`
	Tags                 []Tag                     `json:"Tags"`
	TaskActualTotal      float64                   `json:"TaskActualTotal"`
	TaskEstimateTotal    float64                   `json:"TaskEstimateTotal"`
	TaskRemainingTotal   float64                   `json:"TaskRemainingTotal"`
	TaskStatus           string                    `json:"TaskStatus"`
	Tasks                []Task                    `json:"Tasks"`
	TestCaseCount        float64                   `json:"TestCaseCount"`
	TestCaseStatus       string                    `json:"TestCaseStatus"`
	TestCases            map[string]interface{}    `json:"TestCases"`
	VersionID            string                    `json:"VersionId"`
	Workspace            Workspace                 `json:"Workspace"`
	CreatedAt            string                    `json:"_CreatedAt"`
	ObjectVersion        string                    `json:"_objectVersion"`
	APIMajor             string                    `json:"_rallyAPIMajor"`
	APIMinor             string                    `json:"_rallyAPIMinor"`
	URI                  string                    `json:"_ref"`
	RefObjectName        string                    `json:"_refObjectName"`
	RefObjectUUID        string                    `json:"_refObjectUUID"`
	RefType              string                    `json:"_type"`
	CustomAttributes     map[string]interface{}    `json:"CustomAttributes"`
}

//Initiative auto-generated
type Initiative struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal"`
	Attachments                        []Attachment           `json:"Attachments"`
	Changesets                         map[string]interface{} `json:"Changesets"`
	Children                           map[string]interface{} `json:"Children"`
	Collaborators                      []User                 `json:"Collaborators"`
	CreationDate                       string                 `json:"CreationDate"`
	Description                        string                 `json:"Description"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount"`
	Discussion                         map[string]interface{} `json:"Discussion"`
	DragAndDropRank                    string                 `json:"DragAndDropRank"`
	Errors                             []interface{}          `json:"Errors"`
	FormattedID                        string                 `json:"FormattedID"`
	InvestmentCategory                 string                 `json:"InvestmentCategory"`
	JobSize                            float64                `json:"JobSize"`
	LastUpdateDate                     string                 `json:"LastUpdateDate"`
	LeafStoryCount                     float64                `json:"LeafStoryCount"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal"`
	Milestones                         map[string]interface{} `json:"Milestones"`
	Name                               string                 `json:"Name"`
	Notes                              string                 `json:"Notes"`
	ObjectID                           float64                `json:"ObjectID"`
	ObjectUUID                         string                 `json:"ObjectUUID"`
	Owner                              User                   `json:"Owner"`
	Parent                             map[string]interface{} `json:"Parent"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate"`
	PortfolioItemType                  TypeDefinition         `json:"PortfolioItemType"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName"`
	Predecessors                       map[string]interface{} `json:"Predecessors"`
	Project                            Project                `json:"Project"`
	RROEValue                          float64                `json:"RROEValue"`
	RefinedEstimate                    float64                `json:"RefinedEstimate"`
	RevisionHistory                    RevisionHistory        `json:"RevisionHistory"`
	RiskScore                          float64                `json:"RiskScore"`
	StateChangedDate                   string                 `json:"StateChangedDate"`
	Subscription                       Subscription           `json:"Subscription"`
	Successors                         map[string]interface{} `json:"Successors"`
	Tags                               []Tag                  `json:"Tags"`
	TimeCriticality                    float64                `json:"TimeCriticality"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount"`
	UserBusinessValue                  float64                `json:"UserBusinessValue"`
	ValueScore                         float64                `json:"ValueScore"`
	VersionID                          string                 `json:"VersionId"`
	WSJFScore                          float64                `json:"WSJFScore"`
	Warnings                           []interface{}          `json:"Warnings"`
	Workspace                          Workspace              `json:"Workspace"`
	CreatedAt                          string                 `json:"_CreatedAt"`
	ObjectVersion                      string                 `json:"_objectVersion"`
	APIMajor                           string                 `json:"_rallyAPIMajor"`
	APIMinor                           string                 `json:"_rallyAPIMinor"`
	URI                                string                 `json:"_ref"`
	RefObjectName                      string                 `json:"_refObjectName"`
	RefObjectUUID                      string                 `json:"_refObjectUUID"`
	RefType                            string                 `json:"_type"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//Iteration auto-generated
type Iteration struct {
	CreationDate            string                  `json:"CreationDate"`
	EndDate                 string                  `json:"EndDate"`
	Name                    string                  `json:"Name"`
	Notes                   string                  `json:"Notes"`
	ObjectID                float64                 `json:"ObjectID"`
	ObjectUUID              string                  `json:"ObjectUUID"`
	Project                 Project                 `json:"Project"`
	RevisionHistory         RevisionHistory         `json:"RevisionHistory"`
	StartDate               string                  `json:"StartDate"`
	State                   string                  `json:"State"`
	Subscription            Subscription            `json:"Subscription"`
	Theme                   string                  `json:"Theme"`
	UserIterationCapacities []UserIterationCapacity `json:"UserIterationCapacities"`
	VersionID               string                  `json:"VersionId"`
	Workspace               Workspace               `json:"Workspace"`
	CreatedAt               string                  `json:"_CreatedAt"`
	ObjectVersion           string                  `json:"_objectVersion"`
	APIMajor                string                  `json:"_rallyAPIMajor"`
	APIMinor                string                  `json:"_rallyAPIMinor"`
	URI                     string                  `json:"_ref"`
	RefObjectName           string                  `json:"_refObjectName"`
	RefObjectUUID           string                  `json:"_refObjectUUID"`
	RefType                 string                  `json:"_type"`
	CustomAttributes        map[string]interface{}  `json:"CustomAttributes"`
}

//OperationResult auto-generated
type OperationResult struct {
	Errors           []interface{}          `json:"Errors"`
	Warnings         []interface{}          `json:"Warnings"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//PortfolioItemFeature auto-generated
type PortfolioItemFeature struct {
	AcceptedLeafStoryCount             float64                   `json:"AcceptedLeafStoryCount"`
	AcceptedLeafStoryPlanEstimateTotal float64                   `json:"AcceptedLeafStoryPlanEstimateTotal"`
	Attachments                        []Attachment              `json:"Attachments"`
	Changesets                         map[string]interface{}    `json:"Changesets"`
	Collaborators                      []User                    `json:"Collaborators"`
	CreationDate                       string                    `json:"CreationDate"`
	Description                        string                    `json:"Description"`
	DirectChildrenCount                float64                   `json:"DirectChildrenCount"`
	Discussion                         map[string]interface{}    `json:"Discussion"`
	DisplayColor                       string                    `json:"DisplayColor"`
	DragAndDropRank                    string                    `json:"DragAndDropRank"`
	FormattedID                        string                    `json:"FormattedID"`
	InvestmentCategory                 string                    `json:"InvestmentCategory"`
	JobSize                            float64                   `json:"JobSize"`
	LastUpdateDate                     string                    `json:"LastUpdateDate"`
	LateChildCount                     float64                   `json:"LateChildCount"`
	LeafStoryCount                     float64                   `json:"LeafStoryCount"`
	LeafStoryPlanEstimateTotal         float64                   `json:"LeafStoryPlanEstimateTotal"`
	Milestones                         map[string]interface{}    `json:"Milestones"`
	Name                               string                    `json:"Name"`
	Notes                              string                    `json:"Notes"`
	ObjectID                           float64                   `json:"ObjectID"`
	ObjectUUID                         string                    `json:"ObjectUUID"`
	Owner                              User                      `json:"Owner"`
	Parent                             map[string]interface{}    `json:"Parent"`
	PercentDoneByStoryCount            float64                   `json:"PercentDoneByStoryCount"`
	PercentDoneByStoryPlanEstimate     float64                   `json:"PercentDoneByStoryPlanEstimate"`
	PortfolioItemType                  TypeDefinition            `json:"PortfolioItemType"`
	PortfolioItemTypeName              string                    `json:"PortfolioItemTypeName"`
	Predecessors                       map[string]interface{}    `json:"Predecessors"`
	Project                            Project                   `json:"Project"`
	RROEValue                          float64                   `json:"RROEValue"`
	RefinedEstimate                    float64                   `json:"RefinedEstimate"`
	RevisionHistory                    RevisionHistory           `json:"RevisionHistory"`
	RiskScore                          float64                   `json:"RiskScore"`
	State                              State                     `json:"State"`
	StateChangedDate                   string                    `json:"StateChangedDate"`
	Subscription                       Subscription              `json:"Subscription"`
	Successors                         map[string]interface{}    `json:"Successors"`
	Tags                               []Tag                     `json:"Tags"`
	TimeCriticality                    float64                   `json:"TimeCriticality"`
	UnEstimatedLeafStoryCount          float64                   `json:"UnEstimatedLeafStoryCount"`
	UserBusinessValue                  float64                   `json:"UserBusinessValue"`
	UserStories                        []HierarchicalRequirement `json:"UserStories"`
	ValueScore                         float64                   `json:"ValueScore"`
	VersionID                          string                    `json:"VersionId"`
	WSJFScore                          float64                   `json:"WSJFScore"`
	Workspace                          Workspace                 `json:"Workspace"`
	CreatedAt                          string                    `json:"_CreatedAt"`
	ObjectVersion                      string                    `json:"_objectVersion"`
	APIMajor                           string                    `json:"_rallyAPIMajor"`
	APIMinor                           string                    `json:"_rallyAPIMinor"`
	URI                                string                    `json:"_ref"`
	RefObjectName                      string                    `json:"_refObjectName"`
	RefObjectUUID                      string                    `json:"_refObjectUUID"`
	RefType                            string                    `json:"_type"`
	CustomAttributes                   map[string]interface{}    `json:"CustomAttributes"`
}

//PortfolioItemInitiative auto-generated
type PortfolioItemInitiative struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal"`
	Attachments                        []Attachment           `json:"Attachments"`
	Changesets                         map[string]interface{} `json:"Changesets"`
	Children                           map[string]interface{} `json:"Children"`
	Collaborators                      []User                 `json:"Collaborators"`
	CreationDate                       string                 `json:"CreationDate"`
	Description                        string                 `json:"Description"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount"`
	Discussion                         map[string]interface{} `json:"Discussion"`
	DragAndDropRank                    string                 `json:"DragAndDropRank"`
	FormattedID                        string                 `json:"FormattedID"`
	InvestmentCategory                 string                 `json:"InvestmentCategory"`
	JobSize                            float64                `json:"JobSize"`
	LastUpdateDate                     string                 `json:"LastUpdateDate"`
	LeafStoryCount                     float64                `json:"LeafStoryCount"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal"`
	Milestones                         map[string]interface{} `json:"Milestones"`
	Name                               string                 `json:"Name"`
	Notes                              string                 `json:"Notes"`
	ObjectID                           float64                `json:"ObjectID"`
	ObjectUUID                         string                 `json:"ObjectUUID"`
	Owner                              User                   `json:"Owner"`
	Parent                             map[string]interface{} `json:"Parent"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate"`
	PortfolioItemType                  TypeDefinition         `json:"PortfolioItemType"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName"`
	Predecessors                       map[string]interface{} `json:"Predecessors"`
	Project                            Project                `json:"Project"`
	RROEValue                          float64                `json:"RROEValue"`
	RefinedEstimate                    float64                `json:"RefinedEstimate"`
	RevisionHistory                    RevisionHistory        `json:"RevisionHistory"`
	RiskScore                          float64                `json:"RiskScore"`
	StateChangedDate                   string                 `json:"StateChangedDate"`
	Subscription                       Subscription           `json:"Subscription"`
	Successors                         map[string]interface{} `json:"Successors"`
	Tags                               []Tag                  `json:"Tags"`
	TimeCriticality                    float64                `json:"TimeCriticality"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount"`
	UserBusinessValue                  float64                `json:"UserBusinessValue"`
	ValueScore                         float64                `json:"ValueScore"`
	VersionID                          string                 `json:"VersionId"`
	WSJFScore                          float64                `json:"WSJFScore"`
	Workspace                          Workspace              `json:"Workspace"`
	CreatedAt                          string                 `json:"_CreatedAt"`
	ObjectVersion                      string                 `json:"_objectVersion"`
	APIMajor                           string                 `json:"_rallyAPIMajor"`
	APIMinor                           string                 `json:"_rallyAPIMinor"`
	URI                                string                 `json:"_ref"`
	RefObjectName                      string                 `json:"_refObjectName"`
	RefObjectUUID                      string                 `json:"_refObjectUUID"`
	RefType                            string                 `json:"_type"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//PortfolioItemTheme auto-generated
type PortfolioItemTheme struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal"`
	Attachments                        []Attachment           `json:"Attachments"`
	Changesets                         map[string]interface{} `json:"Changesets"`
	Children                           map[string]interface{} `json:"Children"`
	Collaborators                      []User                 `json:"Collaborators"`
	CreationDate                       string                 `json:"CreationDate"`
	Description                        string                 `json:"Description"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount"`
	Discussion                         map[string]interface{} `json:"Discussion"`
	DragAndDropRank                    string                 `json:"DragAndDropRank"`
	FormattedID                        string                 `json:"FormattedID"`
	InvestmentCategory                 string                 `json:"InvestmentCategory"`
	JobSize                            float64                `json:"JobSize"`
	LastUpdateDate                     string                 `json:"LastUpdateDate"`
	LeafStoryCount                     float64                `json:"LeafStoryCount"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal"`
	Milestones                         map[string]interface{} `json:"Milestones"`
	Name                               string                 `json:"Name"`
	Notes                              string                 `json:"Notes"`
	ObjectID                           float64                `json:"ObjectID"`
	ObjectUUID                         string                 `json:"ObjectUUID"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate"`
	PortfolioItemType                  TypeDefinition         `json:"PortfolioItemType"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName"`
	Predecessors                       map[string]interface{} `json:"Predecessors"`
	Project                            Project                `json:"Project"`
	RROEValue                          float64                `json:"RROEValue"`
	RefinedEstimate                    float64                `json:"RefinedEstimate"`
	RevisionHistory                    RevisionHistory        `json:"RevisionHistory"`
	RiskScore                          float64                `json:"RiskScore"`
	StateChangedDate                   string                 `json:"StateChangedDate"`
	Subscription                       Subscription           `json:"Subscription"`
	Successors                         map[string]interface{} `json:"Successors"`
	Tags                               []Tag                  `json:"Tags"`
	TimeCriticality                    float64                `json:"TimeCriticality"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount"`
	UserBusinessValue                  float64                `json:"UserBusinessValue"`
	ValueScore                         float64                `json:"ValueScore"`
	VersionID                          string                 `json:"VersionId"`
	WSJFScore                          float64                `json:"WSJFScore"`
	Workspace                          Workspace              `json:"Workspace"`
	CreatedAt                          string                 `json:"_CreatedAt"`
	ObjectVersion                      string                 `json:"_objectVersion"`
	APIMajor                           string                 `json:"_rallyAPIMajor"`
	APIMinor                           string                 `json:"_rallyAPIMinor"`
	URI                                string                 `json:"_ref"`
	RefObjectName                      string                 `json:"_refObjectName"`
	RefObjectUUID                      string                 `json:"_refObjectUUID"`
	RefType                            string                 `json:"_type"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//PreliminaryEstimate auto-generated
type PreliminaryEstimate struct {
	CreationDate     string                 `json:"CreationDate"`
	Description      string                 `json:"Description"`
	Errors           []interface{}          `json:"Errors"`
	Name             string                 `json:"Name"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	RevisionHistory  RevisionHistory        `json:"RevisionHistory"`
	Subscription     Subscription           `json:"Subscription"`
	Value            float64                `json:"Value"`
	VersionID        string                 `json:"VersionId"`
	Warnings         []interface{}          `json:"Warnings"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Project auto-generated
type Project struct {
	BuildDefinitions []BuildDefinition      `json:"BuildDefinitions"`
	Children         []Project              `json:"Children"`
	CreationDate     string                 `json:"CreationDate"`
	Description      string                 `json:"Description"`
	Editors          []User                 `json:"Editors"`
	Iterations       []Iteration            `json:"Iterations"`
	Name             string                 `json:"Name"`
	Notes            string                 `json:"Notes"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Owner            User                   `json:"Owner"`
	Releases         []Release              `json:"Releases"`
	SchemaVersion    string                 `json:"SchemaVersion"`
	State            string                 `json:"State"`
	Subscription     Subscription           `json:"Subscription"`
	TeamMembers      []User                 `json:"TeamMembers"`
	VersionID        string                 `json:"VersionId"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//ProjectPermission auto-generated
type ProjectPermission struct {
	CustomObjectID   string                 `json:"CustomObjectID"`
	Name             string                 `json:"Name"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Project          Project                `json:"Project"`
	Role             string                 `json:"Role"`
	Subscription     Subscription           `json:"Subscription"`
	User             User                   `json:"User"`
	VersionID        string                 `json:"VersionId"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Release auto-generated
type Release struct {
	CreationDate                 string                 `json:"CreationDate"`
	GrossEstimateConversionRatio float64                `json:"GrossEstimateConversionRatio"`
	Name                         string                 `json:"Name"`
	Notes                        string                 `json:"Notes"`
	ObjectID                     float64                `json:"ObjectID"`
	ObjectUUID                   string                 `json:"ObjectUUID"`
	Project                      Project                `json:"Project"`
	ReleaseDate                  string                 `json:"ReleaseDate"`
	ReleaseStartDate             string                 `json:"ReleaseStartDate"`
	RevisionHistory              RevisionHistory        `json:"RevisionHistory"`
	State                        string                 `json:"State"`
	Subscription                 Subscription           `json:"Subscription"`
	Theme                        string                 `json:"Theme"`
	VersionID                    string                 `json:"VersionId"`
	Workspace                    Workspace              `json:"Workspace"`
	CreatedAt                    string                 `json:"_CreatedAt"`
	ObjectVersion                string                 `json:"_objectVersion"`
	APIMajor                     string                 `json:"_rallyAPIMajor"`
	APIMinor                     string                 `json:"_rallyAPIMinor"`
	URI                          string                 `json:"_ref"`
	RefObjectName                string                 `json:"_refObjectName"`
	RefObjectUUID                string                 `json:"_refObjectUUID"`
	RefType                      string                 `json:"_type"`
	CustomAttributes             map[string]interface{} `json:"CustomAttributes"`
}

//Revision auto-generated
type Revision struct {
	CreationDate     string                 `json:"CreationDate"`
	Description      string                 `json:"Description"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	RevisionHistory  RevisionHistory        `json:"RevisionHistory"`
	RevisionNumber   float64                `json:"RevisionNumber"`
	Subscription     Subscription           `json:"Subscription"`
	User             User                   `json:"User"`
	VersionID        string                 `json:"VersionId"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//RevisionHistory auto-generated
type RevisionHistory struct {
	CreationDate     string                 `json:"CreationDate"`
	Errors           []interface{}          `json:"Errors"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Revisions        []Revision             `json:"Revisions"`
	Subscription     Base                   `json:"Subscription"`
	VersionID        string                 `json:"VersionId"`
	Warnings         []interface{}          `json:"Warnings"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//State auto-generated
type State struct {
	CreationDate     string                 `json:"CreationDate"`
	Description      string                 `json:"Description"`
	Errors           []interface{}          `json:"Errors"`
	Name             string                 `json:"Name"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	OrderIndex       float64                `json:"OrderIndex"`
	RevisionHistory  RevisionHistory        `json:"RevisionHistory"`
	Subscription     Subscription           `json:"Subscription"`
	TypeDef          TypeDefinition         `json:"TypeDef"`
	VersionID        string                 `json:"VersionId"`
	WIPLimit         float64                `json:"WIPLimit"`
	Warnings         []interface{}          `json:"Warnings"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Subscription auto-generated
type Subscription struct {
	CreationDate            string                 `json:"CreationDate"`
	Errors                  []interface{}          `json:"Errors"`
	MaximumCustomUserFields float64                `json:"MaximumCustomUserFields"`
	MaximumProjects         float64                `json:"MaximumProjects"`
	Modules                 string                 `json:"Modules"`
	Name                    string                 `json:"Name"`
	ObjectID                float64                `json:"ObjectID"`
	ObjectUUID              string                 `json:"ObjectUUID"`
	PasswordExpirationDays  float64                `json:"PasswordExpirationDays"`
	PreviousPasswordCount   float64                `json:"PreviousPasswordCount"`
	RevisionHistory         RevisionHistory        `json:"RevisionHistory"`
	StoryHierarchyType      string                 `json:"StoryHierarchyType"`
	SubscriptionID          float64                `json:"SubscriptionID"`
	SubscriptionType        string                 `json:"SubscriptionType"`
	VersionID               string                 `json:"VersionId"`
	Warnings                []interface{}          `json:"Warnings"`
	Workspaces              []Workspace            `json:"Workspaces"`
	CreatedAt               string                 `json:"_CreatedAt"`
	ObjectVersion           string                 `json:"_objectVersion"`
	APIMajor                string                 `json:"_rallyAPIMajor"`
	APIMinor                string                 `json:"_rallyAPIMinor"`
	URI                     string                 `json:"_ref"`
	RefObjectName           string                 `json:"_refObjectName"`
	RefObjectUUID           string                 `json:"_refObjectUUID"`
	CustomAttributes        map[string]interface{} `json:"CustomAttributes"`
}

//Tag auto-generated
type Tag struct {
	CreationDate     string                 `json:"CreationDate"`
	Name             string                 `json:"Name"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Subscription     Subscription           `json:"Subscription"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//Task auto-generated
type Task struct {
	Attachments      []Attachment            `json:"Attachments"`
	Changesets       map[string]interface{}  `json:"Changesets"`
	CreationDate     string                  `json:"CreationDate"`
	Description      string                  `json:"Description"`
	Discussion       map[string]interface{}  `json:"Discussion"`
	DragAndDropRank  string                  `json:"DragAndDropRank"`
	Estimate         float64                 `json:"Estimate"`
	FormattedID      string                  `json:"FormattedID"`
	Iteration        Iteration               `json:"Iteration"`
	LastUpdateDate   string                  `json:"LastUpdateDate"`
	Milestones       map[string]interface{}  `json:"Milestones"`
	Name             string                  `json:"Name"`
	Notes            string                  `json:"Notes"`
	ObjectID         float64                 `json:"ObjectID"`
	ObjectUUID       string                  `json:"ObjectUUID"`
	Owner            User                    `json:"Owner"`
	Project          Project                 `json:"Project"`
	Release          Release                 `json:"Release"`
	RevisionHistory  RevisionHistory         `json:"RevisionHistory"`
	State            string                  `json:"State"`
	Subscription     Subscription            `json:"Subscription"`
	Tags             []Tag                   `json:"Tags"`
	TaskIndex        float64                 `json:"TaskIndex"`
	TimeSpent        float64                 `json:"TimeSpent"`
	ToDo             float64                 `json:"ToDo"`
	VersionID        string                  `json:"VersionId"`
	WorkProduct      HierarchicalRequirement `json:"WorkProduct"`
	Workspace        Workspace               `json:"Workspace"`
	CreatedAt        string                  `json:"_CreatedAt"`
	ObjectVersion    string                  `json:"_objectVersion"`
	APIMajor         string                  `json:"_rallyAPIMajor"`
	APIMinor         string                  `json:"_rallyAPIMinor"`
	URI              string                  `json:"_ref"`
	RefObjectName    string                  `json:"_refObjectName"`
	RefObjectUUID    string                  `json:"_refObjectUUID"`
	RefType          string                  `json:"_type"`
	CustomAttributes map[string]interface{}  `json:"CustomAttributes"`
}

//Theme auto-generated
type Theme struct {
	AcceptedLeafStoryCount             float64                `json:"AcceptedLeafStoryCount"`
	AcceptedLeafStoryPlanEstimateTotal float64                `json:"AcceptedLeafStoryPlanEstimateTotal"`
	Attachments                        []Attachment           `json:"Attachments"`
	Changesets                         map[string]interface{} `json:"Changesets"`
	Children                           map[string]interface{} `json:"Children"`
	Collaborators                      []User                 `json:"Collaborators"`
	CreationDate                       string                 `json:"CreationDate"`
	Description                        string                 `json:"Description"`
	DirectChildrenCount                float64                `json:"DirectChildrenCount"`
	Discussion                         map[string]interface{} `json:"Discussion"`
	DragAndDropRank                    string                 `json:"DragAndDropRank"`
	Errors                             []interface{}          `json:"Errors"`
	FormattedID                        string                 `json:"FormattedID"`
	InvestmentCategory                 string                 `json:"InvestmentCategory"`
	JobSize                            float64                `json:"JobSize"`
	LastUpdateDate                     string                 `json:"LastUpdateDate"`
	LeafStoryCount                     float64                `json:"LeafStoryCount"`
	LeafStoryPlanEstimateTotal         float64                `json:"LeafStoryPlanEstimateTotal"`
	Milestones                         map[string]interface{} `json:"Milestones"`
	Name                               string                 `json:"Name"`
	Notes                              string                 `json:"Notes"`
	ObjectID                           float64                `json:"ObjectID"`
	ObjectUUID                         string                 `json:"ObjectUUID"`
	PercentDoneByStoryCount            float64                `json:"PercentDoneByStoryCount"`
	PercentDoneByStoryPlanEstimate     float64                `json:"PercentDoneByStoryPlanEstimate"`
	PortfolioItemType                  TypeDefinition         `json:"PortfolioItemType"`
	PortfolioItemTypeName              string                 `json:"PortfolioItemTypeName"`
	Predecessors                       map[string]interface{} `json:"Predecessors"`
	Project                            Project                `json:"Project"`
	RROEValue                          float64                `json:"RROEValue"`
	RefinedEstimate                    float64                `json:"RefinedEstimate"`
	RevisionHistory                    RevisionHistory        `json:"RevisionHistory"`
	RiskScore                          float64                `json:"RiskScore"`
	StateChangedDate                   string                 `json:"StateChangedDate"`
	Subscription                       Subscription           `json:"Subscription"`
	Successors                         map[string]interface{} `json:"Successors"`
	Tags                               []Tag                  `json:"Tags"`
	TimeCriticality                    float64                `json:"TimeCriticality"`
	UnEstimatedLeafStoryCount          float64                `json:"UnEstimatedLeafStoryCount"`
	UserBusinessValue                  float64                `json:"UserBusinessValue"`
	ValueScore                         float64                `json:"ValueScore"`
	VersionID                          string                 `json:"VersionId"`
	WSJFScore                          float64                `json:"WSJFScore"`
	Warnings                           []interface{}          `json:"Warnings"`
	Workspace                          Workspace              `json:"Workspace"`
	CreatedAt                          string                 `json:"_CreatedAt"`
	ObjectVersion                      string                 `json:"_objectVersion"`
	APIMajor                           string                 `json:"_rallyAPIMajor"`
	APIMinor                           string                 `json:"_rallyAPIMinor"`
	URI                                string                 `json:"_ref"`
	RefObjectName                      string                 `json:"_refObjectName"`
	RefObjectUUID                      string                 `json:"_refObjectUUID"`
	RefType                            string                 `json:"_type"`
	CustomAttributes                   map[string]interface{} `json:"CustomAttributes"`
}

//TypeDefinition auto-generated
type TypeDefinition struct {
	Attributes       []AttributeDefinition  `json:"Attributes"`
	CreationDate     string                 `json:"CreationDate"`
	DisplayName      string                 `json:"DisplayName"`
	ElementName      string                 `json:"ElementName"`
	IDPrefix         string                 `json:"IDPrefix"`
	Name             string                 `json:"Name"`
	Note             string                 `json:"Note"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Ordinal          float64                `json:"Ordinal"`
	Parent           Base                   `json:"Parent"`
	RevisionHistory  RevisionHistory        `json:"RevisionHistory"`
	Subscription     Subscription           `json:"Subscription"`
	TypePath         string                 `json:"TypePath"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//User auto-generated
type User struct {
	CostCenter             string                 `json:"CostCenter"`
	CreationDate           string                 `json:"CreationDate"`
	Department             string                 `json:"Department"`
	DisplayName            string                 `json:"DisplayName"`
	EmailAddress           string                 `json:"EmailAddress"`
	Errors                 []interface{}          `json:"Errors"`
	FirstName              string                 `json:"FirstName"`
	LandingPage            string                 `json:"LandingPage"`
	LastLoginDate          string                 `json:"LastLoginDate"`
	LastName               string                 `json:"LastName"`
	LastPasswordUpdateDate string                 `json:"LastPasswordUpdateDate"`
	LastSystemTimeZoneName string                 `json:"LastSystemTimeZoneName"`
	ObjectID               float64                `json:"ObjectID"`
	ObjectUUID             string                 `json:"ObjectUUID"`
	OfficeLocation         string                 `json:"OfficeLocation"`
	RevisionHistory        RevisionHistory        `json:"RevisionHistory"`
	Role                   string                 `json:"Role"`
	Subscription           Subscription           `json:"Subscription"`
	SubscriptionID         float64                `json:"SubscriptionID"`
	SubscriptionPermission string                 `json:"SubscriptionPermission"`
	TeamMemberships        []Project              `json:"TeamMemberships"`
	UserName               string                 `json:"UserName"`
	UserPermissions        map[string]interface{} `json:"UserPermissions"`
	UserProfile            UserProfile            `json:"UserProfile"`
	VersionID              string                 `json:"VersionId"`
	Warnings               []interface{}          `json:"Warnings"`
	WorkspacePermission    string                 `json:"WorkspacePermission"`
	CreatedAt              string                 `json:"_CreatedAt"`
	ObjectVersion          string                 `json:"_objectVersion"`
	APIMajor               string                 `json:"_rallyAPIMajor"`
	APIMinor               string                 `json:"_rallyAPIMinor"`
	URI                    string                 `json:"_ref"`
	RefObjectName          string                 `json:"_refObjectName"`
	RefObjectUUID          string                 `json:"_refObjectUUID"`
	RefType                string                 `json:"_type"`
	CustomAttributes       map[string]interface{} `json:"CustomAttributes"`
}

//UserIterationCapacity auto-generated
type UserIterationCapacity struct {
	CreationDate     string                 `json:"CreationDate"`
	Iteration        Iteration              `json:"Iteration"`
	ObjectID         float64                `json:"ObjectID"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Project          Project                `json:"Project"`
	Subscription     Subscription           `json:"Subscription"`
	User             User                   `json:"User"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}

//UserProfile auto-generated
type UserProfile struct {
	CreationDate          string                 `json:"CreationDate"`
	DateFormat            string                 `json:"DateFormat"`
	DateTimeFormat        string                 `json:"DateTimeFormat"`
	Errors                []interface{}          `json:"Errors"`
	ObjectID              float64                `json:"ObjectID"`
	ObjectUUID            string                 `json:"ObjectUUID"`
	SessionTimeoutSeconds float64                `json:"SessionTimeoutSeconds"`
	Subscription          Subscription           `json:"Subscription"`
	TimeZone              string                 `json:"TimeZone"`
	VersionID             string                 `json:"VersionId"`
	Warnings              []interface{}          `json:"Warnings"`
	CreatedAt             string                 `json:"_CreatedAt"`
	ObjectVersion         string                 `json:"_objectVersion"`
	APIMajor              string                 `json:"_rallyAPIMajor"`
	APIMinor              string                 `json:"_rallyAPIMinor"`
	URI                   string                 `json:"_ref"`
	RefObjectUUID         string                 `json:"_refObjectUUID"`
	RefType               string                 `json:"_type"`
	CustomAttributes      map[string]interface{} `json:"CustomAttributes"`
}

//Workspace auto-generated
type Workspace struct {
	Children               []Project              `json:"Children"`
	CreationDate           string                 `json:"CreationDate"`
	Description            string                 `json:"Description"`
	Errors                 []interface{}          `json:"Errors"`
	Name                   string                 `json:"Name"`
	Notes                  string                 `json:"Notes"`
	ObjectID               float64                `json:"ObjectID"`
	ObjectUUID             string                 `json:"ObjectUUID"`
	Owner                  User                   `json:"Owner"`
	Projects               []Project              `json:"Projects"`
	RevisionHistory        RevisionHistory        `json:"RevisionHistory"`
	SchemaVersion          string                 `json:"SchemaVersion"`
	State                  string                 `json:"State"`
	Style                  string                 `json:"Style"`
	Subscription           Subscription           `json:"Subscription"`
	TypeDefinitions        []TypeDefinition       `json:"TypeDefinitions"`
	VersionID              string                 `json:"VersionId"`
	Warnings               []interface{}          `json:"Warnings"`
	WorkspaceConfiguration Base                   `json:"WorkspaceConfiguration"`
	CreatedAt              string                 `json:"_CreatedAt"`
	ObjectVersion          string                 `json:"_objectVersion"`
	APIMajor               string                 `json:"_rallyAPIMajor"`
	APIMinor               string                 `json:"_rallyAPIMinor"`
	URI                    string                 `json:"_ref"`
	RefObjectName          string                 `json:"_refObjectName"`
	RefObjectUUID          string                 `json:"_refObjectUUID"`
	RefType                string                 `json:"_type"`
	CustomAttributes       map[string]interface{} `json:"CustomAttributes"`
}

//WorkspaceConfiguration auto-generated
type WorkspaceConfiguration struct {
	CreationDate              string                 `json:"CreationDate"`
	DateFormat                string                 `json:"DateFormat"`
	DateTimeFormat            string                 `json:"DateTimeFormat"`
	Errors                    []interface{}          `json:"Errors"`
	IterationEstimateUnitName string                 `json:"IterationEstimateUnitName"`
	ObjectID                  float64                `json:"ObjectID"`
	ObjectUUID                string                 `json:"ObjectUUID"`
	ReleaseEstimateUnitName   string                 `json:"ReleaseEstimateUnitName"`
	Subscription              Subscription           `json:"Subscription"`
	TaskUnitName              string                 `json:"TaskUnitName"`
	TimeZone                  string                 `json:"TimeZone"`
	VersionID                 string                 `json:"VersionId"`
	Warnings                  []interface{}          `json:"Warnings"`
	WorkDays                  string                 `json:"WorkDays"`
	Workspace                 Workspace              `json:"Workspace"`
	CreatedAt                 string                 `json:"_CreatedAt"`
	ObjectVersion             string                 `json:"_objectVersion"`
	APIMajor                  string                 `json:"_rallyAPIMajor"`
	APIMinor                  string                 `json:"_rallyAPIMinor"`
	URI                       string                 `json:"_ref"`
	RefObjectUUID             string                 `json:"_refObjectUUID"`
	RefType                   string                 `json:"_type"`
	CustomAttributes          map[string]interface{} `json:"CustomAttributes"`
}

//WorkspacePermission auto-generated
type WorkspacePermission struct {
	CustomObjectID   string                 `json:"CustomObjectID"`
	Name             string                 `json:"Name"`
	ObjectUUID       string                 `json:"ObjectUUID"`
	Role             string                 `json:"Role"`
	Subscription     Subscription           `json:"Subscription"`
	User             User                   `json:"User"`
	VersionID        string                 `json:"VersionId"`
	Workspace        Workspace              `json:"Workspace"`
	CreatedAt        string                 `json:"_CreatedAt"`
	ObjectVersion    string                 `json:"_objectVersion"`
	APIMajor         string                 `json:"_rallyAPIMajor"`
	APIMinor         string                 `json:"_rallyAPIMinor"`
	URI              string                 `json:"_ref"`
	RefObjectName    string                 `json:"_refObjectName"`
	RefObjectUUID    string                 `json:"_refObjectUUID"`
	RefType          string                 `json:"_type"`
	CustomAttributes map[string]interface{} `json:"CustomAttributes"`
}
