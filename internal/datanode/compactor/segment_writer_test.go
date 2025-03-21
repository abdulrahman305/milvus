// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compactor

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/v2/util/paramtable"
	"github.com/milvus-io/milvus/pkg/v2/util/tsoutil"
)

func TestSegmentWriterSuite(t *testing.T) {
	suite.Run(t, new(SegmentWriteSuite))
}

type SegmentWriteSuite struct {
	suite.Suite
	collectionID int64
	parititonID  int64
}

func (s *SegmentWriteSuite) SetupSuite() {
	s.collectionID = 100
	s.parititonID = 101
}

func (s *SegmentWriteSuite) TestWriteFailed() {
	paramtable.Init()
	s.Run("get bm25 field failed", func() {
		schema := genCollectionSchemaWithBM25()
		// init segment writer with invalid bm25 fieldID
		writer, err := NewSegmentWriter(schema, 1024, compactionBatchSize, 1, s.parititonID, s.collectionID, []int64{1000})
		s.Require().NoError(err)

		v := storage.Value{
			PK:        storage.NewInt64PrimaryKey(int64(0)),
			Timestamp: int64(tsoutil.ComposeTSByTime(getMilvusBirthday(), 0)),
			Value:     genRowWithBM25(int64(0)),
		}
		err = writer.Write(&v)
		s.Error(err)
	})

	s.Run("parse bm25 field data failed", func() {
		schema := genCollectionSchemaWithBM25()
		// init segment writer with wrong field as bm25 sparse field
		writer, err := NewSegmentWriter(schema, 1024, compactionBatchSize, 1, s.parititonID, s.collectionID, []int64{101})
		s.Require().NoError(err)

		v := storage.Value{
			PK:        storage.NewInt64PrimaryKey(int64(0)),
			Timestamp: int64(tsoutil.ComposeTSByTime(getMilvusBirthday(), 0)),
			Value:     genRowWithBM25(int64(0)),
		}
		err = writer.Write(&v)
		s.Error(err)
	})
}
