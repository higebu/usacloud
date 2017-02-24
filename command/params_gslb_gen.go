// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package command

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListGSLBParam is input parameters for the sacloud API
type ListGSLBParam struct {
	Id   []int64
	From int
	Max  int
	Sort []string
	Name []string
}

// NewListGSLBParam return new ListGSLBParam
func NewListGSLBParam() *ListGSLBParam {
	return &ListGSLBParam{}
}

// Validate checks current values in model
func (p *ListGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["GSLB"].Commands["list"].Params["id"].ValidateFunc
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

func (p *ListGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ListGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["list"]
}

func (p *ListGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ListGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ListGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ListGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ListGSLBParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListGSLBParam) GetId() []int64 {
	return p.Id
}
func (p *ListGSLBParam) SetFrom(v int) {
	p.From = v
}

func (p *ListGSLBParam) GetFrom() int {
	return p.From
}
func (p *ListGSLBParam) SetMax(v int) {
	p.Max = v
}

func (p *ListGSLBParam) GetMax() int {
	return p.Max
}
func (p *ListGSLBParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListGSLBParam) GetSort() []string {
	return p.Sort
}
func (p *ListGSLBParam) SetName(v []string) {
	p.Name = v
}

func (p *ListGSLBParam) GetName() []string {
	return p.Name
}

// CreateGSLBParam is input parameters for the sacloud API
type CreateGSLBParam struct {
	Name         string
	Description  string
	Port         int
	DelayLoop    int
	Weighted     bool
	SorryServer  string
	Tags         []string
	IconId       int64
	Protocol     string
	HostHeader   string
	Path         string
	ResponseCode int
}

// NewCreateGSLBParam return new CreateGSLBParam
func NewCreateGSLBParam() *CreateGSLBParam {
	return &CreateGSLBParam{

		DelayLoop: 10,

		Weighted: true,

		Protocol: "ping",

		Path: "/",

		ResponseCode: 200,
	}
}

// Validate checks current values in model
func (p *CreateGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["port"].ValidateFunc
		errs := validator("--port", p.Port)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["delay-loop"].ValidateFunc
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["create"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *CreateGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *CreateGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["create"]
}

func (p *CreateGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *CreateGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *CreateGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *CreateGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *CreateGSLBParam) SetName(v string) {
	p.Name = v
}

func (p *CreateGSLBParam) GetName() string {
	return p.Name
}
func (p *CreateGSLBParam) SetDescription(v string) {
	p.Description = v
}

func (p *CreateGSLBParam) GetDescription() string {
	return p.Description
}
func (p *CreateGSLBParam) SetPort(v int) {
	p.Port = v
}

func (p *CreateGSLBParam) GetPort() int {
	return p.Port
}
func (p *CreateGSLBParam) SetDelayLoop(v int) {
	p.DelayLoop = v
}

func (p *CreateGSLBParam) GetDelayLoop() int {
	return p.DelayLoop
}
func (p *CreateGSLBParam) SetWeighted(v bool) {
	p.Weighted = v
}

func (p *CreateGSLBParam) GetWeighted() bool {
	return p.Weighted
}
func (p *CreateGSLBParam) SetSorryServer(v string) {
	p.SorryServer = v
}

func (p *CreateGSLBParam) GetSorryServer() string {
	return p.SorryServer
}
func (p *CreateGSLBParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateGSLBParam) GetTags() []string {
	return p.Tags
}
func (p *CreateGSLBParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *CreateGSLBParam) GetIconId() int64 {
	return p.IconId
}
func (p *CreateGSLBParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *CreateGSLBParam) GetProtocol() string {
	return p.Protocol
}
func (p *CreateGSLBParam) SetHostHeader(v string) {
	p.HostHeader = v
}

func (p *CreateGSLBParam) GetHostHeader() string {
	return p.HostHeader
}
func (p *CreateGSLBParam) SetPath(v string) {
	p.Path = v
}

func (p *CreateGSLBParam) GetPath() string {
	return p.Path
}
func (p *CreateGSLBParam) SetResponseCode(v int) {
	p.ResponseCode = v
}

func (p *CreateGSLBParam) GetResponseCode() int {
	return p.ResponseCode
}

// ReadGSLBParam is input parameters for the sacloud API
type ReadGSLBParam struct {
	Id int64
}

// NewReadGSLBParam return new ReadGSLBParam
func NewReadGSLBParam() *ReadGSLBParam {
	return &ReadGSLBParam{}
}

// Validate checks current values in model
func (p *ReadGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ReadGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ReadGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["read"]
}

func (p *ReadGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ReadGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ReadGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ReadGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ReadGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadGSLBParam) GetId() int64 {
	return p.Id
}

// UpdateGSLBParam is input parameters for the sacloud API
type UpdateGSLBParam struct {
	ResponseCode int
	Description  string
	IconId       int64
	Port         int
	Weighted     bool
	Id           int64
	Protocol     string
	HostHeader   string
	Name         string
	Path         string
	DelayLoop    int
	SorryServer  string
	Tags         []string
}

// NewUpdateGSLBParam return new UpdateGSLBParam
func NewUpdateGSLBParam() *UpdateGSLBParam {
	return &UpdateGSLBParam{}
}

// Validate checks current values in model
func (p *UpdateGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["description"].ValidateFunc
		errs := validator("--description", p.Description)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["icon-id"].ValidateFunc
		errs := validator("--icon-id", p.IconId)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["port"].ValidateFunc
		errs := validator("--port", p.Port)
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
		validator := define.Resources["GSLB"].Commands["update"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["protocol"].ValidateFunc
		errs := validator("--protocol", p.Protocol)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["delay-loop"].ValidateFunc
		errs := validator("--delay-loop", p.DelayLoop)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *UpdateGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *UpdateGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["update"]
}

func (p *UpdateGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *UpdateGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *UpdateGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *UpdateGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *UpdateGSLBParam) SetResponseCode(v int) {
	p.ResponseCode = v
}

func (p *UpdateGSLBParam) GetResponseCode() int {
	return p.ResponseCode
}
func (p *UpdateGSLBParam) SetDescription(v string) {
	p.Description = v
}

func (p *UpdateGSLBParam) GetDescription() string {
	return p.Description
}
func (p *UpdateGSLBParam) SetIconId(v int64) {
	p.IconId = v
}

func (p *UpdateGSLBParam) GetIconId() int64 {
	return p.IconId
}
func (p *UpdateGSLBParam) SetPort(v int) {
	p.Port = v
}

func (p *UpdateGSLBParam) GetPort() int {
	return p.Port
}
func (p *UpdateGSLBParam) SetWeighted(v bool) {
	p.Weighted = v
}

func (p *UpdateGSLBParam) GetWeighted() bool {
	return p.Weighted
}
func (p *UpdateGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateGSLBParam) GetId() int64 {
	return p.Id
}
func (p *UpdateGSLBParam) SetProtocol(v string) {
	p.Protocol = v
}

func (p *UpdateGSLBParam) GetProtocol() string {
	return p.Protocol
}
func (p *UpdateGSLBParam) SetHostHeader(v string) {
	p.HostHeader = v
}

func (p *UpdateGSLBParam) GetHostHeader() string {
	return p.HostHeader
}
func (p *UpdateGSLBParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateGSLBParam) GetName() string {
	return p.Name
}
func (p *UpdateGSLBParam) SetPath(v string) {
	p.Path = v
}

func (p *UpdateGSLBParam) GetPath() string {
	return p.Path
}
func (p *UpdateGSLBParam) SetDelayLoop(v int) {
	p.DelayLoop = v
}

func (p *UpdateGSLBParam) GetDelayLoop() int {
	return p.DelayLoop
}
func (p *UpdateGSLBParam) SetSorryServer(v string) {
	p.SorryServer = v
}

func (p *UpdateGSLBParam) GetSorryServer() string {
	return p.SorryServer
}
func (p *UpdateGSLBParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateGSLBParam) GetTags() []string {
	return p.Tags
}

// ServerListGSLBParam is input parameters for the sacloud API
type ServerListGSLBParam struct {
	Id int64
}

// NewServerListGSLBParam return new ServerListGSLBParam
func NewServerListGSLBParam() *ServerListGSLBParam {
	return &ServerListGSLBParam{}
}

// Validate checks current values in model
func (p *ServerListGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["server-list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ServerListGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ServerListGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-list"]
}

func (p *ServerListGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerListGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerListGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerListGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerListGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerListGSLBParam) GetId() int64 {
	return p.Id
}

// DeleteGSLBParam is input parameters for the sacloud API
type DeleteGSLBParam struct {
	Id int64
}

// NewDeleteGSLBParam return new DeleteGSLBParam
func NewDeleteGSLBParam() *DeleteGSLBParam {
	return &DeleteGSLBParam{}
}

// Validate checks current values in model
func (p *DeleteGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["delete"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *DeleteGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *DeleteGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["delete"]
}

func (p *DeleteGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *DeleteGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *DeleteGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *DeleteGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *DeleteGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteGSLBParam) GetId() int64 {
	return p.Id
}

// ServerAddGSLBParam is input parameters for the sacloud API
type ServerAddGSLBParam struct {
	Enabled   bool
	Weight    int
	Id        int64
	Ipaddress string
}

// NewServerAddGSLBParam return new ServerAddGSLBParam
func NewServerAddGSLBParam() *ServerAddGSLBParam {
	return &ServerAddGSLBParam{

		Enabled: true,
	}
}

// Validate checks current values in model
func (p *ServerAddGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["GSLB"].Commands["server-add"].Params["weight"].ValidateFunc
		errs := validator("--weight", p.Weight)
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
		validator := define.Resources["GSLB"].Commands["server-add"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["server-add"].Params["ipaddress"].ValidateFunc
		errs := validator("--ipaddress", p.Ipaddress)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ServerAddGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ServerAddGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-add"]
}

func (p *ServerAddGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerAddGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerAddGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerAddGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerAddGSLBParam) SetEnabled(v bool) {
	p.Enabled = v
}

func (p *ServerAddGSLBParam) GetEnabled() bool {
	return p.Enabled
}
func (p *ServerAddGSLBParam) SetWeight(v int) {
	p.Weight = v
}

func (p *ServerAddGSLBParam) GetWeight() int {
	return p.Weight
}
func (p *ServerAddGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerAddGSLBParam) GetId() int64 {
	return p.Id
}
func (p *ServerAddGSLBParam) SetIpaddress(v string) {
	p.Ipaddress = v
}

func (p *ServerAddGSLBParam) GetIpaddress() string {
	return p.Ipaddress
}

// ServerUpdateGSLBParam is input parameters for the sacloud API
type ServerUpdateGSLBParam struct {
	Ipaddress string
	Enabled   bool
	Weight    int
	Id        int64
	Index     int
}

// NewServerUpdateGSLBParam return new ServerUpdateGSLBParam
func NewServerUpdateGSLBParam() *ServerUpdateGSLBParam {
	return &ServerUpdateGSLBParam{}
}

// Validate checks current values in model
func (p *ServerUpdateGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["GSLB"].Commands["server-update"].Params["ipaddress"].ValidateFunc
		errs := validator("--ipaddress", p.Ipaddress)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["server-update"].Params["weight"].ValidateFunc
		errs := validator("--weight", p.Weight)
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
		validator := define.Resources["GSLB"].Commands["server-update"].Params["id"].ValidateFunc
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

func (p *ServerUpdateGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ServerUpdateGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-update"]
}

func (p *ServerUpdateGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerUpdateGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerUpdateGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerUpdateGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerUpdateGSLBParam) SetIpaddress(v string) {
	p.Ipaddress = v
}

func (p *ServerUpdateGSLBParam) GetIpaddress() string {
	return p.Ipaddress
}
func (p *ServerUpdateGSLBParam) SetEnabled(v bool) {
	p.Enabled = v
}

func (p *ServerUpdateGSLBParam) GetEnabled() bool {
	return p.Enabled
}
func (p *ServerUpdateGSLBParam) SetWeight(v int) {
	p.Weight = v
}

func (p *ServerUpdateGSLBParam) GetWeight() int {
	return p.Weight
}
func (p *ServerUpdateGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerUpdateGSLBParam) GetId() int64 {
	return p.Id
}
func (p *ServerUpdateGSLBParam) SetIndex(v int) {
	p.Index = v
}

func (p *ServerUpdateGSLBParam) GetIndex() int {
	return p.Index
}

// ServerDeleteGSLBParam is input parameters for the sacloud API
type ServerDeleteGSLBParam struct {
	Id    int64
	Index int
}

// NewServerDeleteGSLBParam return new ServerDeleteGSLBParam
func NewServerDeleteGSLBParam() *ServerDeleteGSLBParam {
	return &ServerDeleteGSLBParam{}
}

// Validate checks current values in model
func (p *ServerDeleteGSLBParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["GSLB"].Commands["server-delete"].Params["id"].ValidateFunc
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

func (p *ServerDeleteGSLBParam) getResourceDef() *schema.Resource {
	return define.Resources["GSLB"]
}

func (p *ServerDeleteGSLBParam) getCommandDef() *schema.Command {
	return p.getResourceDef().Commands["server-delete"]
}

func (p *ServerDeleteGSLBParam) GetIncludeFields() []string {
	return p.getCommandDef().IncludeFields
}

func (p *ServerDeleteGSLBParam) GetExcludeFields() []string {
	return p.getCommandDef().ExcludeFields
}

func (p *ServerDeleteGSLBParam) GetTableType() output.OutputTableType {
	return p.getCommandDef().TableType
}

func (p *ServerDeleteGSLBParam) GetColumnDefs() []output.ColumnDef {
	return p.getCommandDef().TableColumnDefines
}

func (p *ServerDeleteGSLBParam) SetId(v int64) {
	p.Id = v
}

func (p *ServerDeleteGSLBParam) GetId() int64 {
	return p.Id
}
func (p *ServerDeleteGSLBParam) SetIndex(v int) {
	p.Index = v
}

func (p *ServerDeleteGSLBParam) GetIndex() int {
	return p.Index
}
