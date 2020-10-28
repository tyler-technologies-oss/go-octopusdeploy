package examples

import (
	"fmt"
	"net/url"

	"github.com/OctopusDeploy/go-octopusdeploy/octopusdeploy"
)

func UpdateLifecycleExample() {
	var (
		apiKey     string = "API-YOUR_API_KEY"
		octopusURL string = "https://your_octopus_url"
		spaceID    string = "space-id"

		// lifecycle values
		lifecycleID string = "lifecycle-id"
	)

	apiURL, err := url.Parse(octopusURL)
	if err != nil {
		_ = fmt.Errorf("error parsing URL for Octopus API: %v", err)
		return
	}

	client, err := octopusdeploy.NewClient(nil, apiURL, apiKey, spaceID)
	if err != nil {
		_ = fmt.Errorf("error creating API client: %v", err)
		return
	}

	// get lifecycle by its ID
	lifecycle, err := client.Lifecycles.GetByID(lifecycleID)
	if err != nil {
		_ = fmt.Errorf("error getting lifecycle: %v", err)
		return
	}

	// update lifecycle fields
	lifecycle.Description = "new-description"

	// update lifecycle
	updatedLifecycle, err := client.Lifecycles.Update(lifecycle)
	if err != nil {
		_ = fmt.Errorf("error updating lifecycle: %v", err)
		return
	}

	fmt.Printf("lifecycle updated: (%s)\n", updatedLifecycle.GetModifiedOn())
}
