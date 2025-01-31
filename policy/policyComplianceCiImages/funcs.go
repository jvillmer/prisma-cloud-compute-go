package policyComplianceCiImages

import (
	"fmt"
	"strings"

	pc "github.com/paloaltonetworks/prisma-cloud-compute-go"
)

// Get returns the policy that has the specified ID.
func Get(c pc.PrismaCloudClient) (Policy, error) {
	c.Log(pc.LogAction, "(get) %s", singular)

	var ans Policy
	_, err := c.Communicate("GET", Suffix, nil, nil, &ans)

	return ans, err
}

// Create adds a new policy.
func Create(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(false, c, policy)
}

// Update modifies the existing policy.
func Update(c pc.PrismaCloudClient, policy Policy) error {
	return createUpdate(true, c, policy)
}

func createUpdate(exists bool, c pc.PrismaCloudClient, policy Policy) error {
	var (
		logMsg strings.Builder
		method string
	)

	logMsg.Grow(30)
	logMsg.WriteString("(")
	method = "PUT"
	logMsg.WriteString(") ")

	logMsg.WriteString(singular)
	if exists {
		fmt.Fprintf(&logMsg, ":%s", policy.PolicyId)
	}

	c.Log(pc.LogAction, logMsg.String())

	_, err := c.Communicate(method, Suffix, nil, policy, nil)
	return err
}
