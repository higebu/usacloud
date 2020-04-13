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
	"fmt"
	"io"

	"github.com/sacloud/libsacloud/sacloud"
	v0params "github.com/sacloud/usacloud/command/params"
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/pkg/utils"
	"github.com/sacloud/usacloud/pkg/validation"
	"github.com/sacloud/usacloud/schema"
)

// ListProductLicenseParam is input parameters for the sacloud API
type ListProductLicenseParam struct {
	Name              []string
	Id                []sacloud.ID
	From              int
	Max               int
	Sort              []string
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string

	input Input
}

// NewListProductLicenseParam return new ListProductLicenseParam
func NewListProductLicenseParam() *ListProductLicenseParam {
	return &ListProductLicenseParam{}
}

// Initialize init ListProductLicenseParam
func (p *ListProductLicenseParam) Initialize(in Input, args []string) error {
	p.input = in
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ListProductLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ListProductLicenseParam) FillValueToSkeleton() {
	if utils.IsEmpty(p.Name) {
		p.Name = []string{""}
	}
	if utils.IsEmpty(p.Id) {
		p.Id = []sacloud.ID{}
	}
	if utils.IsEmpty(p.From) {
		p.From = 0
	}
	if utils.IsEmpty(p.Max) {
		p.Max = 0
	}
	if utils.IsEmpty(p.Sort) {
		p.Sort = []string{""}
	}
	if utils.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if utils.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if utils.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if utils.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if utils.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if utils.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if utils.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if utils.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if utils.IsEmpty(p.Format) {
		p.Format = ""
	}
	if utils.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if utils.IsEmpty(p.Query) {
		p.Query = ""
	}
	if utils.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}

}

func (p *ListProductLicenseParam) validate() error {
	var errors []error

	{
		errs := validation.ConflictsWith("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		validator := define.Resources["ProductLicense"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validation.ConflictsWith("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
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
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *ListProductLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ListProductLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ListProductLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ListProductLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["list"]
}

func (p *ListProductLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ListProductLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ListProductLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ListProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ListProductLicenseParam) SetName(v []string) {
	p.Name = v
}

func (p *ListProductLicenseParam) GetName() []string {
	return p.Name
}
func (p *ListProductLicenseParam) SetId(v []sacloud.ID) {
	p.Id = v
}

func (p *ListProductLicenseParam) GetId() []sacloud.ID {
	return p.Id
}
func (p *ListProductLicenseParam) SetFrom(v int) {
	p.From = v
}

func (p *ListProductLicenseParam) GetFrom() int {
	return p.From
}
func (p *ListProductLicenseParam) SetMax(v int) {
	p.Max = v
}

func (p *ListProductLicenseParam) GetMax() int {
	return p.Max
}
func (p *ListProductLicenseParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListProductLicenseParam) GetSort() []string {
	return p.Sort
}
func (p *ListProductLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListProductLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListProductLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ListProductLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *ListProductLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListProductLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListProductLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ListProductLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ListProductLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListProductLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ListProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListProductLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ListProductLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListProductLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ListProductLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ListProductLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ListProductLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ListProductLicenseParam) GetQueryFile() string {
	return p.QueryFile
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ListProductLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *ListProductLicenseParam) ToV0() *v0params.ListProductLicenseParam {
	return &v0params.ListProductLicenseParam{
		Name:              p.Name,
		Id:                p.Id,
		From:              p.From,
		Max:               p.Max,
		Sort:              p.Sort,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
	}
}

// ReadProductLicenseParam is input parameters for the sacloud API
type ReadProductLicenseParam struct {
	Assumeyes         bool
	ParamTemplate     string
	Parameters        string
	ParamTemplateFile string
	ParameterFile     string
	GenerateSkeleton  bool
	OutputType        string
	Column            []string
	Quiet             bool
	Format            string
	FormatFile        string
	Query             string
	QueryFile         string
	Id                sacloud.ID

	input Input
}

// NewReadProductLicenseParam return new ReadProductLicenseParam
func NewReadProductLicenseParam() *ReadProductLicenseParam {
	return &ReadProductLicenseParam{}
}

// Initialize init ReadProductLicenseParam
func (p *ReadProductLicenseParam) Initialize(in Input, args []string) error {
	p.input = in

	if len(args) == 0 {
		return fmt.Errorf("argument <ID> is required")
	}
	p.Id = sacloud.StringID(args[0])
	if p.Id.IsEmpty() {
		return fmt.Errorf("argument <ID> is required")
	}
	if err := p.validate(); err != nil {
		return err
	}
	return loadParameters(p)
}

// WriteSkeleton writes skeleton of JSON encoded parameters to specified writer
func (p *ReadProductLicenseParam) WriteSkeleton(writer io.Writer) error {
	return writeSkeleton(p, writer)
}

// FillValueToSkeleton fills empty value to the parameter
func (p *ReadProductLicenseParam) FillValueToSkeleton() {
	if utils.IsEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
	if utils.IsEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if utils.IsEmpty(p.Parameters) {
		p.Parameters = ""
	}
	if utils.IsEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if utils.IsEmpty(p.ParameterFile) {
		p.ParameterFile = ""
	}
	if utils.IsEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if utils.IsEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if utils.IsEmpty(p.Column) {
		p.Column = []string{""}
	}
	if utils.IsEmpty(p.Quiet) {
		p.Quiet = false
	}
	if utils.IsEmpty(p.Format) {
		p.Format = ""
	}
	if utils.IsEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if utils.IsEmpty(p.Query) {
		p.Query = ""
	}
	if utils.IsEmpty(p.QueryFile) {
		p.QueryFile = ""
	}
	if utils.IsEmpty(p.Id) {
		p.Id = sacloud.ID(0)
	}

}

func (p *ReadProductLicenseParam) validate() error {
	var errors []error

	{
		validator := validateRequired
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["ProductLicense"].Commands["read"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	{
		errs := validateParameterOptions(p)
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
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	return utils.FlattenErrors(errors)
}

func (p *ReadProductLicenseParam) ResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ReadProductLicenseParam) CommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductLicenseParam) IncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductLicenseParam) ExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductLicenseParam) TableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductLicenseParam) ColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

