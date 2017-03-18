// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CreateISOImageParam is input parameters for the sacloud API
type CreateISOImageParam struct {
	Description string
	IconId      int64
	IsoFile     string
	Name        string
	Size        int
	Tags        []string
	OutputType  string
	Column      []string
	Quiet       bool
	Format      string
}

// NewCreateISOImageParam return new CreateISOImageParam
func NewCreateISOImageParam() *CreateISOImageParam {
	return &CreateISOImageParam{

		Size: 5,
	}
}

// Validate checks current values in model
func (p *CreateISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--iso-file", p.IsoFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["iso-file"].ValidateFunc
		errs := validator("--iso-file", p.IsoFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["size"].ValidateFunc
		errs := validator("--size", p.Size)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *CreateISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *CreateISOImageParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateISOImageParam) GetDescription() string {
	return p.Description
}
func (p *CreateISOImageParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateISOImageParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateISOImageParam) SetIsoFile(v string) {
	p.IsoFile = v
}

func (p *CreateISOImageParam) GetIsoFile() string {
	return p.IsoFile
}
func (p *CreateISOImageParam) SetName(v string) {
	p.Name = v
}

func (p *CreateISOImageParam) GetName() string {
	return p.Name
}
func (p *CreateISOImageParam) SetSize(v int) {
	p.Size = v
}

func (p *CreateISOImageParam) GetSize() int {
	return p.Size
}
func (p *CreateISOImageParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateISOImageParam) GetTags() []string {
	return p.Tags
}
func (p *CreateISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *CreateISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateISOImageParam) GetFormat() string {
	return p.Format
}

// DeleteISOImageParam is input parameters for the sacloud API
type DeleteISOImageParam struct {
	Force      bool
	Id         int64
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewDeleteISOImageParam return new DeleteISOImageParam
func NewDeleteISOImageParam() *DeleteISOImageParam {
	return &DeleteISOImageParam{}
}

// Validate checks current values in model
func (p *DeleteISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *DeleteISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteISOImageParam) SetForce(v bool) {
	p.Force = v
}

func (p *DeleteISOImageParam) GetForce() bool {
	return p.Force
}
func (p *DeleteISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteISOImageParam) GetId() int64 {
	return p.Id
}
func (p *DeleteISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteISOImageParam) GetFormat() string {
	return p.Format
}

// DownloadISOImageParam is input parameters for the sacloud API
type DownloadISOImageParam struct {
	FileDestination string
	Id              int64
}

// NewDownloadISOImageParam return new DownloadISOImageParam
func NewDownloadISOImageParam() *DownloadISOImageParam {
	return &DownloadISOImageParam{}
}

// Validate checks current values in model
func (p *DownloadISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--file-destination", p.FileDestination)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["download"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DownloadISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *DownloadISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["download"]
}

func (p *DownloadISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DownloadISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DownloadISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DownloadISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DownloadISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *DownloadISOImageParam) SetFileDestination(v string) {
	p.FileDestination = v
}

func (p *DownloadISOImageParam) GetFileDestination() string {
	return p.FileDestination
}
func (p *DownloadISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *DownloadISOImageParam) GetId() int64 {
	return p.Id
}

// FtpCloseISOImageParam is input parameters for the sacloud API
type FtpCloseISOImageParam struct {
	Id int64
}

// NewFtpCloseISOImageParam return new FtpCloseISOImageParam
func NewFtpCloseISOImageParam() *FtpCloseISOImageParam {
	return &FtpCloseISOImageParam{}
}

// Validate checks current values in model
func (p *FtpCloseISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["ftp-close"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *FtpCloseISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *FtpCloseISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["ftp-close"]
}

func (p *FtpCloseISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *FtpCloseISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *FtpCloseISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *FtpCloseISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *FtpCloseISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *FtpCloseISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *FtpCloseISOImageParam) GetId() int64 {
	return p.Id
}

// FtpOpenISOImageParam is input parameters for the sacloud API
type FtpOpenISOImageParam struct {
	Id         int64
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewFtpOpenISOImageParam return new FtpOpenISOImageParam
func NewFtpOpenISOImageParam() *FtpOpenISOImageParam {
	return &FtpOpenISOImageParam{}
}

// Validate checks current values in model
func (p *FtpOpenISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["ftp-open"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *FtpOpenISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *FtpOpenISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["ftp-open"]
}

func (p *FtpOpenISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *FtpOpenISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *FtpOpenISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *FtpOpenISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *FtpOpenISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *FtpOpenISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *FtpOpenISOImageParam) GetId() int64 {
	return p.Id
}
func (p *FtpOpenISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *FtpOpenISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *FtpOpenISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *FtpOpenISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *FtpOpenISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *FtpOpenISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *FtpOpenISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *FtpOpenISOImageParam) GetFormat() string {
	return p.Format
}

// ListISOImageParam is input parameters for the sacloud API
type ListISOImageParam struct {
	From       int
	Id         []int64
	Max        int
	Name       []string
	Scope      string
	Sort       []string
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewListISOImageParam return new ListISOImageParam
func NewListISOImageParam() *ListISOImageParam {
	return &ListISOImageParam{}
}

// Validate checks current values in model
func (p *ListISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["ISOImage"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ListISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *ListISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *ListISOImageParam) SetFrom(v int) {
	p.From = v
}

func (p *ListISOImageParam) GetFrom() int {
	return p.From
}
func (p *ListISOImageParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListISOImageParam) GetId() []int64 {
	return p.Id
}
func (p *ListISOImageParam) SetMax(v int) {
	p.Max = v
}

func (p *ListISOImageParam) GetMax() int {
	return p.Max
}
func (p *ListISOImageParam) SetName(v []string) {
	p.Name = v
}

func (p *ListISOImageParam) GetName() []string {
	return p.Name
}
func (p *ListISOImageParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListISOImageParam) GetScope() string {
	return p.Scope
}
func (p *ListISOImageParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListISOImageParam) GetSort() []string {
	return p.Sort
}
func (p *ListISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *ListISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListISOImageParam) GetFormat() string {
	return p.Format
}

// ReadISOImageParam is input parameters for the sacloud API
type ReadISOImageParam struct {
	Id         int64
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewReadISOImageParam return new ReadISOImageParam
func NewReadISOImageParam() *ReadISOImageParam {
	return &ReadISOImageParam{}
}

// Validate checks current values in model
func (p *ReadISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *ReadISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadISOImageParam) GetId() int64 {
	return p.Id
}
func (p *ReadISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *ReadISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadISOImageParam) GetFormat() string {
	return p.Format
}

// UpdateISOImageParam is input parameters for the sacloud API
type UpdateISOImageParam struct {
	Description string
	IconId      int64
	Id          int64
	Name        string
	Tags        []string
	OutputType  string
	Column      []string
	Quiet       bool
	Format      string
}

// NewUpdateISOImageParam return new UpdateISOImageParam
func NewUpdateISOImageParam() *UpdateISOImageParam {
	return &UpdateISOImageParam{}
}

// Validate checks current values in model
func (p *UpdateISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *UpdateISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdateISOImageParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateISOImageParam) GetDescription() string {
	return p.Description
}
func (p *UpdateISOImageParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateISOImageParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateISOImageParam) GetId() int64 {
	return p.Id
}
func (p *UpdateISOImageParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateISOImageParam) GetName() string {
	return p.Name
}
func (p *UpdateISOImageParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateISOImageParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateISOImageParam) GetFormat() string {
	return p.Format
}

// UploadISOImageParam is input parameters for the sacloud API
type UploadISOImageParam struct {
	Id         int64
	IsoFile    string
	OutputType string
	Column     []string
	Quiet      bool
	Format     string
}

// NewUploadISOImageParam return new UploadISOImageParam
func NewUploadISOImageParam() *UploadISOImageParam {
	return &UploadISOImageParam{}
}

// Validate checks current values in model
func (p *UploadISOImageParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["upload"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--iso-file", p.IsoFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ISOImage"].Commands["upload"].Params["iso-file"].ValidateFunc
		errs := validator("--iso-file", p.IsoFile)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UploadISOImageParam) getResourceDef() *schema.Resource {
	return define.Resources["ISOImage"]
}

func (p *UploadISOImageParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["upload"]
}

func (p *UploadISOImageParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UploadISOImageParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UploadISOImageParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UploadISOImageParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UploadISOImageParam) GetOutputFormat() string {
	return "table"
}

func (p *UploadISOImageParam) SetId(v int64) {
	p.Id = v
}

func (p *UploadISOImageParam) GetId() int64 {
	return p.Id
}
func (p *UploadISOImageParam) SetIsoFile(v string) {
	p.IsoFile = v
}

func (p *UploadISOImageParam) GetIsoFile() string {
	return p.IsoFile
}
func (p *UploadISOImageParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UploadISOImageParam) GetOutputType() string {
	return p.OutputType
}
func (p *UploadISOImageParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UploadISOImageParam) GetColumn() []string {
	return p.Column
}
func (p *UploadISOImageParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UploadISOImageParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UploadISOImageParam) SetFormat(v string) {
	p.Format = v
}

func (p *UploadISOImageParam) GetFormat() string {
	return p.Format
}
