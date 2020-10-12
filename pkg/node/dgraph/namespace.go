//--------------------------------------------------------------------------
// Copyright 2018 Infinite Devices GmbH
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package dgraph

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgraph-io/dgo/protos/api"
	"github.com/infinimesh/infinimesh/pkg/node/nodepb"
	"google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//CreateNamespace is a method to create Namespaces
func (s *DGraphRepo) CreateNamespace(ctx context.Context, name string) (id string, err error) {

	//JSON for creating the node in Dgraph DB
	ns := &Namespace{
		Node: Node{
			Type: "namespace",
			UID:  "_:namespace",
		},
		Name:                 name,
		MarkForDeletion:      false,
		DeleteInitiationTime: "0000-01-01T00:00:00Z",
	}

	txn := s.Dg.NewTxn()
	js, err := json.Marshal(&ns)
	if err != nil {
		return "", err
	}

	assigned, err := txn.Mutate(ctx, &api.Mutation{
		SetJson:   js,
		CommitNow: true,
	})
	if err != nil {
		return "", err
	}

	return assigned.GetUids()["namespace"], nil
}

//GetNamespace is a method to get the namespace based on Name
func (s *DGraphRepo) GetNamespace(ctx context.Context, namespacename string) (namespace *nodepb.Namespace, err error) {
	const q = `query getNamespaces($namespace: string) {
                     namespaces(func: eq(name, $namespace)) @filter(eq(type, "namespace"))  {
	               	uid
					name
					markfordeletion
					deleteinitiationtime
	             }
                   }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$namespace": namespacename})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	if len(resultSet.Namespaces) > 0 {
		return &nodepb.Namespace{
			Id:                   resultSet.Namespaces[0].UID,
			Name:                 resultSet.Namespaces[0].Name,
			Markfordeletion:      resultSet.Namespaces[0].MarkForDeletion,
			Deleteinitiationtime: resultSet.Namespaces[0].DeleteInitiationTime,
		}, nil
	}

	return nil, errors.New("The Namespace is not found")
}

//GetNamespaceID is a method to get the namespace based on ID
func (s *DGraphRepo) GetNamespaceID(ctx context.Context, namespaceID string) (namespace *nodepb.Namespace, err error) {
	const q = `query getNamespaces($namespaceid: string) {
                     namespaces(func: uid($namespaceid)) @filter(eq(type, "namespace"))  {
						uid
						name
						markfordeletion
						deleteinitiationtime
	             }
                   }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$namespaceid": namespaceID})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	if len(resultSet.Namespaces) > 0 {
		return &nodepb.Namespace{
			Id:                   resultSet.Namespaces[0].UID,
			Name:                 resultSet.Namespaces[0].Name,
			Markfordeletion:      resultSet.Namespaces[0].MarkForDeletion,
			Deleteinitiationtime: resultSet.Namespaces[0].DeleteInitiationTime,
		}, nil
	}

	return nil, errors.New("The Namespace is not found")
}

//ListNamespaces is a method to List details of all Namespaces
func (s *DGraphRepo) ListNamespaces(ctx context.Context) (namespaces []*nodepb.Namespace, err error) {
	const q = `{
                     namespaces(func: eq(type, "namespace")) {
						uid
						name
						markfordeletion
						deleteinitiationtime
	             }
                   }`

	res, err := s.Dg.NewReadOnlyTxn().Query(ctx, q)
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, namespace := range resultSet.Namespaces {
		namespaces = append(namespaces, &nodepb.Namespace{
			Id:                   namespace.UID,
			Name:                 namespace.Name,
			Markfordeletion:      namespace.MarkForDeletion,
			Deleteinitiationtime: namespace.DeleteInitiationTime,
		})
	}

	return namespaces, nil
}

