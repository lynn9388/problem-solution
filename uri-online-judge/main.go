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

	"github.com/dedis/student_18/dgcosi/code/onet/log"
	"github.com/lynn9388/problem-solution/uri-online-judge/urioj"
)

func main() {
	var id int
	for {
		fmt.Print("Please input the problem id:")
		fmt.Scan(&id)

		p, err := urioj.NewProblem(id)
		if err != nil {
			log.Error(err)
		}

		path := fmt.Sprintf("%v %v/main.go", p.ID, p.Name)
		if err := urioj.NewDescriptionFile(id, path); err != nil {
			log.Error(err)
		}

		description, err := urioj.NewDescription(id)
		if err != nil {
			log.Error(err)
		}
		fmt.Println(description)
	}
}
