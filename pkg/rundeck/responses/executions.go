package responses

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
)

// JobExecutionsResponse is the response for listing the executions for a job
type JobExecutionsResponse ListRunningExecutionsResponse

func (a JobExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a JobExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a JobExecutionsResponse) deprecated() bool { return false }

// ListRunningExecutionsResponseTestFile is the test data for JobExecutionResponse
const ListRunningExecutionsResponseTestFile = "executions.json"

// ListRunningExecutionsResponse is the response for listing the running executions for a project
type ListRunningExecutionsResponse struct {
	Paging     PagingResponse      `json:"paging"`
	Executions []ExecutionResponse `json:"executions"`
}

func (a ListRunningExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ListRunningExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a ListRunningExecutionsResponse) deprecated() bool { return false }

// FromReader returns an ListRunningExecutionsResponse from an io.Reader
func (a *ListRunningExecutionsResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ListRunningExectionsResponse from a file
func (a *ListRunningExecutionsResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ListRunningExecutionsResponse from a byte slice
func (a *ListRunningExecutionsResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionResponseTestFile is the test data for ExecutionResponse
const ExecutionResponseTestFile = "execution.json"

// ExecutionResponse represents an individual execution response
type ExecutionResponse struct {
	ID           int    `json:"id"`
	HRef         string `json:"href"`
	Permalink    string `json:"permalink"`
	Status       string `json:"status"`
	CustomStatus string `json:"customStatus"`
	Project      string `json:"project"`
	User         string `json:"user"`
	ServerUUID   string `json:"serverUUID"`
	DateStarted  struct {
		UnixTime int64     `json:"unixtime"`
		Date     *JSONTime `json:"date"`
	} `json:"date-started"`
	DateEnded struct {
		UnixTime int64     `json:"unixtime"`
		Date     *JSONTime `json:"date"`
	} `json:"date-ended"`
	Job             ExecutionJobEntryResponse `json:"job"`
	Description     string                    `json:"description"`
	ArgString       string                    `json:"argstring"`
	SuccessfulNodes []string                  `json:"successfulNodes"`
	FailedNodes     []string                  `json:"failedNodes"`
}

func (a ExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionResponse) deprecated() bool { return false }

// FromReader returns an ExecutionResponse from an io.Reader
func (a *ExecutionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExectionResponse from a file
func (a *ExecutionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionResponse from a byte slice
func (a *ExecutionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionJobEntryResponse represents an individual job execution entry response
type ExecutionJobEntryResponse struct {
	ID              string            `json:"id"`
	AverageDuration int64             `json:"averageDuration"`
	Name            string            `json:"name"`
	Group           string            `json:"group"`
	Project         string            `json:"project"`
	Description     string            `json:"description"`
	HRef            string            `json:"href"`
	Permalink       string            `json:"permalink"`
	Options         map[string]string `json:"options"`
}

func (a ExecutionJobEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionJobEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionJobEntryResponse) deprecated() bool { return false }

// ExecutionInputFileResponse is an individual execution input file entry response
type ExecutionInputFileResponse struct {
	ID             string    `json:"id"`
	User           string    `json:"user"`
	FileState      string    `json:"fileState"`
	SHA            string    `json:"sha"`
	JobID          string    `json:"jobId"`
	DateCreated    *JSONTime `json:"dateCreated"`
	ServerNodeUUID string    `json:"serverNodeUUID"`
	FileName       string    `json:"fileName"`
	Size           int64     `json:"size"`
	ExpirationDate *JSONTime `json:"expirationDate"`
	ExecID         int       `json:"execId"`
}

func (a ExecutionInputFileResponse) minVersion() int  { return 19 }
func (a ExecutionInputFileResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionInputFileResponse) deprecated() bool { return false }

// ExecutionInputFilesResponseTestFile is test data for an ExecutionInputFileResponse
const ExecutionInputFilesResponseTestFile = "execution_input_files.json"

// ExecutionInputFilesResponse is a response for listing execution input files
type ExecutionInputFilesResponse struct {
	Files []ExecutionInputFileResponse `json:"files"`
}

func (a ExecutionInputFilesResponse) minVersion() int  { return 19 }
func (a ExecutionInputFilesResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionInputFilesResponse) deprecated() bool { return false }

// FromReader returns an ExecutionInputFilesResponse from an io.Reader
func (a *ExecutionInputFilesResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExecutionInputFilesResponse from a file
func (a *ExecutionInputFilesResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionInputFilesResponse from a byte slice
func (a *ExecutionInputFilesResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// BulkDeleteExecutionsResponseTestFile is test data for an ExecutionInputFileResponse
const BulkDeleteExecutionsResponseTestFile = "bulk_delete_executions.json"

// BulkDeleteExecutionsResponse represents a bulk delete execution response
type BulkDeleteExecutionsResponse struct {
	FailedCount   int                                  `json:"failedCount"`
	SuccessCount  int                                  `json:"successCount"`
	AllSuccessful bool                                 `json:"allsuccessful"`
	RequestCount  int                                  `json:"requestCount"`
	Failures      []BulkDeleteExecutionFailureResponse `json:"failures"`
}

func (a BulkDeleteExecutionsResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a BulkDeleteExecutionsResponse) maxVersion() int  { return CurrentVersion }
func (a BulkDeleteExecutionsResponse) deprecated() bool { return false }

// FromReader returns an BulkDeleteExecutionsResponse from an io.Reader
func (a *BulkDeleteExecutionsResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an BulkDeleteExecutionsResponse from a file
func (a *BulkDeleteExecutionsResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a BulkDeleteExecutionsResponse from a byte slice
func (a *BulkDeleteExecutionsResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// BulkDeleteExecutionFailureResponse represents an individual bulk delete executions failure entry
type BulkDeleteExecutionFailureResponse struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// ExecutionStateResponseTestFile is the test data for ExecutionStateResponse
const ExecutionStateResponseTestFile = "execution_state.json"

// ExecutionStateResponse is an execution state response
type ExecutionStateResponse struct {
	Completed      bool                                         `json:"completed"`
	ExecutionState string                                       `json:"executionState"`
	EndTime        *JSONTime                                    `json:"endTime"`
	ServerNode     string                                       `json:"serverNode"`
	StartTime      *JSONTime                                    `json:"startTime"`
	UpdateTime     *JSONTime                                    `json:"updateTime"`
	StepCount      int                                          `json:"stepCount"`
	AllNodes       []string                                     `json:"allNodes"`
	TargetNodes    []string                                     `json:"targetNodes"`
	Nodes          map[string][]ExecutionStateNodeEntryResponse `json:"nodes"`
	ExecutionID    int                                          `json:"executionId"`
	Steps          []json.RawMessage                            `json:"steps"`
}

func (a ExecutionStateResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStateResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStateResponse) deprecated() bool { return false }

// FromReader returns an ExecutionStateResponse from an io.Reader
func (a *ExecutionStateResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an ExecutionStateResponse from a file
func (a *ExecutionStateResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a ExecutionStateResponse from a byte slice
func (a *ExecutionStateResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// ExecutionStepResponse represents an execution step
type ExecutionStepResponse struct {
	ExecutionState string                        `json:"executionState"`
	EndTime        *JSONTime                     `json:"endTime"`
	NodeStates     map[string]*NodeStateResponse `json:"nodeStates"`
	UpdateTime     *JSONTime                     `json:"updateTime"`
	NodeStep       bool                          `json:"nodeStep"`
	ID             string                        `json:"id"`
	StartTime      *JSONTime                     `json:"startTime"`
}

func (a ExecutionStepResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStepResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStepResponse) deprecated() bool { return false }

// WorkflowResponse represents a workflow response
type WorkflowResponse struct {
	Completed      bool              `json:"completed"`
	EndTime        *JSONTime         `json:"endTime"`
	StartTime      *JSONTime         `json:"startTime"`
	UpdateTime     *JSONTime         `json:"updateTime"`
	StepCount      int               `json:"stepCount"`
	AllNodes       []string          `json:"allNodes"`
	TargetNodes    []string          `json:"targetNodes"`
	ExecutionState string            `json:"executionState"`
	Steps          []json.RawMessage `json:"steps"`
}

func (a WorkflowResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a WorkflowResponse) maxVersion() int  { return CurrentVersion }
func (a WorkflowResponse) deprecated() bool { return false }

// WorkflowStepResponse represents a workflow step response
type WorkflowStepResponse struct {
	Workflow       *WorkflowResponse `json:"workflow"`
	ExecutionState string            `json:"executionState"`
	EndTime        *JSONTime         `json:"endTime"`
	StartTime      *JSONTime         `json:"startTime"`
	UpdateTime     *JSONTime         `json:"updateTime"`
	HasSubworkFlow bool              `json:"hasSubworkFlow"`
	NodeStep       bool              `json:"nodeStep"`
	ID             string            `json:"id"`
}

func (a WorkflowStepResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a WorkflowStepResponse) maxVersion() int  { return CurrentVersion }
func (a WorkflowStepResponse) deprecated() bool { return false }

// NodeStateResponse represents a nodeState response
type NodeStateResponse struct {
	ExecutionState string    `json:"executionState"`
	EndTime        *JSONTime `json:"endTime"`
	UpdateTime     *JSONTime `json:"updateTime"`
	StartTime      *JSONTime `json:"startTime"`
}

func (a NodeStateResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a NodeStateResponse) maxVersion() int  { return CurrentVersion }
func (a NodeStateResponse) deprecated() bool { return false }

// ExecutionStateNodeEntryResponse represents an individual node entry response
type ExecutionStateNodeEntryResponse struct {
	ExecutionState string `json:"executionState"`
	StepCtx        string `json:"stepctx"`
}

func (a ExecutionStateNodeEntryResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a ExecutionStateNodeEntryResponse) maxVersion() int  { return CurrentVersion }
func (a ExecutionStateNodeEntryResponse) deprecated() bool { return false }

// AdHocExecutionResponseTestFile is the test data for an AdHocExecutionResponse
const AdHocExecutionResponseTestFile = "execution_adhoc.json"

// AdHocExecutionResponse is the response for an running and adhoc command
type AdHocExecutionResponse struct {
	Message   string                     `json:"message"`
	Execution AdHocExecutionItemResponse `json:"execution"`
}

func (a AdHocExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AdHocExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a AdHocExecutionResponse) deprecated() bool { return false }

// FromReader returns an AdHocExecutionResponse from an io.Reader
func (a *AdHocExecutionResponse) FromReader(i io.Reader) error {
	b, err := ioutil.ReadAll(i)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, a)
}

// FromFile returns an AdHocExecutionResponse from a file
func (a *AdHocExecutionResponse) FromFile(f string) error {
	file, err := os.Open(f)
	defer func() { _ = file.Close() }()
	if err != nil {
		return err
	}
	return a.FromReader(file)
}

// FromBytes returns a AdHocExecutionResponse from a byte slice
func (a *AdHocExecutionResponse) FromBytes(f []byte) error {
	file := bytes.NewReader(f)
	return a.FromReader(file)
}

// AdHocExecutionItemResponse is an individual adhoc execution response
type AdHocExecutionItemResponse struct {
	ID        int    `json:"id"`
	HRef      string `json:"href"`
	Permalink string `json:"permalink"`
}

func (a AdHocExecutionItemResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AdHocExecutionItemResponse) maxVersion() int  { return CurrentVersion }
func (a AdHocExecutionItemResponse) deprecated() bool { return false }

// AbortExecutionResponse is the response for aborting an execution
type AbortExecutionResponse struct {
	Abort struct {
		Status string `json:"status"`
		Reason string `json:"reason"`
	} `json:"abort"`
	Execution struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		HRef   string `json:"href"`
	} `json:"execution"`
}

func (a AbortExecutionResponse) minVersion() int  { return AbsoluteMinimumVersion }
func (a AbortExecutionResponse) maxVersion() int  { return CurrentVersion }
func (a AbortExecutionResponse) deprecated() bool { return false }
