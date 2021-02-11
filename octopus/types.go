package octopus

import (
	"time"
)

type ProgressResponse struct {
	Environments        []EnvironmentsSummary           `json:"Environments"`
	ChannelEnvironments map[string][]ChannelEnvironment `json:"ChannelEnvironments"`
	Releases            []Releases                      `json:"Releases"`
	Links               Links                           `json:"Links,omitempty"`
}

type EnvironmentsSummary struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type ChannelEnvironment struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type SelectedPackages struct {
	StepName             string `json:"StepName"`
	ActionName           string `json:"ActionName"`
	Version              string `json:"Version"`
	PackageReferenceName string `json:"PackageReferenceName"`
}
type ReleaseLinks struct {
	Self                             string `json:"Self"`
	Project                          string `json:"Project"`
	Progression                      string `json:"Progression"`
	Deployments                      string `json:"Deployments"`
	DeploymentTemplate               string `json:"DeploymentTemplate"`
	Artifacts                        string `json:"Artifacts"`
	ProjectVariableSnapshot          string `json:"ProjectVariableSnapshot"`
	ProjectDeploymentProcessSnapshot string `json:"ProjectDeploymentProcessSnapshot"`
	Web                              string `json:"Web"`
	SnapshotVariables                string `json:"SnapshotVariables"`
	Defects                          string `json:"Defects"`
	ReportDefect                     string `json:"ReportDefect"`
	ResolveDefect                    string `json:"ResolveDefect"`
	DeploymentPreviews               string `json:"DeploymentPreviews"`
}
type Release struct {
	ID                                 string             `json:"Id"`
	Version                            string             `json:"Version"`
	ChannelID                          string             `json:"ChannelId"`
	ReleaseNotes                       string             `json:"ReleaseNotes"`
	ProjectDeploymentProcessSnapshotID string             `json:"ProjectDeploymentProcessSnapshotId"`
	IgnoreChannelRules                 bool               `json:"IgnoreChannelRules"`
	BuildInformation                   []interface{}      `json:"BuildInformation"`
	Assembled                          time.Time          `json:"Assembled"`
	ProjectID                          string             `json:"ProjectId"`
	LibraryVariableSetSnapshotIds      []string           `json:"LibraryVariableSetSnapshotIds"`
	SelectedPackages                   []SelectedPackages `json:"SelectedPackages"`
	ProjectVariableSetSnapshotID       string             `json:"ProjectVariableSetSnapshotId"`
	SpaceID                            string             `json:"SpaceId"`
	Links                              Links              `json:"Links"`
}
type ActionPackages struct {
	DeploymentAction string `json:"DeploymentAction"`
	PackageReference string `json:"PackageReference"`
}
type Rules struct {
	ID             string           `json:"Id"`
	VersionRange   string           `json:"VersionRange"`
	Tag            string           `json:"Tag"`
	ActionPackages []ActionPackages `json:"ActionPackages"`
	Links          Links            `json:"Links"`
	Actions        []string         `json:"Actions"`
}
type ChannelLinks struct {
	Self     string `json:"Self"`
	Releases string `json:"Releases"`
	Project  string `json:"Project"`
}
type Channel struct {
	ID          string        `json:"Id"`
	Name        string        `json:"Name"`
	Description interface{}   `json:"Description"`
	ProjectID   string        `json:"ProjectId"`
	LifecycleID interface{}   `json:"LifecycleId"`
	IsDefault   bool          `json:"IsDefault"`
	Rules       []Rules       `json:"Rules"`
	TenantTags  []interface{} `json:"TenantTags"`
	SpaceID     string        `json:"SpaceId"`
	Links       Links         `json:"Links"`
}
type EnvironmentLinks struct {
	Self    string `json:"Self"`
	Release string `json:"Release"`
	Tenant  string `json:"Tenant"`
	Task    string `json:"Task"`
}
type Environment struct {
	ID                      string           `json:"Id"`
	ProjectID               string           `json:"ProjectId"`
	EnvironmentID           string           `json:"EnvironmentId"`
	ReleaseID               string           `json:"ReleaseId"`
	DeploymentID            string           `json:"DeploymentId"`
	TaskID                  string           `json:"TaskId"`
	TenantID                interface{}      `json:"TenantId"`
	ChannelID               string           `json:"ChannelId"`
	ReleaseVersion          string           `json:"ReleaseVersion"`
	Created                 time.Time        `json:"Created"`
	QueueTime               time.Time        `json:"QueueTime"`
	StartTime               time.Time        `json:"StartTime"`
	CompletedTime           time.Time        `json:"CompletedTime"`
	State                   string           `json:"State"`
	HasPendingInterruptions bool             `json:"HasPendingInterruptions"`
	HasWarningsOrErrors     bool             `json:"HasWarningsOrErrors"`
	ErrorMessage            string           `json:"ErrorMessage"`
	Duration                string           `json:"Duration"`
	IsCurrent               bool             `json:"IsCurrent"`
	IsPrevious              bool             `json:"IsPrevious"`
	IsCompleted             bool             `json:"IsCompleted"`
	Links                   EnvironmentLinks `json:"Links"`
}

