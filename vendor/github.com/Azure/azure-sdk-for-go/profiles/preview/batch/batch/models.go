// +build go1.9

// Copyright 2018 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package batch

import original "github.com/Azure/azure-sdk-for-go/services/batch/2018-03-01.6.1/batch"

type AccountClient = original.AccountClient

func NewAccountClient() AccountClient {
	return original.NewAccountClient()
}
func NewAccountClientWithBaseURI(baseURI string) AccountClient {
	return original.NewAccountClientWithBaseURI(baseURI)
}

type ApplicationClient = original.ApplicationClient

func NewApplicationClient() ApplicationClient {
	return original.NewApplicationClient()
}
func NewApplicationClientWithBaseURI(baseURI string) ApplicationClient {
	return original.NewApplicationClientWithBaseURI(baseURI)
}

type CertificateClient = original.CertificateClient

func NewCertificateClient() CertificateClient {
	return original.NewCertificateClient()
}
func NewCertificateClientWithBaseURI(baseURI string) CertificateClient {
	return original.NewCertificateClientWithBaseURI(baseURI)
}

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient

func New() BaseClient {
	return original.New()
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}

type ComputeNodeClient = original.ComputeNodeClient

func NewComputeNodeClient() ComputeNodeClient {
	return original.NewComputeNodeClient()
}
func NewComputeNodeClientWithBaseURI(baseURI string) ComputeNodeClient {
	return original.NewComputeNodeClientWithBaseURI(baseURI)
}

type FileClient = original.FileClient

func NewFileClient() FileClient {
	return original.NewFileClient()
}
func NewFileClientWithBaseURI(baseURI string) FileClient {
	return original.NewFileClientWithBaseURI(baseURI)
}

type JobClient = original.JobClient

func NewJobClient() JobClient {
	return original.NewJobClient()
}
func NewJobClientWithBaseURI(baseURI string) JobClient {
	return original.NewJobClientWithBaseURI(baseURI)
}

type JobScheduleClient = original.JobScheduleClient

func NewJobScheduleClient() JobScheduleClient {
	return original.NewJobScheduleClient()
}
func NewJobScheduleClientWithBaseURI(baseURI string) JobScheduleClient {
	return original.NewJobScheduleClientWithBaseURI(baseURI)
}

type AccessScope = original.AccessScope

const (
	Job AccessScope = original.Job
)

func PossibleAccessScopeValues() []AccessScope {
	return original.PossibleAccessScopeValues()
}

type AllocationState = original.AllocationState

const (
	Resizing AllocationState = original.Resizing
	Steady   AllocationState = original.Steady
	Stopping AllocationState = original.Stopping
)

func PossibleAllocationStateValues() []AllocationState {
	return original.PossibleAllocationStateValues()
}

type AutoUserScope = original.AutoUserScope

const (
	Pool AutoUserScope = original.Pool
	Task AutoUserScope = original.Task
)

func PossibleAutoUserScopeValues() []AutoUserScope {
	return original.PossibleAutoUserScopeValues()
}

type CachingType = original.CachingType

const (
	None      CachingType = original.None
	ReadOnly  CachingType = original.ReadOnly
	ReadWrite CachingType = original.ReadWrite
)

func PossibleCachingTypeValues() []CachingType {
	return original.PossibleCachingTypeValues()
}

type CertificateFormat = original.CertificateFormat

const (
	Cer CertificateFormat = original.Cer
	Pfx CertificateFormat = original.Pfx
)

func PossibleCertificateFormatValues() []CertificateFormat {
	return original.PossibleCertificateFormatValues()
}

type CertificateState = original.CertificateState

const (
	Active       CertificateState = original.Active
	DeleteFailed CertificateState = original.DeleteFailed
	Deleting     CertificateState = original.Deleting
)

func PossibleCertificateStateValues() []CertificateState {
	return original.PossibleCertificateStateValues()
}

type CertificateStoreLocation = original.CertificateStoreLocation

const (
	CurrentUser  CertificateStoreLocation = original.CurrentUser
	LocalMachine CertificateStoreLocation = original.LocalMachine
)

