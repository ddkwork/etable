// Copyright (c) 2019, The eTable Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etable

import "github.com/emer/etable/etensor"

// Column specifies everything about a column -- can be used for constructing tables
type Column struct {

	// name of column -- must be unique for a table
	Name string `desc:"name of column -- must be unique for a table"`

	// data type, using etensor types which are isomorphic with arrow.Type
	Type etensor.Type `desc:"data type, using etensor types which are isomorphic with arrow.Type"`

	// shape of a single cell in the column (i.e., without the row dimension) -- for scalars this is nil -- tensor column will add the outer row dimension to this shape
	CellShape []int `desc:"shape of a single cell in the column (i.e., without the row dimension) -- for scalars this is nil -- tensor column will add the outer row dimension to this shape"`

	// names of the dimensions within the CellShape -- 'Row' will be added to outer dimension
	DimNames []string `desc:"names of the dimensions within the CellShape -- 'Row' will be added to outer dimension"`
}

// Schema specifies all of the columns of a table, sufficient to create the table
// It is just a slice list of Columns
type Schema []Column