type Deployments struct {
	Environments []Environment `json:"Environments"`
}
type Environments2 struct {
	ID                      string      `json:"Id"`
	ProjectID               string      `json:"ProjectId"`
	EnvironmentID           string      `json:"EnvironmentId"`
	ReleaseID               string      `json:"ReleaseId"`
	DeploymentID            string      `json:"DeploymentId"`
	TaskID                  string      `json:"TaskId"`
	TenantID                interface{} `json:"TenantId"`
	ChannelID               string      `json:"ChannelId"`
	ReleaseVersion          string      `json:"ReleaseVersion"`
	Created                 time.Time   `json:"Created"`
	QueueTime               time.Time   `json:"QueueTime"`
	StartTime               time.Time   `json:"StartTime"`
	CompletedTime           time.Time   `json:"CompletedTime"`
	State                   string      `json:"State"`
	HasPendingInterruptions bool        `json:"HasPendingInterruptions"`
	HasWarningsOrErrors     bool        `json:"HasWarningsOrErrors"`
	ErrorMessage            string      `json:"ErrorMessage"`
	Duration                string      `json:"Duration"`
	IsCurrent               bool        `json:"IsCurrent"`
	IsPrevious              bool        `json:"IsPrevious"`
	IsCompleted             bool        `json:"IsCompleted"`
	Links                   Links       `json:"Links"`
}

type Releases struct {
	Release                 Release                  `json:"Release"`
	Channel                 Channel                  `json:"Channel"`
	Deployments             map[string][]Environment `json:"Deployments,omitempty"`
	NextDeployments         []interface{}            `json:"NextDeployments"`
	HasUnresolvedDefect     bool                     `json:"HasUnresolvedDefect"`
	ReleaseRetentionPeriod  interface{}              `json:"ReleaseRetentionPeriod"`
	TentacleRetentionPeriod interface{}              `json:"TentacleRetentionPeriod"`
}
type Links struct {
}

