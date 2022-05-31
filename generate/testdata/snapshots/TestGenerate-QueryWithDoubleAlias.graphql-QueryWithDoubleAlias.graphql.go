// Code generated by github.com/opiuman/genqlient, DO NOT EDIT.

package test

import (
	"github.com/opiuman/genqlient/graphql"
	"github.com/opiuman/genqlient/internal/testutil"
)

// QueryWithDoubleAliasResponse is returned by QueryWithDoubleAlias on success.
type QueryWithDoubleAliasResponse struct {
	// user looks up a user by some stuff.
	//
	// See UserQueryInput for what stuff is supported.
	// If query is null, returns the current user.
	User QueryWithDoubleAliasUser `json:"user"`
}

// GetUser returns QueryWithDoubleAliasResponse.User, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasResponse) GetUser() QueryWithDoubleAliasUser { return v.User }

// QueryWithDoubleAliasUser includes the requested fields of the GraphQL type User.
// The GraphQL type's documentation follows.
//
// A User is a user!
type QueryWithDoubleAliasUser struct {
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	ID testutil.ID `json:"ID"`
	// id is the user's ID.
	//
	// It is stable, unique, and opaque, like all good IDs.
	AlsoID testutil.ID `json:"AlsoID"`
}

// GetID returns QueryWithDoubleAliasUser.ID, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasUser) GetID() testutil.ID { return v.ID }

// GetAlsoID returns QueryWithDoubleAliasUser.AlsoID, and is useful for accessing the field via an interface.
func (v *QueryWithDoubleAliasUser) GetAlsoID() testutil.ID { return v.AlsoID }

func QueryWithDoubleAlias(
	client graphql.Client,
	opts ...graphql.RequestOption,
) (*QueryWithDoubleAliasResponse, error) {
	req := &graphql.Request{
		OpName: "QueryWithDoubleAlias",
		Query: `
query QueryWithDoubleAlias {
	user {
		ID: id
		AlsoID: id
	}
}
`,
	}
	var err error

	var data QueryWithDoubleAliasResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		nil,
		req,
		resp,
		opts...,
	)

	return &data, err
}