func PossibleCertificateStoreLocationValues() []CertificateStoreLocation {
	return original.PossibleCertificateStoreLocationValues()
}

type CertificateVisibility = original.CertificateVisibility

const (
	CertificateVisibilityRemoteUser CertificateVisibility = original.CertificateVisibilityRemoteUser
	CertificateVisibilityStartTask  CertificateVisibility = original.CertificateVisibilityStartTask
	CertificateVisibilityTask       CertificateVisibility = original.CertificateVisibilityTask
)

func PossibleCertificateVisibilityValues() []CertificateVisibility {
	return original.PossibleCertificateVisibilityValues()
}

type ComputeNodeDeallocationOption = original.ComputeNodeDeallocationOption

const (
	Requeue        ComputeNodeDeallocationOption = original.Requeue
	RetainedData   ComputeNodeDeallocationOption = original.RetainedData
	TaskCompletion ComputeNodeDeallocationOption = original.TaskCompletion
	Terminate      ComputeNodeDeallocationOption = original.Terminate
)

func PossibleComputeNodeDeallocationOptionValues() []ComputeNodeDeallocationOption {
	return original.PossibleComputeNodeDeallocationOptionValues()
}

type ComputeNodeFillType = original.ComputeNodeFillType

const (
	Pack   ComputeNodeFillType = original.Pack
	Spread ComputeNodeFillType = original.Spread
)

func PossibleComputeNodeFillTypeValues() []ComputeNodeFillType {
	return original.PossibleComputeNodeFillTypeValues()
}

type ComputeNodeRebootOption = original.ComputeNodeRebootOption

const (
	ComputeNodeRebootOptionRequeue        ComputeNodeRebootOption = original.ComputeNodeRebootOptionRequeue
	ComputeNodeRebootOptionRetainedData   ComputeNodeRebootOption = original.ComputeNodeRebootOptionRetainedData
	ComputeNodeRebootOptionTaskCompletion ComputeNodeRebootOption = original.ComputeNodeRebootOptionTaskCompletion
	ComputeNodeRebootOptionTerminate      ComputeNodeRebootOption = original.ComputeNodeRebootOptionTerminate
)

func PossibleComputeNodeRebootOptionValues() []ComputeNodeRebootOption {
	return original.PossibleComputeNodeRebootOptionValues()
}

type ComputeNodeReimageOption = original.ComputeNodeReimageOption

const (
	ComputeNodeReimageOptionRequeue        ComputeNodeReimageOption = original.ComputeNodeReimageOptionRequeue
	ComputeNodeReimageOptionRetainedData   ComputeNodeReimageOption = original.ComputeNodeReimageOptionRetainedData
	ComputeNodeReimageOptionTaskCompletion ComputeNodeReimageOption = original.ComputeNodeReimageOptionTaskCompletion
	ComputeNodeReimageOptionTerminate      ComputeNodeReimageOption = original.ComputeNodeReimageOptionTerminate
)

func PossibleComputeNodeReimageOptionValues() []ComputeNodeReimageOption {
	return original.PossibleComputeNodeReimageOptionValues()
}

type ComputeNodeState = original.ComputeNodeState

const (
	Creating            ComputeNodeState = original.Creating
	Idle                ComputeNodeState = original.Idle
	LeavingPool         ComputeNodeState = original.LeavingPool
	Offline             ComputeNodeState = original.Offline
	Preempted           ComputeNodeState = original.Preempted
	Rebooting           ComputeNodeState = original.Rebooting
	Reimaging           ComputeNodeState = original.Reimaging
	Running             ComputeNodeState = original.Running
	Starting            ComputeNodeState = original.Starting
	StartTaskFailed     ComputeNodeState = original.StartTaskFailed
	Unknown             ComputeNodeState = original.Unknown
	Unusable            ComputeNodeState = original.Unusable
	WaitingForStartTask ComputeNodeState = original.WaitingForStartTask
)

func PossibleComputeNodeStateValues() []ComputeNodeState {
	return original.PossibleComputeNodeStateValues()
}

type DependencyAction = original.DependencyAction

