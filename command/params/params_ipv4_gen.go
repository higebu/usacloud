// Copyright 2017-2020 The Usacloud Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/libsacloud/sacloud"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListIPv4Param is input parameters for the sacloud API
type ListIPv4Param struct {
	Name              []string     `json:"name"`
	Id                []sacloud.ID `json:"id"`
	From              int          `json:"from"`
	Max               int          `json:"max"`
	Sort              []string     `json:"sort"`
	ParamTemplate     string       `json:"param-template"`
	Parameters        string       `json:"parameters"`
	ParamTemplateFile string       `json:"param-template-file"`
	ParameterFile     string       `json:"parameter-file"`
	GenerateSkeleton  bool         `json:"generate-skeleton"`
	OutputType        string       `json:"output-type"`
	Column            []string     `json:"column"`
	Quiet             bool         `json:"quiet"`
	Format            string       `json:"format"`
	FormatFile        string       `json:"format-file"`
	Query             string       `json:"query"`
	QueryFile         string       `json:"query-file"`
}

// NewListIPv4Param return new ListIPv4Param
func NewListIPv4Param() *ListIPv4Param {
	return &ListIPv4Param{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListIPv4Param) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *ListIPv4Param) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["IPv4"].Commands["list"].Params["id"].ValidateFunc
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
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *ListIPv4Param) GetResourceDef() *schema.Resource {
	return define.Resources["IPv4"]
}

func (p *ListIPv4Param) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListIPv4Param) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListIPv4Param) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListIPv4Param) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListIPv4Param) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListIPv4Param) SetName(v []string) {
	p.Name = v
}

func (p *ListIPv4Param) GetName() []string {
	return p.Name
}
func (p *ListIPv4Param) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListIPv4Param) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListIPv4Param) SetFrom(v int) {
	p.From = v
}

func (p *ListIPv4Param) GetFrom() int {
	return p.From
}
func (p *ListIPv4Param) SetMax(v int) {
	p.Max = v
}

func (p *ListIPv4Param) GetMax() int {
	return p.Max
}
func (p *ListIPv4Param) SetSort(v []string) {
	p.Sort = v
}

func (p *ListIPv4Param) GetSort() []string {
	return p.Sort
}
func (p *ListIPv4Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListIPv4Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListIPv4Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListIPv4Param) GetParameters() string {
	return p.Parameters
}
func (p *ListIPv4Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListIPv4Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListIPv4Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListIPv4Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListIPv4Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListIPv4Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListIPv4Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListIPv4Param) GetOutputType() string {
	return p.OutputType
}
func (p *ListIPv4Param) SetColumn(v []string) {
	p.Column = v
}

func (p *ListIPv4Param) GetColumn() []string {
	return p.Column
}
func (p *ListIPv4Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListIPv4Param) GetQuiet() bool {
	return p.Quiet
}
func (p *ListIPv4Param) SetFormat(v string) {
	p.Format = v
}

func (p *ListIPv4Param) GetFormat() string {
	return p.Format
}
func (p *ListIPv4Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListIPv4Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListIPv4Param) SetQuery(v string) {
	p.Query = v
}

