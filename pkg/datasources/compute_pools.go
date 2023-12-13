package datasources

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/snowflake"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var computePoolsSchema = map[string]*schema.Schema{
	"compute_pools": {
		Type:        schema.TypeList,
		Computed:    true,
		Description: "The compute pools in your account",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"state": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"instance_family": {
					Type:     schema.TypeString,
					Computed: true,
				},
				"comment": {
					Type:     schema.TypeString,
					Computed: true,
				},
			},
		},
	},
}

func ComputePools() *schema.Resource {
	return &schema.Resource{
		Read:   ReadComputePools,
		Schema: computePoolsSchema,
	}
}

func ReadComputePools(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	account, err := snowflake.ReadCurrentAccount(db)
	if err != nil {
		d.SetId("")
		return nil
	}
	d.SetId(fmt.Sprintf("%s.%s", account.Account, account.Region))

	extractedComputePools, err := client.ComputePools.Show(ctx, sdk.NewShowComputePoolRequest())
	if err != nil {
		// If not found, mark resource to be removed from state file during apply or refresh
		log.Printf("[DEBUG] compute pools in id (%s) not found", d.Id())
		d.SetId("")
		return nil
	}

	computePools := make([]map[string]any, 0, len(extractedComputePools))
	for i, computePool := range extractedComputePools {
		computePools[i] = map[string]any{
			"name":            computePool.Name,
			"state":           computePool.State,
			"instance_family": computePool.InstanceFamily,
			"comment":         computePool.Comment,
		}
	}

	return d.Set("compute_pools", computePools)
}