const (
	Block   DependencyAction = original.Block
	Satisfy DependencyAction = original.Satisfy
)

func PossibleDependencyActionValues() []DependencyAction {
	return original.PossibleDependencyActionValues()
}

type DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOption

const (
	DisableComputeNodeSchedulingOptionRequeue        DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionRequeue
	DisableComputeNodeSchedulingOptionTaskCompletion DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionTaskCompletion
	DisableComputeNodeSchedulingOptionTerminate      DisableComputeNodeSchedulingOption = original.DisableComputeNodeSchedulingOptionTerminate
)

func PossibleDisableComputeNodeSchedulingOptionValues() []DisableComputeNodeSchedulingOption {
	return original.PossibleDisableComputeNodeSchedulingOptionValues()
}

type DisableJobOption = original.DisableJobOption

const (
	DisableJobOptionRequeue   DisableJobOption = original.DisableJobOptionRequeue
	DisableJobOptionTerminate DisableJobOption = original.DisableJobOptionTerminate
	DisableJobOptionWait      DisableJobOption = original.DisableJobOptionWait
)

func PossibleDisableJobOptionValues() []DisableJobOption {
	return original.PossibleDisableJobOptionValues()
}

type ElevationLevel = original.ElevationLevel

const (
	Admin    ElevationLevel = original.Admin
	NonAdmin ElevationLevel = original.NonAdmin
)

func PossibleElevationLevelValues() []ElevationLevel {
	return original.PossibleElevationLevelValues()
}

type ErrorCategory = original.ErrorCategory

const (
	ServerError ErrorCategory = original.ServerError
	UserError   ErrorCategory = original.UserError
)

func PossibleErrorCategoryValues() []ErrorCategory {
	return original.PossibleErrorCategoryValues()
}

type InboundEndpointProtocol = original.InboundEndpointProtocol

const (
	TCP InboundEndpointProtocol = original.TCP
	UDP InboundEndpointProtocol = original.UDP
)

func PossibleInboundEndpointProtocolValues() []InboundEndpointProtocol {
	return original.PossibleInboundEndpointProtocolValues()
}

type JobAction = original.JobAction

const (
	JobActionDisable   JobAction = original.JobActionDisable
	JobActionNone      JobAction = original.JobActionNone
	JobActionTerminate JobAction = original.JobActionTerminate
)

func PossibleJobActionValues() []JobAction {
	return original.PossibleJobActionValues()
}

type JobPreparationTaskState = original.JobPreparationTaskState

const (
	JobPreparationTaskStateCompleted JobPreparationTaskState = original.JobPreparationTaskStateCompleted
	JobPreparationTaskStateRunning   JobPreparationTaskState = original.JobPreparationTaskStateRunning
)

func PossibleJobPreparationTaskStateValues() []JobPreparationTaskState {
	return original.PossibleJobPreparationTaskStateValues()
}

type JobReleaseTaskState = original.JobReleaseTaskState

const (
	JobReleaseTaskStateCompleted JobReleaseTaskState = original.JobReleaseTaskStateCompleted
	JobReleaseTaskStateRunning   JobReleaseTaskState = original.JobReleaseTaskStateRunning
)

func PossibleJobReleaseTaskStateValues() []JobReleaseTaskState {
	return original.PossibleJobReleaseTaskStateValues()
}

type JobScheduleState = original.JobScheduleState

const (
	JobScheduleStateActive      JobScheduleState = original.JobScheduleStateActive
	JobScheduleStateCompleted   JobScheduleState = original.JobScheduleStateCompleted
	JobScheduleStateDeleting    JobScheduleState = original.JobScheduleStateDeleting
	JobScheduleStateDisabled    JobScheduleState = original.JobScheduleStateDisabled
	JobScheduleStateTerminating JobScheduleState = original.JobScheduleStateTerminating
)

func PossibleJobScheduleStateValues() []JobScheduleState {
	return original.PossibleJobScheduleStateValues()
}

type JobState = original.JobState

