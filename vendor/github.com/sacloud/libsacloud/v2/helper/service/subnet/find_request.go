// Copyright 2016-2020 The Libsacloud Authors
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

package subnet

import (
	"github.com/sacloud/libsacloud/v2/helper/service"
	"github.com/sacloud/libsacloud/v2/helper/validate"
	"github.com/sacloud/libsacloud/v2/sacloud"
	"github.com/sacloud/libsacloud/v2/sacloud/search"
)

type FindRequest struct {
	Zone string `request:"-" validate:"required"`

	Sort  search.SortKeys
	Count int
	From  int
}

func (req *FindRequest) Validate() error {
	return validate.Struct(req)
}

func (req *FindRequest) ToRequestParameter() (*sacloud.FindCondition, error) {
	condition := &sacloud.FindCondition{
		Filter: map[search.FilterKey]interface{}{},
	}
	if err := service.RequestConvertTo(req, condition); err != nil {
		return nil, err
	}
	return condition, nil
}
