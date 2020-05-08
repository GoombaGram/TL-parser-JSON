/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package parser

import (
	"bufio"
	"encoding/json"
	"github.com/ErikPelli/TL-parser-JSON/parser/jsonStruct"
	"os"
)

func Parse (file *os.File) error {
	// Remove old schema or create a new folder if not exists
	_ = os.Remove("result/schema.json")
	_ = os.Mkdir("result", os.ModeDir)

	// Create a new output file
	output, err := os.OpenFile("result/schema.json", os.O_APPEND|os.O_RDWR|os.O_CREATE, 666)
	if err != nil {
		return err
	}
	defer output.Close()

	// Slices to convert to jsonStruct
	constructors := make([]*jsonStruct.Constructor, 0)
	methods := make([]*jsonStruct.Method, 0)

	// Buffer I/O Scanner from TL input file
	scanner := bufio.NewScanner(file)

	// Current mode (constructor/method)
	methodMode := false // First mode is always constructor

	// Parse every line
	for scanner.Scan() {
		// Get line as string
		line := scanner.Text()

		// Iteration check
		if line == "---functions---" {
			// Set mode to function
			methodMode = true
			continue
		} else if line == "---types---" {
			// Set mode to constructor
			methodMode = false
			continue
		} else if line == "" || (line[0] == '/' && line[1] == '/') {
			// If this line is useless (empty or a comment), skip it
			continue
		}

		// Parse line to single jsonStruct object and add it to corresponding slice
		if methodMode {
			// Create a new method
			method := new(jsonStruct.Method)

			// Fill method
			method.Id = getID(line)
			method.Method = getName(line)
			method.Params = getParams(line)
			method.Type = getType(line)

			// Add method to slice
			methods = append(methods, method)
		} else {
			// Create a new constructor
			constructor := new(jsonStruct.Constructor)

			// Fill constructor
			constructor.Id = getID(line)
			constructor.Predicate = getName(line)
			constructor.Params = getParams(line)
			constructor.Type = getType(line)

			// Add constructor to slice
			constructors = append(constructors, constructor)
		}
	}

	// Create a Schema struct
	TLstruct := new(jsonStruct.Schema)
	TLstruct.Constructors = constructors
	TLstruct.Methods = methods

	// Convert struct to json
	jsonBytes, err := json.Marshal(TLstruct); if err != nil {return err}

	// Write json to file
	_, err = output.Write(jsonBytes); if err != nil {return err}

	return nil
}