const (
	JobStateActive      JobState = original.JobStateActive
	JobStateCompleted   JobState = original.JobStateCompleted
	JobStateDeleting    JobState = original.JobStateDeleting
	JobStateDisabled    JobState = original.JobStateDisabled
	JobStateDisabling   JobState = original.JobStateDisabling
	JobStateEnabling    JobState = original.JobStateEnabling
	JobStateTerminating JobState = original.JobStateTerminating
)

func PossibleJobStateValues() []JobState {
	return original.PossibleJobStateValues()
}

type NetworkSecurityGroupRuleAccess = original.NetworkSecurityGroupRuleAccess

const (
	Allow NetworkSecurityGroupRuleAccess = original.Allow
	Deny  NetworkSecurityGroupRuleAccess = original.Deny
)

func PossibleNetworkSecurityGroupRuleAccessValues() []NetworkSecurityGroupRuleAccess {
	return original.PossibleNetworkSecurityGroupRuleAccessValues()
}

type OnAllTasksComplete = original.OnAllTasksComplete

const (
	NoAction     OnAllTasksComplete = original.NoAction
	TerminateJob OnAllTasksComplete = original.TerminateJob
)

func PossibleOnAllTasksCompleteValues() []OnAllTasksComplete {
	return original.PossibleOnAllTasksCompleteValues()
}

type OnTaskFailure = original.OnTaskFailure

const (
	OnTaskFailureNoAction                    OnTaskFailure = original.OnTaskFailureNoAction
	OnTaskFailurePerformExitOptionsJobAction OnTaskFailure = original.OnTaskFailurePerformExitOptionsJobAction
)

func PossibleOnTaskFailureValues() []OnTaskFailure {
	return original.PossibleOnTaskFailureValues()
}

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}

type OutputFileUploadCondition = original.OutputFileUploadCondition

const (
	OutputFileUploadConditionTaskCompletion OutputFileUploadCondition = original.OutputFileUploadConditionTaskCompletion
	OutputFileUploadConditionTaskFailure    OutputFileUploadCondition = original.OutputFileUploadConditionTaskFailure
	OutputFileUploadConditionTaskSuccess    OutputFileUploadCondition = original.OutputFileUploadConditionTaskSuccess
)

func PossibleOutputFileUploadConditionValues() []OutputFileUploadCondition {
	return original.PossibleOutputFileUploadConditionValues()
}

type PoolLifetimeOption = original.PoolLifetimeOption

const (
	PoolLifetimeOptionJob         PoolLifetimeOption = original.PoolLifetimeOptionJob
	PoolLifetimeOptionJobSchedule PoolLifetimeOption = original.PoolLifetimeOptionJobSchedule
)

func PossiblePoolLifetimeOptionValues() []PoolLifetimeOption {
	return original.PossiblePoolLifetimeOptionValues()
}

type PoolState = original.PoolState

const (
	PoolStateActive    PoolState = original.PoolStateActive
	PoolStateDeleting  PoolState = original.PoolStateDeleting
	PoolStateUpgrading PoolState = original.PoolStateUpgrading
)

func PossiblePoolStateValues() []PoolState {
	return original.PossiblePoolStateValues()
}

type SchedulingState = original.SchedulingState

const (
	Disabled SchedulingState = original.Disabled
	Enabled  SchedulingState = original.Enabled
)

func PossibleSchedulingStateValues() []SchedulingState {
	return original.PossibleSchedulingStateValues()
}

type StartTaskState = original.StartTaskState

const (
	StartTaskStateCompleted StartTaskState = original.StartTaskStateCompleted
	StartTaskStateRunning   StartTaskState = original.StartTaskStateRunning
)

func PossibleStartTaskStateValues() []StartTaskState {
	return original.PossibleStartTaskStateValues()
}

type StorageAccountType = original.StorageAccountType

const (
	PremiumLRS  StorageAccountType = original.PremiumLRS
	StandardLRS StorageAccountType = original.StandardLRS
)

func PossibleStorageAccountTypeValues() []StorageAccountType {
	return original.PossibleStorageAccountTypeValues()
}

type SubtaskState = original.SubtaskState