func (p *ListIPv4Param) GetQuery() string {
	return p.Query
}
func (p *ListIPv4Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListIPv4Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrAddIPv4Param is input parameters for the sacloud API
type PtrAddIPv4Param struct {
	Hostname          string   `json:"hostname"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	Parameters        string   `json:"parameters"`
	ParamTemplateFile string   `json:"param-template-file"`
	ParameterFile     string   `json:"parameter-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewPtrAddIPv4Param return new PtrAddIPv4Param
func NewPtrAddIPv4Param() *PtrAddIPv4Param {
	return &PtrAddIPv4Param{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PtrAddIPv4Param) FillValueToSkeleton() {
	if isEmpty(p.Hostname) {
		p.Hostname = ""
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *PtrAddIPv4Param) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *PtrAddIPv4Param) GetResourceDef() *schema.Resource {
	return define.Resources["IPv4"]
}

func (p *PtrAddIPv4Param) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["ptr-add"]
}

func (p *PtrAddIPv4Param) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PtrAddIPv4Param) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PtrAddIPv4Param) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PtrAddIPv4Param) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PtrAddIPv4Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrAddIPv4Param) GetHostname() string {
	return p.Hostname
}
func (p *PtrAddIPv4Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrAddIPv4Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrAddIPv4Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrAddIPv4Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrAddIPv4Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrAddIPv4Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrAddIPv4Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrAddIPv4Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrAddIPv4Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrAddIPv4Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrAddIPv4Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrAddIPv4Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrAddIPv4Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrAddIPv4Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrAddIPv4Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrAddIPv4Param) GetColumn() []string {
	return p.Column
}
func (p *PtrAddIPv4Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrAddIPv4Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrAddIPv4Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrAddIPv4Param) GetFormat() string {
	return p.Format
}
func (p *PtrAddIPv4Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrAddIPv4Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrAddIPv4Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrAddIPv4Param) GetQuery() string {
	return p.Query
}
func (p *PtrAddIPv4Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrAddIPv4Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrReadIPv4Param is input parameters for the sacloud API
type PtrReadIPv4Param struct {
	ParamTemplate     string   `json:"param-template"`
	Parameters        string   `json:"parameters"`
	ParamTemplateFile string   `json:"param-template-file"`
	ParameterFile     string   `json:"parameter-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewPtrReadIPv4Param return new PtrReadIPv4Param
func NewPtrReadIPv4Param() *PtrReadIPv4Param {
	return &PtrReadIPv4Param{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PtrReadIPv4Param) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *PtrReadIPv4Param) Validate() []error {
	errors := []error{}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *PtrReadIPv4Param) GetResourceDef() *schema.Resource {
	return define.Resources["IPv4"]
}

func (p *PtrReadIPv4Param) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["ptr-read"]
}

func (p *PtrReadIPv4Param) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PtrReadIPv4Param) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PtrReadIPv4Param) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PtrReadIPv4Param) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PtrReadIPv4Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrReadIPv4Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrReadIPv4Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrReadIPv4Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrReadIPv4Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrReadIPv4Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrReadIPv4Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrReadIPv4Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrReadIPv4Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrReadIPv4Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrReadIPv4Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrReadIPv4Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrReadIPv4Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrReadIPv4Param) GetColumn() []string {
	return p.Column
}
func (p *PtrReadIPv4Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrReadIPv4Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrReadIPv4Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrReadIPv4Param) GetFormat() string {
	return p.Format
}
func (p *PtrReadIPv4Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrReadIPv4Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrReadIPv4Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrReadIPv4Param) GetQuery() string {
	return p.Query
}
func (p *PtrReadIPv4Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrReadIPv4Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrUpdateIPv4Param is input parameters for the sacloud API
type PtrUpdateIPv4Param struct {
	Hostname          string   `json:"hostname"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	Parameters        string   `json:"parameters"`
	ParamTemplateFile string   `json:"param-template-file"`
	ParameterFile     string   `json:"parameter-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewPtrUpdateIPv4Param return new PtrUpdateIPv4Param
func NewPtrUpdateIPv4Param() *PtrUpdateIPv4Param {
	return &PtrUpdateIPv4Param{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PtrUpdateIPv4Param) FillValueToSkeleton() {
	if isEmpty(p.Hostname) {
		p.Hostname = ""
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *PtrUpdateIPv4Param) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--hostname", p.Hostname)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *PtrUpdateIPv4Param) GetResourceDef() *schema.Resource {
	return define.Resources["IPv4"]
}

func (p *PtrUpdateIPv4Param) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["ptr-update"]
}

func (p *PtrUpdateIPv4Param) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PtrUpdateIPv4Param) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PtrUpdateIPv4Param) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PtrUpdateIPv4Param) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PtrUpdateIPv4Param) SetHostname(v string) {
	p.Hostname = v
}

func (p *PtrUpdateIPv4Param) GetHostname() string {
	return p.Hostname
}
func (p *PtrUpdateIPv4Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrUpdateIPv4Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrUpdateIPv4Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrUpdateIPv4Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrUpdateIPv4Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrUpdateIPv4Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrUpdateIPv4Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrUpdateIPv4Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrUpdateIPv4Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrUpdateIPv4Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrUpdateIPv4Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrUpdateIPv4Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrUpdateIPv4Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrUpdateIPv4Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrUpdateIPv4Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrUpdateIPv4Param) GetColumn() []string {
	return p.Column
}
func (p *PtrUpdateIPv4Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrUpdateIPv4Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrUpdateIPv4Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrUpdateIPv4Param) GetFormat() string {
	return p.Format
}
func (p *PtrUpdateIPv4Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrUpdateIPv4Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrUpdateIPv4Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrUpdateIPv4Param) GetQuery() string {
	return p.Query
}
func (p *PtrUpdateIPv4Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrUpdateIPv4Param) GetQueryFile() string {
	return p.QueryFile
}

// PtrDeleteIPv4Param is input parameters for the sacloud API
type PtrDeleteIPv4Param struct {
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	Parameters        string   `json:"parameters"`
	ParamTemplateFile string   `json:"param-template-file"`
	ParameterFile     string   `json:"parameter-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Query             string   `json:"query"`
	QueryFile         string   `json:"query-file"`
}

// NewPtrDeleteIPv4Param return new PtrDeleteIPv4Param
func NewPtrDeleteIPv4Param() *PtrDeleteIPv4Param {
	return &PtrDeleteIPv4Param{}
}

// FillValueToSkeleton fill values to empty fields
func (p *PtrDeleteIPv4Param) FillValueToSkeleton() {
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.Query) {
		p.Query = ""
	}
	if isEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

// Validate checks current values in model
func (p *PtrDeleteIPv4Param) Validate() []error {
	errors := []error{}

	{
		validator := schema.ValidateInStrValues(define.AllowOutputTypes...)
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
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

func (p *PtrDeleteIPv4Param) GetResourceDef() *schema.Resource {
	return define.Resources["IPv4"]
}

func (p *PtrDeleteIPv4Param) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["ptr-delete"]
}

func (p *PtrDeleteIPv4Param) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *PtrDeleteIPv4Param) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *PtrDeleteIPv4Param) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *PtrDeleteIPv4Param) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *PtrDeleteIPv4Param) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *PtrDeleteIPv4Param) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *PtrDeleteIPv4Param) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *PtrDeleteIPv4Param) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *PtrDeleteIPv4Param) SetParameters(v string) {
	p.Parameters = v
}