/*
 * v0系との互換性維持のための実装
 */
func (p *ReadProductLicenseParam) GetResourceDef() *schema.Resource {
	return define.Resources["ProductLicense"]
}

func (p *ReadProductLicenseParam) GetCommandDef() *schema.Command {
	return p.ResourceDef().Commands["read"]
}

func (p *ReadProductLicenseParam) GetIncludeFields() []string {
	return p.CommandDef().IncludeFields
}

func (p *ReadProductLicenseParam) GetExcludeFields() []string {
	return p.CommandDef().ExcludeFields
}

func (p *ReadProductLicenseParam) GetTableType() output.TableType {
	return p.CommandDef().TableType
}

func (p *ReadProductLicenseParam) GetColumnDefs() []output.ColumnDef {
	return p.CommandDef().TableColumnDefines
}

func (p *ReadProductLicenseParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *ReadProductLicenseParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *ReadProductLicenseParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadProductLicenseParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadProductLicenseParam) SetParameters(v string) {
	p.Parameters = v
}

func (p *ReadProductLicenseParam) GetParameters() string {
	return p.Parameters
}
func (p *ReadProductLicenseParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadProductLicenseParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadProductLicenseParam) SetParameterFile(v string) {
	p.ParameterFile = v
}

func (p *ReadProductLicenseParam) GetParameterFile() string {
	return p.ParameterFile
}
func (p *ReadProductLicenseParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadProductLicenseParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadProductLicenseParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadProductLicenseParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadProductLicenseParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadProductLicenseParam) GetColumn() []string {
	return p.Column
}
func (p *ReadProductLicenseParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadProductLicenseParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadProductLicenseParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadProductLicenseParam) GetFormat() string {
	return p.Format
}
func (p *ReadProductLicenseParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadProductLicenseParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadProductLicenseParam) SetQuery(v string) {
	p.Query = v
}

func (p *ReadProductLicenseParam) GetQuery() string {
	return p.Query
}
func (p *ReadProductLicenseParam) SetQueryFile(v string) {
	p.QueryFile = v
}

func (p *ReadProductLicenseParam) GetQueryFile() string {
	return p.QueryFile
}
func (p *ReadProductLicenseParam) SetId(v sacloud.ID) {
	p.Id = v
}

func (p *ReadProductLicenseParam) GetId() sacloud.ID {
	return p.Id
}

// Changed usacloud v0系との互換性維持のための実装
func (p *ReadProductLicenseParam) Changed(name string) bool {
	return p.input.Changed(name)
}

func (p *ReadProductLicenseParam) ToV0() *v0params.ReadProductLicenseParam {
	return &v0params.ReadProductLicenseParam{
		Assumeyes:         p.Assumeyes,
		ParamTemplate:     p.ParamTemplate,
		Parameters:        p.Parameters,
		ParamTemplateFile: p.ParamTemplateFile,
		ParameterFile:     p.ParameterFile,
		GenerateSkeleton:  p.GenerateSkeleton,
		OutputType:        p.OutputType,
		Column:            p.Column,
		Quiet:             p.Quiet,
		Format:            p.Format,
		FormatFile:        p.FormatFile,
		Query:             p.Query,
		QueryFile:         p.QueryFile,
		Id:                p.Id,
	}
}