//Project
type Project struct {
	ID                              string   `json:"Id"`
	VariableSetID                   string   `json:"VariableSetId"`
	DeploymentProcessID             string   `json:"DeploymentProcessId"`
	ClonedFromProjectID             string   `json:"ClonedFromProjectId"`
	DiscreteChannelRelease          bool     `json:"DiscreteChannelRelease"`
	IncludedLibraryVariableSetIds   []string `json:"IncludedLibraryVariableSetIds"`
	DefaultToSkipIfAlreadyInstalled bool     `json:"DefaultToSkipIfAlreadyInstalled"`
	TenantedDeploymentMode          string   `json:"TenantedDeploymentMode"`
	DefaultGuidedFailureMode        string   `json:"DefaultGuidedFailureMode"`
	VersioningStrategy              struct {
		Template     string `json:"Template"`
		DonorPackage struct {
			DeploymentAction string `json:"DeploymentAction"`
			PackageReference string `json:"PackageReference"`
		} `json:"DonorPackage"`
	} `json:"VersioningStrategy"`
	ReleaseCreationStrategy struct {
		ChannelID              string `json:"ChannelId"`
		ReleaseCreationPackage struct {
			DeploymentAction string `json:"DeploymentAction"`
			PackageReference string `json:"PackageReference"`
		} `json:"ReleaseCreationPackage"`
	} `json:"ReleaseCreationStrategy"`
	Templates []struct {
		ID           string `json:"Id"`
		Name         string `json:"Name"`
		Label        string `json:"Label"`
		HelpText     string `json:"HelpText"`
		DefaultValue struct {
			IsSensitive    bool   `json:"IsSensitive"`
			Value          string `json:"Value"`
			SensitiveValue struct {
				HasValue bool   `json:"HasValue"`
				NewValue string `json:"NewValue"`
			} `json:"SensitiveValue"`
		} `json:"DefaultValue"`
		DisplaySettings struct {
			AdditionalProp1 string `json:"additionalProp1"`
			AdditionalProp2 string `json:"additionalProp2"`
			AdditionalProp3 string `json:"additionalProp3"`
		} `json:"DisplaySettings"`
	} `json:"Templates"`
	AutoDeployReleaseOverrides []struct {
		EnvironmentID string `json:"EnvironmentId"`
		TenantID      string `json:"TenantId"`
		ReleaseID     string `json:"ReleaseId"`
	} `json:"AutoDeployReleaseOverrides"`
	ReleaseNotesTemplate      string `json:"ReleaseNotesTemplate"`
	DeploymentChangesTemplate string `json:"DeploymentChangesTemplate"`
	SpaceID                   string `json:"SpaceId"`
	ExtensionSettings         []struct {
		ExtensionID string `json:"ExtensionId"`
		Values      struct {
		} `json:"Values"`
	} `json:"ExtensionSettings"`
	Name                      string `json:"Name"`
	Slug                      string `json:"Slug"`
	Description               string `json:"Description"`
	IsDisabled                bool   `json:"IsDisabled"`
	ProjectGroupID            string `json:"ProjectGroupId"`
	LifecycleID               string `json:"LifecycleId"`
	AutoCreateRelease         bool   `json:"AutoCreateRelease"`
	ProjectConnectivityPolicy struct {
		SkipMachineBehavior         string   `json:"SkipMachineBehavior"`
		TargetRoles                 []string `json:"TargetRoles"`
		AllowDeploymentsToNoTargets bool     `json:"AllowDeploymentsToNoTargets"`
		ExcludeUnhealthyTargets     bool     `json:"ExcludeUnhealthyTargets"`
	} `json:"ProjectConnectivityPolicy"`
	LastModifiedOn time.Time `json:"LastModifiedOn"`
	LastModifiedBy string    `json:"LastModifiedBy"`
	Links          struct {
		AdditionalProp1 string `json:"additionalProp1"`
		AdditionalProp2 string `json:"additionalProp2"`
		AdditionalProp3 string `json:"additionalProp3"`
	} `json:"Links"`
}

type ChannelsResponse struct {
	ItemType       string `json:"ItemType"`
	TotalResults   int    `json:"TotalResults"`
	ItemsPerPage   int    `json:"ItemsPerPage"`
	NumberOfPages  int    `json:"NumberOfPages"`
	LastPageNumber int    `json:"LastPageNumber"`
	Items          []struct {
		ID          string        `json:"Id"`
		Name        string        `json:"Name"`
		Description interface{}   `json:"Description"`
		ProjectID   string        `json:"ProjectId"`
		LifecycleID interface{}   `json:"LifecycleId"`
		IsDefault   bool          `json:"IsDefault"`
		Rules       []interface{} `json:"Rules"`
		TenantTags  []interface{} `json:"TenantTags"`
		SpaceID     string        `json:"SpaceId"`
		Links       struct {
			Self     string `json:"Self"`
			Releases string `json:"Releases"`
			Project  string `json:"Project"`
		} `json:"Links"`
	} `json:"Items"`
	Links struct {
		Self        string `json:"Self"`
		Template    string `json:"Template"`
		PageAll     string `json:"Page.All"`
		PageCurrent string `json:"Page.Current"`
		PageLast    string `json:"Page.Last"`
	} `json:"Links"`
}

