/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/lynn9388/problem-solution/urioj"
)

func main() {
	var id int
	for {
		fmt.Print("Please input the problem id:")
		fmt.Scan(&id)

		p, err := urioj.NewProblem(id)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(p)

		path := fmt.Sprintf("%v %v/main.go", p.ID, p.Name)
		if err := newSourceFile(p.String(), path); err != nil {
			log.Print(err)
		}
	}
}

func newSourceFile(content string, path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := ioutil.WriteFile(path, []byte(content), 0664); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("file already exists: %v", path)
	}

	return nil
}
