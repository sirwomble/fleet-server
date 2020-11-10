// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package dsl

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Emulate agent_id saved object query
func makeQuery(leaf interface{}) *Node {
	const ty = "fleet-agents"

	root := NewRoot()
	mustNode := root.Query().Bool().Must()
	mustNode.Term("type", ty, nil)
	mustNode.Term("fleet-agents.access_api_key_id", leaf, nil)
	return root
}

func makeQuery2(leaf1 interface{}, leaf2 interface{}) *Node {
	const ty = "fleet-agent-actions"

	root := NewRoot()
	root.Size(1)
	root.Sort().SortOrder("fleet-agent-actions.created_at", SortDescend)

	mustNode := root.Query().Bool().Must()
	mustNode.Term("type", ty, nil)
	mustNode.Term("fleet-agent-actions.policy_id", leaf1, nil)
	mustNode.Range("fleet-agent-actions.policy_revision", WithRangeGT(leaf2))
	return root
}

func BenchmarkRenderOne(b *testing.B) {

	const kName = "api_key"
	tmpl := NewTmpl()
	token := tmpl.Bind(kName)

	query := makeQuery(token)

	if err := tmpl.Resolve(query); err != nil {
		panic(err)
	}

	// run the RenderOne function b.N times
	for n := 0; n < b.N; n++ {
		tmpl.RenderOne(kName, "2Ye0F3UByTc0c1e9OeMO")
	}
}

func BenchmarkRender(b *testing.B) {

	const kName = "api_key"
	tmpl := NewTmpl()
	token := tmpl.Bind(kName)

	query := makeQuery(token)

	if err := tmpl.Resolve(query); err != nil {
		panic(err)
	}

	v := "2Ye0F3UByTc0c1e9OeMO"

	// run the RenderOne function b.N times
	for n := 0; n < b.N; n++ {
		tmpl.Render(map[string]interface{}{
			kName: v,
		})
	}
}

func BenchmarkMarshalNode(b *testing.B) {

	// run the RenderOne function b.N times
	for n := 0; n < b.N; n++ {
		query := makeQuery("2Ye0F3UByTc0c1e9OeMO")
		json.Marshal(query)
	}
}

func BenchmarkMarshalNode2(b *testing.B) {

	// run the RenderOne function b.N times
	for n := 0; n < b.N; n++ {
		query := makeQuery2("27e58fc0-09a2-11eb-a8cd-57e98f140de5", 3)
		json.Marshal(query)
	}
}

var ssprintres string

func BenchmarkSprintf(b *testing.B) {
	queryTmpl := `{"size": 1,"sort": [{"fleet-agent-actions.created_at": {"order": "DESC"}}],"query": {"bool": {"must": [{"term": {"type": "fleet-agent-actions"}},{"term": {"fleet-agent-actions.policy_id": "%s"}},{"range": {"fleet-agent-actions.policy_revision": {"gt": %d}}}]}}}`

	policyId := "27e58fc0-09a2-11eb-a8cd-57e98f140de5"
	policyRev := 3

	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(queryTmpl, policyId, policyRev)
	}
	ssprintres = s
}

func BenchmarkRender2(b *testing.B) {

	const kName1 = "policyId"
	const kName2 = "policyRev"

	tmpl := NewTmpl()
	token1 := tmpl.Bind(kName1)
	token2 := tmpl.Bind(kName2)

	query := makeQuery2(token1, token2)

	if err := tmpl.Resolve(query); err != nil {
		panic(err)
	}

	// run the RenderOne function b.N times
	for n := 0; n < b.N; n++ {
		m := map[string]interface{}{
			kName1: "27e58fc0-09a2-11eb-a8cd-57e98f140de5",
			kName2: 3,
		}

		tmpl.Render(m)
	}
}
