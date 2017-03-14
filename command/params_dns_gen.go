// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// CreateDNSParam is input parameters for the sacloud API
type CreateDNSParam struct {
	Description string
	IconId      int64
	Name        string
	Tags        []string
}

// NewCreateDNSParam return new CreateDNSParam
func NewCreateDNSParam() *CreateDNSParam {
	return &CreateDNSParam{}
}

// Validate checks current values in model
func (p *CreateDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["DNS"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
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
		validator := define.Resources["DNS"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *CreateDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateDNSParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateDNSParam) GetDescription() string {
	return p.Description
}
func (p *CreateDNSParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateDNSParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateDNSParam) SetName(v string) {
	p.Name = v
}

func (p *CreateDNSParam) GetName() string {
	return p.Name
}
func (p *CreateDNSParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateDNSParam) GetTags() []string {
	return p.Tags
}

// DeleteDNSParam is input parameters for the sacloud API
type DeleteDNSParam struct {
	Force bool
	Id    int64
}

// NewDeleteDNSParam return new DeleteDNSParam
func NewDeleteDNSParam() *DeleteDNSParam {
	return &DeleteDNSParam{}
}

// Validate checks current values in model
func (p *DeleteDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *DeleteDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteDNSParam) SetForce(v bool) {
	p.Force = v
}

func (p *DeleteDNSParam) GetForce() bool {
	return p.Force
}
func (p *DeleteDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteDNSParam) GetId() int64 {
	return p.Id
}

// ListDNSParam is input parameters for the sacloud API
type ListDNSParam struct {
	From int
	Id   []int64
	Max  int
	Name []string
	Sort []string
}

// NewListDNSParam return new ListDNSParam
func NewListDNSParam() *ListDNSParam {
	return &ListDNSParam{}
}

// Validate checks current values in model
func (p *ListDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["DNS"].Commands["list"].Params["id"].ValidateFunc
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

	return errors
}

func (p *ListDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *ListDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListDNSParam) SetFrom(v int) {
	p.From = v
}

func (p *ListDNSParam) GetFrom() int {
	return p.From
}
func (p *ListDNSParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListDNSParam) GetId() []int64 {
	return p.Id
}
func (p *ListDNSParam) SetMax(v int) {
	p.Max = v
}

func (p *ListDNSParam) GetMax() int {
	return p.Max
}
func (p *ListDNSParam) SetName(v []string) {
	p.Name = v
}

func (p *ListDNSParam) GetName() []string {
	return p.Name
}
func (p *ListDNSParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListDNSParam) GetSort() []string {
	return p.Sort
}

// ReadDNSParam is input parameters for the sacloud API
type ReadDNSParam struct {
	Id int64
}

// NewReadDNSParam return new ReadDNSParam
func NewReadDNSParam() *ReadDNSParam {
	return &ReadDNSParam{}
}

// Validate checks current values in model
func (p *ReadDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *ReadDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadDNSParam) GetId() int64 {
	return p.Id
}

// RecordAddDNSParam is input parameters for the sacloud API
type RecordAddDNSParam struct {
	Id          int64
	MxPriority  int
	Name        string
	SrvPort     int
	SrvPriority int
	SrvTarget   string
	SrvWeight   int
	Ttl         int
	Type        string
	Value       string
}

// NewRecordAddDNSParam return new RecordAddDNSParam
func NewRecordAddDNSParam() *RecordAddDNSParam {
	return &RecordAddDNSParam{

		MxPriority: 10,

		Ttl: 3600,
	}
}

// Validate checks current values in model
func (p *RecordAddDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["mx-priority"].ValidateFunc
		errs := validator("--mx-priority", p.MxPriority)
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
		validator := define.Resources["DNS"].Commands["record-add"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-port"].ValidateFunc
		errs := validator("--srv-port", p.SrvPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-priority"].ValidateFunc
		errs := validator("--srv-priority", p.SrvPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-target"].ValidateFunc
		errs := validator("--srv-target", p.SrvTarget)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["srv-weight"].ValidateFunc
		errs := validator("--srv-weight", p.SrvWeight)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["ttl"].ValidateFunc
		errs := validator("--ttl", p.Ttl)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-add"].Params["type"].ValidateFunc
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RecordAddDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordAddDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["record-add"]
}

func (p *RecordAddDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RecordAddDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RecordAddDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RecordAddDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RecordAddDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *RecordAddDNSParam) GetId() int64 {
	return p.Id
}
func (p *RecordAddDNSParam) SetMxPriority(v int) {
	p.MxPriority = v
}

func (p *RecordAddDNSParam) GetMxPriority() int {
	return p.MxPriority
}
func (p *RecordAddDNSParam) SetName(v string) {
	p.Name = v
}

func (p *RecordAddDNSParam) GetName() string {
	return p.Name
}
func (p *RecordAddDNSParam) SetSrvPort(v int) {
	p.SrvPort = v
}

func (p *RecordAddDNSParam) GetSrvPort() int {
	return p.SrvPort
}
func (p *RecordAddDNSParam) SetSrvPriority(v int) {
	p.SrvPriority = v
}

func (p *RecordAddDNSParam) GetSrvPriority() int {
	return p.SrvPriority
}
func (p *RecordAddDNSParam) SetSrvTarget(v string) {
	p.SrvTarget = v
}

func (p *RecordAddDNSParam) GetSrvTarget() string {
	return p.SrvTarget
}
func (p *RecordAddDNSParam) SetSrvWeight(v int) {
	p.SrvWeight = v
}

func (p *RecordAddDNSParam) GetSrvWeight() int {
	return p.SrvWeight
}
func (p *RecordAddDNSParam) SetTtl(v int) {
	p.Ttl = v
}

func (p *RecordAddDNSParam) GetTtl() int {
	return p.Ttl
}
func (p *RecordAddDNSParam) SetType(v string) {
	p.Type = v
}

func (p *RecordAddDNSParam) GetType() string {
	return p.Type
}
func (p *RecordAddDNSParam) SetValue(v string) {
	p.Value = v
}

func (p *RecordAddDNSParam) GetValue() string {
	return p.Value
}

// RecordDeleteDNSParam is input parameters for the sacloud API
type RecordDeleteDNSParam struct {
	Force bool
	Id    int64
	Index int
}

// NewRecordDeleteDNSParam return new RecordDeleteDNSParam
func NewRecordDeleteDNSParam() *RecordDeleteDNSParam {
	return &RecordDeleteDNSParam{}
}

// Validate checks current values in model
func (p *RecordDeleteDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RecordDeleteDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordDeleteDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["record-delete"]
}

func (p *RecordDeleteDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RecordDeleteDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RecordDeleteDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RecordDeleteDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RecordDeleteDNSParam) SetForce(v bool) {
	p.Force = v
}

func (p *RecordDeleteDNSParam) GetForce() bool {
	return p.Force
}
func (p *RecordDeleteDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *RecordDeleteDNSParam) GetId() int64 {
	return p.Id
}
func (p *RecordDeleteDNSParam) SetIndex(v int) {
	p.Index = v
}

func (p *RecordDeleteDNSParam) GetIndex() int {
	return p.Index
}

// RecordListDNSParam is input parameters for the sacloud API
type RecordListDNSParam struct {
	Id int64
}

// NewRecordListDNSParam return new RecordListDNSParam
func NewRecordListDNSParam() *RecordListDNSParam {
	return &RecordListDNSParam{}
}

// Validate checks current values in model
func (p *RecordListDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RecordListDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordListDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["record-list"]
}

func (p *RecordListDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RecordListDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RecordListDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RecordListDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RecordListDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *RecordListDNSParam) GetId() int64 {
	return p.Id
}

// RecordUpdateDNSParam is input parameters for the sacloud API
type RecordUpdateDNSParam struct {
	Id          int64
	Index       int
	MxPriority  int
	Name        string
	SrvPort     int
	SrvPriority int
	SrvTarget   string
	SrvWeight   int
	Ttl         int
	Type        string
	Value       string
}

// NewRecordUpdateDNSParam return new RecordUpdateDNSParam
func NewRecordUpdateDNSParam() *RecordUpdateDNSParam {
	return &RecordUpdateDNSParam{}
}

// Validate checks current values in model
func (p *RecordUpdateDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--index", p.Index)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["mx-priority"].ValidateFunc
		errs := validator("--mx-priority", p.MxPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-port"].ValidateFunc
		errs := validator("--srv-port", p.SrvPort)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-priority"].ValidateFunc
		errs := validator("--srv-priority", p.SrvPriority)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-target"].ValidateFunc
		errs := validator("--srv-target", p.SrvTarget)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["srv-weight"].ValidateFunc
		errs := validator("--srv-weight", p.SrvWeight)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["ttl"].ValidateFunc
		errs := validator("--ttl", p.Ttl)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["record-update"].Params["type"].ValidateFunc
		errs := validator("--type", p.Type)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *RecordUpdateDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *RecordUpdateDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["record-update"]
}

func (p *RecordUpdateDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *RecordUpdateDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *RecordUpdateDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *RecordUpdateDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *RecordUpdateDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *RecordUpdateDNSParam) GetId() int64 {
	return p.Id
}
func (p *RecordUpdateDNSParam) SetIndex(v int) {
	p.Index = v
}

func (p *RecordUpdateDNSParam) GetIndex() int {
	return p.Index
}
func (p *RecordUpdateDNSParam) SetMxPriority(v int) {
	p.MxPriority = v
}

func (p *RecordUpdateDNSParam) GetMxPriority() int {
	return p.MxPriority
}
func (p *RecordUpdateDNSParam) SetName(v string) {
	p.Name = v
}

func (p *RecordUpdateDNSParam) GetName() string {
	return p.Name
}
func (p *RecordUpdateDNSParam) SetSrvPort(v int) {
	p.SrvPort = v
}

func (p *RecordUpdateDNSParam) GetSrvPort() int {
	return p.SrvPort
}
func (p *RecordUpdateDNSParam) SetSrvPriority(v int) {
	p.SrvPriority = v
}

func (p *RecordUpdateDNSParam) GetSrvPriority() int {
	return p.SrvPriority
}
func (p *RecordUpdateDNSParam) SetSrvTarget(v string) {
	p.SrvTarget = v
}

func (p *RecordUpdateDNSParam) GetSrvTarget() string {
	return p.SrvTarget
}
func (p *RecordUpdateDNSParam) SetSrvWeight(v int) {
	p.SrvWeight = v
}

func (p *RecordUpdateDNSParam) GetSrvWeight() int {
	return p.SrvWeight
}
func (p *RecordUpdateDNSParam) SetTtl(v int) {
	p.Ttl = v
}

func (p *RecordUpdateDNSParam) GetTtl() int {
	return p.Ttl
}
func (p *RecordUpdateDNSParam) SetType(v string) {
	p.Type = v
}

func (p *RecordUpdateDNSParam) GetType() string {
	return p.Type
}
func (p *RecordUpdateDNSParam) SetValue(v string) {
	p.Value = v
}

func (p *RecordUpdateDNSParam) GetValue() string {
	return p.Value
}

// UpdateDNSParam is input parameters for the sacloud API
type UpdateDNSParam struct {
	Description string
	IconId      int64
	Id          int64
	Tags        []string
}

// NewUpdateDNSParam return new UpdateDNSParam
func NewUpdateDNSParam() *UpdateDNSParam {
	return &UpdateDNSParam{}
}

// Validate checks current values in model
func (p *UpdateDNSParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["DNS"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["update"].Params["icon-id"].ValidateFunc
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
		validator := define.Resources["DNS"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["DNS"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateDNSParam) getResourceDef() *schema.Resource {
	return define.Resources["DNS"]
}

func (p *UpdateDNSParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateDNSParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateDNSParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateDNSParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateDNSParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateDNSParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateDNSParam) GetDescription() string {
	return p.Description
}
func (p *UpdateDNSParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateDNSParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateDNSParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateDNSParam) GetId() int64 {
	return p.Id
}
func (p *UpdateDNSParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateDNSParam) GetTags() []string {
	return p.Tags
}
