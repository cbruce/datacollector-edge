/*
 * Copyright 2017 StreamSets Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package jsonrecord

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestWriteMapRecord(t *testing.T) {
	stageContext := CreateStageContext()
	commits := map[string]interface{}{
		"rsc": 3711,
		"r":   2138,
		"gri": 1908,
		"adg": 912,
	}
	record1, err := stageContext.CreateRecord("Id1", commits)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")

	bufferWriter := bytes.NewBuffer([]byte{})

	recordWriterFactory := &JsonWriterFactoryImpl{}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)

	if err != nil {
		t.Fatal(err)
	}

	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}

	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var recordObject = make(map[string]int)
	err = decoder.Decode(&recordObject)
	if err != nil {
		t.Fatal(err)
	}

	if recordObject["rsc"] != commits["rsc"] {
		t.Errorf("Excepted: %d, but got: %d", commits["rsc"], recordObject["rsc"])
	}

	if recordObject["r"] != commits["r"] {
		t.Errorf("Excepted: %d, but got: %d", commits["r"], recordObject["r"])
	}

	if recordObject["gri"] != commits["gri"] {
		t.Errorf("Excepted: %d, but got: %d", commits["gri"], recordObject["gri"])
	}

	if recordObject["adg"] != commits["adg"] {
		t.Errorf("Excepted: %d, but got: %d", commits["adg"], recordObject["adg"])
	}
}

func TestWriteListRecord(t *testing.T) {
	stageContext := CreateStageContext()
	stringSlice := []interface{}{"apple", "orange", "banana"}
	record1, err := stageContext.CreateRecord("Id1", stringSlice)
	if err != nil {
		t.Fatal(err)
	}
	record1.GetHeader().SetAttribute("Sample Attribute", "Sample Value1")
	bufferWriter := bytes.NewBuffer([]byte{})

	//
	recordWriterFactory := &JsonWriterFactoryImpl{}
	recordWriter, err := recordWriterFactory.CreateWriter(stageContext, bufferWriter)
	if err != nil {
		t.Fatal(err)
	}
	err = recordWriter.WriteRecord(record1)
	if err != nil {
		t.Fatal(err)
	}
	recordWriter.Flush()
	recordWriter.Close()

	decoder := json.NewDecoder(bufferWriter)
	var listRecordObject = []interface{}{}
	err = decoder.Decode(&listRecordObject)
	if err != nil {
		t.Fatal(err)
	}

	if len(listRecordObject) != 3 {
		t.Errorf("Excepted length: 3, but got: %d", len(listRecordObject))
	}

	if listRecordObject[0] != stringSlice[0] {
		t.Errorf("Excepted: %s, but got: %s", stringSlice[0], listRecordObject[0])
	}
}