//DeletePermissionInNamespace is a method to delete permissions for a Namespaces for an account
func (s *DGraphRepo) DeletePermissionInNamespace(ctx context.Context, namespaceID, accountID string) (err error) {
	txn := s.Dg.NewTxn()
	const q = `query deletePermissionInNamespace($namespaceID: string, $accountID: string){
  accounts(func: uid($namespaceID)) @filter(eq(type, "namespace")) @cascade @normalize {
    namespace_uid: uid
    ~access.to.namespace @filter(uid($accountID))  {
      account_uid: uid
    }
  }
}`

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
		"$accountID":   accountID,
	})
	if err != nil {
		return err
	}

	var resultSet struct {
		Accounts []*struct {
			AccountUID   string `json:"account_uid"`
			NamespaceUID string `json:"namespace_uid"`
		} `json:"accounts"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return err
	}

	m := &api.Mutation{CommitNow: true}
	for _, account := range resultSet.Accounts {
		m.Del = append(m.Del, &api.NQuad{
			Subject:   account.AccountUID,
			Predicate: "access.to.namespace",
			ObjectId:  account.NamespaceUID,
		})
	}
	_, err = txn.Mutate(ctx, m)
	return err
}

//ListPermissionsInNamespace is a method to list permissions for all accounts for a Namespace
func (s *DGraphRepo) ListPermissionsInNamespace(ctx context.Context, namespaceid string) (permissions []*nodepb.Permission, err error) {
	const q = `query listPermissionsInNamespace($namespaceid: string) {
		accounts(func: uid($namespaceid)) @filter(eq(type, "namespace")) @normalize @cascade  {
		  ~access.to.namespace {
			uid: uid
			name: name
		  } @facets(permission)
		}
	  }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$namespaceid": namespaceid})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Accounts []*struct {
			UID    string `json:"uid"`
			Name   string `json:"name"`
			Action string `json:"~access.to.namespace|permission"`
		} `json:"accounts"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, account := range resultSet.Accounts {
		permissions = append(permissions, &nodepb.Permission{
			Namespace:   namespaceid,
			AccountId:   account.UID,
			AccountName: account.Name,
			Action:      nodepb.Action(nodepb.Action_value[account.Action]),
		})
	}

	return permissions, nil

}

//ListNamespacesForAccount is a method to list Namespaces for an Account
func (s *DGraphRepo) ListNamespacesForAccount(ctx context.Context, accountID string) (namespaces []*nodepb.Namespace, err error) {
	const q = `query listNamespaces($account: string) {
		namespaces(func: uid($account)) @normalize @cascade  {
		  access.to.namespace @filter(eq(type, "namespace")) @facets(NOT eq(permission,"NONE")) {
			uid : uid
			name : name
			markfordeletion : markfordeletion
			deleteinitiationtime : deleteinitiationtime
		  }
		}
	  }`

	res, err := s.Dg.NewReadOnlyTxn().QueryWithVars(ctx, q, map[string]string{"$account": accountID})
	if err != nil {
		return nil, err
	}

	var resultSet struct {
		Namespaces []*Namespace `json:"namespaces"`
	}

	if err := json.Unmarshal(res.Json, &resultSet); err != nil {
		return nil, err
	}

	for _, namespace := range resultSet.Namespaces {
		namespaces = append(namespaces, &nodepb.Namespace{
			Id:                   namespace.UID,
			Name:                 namespace.Name,
			Markfordeletion:      namespace.MarkForDeletion,
			Deleteinitiationtime: namespace.DeleteInitiationTime,
		})
	}

	return namespaces, nil
}

//IsAuthorizedNamespace is a method that returns if the access to the namespace for an account is true or false
func (s *DGraphRepo) IsAuthorizedNamespace(ctx context.Context, namespaceid, account string, action nodepb.Action) (decision bool, err error) {
	acc, err := s.GetAccount(ctx, account)
	if err != nil {
		return false, err
	}

	if acc.IsRoot {
		return true, nil
	}

	params := map[string]string{
		"$namespaceid": namespaceid,
		"$user_id":     account,
	}

	txn := s.Dg.NewReadOnlyTxn()

	const q = `query access($namespaceid: string, $user_id: string){
		access(func: uid($user_id)) @normalize @cascade {
		  name
		  uid
		  access.to.namespace @filter(uid($namespaceid)) @facets(permission,inherit) {
			uid
			name
			type
		  }
		}
	  }