func (p *PtrDeleteIPv4Param) GetParameters() string {
	return p.Parameters
}
func (p *PtrDeleteIPv4Param) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *PtrDeleteIPv4Param) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *PtrDeleteIPv4Param) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *PtrDeleteIPv4Param) GetParameterFile() string {
	return p.ParameterFile
}
func (p *PtrDeleteIPv4Param) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *PtrDeleteIPv4Param) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *PtrDeleteIPv4Param) SetOutputType(v string) {
	p.OutputType = v
}

func (p *PtrDeleteIPv4Param) GetOutputType() string {
	return p.OutputType
}
func (p *PtrDeleteIPv4Param) SetColumn(v []string) {
	p.Column = v
}

func (p *PtrDeleteIPv4Param) GetColumn() []string {
	return p.Column
}
func (p *PtrDeleteIPv4Param) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *PtrDeleteIPv4Param) GetQuiet() bool {
	return p.Quiet
}
func (p *PtrDeleteIPv4Param) SetFormat(v string) {
	p.Format = v
}

func (p *PtrDeleteIPv4Param) GetFormat() string {
	return p.Format
}
func (p *PtrDeleteIPv4Param) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *PtrDeleteIPv4Param) GetFormatFile() string {
	return p.FormatFile
}
func (p *PtrDeleteIPv4Param) SetQuery(v string) {
	p.Query = v
}

func (p *PtrDeleteIPv4Param) GetQuery() string {
	return p.Query
}
func (p *PtrDeleteIPv4Param) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *PtrDeleteIPv4Param) GetQueryFile() string {
	return p.QueryFile
}
