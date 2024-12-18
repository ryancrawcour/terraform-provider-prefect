package testutils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func GetResourceWorkspaceImportStateID(resourceName string) resource.ImportStateIdFunc {
	return func(state *terraform.State) (string, error) {
		workspace, exists := state.RootModule().Resources[WorkspaceResourceName]
		if !exists {
			return "", fmt.Errorf("resource not found in state: %s", WorkspaceResourceName)
		}
		workspaceID, _ := uuid.Parse(workspace.Primary.ID)

		fetchedResource, exists := state.RootModule().Resources[resourceName]
		if !exists {
			return "", fmt.Errorf("resource not found in state: %s", resourceName)
		}

		return fmt.Sprintf("%s,%s", fetchedResource.Primary.Attributes["id"], workspaceID), nil
	}
}
