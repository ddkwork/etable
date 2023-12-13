// Copyright (c) 2019, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package etable provides a DataTable structure (also known as a DataFrame)
which is a collection of columnar data all having the same number of rows.
Each column is an etensor.Tensor.

The following sub-packages are included:

* bitslice is a Go slice of bytes []byte that has methods for setting individual bits,
as if it was a slice of bools, while being 8x more memory efficient.  This is used for
encoding null entries in  etensor, and as a Tensor of bool / bits there as well,
and is generally very useful for binary (boolean) data.

* etensor is the emer implementation of a Tensor (n-dimensional array) object.
etensor.Tensor is an interface that applies to many different type-specific instances,
such as etensor.Float32.  A tensor is just a etensor.Shape plus a slice holding
the specific data type.  Our tensor is based directly on the
[Apache Arrow](https://github.com/apache/arrow/tree/master/go) project's tensor,
and it fully interoperates with it.  Arrow tensors are designed to be read-only,
and we needed some extra support to make our etable.Table work well, so we had to roll
our own.  Our tensors also interoperate fully with Gonum's 2D-specific Matrix type for
the 2D case, and can use the gonum/floats and stats routines for raw arithmetic etc.

* etable is our Go version of DataTable from C++ emergent, which is widely useful
for holding input patterns to present to the network, and logs of output from the network,
among many other uses.  A etable.Table is a collection of etensor.Tensor columns,
that are all aligned along the outer-most *row* dimension.  Index-based
indirection is supported via optional args, but we do not take on the burden of
ensuring full updating of the indexes across all operations, which greatly simplifies things.
The etable.Table should interoperate with the under-development gonum DataFrame
structure among others.  The use of this data structure is always optional and orthogonal
to the core network algorithm code -- in Python the pandas library has a suitable DataFrame
structure that can be used instead.
*/
package etable
