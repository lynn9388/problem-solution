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

package urioj

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var descriptionTests = map[int]string{
	1023: `/*********************************************************************
                               Drought
    https://www.urionlinejudge.com.br/judge/en/problems/view/1023

Due to the continuous drought that happened recently in some regions
of Brazil, the Federal Government created an agency to assess the
consumption of these regions in order to verify the behavior of the
population at the time of rationing. This agency will take some cities
(for sampling) and it will verify the consumption of each of the
townspeople and the average consumption per inhabitant of each town.

Input
*****
The input contains a number of test's cases. The first line of each
case of test contains an integer N (1 ≤ N ≤ 1 * 10 ⁶), indicating the
amount of properties. The following N lines contains a pair of values
X (1 ≤ X ≤ 10) and Y ( 1 ≤ Y ≤ 200) indicating the number of residents
of each property and its total consumption (m³). Surely, no residence
consumes more than 200 m³ per month. The input's end is represented by
zero.

Output
******
For each case of test you must present the message “Cidade# n:”, where
n is the number of the city in the sequence (1, 2, 3, ...), and then
you must list in ascending order of consumption, the people's amount
followed by a hyphen and the consumption of these people, rounding the
value down. In the third line of output you should present the consumption
per person in that town, with two decimal places without rounding,
considering the total real consumption. Print a blank line between two
consecutives test's cases. There is no blank line at the end of output.

+--------------+--------------------------+
| INPUT SAMPLE |      OUTPUT SAMPLE       |
+--------------+--------------------------+
| 3            | Cidade# 1:               |
| 3 22         | 2-5 3-7 3-13             |
| 2 11         | Consumo medio: 9.00 m3.  |
| 3 39         |                          |
| 5            | Cidade# 2:               |
| 1 25         | 5-10 6-11 2-20 1-25      |
| 2 20         | Consumo medio: 13.28 m3. |
| 3 31         |                          |
| 2 40         |                          |
| 6 70         |                          |
| 0            |                          |
+--------------+--------------------------+
*********************************************************************/`,
	1048: `/*********************************************************************
                           Salary Increase
    https://www.urionlinejudge.com.br/judge/en/problems/view/1048

The company ABC decided to give a salary increase to its employees,
according to the following table:

+-------------------+-------------------+
|      SALARY       | READJUSTMENT RATE |
+-------------------+-------------------+
| 0 - 400.00        | 15%               |
| 400.01 - 800.00   | 12%               |
| 800.01 - 1200.00  | 10%               |
| 1200.01 - 2000.00 | 7%                |
| Above 2000.00     | 4%                |
+-------------------+-------------------+

Read the employee's salary, calculate and print the new employee's
salary, as well the money earned and the increase percentual obtained
by the employee, with corresponding messages in Portuguese, as the
below example.

Input
*****
The input contains only a floating-point number, with 2 digits after
the decimal point.

Output
******
Print 3 messages followed by the corresponding numbers (see example)
informing the new salary, the among of money earned and the percentual
obtained by the employee. Note:
Novo salario: means "New Salary"
Reajuste ganho: means "Money earned"
Em percentual: means "In percentage"

+--------------+------------------------+
| INPUT SAMPLE |     OUTPUT SAMPLE      |
+--------------+------------------------+
| 400.00       | Novo salario: 460.00   |
|              | Reajuste ganho: 60.00  |
|              | Em percentual: 15 %    |
+--------------+------------------------+
| 800.01       | Novo salario: 880.01   |
|              | Reajuste ganho: 80.00  |
|              | Em percentual: 10 %    |
+--------------+------------------------+
| 2000.00      | Novo salario: 2140.00  |
|              | Reajuste ganho: 140.00 |
|              | Em percentual: 7 %     |
+--------------+------------------------+
*********************************************************************/`,
}

func TestNewDescription(t *testing.T) {
	for id, expect := range descriptionTests {
		get, err := NewDescription(id)
		if err != nil {
			t.Fatal(err)
		}

		if get != expect {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(expect, get, true)
			t.Error(dmp.DiffPrettyText(diffs))
		}
	}
}

func TestNewDescriptionFile(t *testing.T) {
	err := NewDescriptionFile(1015, "test/main.go")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDescriptions(t *testing.T) {
	files, err := ioutil.ReadDir("../")
	if err != nil {
		t.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			id, err := strconv.Atoi(strings.Split(f.Name(), " ")[0])
			if err != nil {
				continue
			}

			content, err := ioutil.ReadFile("../" + f.Name() + "/main.go")
			if err != nil {
				t.Fatal(err)
			}

			expect := strings.TrimSpace(strings.Split(string(content), "\npackage main")[0])

			get, err := NewDescription(id)
			if err != nil {
				t.Fatal(err)
			}

			if expect != get {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(get, expect, false)
				fmt.Println(dmp.DiffPrettyText(diffs))
				t.Fail()
			}

		}
	}
}