const (
	SubtaskStateCompleted SubtaskState = original.SubtaskStateCompleted
	SubtaskStatePreparing SubtaskState = original.SubtaskStatePreparing
	SubtaskStateRunning   SubtaskState = original.SubtaskStateRunning
)

func PossibleSubtaskStateValues() []SubtaskState {
	return original.PossibleSubtaskStateValues()
}

type TaskAddStatus = original.TaskAddStatus

const (
	TaskAddStatusClientError TaskAddStatus = original.TaskAddStatusClientError
	TaskAddStatusServerError TaskAddStatus = original.TaskAddStatusServerError
	TaskAddStatusSuccess     TaskAddStatus = original.TaskAddStatusSuccess
)

func PossibleTaskAddStatusValues() []TaskAddStatus {
	return original.PossibleTaskAddStatusValues()
}

type TaskCountValidationStatus = original.TaskCountValidationStatus

const (
	Unvalidated TaskCountValidationStatus = original.Unvalidated
	Validated   TaskCountValidationStatus = original.Validated
)

func PossibleTaskCountValidationStatusValues() []TaskCountValidationStatus {
	return original.PossibleTaskCountValidationStatusValues()
}

type TaskExecutionResult = original.TaskExecutionResult

const (
	Failure TaskExecutionResult = original.Failure
	Success TaskExecutionResult = original.Success
)

func PossibleTaskExecutionResultValues() []TaskExecutionResult {
	return original.PossibleTaskExecutionResultValues()
}

type TaskState = original.TaskState

const (
	TaskStateActive    TaskState = original.TaskStateActive
	TaskStateCompleted TaskState = original.TaskStateCompleted
	TaskStatePreparing TaskState = original.TaskStatePreparing
	TaskStateRunning   TaskState = original.TaskStateRunning
)

func PossibleTaskStateValues() []TaskState {
	return original.PossibleTaskStateValues()
}