`

	res, err := txn.QueryWithVars(ctx, q, params)
	if err != nil {
		return false, err
	}
	var access struct {
		Access []Namespace `json:"access"`
	}

	err = json.Unmarshal(res.Json, &access)
	if err != nil {
		return false, err
	}

	actionValue := strings.Split(action.String(), "_")

	if len(access.Access) > 0 {
		if isPermissionSufficient(actionValue[0], access.Access[0].AccessToPermission) {
			return true, nil
		}
	}

	return false, nil
}

//SoftDeleteNamespace is a method that mark the namespace for deletion
func (s *DGraphRepo) SoftDeleteNamespace(ctx context.Context, namespaceID string) (err error) {
	txn := s.Dg.NewReadOnlyTxn()

	const q = `query deleteNodes($namespaceID: string){
        nodes(func: uid($namespaceID)) @filter(eq(type,"namespace")) {
          uid
        }
      }
      `

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
	})
	if err != nil {
		return err
	}

	var result struct {
		Nodes []*Node `json:"nodes"`
	}

	err = json.Unmarshal(res.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Nodes) != 1 {
		return status.Error(codes.NotFound, "The Namespace is not found")
	}

	err = s.UpdateNamespace(ctx, &nodepb.UpdateNamespaceRequest{
		Namespace: &nodepb.Namespace{
			Id:                   namespaceID,
			Markfordeletion:      true,
			Deleteinitiationtime: time.Now().Format(time.RFC3339),
		},
		NamespaceMask: &field_mask.FieldMask{
			Paths: []string{"markfordeletion", "deleteinitiationtime"},
		},
	})
	return err
}

//HardDeleteNamespace is a method that deletes a namespace permantly
func (s *DGraphRepo) HardDeleteNamespace(ctx context.Context, datecondition string) (err error) {
	txn := s.Dg.NewReadOnlyTxn()
	var q = `query deleteNodes{
        nodes(func: eq(type,"namespace")) @filter(eq(markfordeletion,"true") AND lt(deleteinitiationtime,"%v")) @normalize {
          uid
        owns {
          uid
        }
        }
      }
      `

	if datecondition != "" {
		q = fmt.Sprintf(q, datecondition)
	} else {
		q = fmt.Sprintf(q, "")
	}

	res, err := txn.Query(ctx, q)
	if err != nil {
		return err
	}

	var result struct {
		Nodes []*Node `json:"nodes"`
	}

	err = json.Unmarshal(res.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Nodes) < 1 {
		return status.Error(codes.NotFound, "The Namespace is not found")
	}

	for _, item := range result.Nodes {
		err = s.DeleteObject(ctx, item.UID)
	}

	return err
}

//RevokeNamespace is a method that nmarks the Namespace from deletion
func (s *DGraphRepo) RevokeNamespace(ctx context.Context, namespaceID string) (err error) {
	txn := s.Dg.NewReadOnlyTxn()

	const q = `query revokeNodes($namespaceID: string){
        nodes(func: uid($namespaceID)) @filter(eq(type,"namespace")) {
          uid
        }
      }
      `

	res, err := txn.QueryWithVars(ctx, q, map[string]string{
		"$namespaceID": namespaceID,
	})
	if err != nil {
		return err
	}

	var result struct {
		Nodes []*Node `json:"nodes"`
	}

	err = json.Unmarshal(res.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Nodes) != 1 {
		return status.Error(codes.NotFound, "The Namespace is not found")
	}

	err = s.UpdateNamespace(ctx, &nodepb.UpdateNamespaceRequest{
		Namespace: &nodepb.Namespace{
			Id:                   namespaceID,
			Markfordeletion:      false,
			Deleteinitiationtime: "0000",
		},
		NamespaceMask: &field_mask.FieldMask{
			Paths: []string{"markfordeletion", "deleteinitiationtime"},
		},
	})
	return err
}

//UpdateNamespace is a method to Udpdate details of an Namespace
func (s *DGraphRepo) UpdateNamespace(ctx context.Context, namespace *nodepb.UpdateNamespaceRequest) (err error) {

	txn := s.Dg.NewTxn()
	m := &api.Mutation{CommitNow: true}

	q := `query namespaceExists($namespaceid: string) {
                exists(func: uid($namespaceid)) @filter(eq(type, "namespace")) {
                  uid
                }
              }
             `

	var result struct {
		Exists []map[string]interface{} `json:"exists"`
	}

	resp, err := txn.QueryWithVars(ctx, q, map[string]string{"$namespaceid": namespace.Namespace.Id})
	if err != nil {
		return err
	}
	err = json.Unmarshal(resp.Json, &result)
	if err != nil {
		return err
	}

	if len(result.Exists) == 0 {
		return errors.New("The Namespace is not found")
	}

	//Loop through the field masks and update the required fields
	for _, field := range namespace.NamespaceMask.Paths {
		switch strings.ToLower(field) {
		case "name":
			m.Set = append(m.Set, &api.NQuad{
				Subject:     namespace.Namespace.Id,
				Predicate:   "name",
				ObjectId:    namespace.Namespace.Id,
				ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: namespace.Namespace.Name}},
			})
		case "markfordeletion":
			m.Set = append(m.Set, &api.NQuad{
				Subject:     namespace.Namespace.Id,
				Predicate:   "markfordeletion",
				ObjectId:    namespace.Namespace.Id,
				ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: strconv.FormatBool(namespace.Namespace.Markfordeletion)}},
			})
			//Reset the deleteinitiationtime to 00 when markfordeletion is set to false ie. revoke is issued
			if !namespace.Namespace.Markfordeletion {
				m.Set = append(m.Set, &api.NQuad{
					Subject:     namespace.Namespace.Id,
					Predicate:   "deleteinitiationtime",
					ObjectId:    namespace.Namespace.Id,
					ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: "0000-01-01T00:00:00Z"}},
				})
			}
		case "deleteinitiationtime":
			m.Set = append(m.Set, &api.NQuad{
				Subject:     namespace.Namespace.Id,
				Predicate:   "deleteinitiationtime",
				ObjectId:    namespace.Namespace.Id,
				ObjectValue: &api.Value{Val: &api.Value_DefaultVal{DefaultVal: namespace.Namespace.Deleteinitiationtime}},
			})
		}
	}

	_, err = txn.Mutate(ctx, m)
	if err != nil {
		return err
	}

	return nil
}
