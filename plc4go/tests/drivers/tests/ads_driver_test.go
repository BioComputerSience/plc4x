/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package tests

import (
	"context"
	"testing"

	"github.com/apache/plc4x/plc4go/internal/ads"
	adsIO "github.com/apache/plc4x/plc4go/protocols/ads/readwrite"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/testutils"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

func TestAdsDriver(t *testing.T) {
	t.Skip("I have to port the commands for reading the symbol-table first")
	testutils.SetToTestingLogger(t, readWriteModel.Plc4xModelLog)
	parser := func(readBufferByteBased utils.ReadBufferByteBased) (any, error) {
		return readWriteModel.AmsTCPPacketParseWithBuffer(context.Background(), readBufferByteBased)
	}
	withCustomLogger := options.WithCustomLogger(testutils.ProduceTestingLogger(t))
	testutils.RunDriverTestsuite(
		t,
		ads.NewDriver(withCustomLogger),
		"assets/testing/protocols/ads/DriverTestsuite.xml",
		adsIO.AdsXmlParserHelper{},
		testutils.WithRootTypeParser(parser),
		withCustomLogger,
	)
}