type AccountListNodeAgentSkusResult = original.AccountListNodeAgentSkusResult
type AccountListNodeAgentSkusResultIterator = original.AccountListNodeAgentSkusResultIterator
type AccountListNodeAgentSkusResultPage = original.AccountListNodeAgentSkusResultPage
type AffinityInformation = original.AffinityInformation
type ApplicationListResult = original.ApplicationListResult
type ApplicationListResultIterator = original.ApplicationListResultIterator
type ApplicationListResultPage = original.ApplicationListResultPage
type ApplicationPackageReference = original.ApplicationPackageReference
type ApplicationSummary = original.ApplicationSummary
type AuthenticationTokenSettings = original.AuthenticationTokenSettings
type AutoPoolSpecification = original.AutoPoolSpecification
type AutoScaleRun = original.AutoScaleRun
type AutoScaleRunError = original.AutoScaleRunError
type AutoUserSpecification = original.AutoUserSpecification
type Certificate = original.Certificate
type CertificateAddParameter = original.CertificateAddParameter
type CertificateListResult = original.CertificateListResult
type CertificateListResultIterator = original.CertificateListResultIterator
type CertificateListResultPage = original.CertificateListResultPage
type CertificateReference = original.CertificateReference
type CloudJob = original.CloudJob
type CloudJobListPreparationAndReleaseTaskStatusResult = original.CloudJobListPreparationAndReleaseTaskStatusResult
type CloudJobListPreparationAndReleaseTaskStatusResultIterator = original.CloudJobListPreparationAndReleaseTaskStatusResultIterator
type CloudJobListPreparationAndReleaseTaskStatusResultPage = original.CloudJobListPreparationAndReleaseTaskStatusResultPage
type CloudJobListResult = original.CloudJobListResult
type CloudJobListResultIterator = original.CloudJobListResultIterator
type CloudJobListResultPage = original.CloudJobListResultPage
type CloudJobSchedule = original.CloudJobSchedule
type CloudJobScheduleListResult = original.CloudJobScheduleListResult
type CloudJobScheduleListResultIterator = original.CloudJobScheduleListResultIterator
type CloudJobScheduleListResultPage = original.CloudJobScheduleListResultPage
type CloudPool = original.CloudPool
type CloudPoolListResult = original.CloudPoolListResult
type CloudPoolListResultIterator = original.CloudPoolListResultIterator
type CloudPoolListResultPage = original.CloudPoolListResultPage
type CloudServiceConfiguration = original.CloudServiceConfiguration
type CloudTask = original.CloudTask
type CloudTaskListResult = original.CloudTaskListResult
type CloudTaskListResultIterator = original.CloudTaskListResultIterator
type CloudTaskListResultPage = original.CloudTaskListResultPage
type CloudTaskListSubtasksResult = original.CloudTaskListSubtasksResult
type ComputeNode = original.ComputeNode
type ComputeNodeEndpointConfiguration = original.ComputeNodeEndpointConfiguration
type ComputeNodeError = original.ComputeNodeError
type ComputeNodeGetRemoteLoginSettingsResult = original.ComputeNodeGetRemoteLoginSettingsResult
type ComputeNodeInformation = original.ComputeNodeInformation
type ComputeNodeListResult = original.ComputeNodeListResult
type ComputeNodeListResultIterator = original.ComputeNodeListResultIterator
type ComputeNodeListResultPage = original.ComputeNodeListResultPage
type ComputeNodeUser = original.ComputeNodeUser
type ContainerConfiguration = original.ContainerConfiguration
type ContainerRegistry = original.ContainerRegistry
type DataDisk = original.DataDisk
type DeleteCertificateError = original.DeleteCertificateError
type EnvironmentSetting = original.EnvironmentSetting
type Error = original.Error
type ErrorDetail = original.ErrorDetail
type ErrorMessage = original.ErrorMessage
type ExitCodeMapping = original.ExitCodeMapping
type ExitCodeRangeMapping = original.ExitCodeRangeMapping
type ExitConditions = original.ExitConditions
type ExitOptions = original.ExitOptions
type FileProperties = original.FileProperties
type ImageReference = original.ImageReference
type InboundEndpoint = original.InboundEndpoint
type InboundNATPool = original.InboundNATPool
type JobAddParameter = original.JobAddParameter
type JobConstraints = original.JobConstraints
type JobDisableParameter = original.JobDisableParameter
type JobExecutionInformation = original.JobExecutionInformation
type JobManagerTask = original.JobManagerTask
type JobPatchParameter = original.JobPatchParameter
type JobPreparationAndReleaseTaskExecutionInformation = original.JobPreparationAndReleaseTaskExecutionInformation
type JobPreparationTask = original.JobPreparationTask
type JobPreparationTaskExecutionInformation = original.JobPreparationTaskExecutionInformation
type JobReleaseTask = original.JobReleaseTask
type JobReleaseTaskExecutionInformation = original.JobReleaseTaskExecutionInformation
type JobScheduleAddParameter = original.JobScheduleAddParameter
type JobScheduleExecutionInformation = original.JobScheduleExecutionInformation
type JobSchedulePatchParameter = original.JobSchedulePatchParameter
type JobScheduleStatistics = original.JobScheduleStatistics
type JobScheduleUpdateParameter = original.JobScheduleUpdateParameter
type JobSchedulingError = original.JobSchedulingError
type JobSpecification = original.JobSpecification
type JobStatistics = original.JobStatistics
type JobTerminateParameter = original.JobTerminateParameter
type JobUpdateParameter = original.JobUpdateParameter
type LinuxUserConfiguration = original.LinuxUserConfiguration
type MetadataItem = original.MetadataItem
type MultiInstanceSettings = original.MultiInstanceSettings
type NameValuePair = original.NameValuePair
type NetworkConfiguration = original.NetworkConfiguration
type NetworkSecurityGroupRule = original.NetworkSecurityGroupRule
type NodeAgentSku = original.NodeAgentSku
type NodeCounts = original.NodeCounts
type NodeDisableSchedulingParameter = original.NodeDisableSchedulingParameter
type NodeFile = original.NodeFile
type NodeFileListResult = original.NodeFileListResult
type NodeFileListResultIterator = original.NodeFileListResultIterator
type NodeFileListResultPage = original.NodeFileListResultPage
type NodeRebootParameter = original.NodeRebootParameter
type NodeReimageParameter = original.NodeReimageParameter
type NodeRemoveParameter = original.NodeRemoveParameter
type NodeUpdateUserParameter = original.NodeUpdateUserParameter
type OSDisk = original.OSDisk
type OutputFile = original.OutputFile
type OutputFileBlobContainerDestination = original.OutputFileBlobContainerDestination
type OutputFileDestination = original.OutputFileDestination
type OutputFileUploadOptions = original.OutputFileUploadOptions
type PoolAddParameter = original.PoolAddParameter
type PoolEnableAutoScaleParameter = original.PoolEnableAutoScaleParameter
type PoolEndpointConfiguration = original.PoolEndpointConfiguration
type PoolEvaluateAutoScaleParameter = original.PoolEvaluateAutoScaleParameter
type PoolInformation = original.PoolInformation
type PoolListUsageMetricsResult = original.PoolListUsageMetricsResult
type PoolListUsageMetricsResultIterator = original.PoolListUsageMetricsResultIterator
type PoolListUsageMetricsResultPage = original.PoolListUsageMetricsResultPage
type PoolNodeCounts = original.PoolNodeCounts
type PoolNodeCountsListResult = original.PoolNodeCountsListResult
type PoolNodeCountsListResultIterator = original.PoolNodeCountsListResultIterator
type PoolNodeCountsListResultPage = original.PoolNodeCountsListResultPage
type PoolPatchParameter = original.PoolPatchParameter
type PoolResizeParameter = original.PoolResizeParameter
type PoolSpecification = original.PoolSpecification
type PoolStatistics = original.PoolStatistics
type PoolUpdatePropertiesParameter = original.PoolUpdatePropertiesParameter
type PoolUpgradeOSParameter = original.PoolUpgradeOSParameter
type PoolUsageMetrics = original.PoolUsageMetrics
type ReadCloser = original.ReadCloser
type RecentJob = original.RecentJob
type ResizeError = original.ResizeError
type ResourceFile = original.ResourceFile
type ResourceStatistics = original.ResourceStatistics
type Schedule = original.Schedule
type StartTask = original.StartTask
type StartTaskInformation = original.StartTaskInformation
type SubtaskInformation = original.SubtaskInformation
type TaskAddCollectionParameter = original.TaskAddCollectionParameter
type TaskAddCollectionResult = original.TaskAddCollectionResult
type TaskAddParameter = original.TaskAddParameter
type TaskAddResult = original.TaskAddResult
type TaskConstraints = original.TaskConstraints
type TaskContainerExecutionInformation = original.TaskContainerExecutionInformation
type TaskContainerSettings = original.TaskContainerSettings
type TaskCounts = original.TaskCounts
type TaskDependencies = original.TaskDependencies
type TaskExecutionInformation = original.TaskExecutionInformation
type TaskFailureInformation = original.TaskFailureInformation
type TaskIDRange = original.TaskIDRange
type TaskInformation = original.TaskInformation
type TaskSchedulingPolicy = original.TaskSchedulingPolicy
type TaskStatistics = original.TaskStatistics
type TaskUpdateParameter = original.TaskUpdateParameter
type UploadBatchServiceLogsConfiguration = original.UploadBatchServiceLogsConfiguration
type UploadBatchServiceLogsResult = original.UploadBatchServiceLogsResult
type UsageStatistics = original.UsageStatistics
type UserAccount = original.UserAccount
type UserIdentity = original.UserIdentity
type VirtualMachineConfiguration = original.VirtualMachineConfiguration
type WindowsConfiguration = original.WindowsConfiguration
type PoolClient = original.PoolClient

func NewPoolClient() PoolClient {
	return original.NewPoolClient()
}
func NewPoolClientWithBaseURI(baseURI string) PoolClient {
	return original.NewPoolClientWithBaseURI(baseURI)
}

type TaskClient = original.TaskClient

func NewTaskClient() TaskClient {
	return original.NewTaskClient()
}
func NewTaskClientWithBaseURI(baseURI string) TaskClient {
	return original.NewTaskClientWithBaseURI(baseURI)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