type Project2 struct {
	ID                              string      `json:"Id"`
	VariableSetID                   string      `json:"VariableSetId"`
	DeploymentProcessID             string      `json:"DeploymentProcessId"`
	ClonedFromProjectID             interface{} `json:"ClonedFromProjectId"`
	DiscreteChannelRelease          bool        `json:"DiscreteChannelRelease"`
	IncludedLibraryVariableSetIds   []string    `json:"IncludedLibraryVariableSetIds"`
	DefaultToSkipIfAlreadyInstalled bool        `json:"DefaultToSkipIfAlreadyInstalled"`
	TenantedDeploymentMode          string      `json:"TenantedDeploymentMode"`
	DefaultGuidedFailureMode        string      `json:"DefaultGuidedFailureMode"`
	VersioningStrategy              struct {
		Template           string      `json:"Template"`
		DonorPackage       interface{} `json:"DonorPackage"`
		DonorPackageStepID interface{} `json:"DonorPackageStepId"`
	} `json:"VersioningStrategy"`
	ReleaseCreationStrategy struct {
		ChannelID                    interface{} `json:"ChannelId"`
		ReleaseCreationPackage       interface{} `json:"ReleaseCreationPackage"`
		ReleaseCreationPackageStepID interface{} `json:"ReleaseCreationPackageStepId"`
	} `json:"ReleaseCreationStrategy"`
	Templates                  []interface{} `json:"Templates"`
	AutoDeployReleaseOverrides []interface{} `json:"AutoDeployReleaseOverrides"`
	ReleaseNotesTemplate       interface{}   `json:"ReleaseNotesTemplate"`
	DeploymentChangesTemplate  interface{}   `json:"DeploymentChangesTemplate"`
	SpaceID                    string        `json:"SpaceId"`
	ExtensionSettings          []interface{} `json:"ExtensionSettings"`
	Name                       string        `json:"Name"`
	Slug                       string        `json:"Slug"`
	Description                string        `json:"Description"`
	IsDisabled                 bool          `json:"IsDisabled"`
	ProjectGroupID             string        `json:"ProjectGroupId"`
	LifecycleID                string        `json:"LifecycleId"`
	AutoCreateRelease          bool          `json:"AutoCreateRelease"`
	ProjectConnectivityPolicy  struct {
		SkipMachineBehavior         string        `json:"SkipMachineBehavior"`
		TargetRoles                 []interface{} `json:"TargetRoles"`
		AllowDeploymentsToNoTargets bool          `json:"AllowDeploymentsToNoTargets"`
		ExcludeUnhealthyTargets     bool          `json:"ExcludeUnhealthyTargets"`
	} `json:"ProjectConnectivityPolicy"`
	Links struct {
		Self                                 string `json:"Self"`
		Releases                             string `json:"Releases"`
		Channels                             string `json:"Channels"`
		Triggers                             string `json:"Triggers"`
		ScheduledTriggers                    string `json:"ScheduledTriggers"`
		OrderChannels                        string `json:"OrderChannels"`
		Variables                            string `json:"Variables"`
		Progression                          string `json:"Progression"`
		RunbookTaskRunDashboardItemsTemplate string `json:"RunbookTaskRunDashboardItemsTemplate"`
		DeploymentProcess                    string `json:"DeploymentProcess"`
		Web                                  string `json:"Web"`
		Logo                                 string `json:"Logo"`
		Metadata                             string `json:"Metadata"`
		Runbooks                             string `json:"Runbooks"`
		RunbookSnapshots                     string `json:"RunbookSnapshots"`
	} `json:"Links"`
}