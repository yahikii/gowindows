package accounts_test

import (
	"context"
	"fmt"

	"github.com/d-strobel/gowindows/windows/local/accounts"
)

// We insert numbers into the function names to ensure that
// the test functions for each local_* file run in a specific order.
func (suite *LocalAccTestSuite) TestGroup1Read() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, c := range suite.clients {
		params := accounts.GroupReadParams{
			Name: "Users",
		}
		g, err := c.GroupRead(ctx, params)
		suite.Require().NoError(err)
		suite.Equal(accounts.Group{
			Name:        "Users",
			Description: "Users are prevented from making accidental or intentional system-wide changes and can run most applications",
			SID: accounts.SID{
				Value: "S-1-5-32-545",
			},
		}, g)
	}
}

func (suite *LocalAccTestSuite) TestGroup2List() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, c := range suite.clients {
		g, err := c.GroupList(ctx)
		suite.Require().NoError(err)
		suite.Contains(g, accounts.Group{
			Name:        "Users",
			Description: "Users are prevented from making accidental or intentional system-wide changes and can run most applications",
			SID: accounts.SID{
				Value: "S-1-5-32-545",
			},
		})
		suite.Contains(g, accounts.Group{
			Name:        "Administrators",
			Description: "Administrators have complete and unrestricted access to the computer/domain",
			SID: accounts.SID{
				Value: "S-1-5-32-544",
			},
		})
	}
}

func (suite *LocalAccTestSuite) TestGroup3Create() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, c := range suite.clients {
		params := accounts.GroupCreateParams{
			Name:        fmt.Sprintf("Test-Group-%d", i),
			Description: "This is a test group",
		}
		g, err := c.GroupCreate(ctx, params)
		suite.NoError(err)
		suite.Equal(accounts.Group{Name: fmt.Sprintf("Test-Group-%d", i)}.Name, g.Name)
		suite.Equal(accounts.Group{Description: "This is a test group"}.Description, g.Description)
	}
}

func (suite *LocalAccTestSuite) TestGroup4Update() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, c := range suite.clients {
		params := accounts.GroupUpdateParams{
			Name:        fmt.Sprintf("Test-Group-%d", i),
			Description: "This is a test group updated",
		}
		err := c.GroupUpdate(ctx, params)
		suite.NoError(err)
	}
}

func (suite *LocalAccTestSuite) TestGroup5Delete() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, c := range suite.clients {
		params := accounts.GroupDeleteParams{
			Name: fmt.Sprintf("Test-Group-%d", i),
		}
		err := c.GroupDelete(ctx, params)
		suite.NoError(err)
	}
}
