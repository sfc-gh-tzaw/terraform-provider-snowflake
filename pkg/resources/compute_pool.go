package resources

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/helpers"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var computePoolSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Specifies the identifier for the compute pool; must be unique for the account.",
	},
	"min_nodes": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Minimum number of nodes for the compute pool.",
	},
	"max_nodes": {
		Type:        schema.TypeInt,
		Required:    true,
		Description: "Maximum number of nodes for the compute pool.",
	},
	"instance_family": {
		Type:        schema.TypeString,
		Required:    true,
		ForceNew:    true,
		Description: "Type of machine you want to provision for the nodes.",
		// ExactlyOneOf: []string{""}, // can add the enum restriction here
	},
	"auto_resume": { // technically has a default value, but unsure if we also have to add it to terraform spec
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     true, // if unset, it won't be part of the create sql - but if set & unset, we should just unset target, no? so this is useless
		Description: "Whether to automatically resume a compute pool when a service or job is submitted to it.",
	},
	"initially_suspended": {
		Type:        schema.TypeBool,
		Optional:    true,
		Default:     false,
		ForceNew:    true,
		Description: "Whether the compute pool is created initially in the suspended state.",
	},
	"auto_suspend_secs": {
		Type:        schema.TypeInt,
		Optional:    true,
		Default:     600, // docs claim 3600 but not observed
		Description: "Number of seconds of inactivity after which pool suspends.",
	},
	"comment": {
		Type:        schema.TypeString,
		Optional:    true,
		Description: "Comment for the pool.",
	},
	// "state": {
	// 	Type:        schema.TypeString,
	// 	Computed:    true,
	// 	Description: "State of the pool.",
	// },
	// "num_services": {
	// 	Type:        schema.TypeInt,
	// 	Computed:    true,
	// 	Description: "Number of services running on the pool.",
	// },
	// "num_jobs": {
	// 	Type:        schema.TypeInt,
	// 	Computed:    true,
	// 	Description: "Number of jobs running on the pool.",
	// },
	// "active_nodes": {
	// 	Type:        schema.TypeInt,
	// 	Computed:    true,
	// 	Description: "Number of nodes active in the pool.",
	// },
	// "idle_nodes": {
	// 	Type:        schema.TypeInt,
	// 	Computed:    true,
	// 	Description: "Number of nodes active in the pool.",
	// },
	"created_on": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Time pool created.",
	},
	// "resumed_on": {
	// 	Type:        schema.TypeString,
	// 	Computed:    true,
	// 	Description: "Time pool resumed.",
	// },
	// "updated_on": {
	// 	Type:        schema.TypeString,
	// 	Computed:    true,
	// 	Description: "Time pool updated.",
	// },
	"owner": {
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Owner of the pool.",
	},
}

func ComputePool() *schema.Resource {
	return &schema.Resource{
		Create: CreateComputePool,
		Read:   ReadComputePool,
		Update: UpdateComputePool,
		Delete: DeleteComputePool,

		Schema: computePoolSchema,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

// CreateComputePool implements schema.CreateFunc.
func CreateComputePool(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	name := d.Get("name").(string)
	objectIdentifier := sdk.NewAccountObjectIdentifier(name)

	minNodes := d.Get("min_nodes").(int)
	maxNodes := d.Get("max_nodes").(int)
	instanceFamily := d.Get("instance_family").(string)

	createRequest := sdk.NewCreateComputePoolRequest(objectIdentifier, minNodes, maxNodes, instanceFamily)

	if v, ok := d.GetOk("auto_resume"); ok { // can also just use .Get() since the default will be false, but should be explicit methinks
		createRequest.WithAutoResume(sdk.Bool(v.(bool)))
	}

	if v, ok := d.GetOk("initially_suspended"); ok {
		createRequest.WithInitiallySuspended(sdk.Bool(v.(bool)))
	}

	if v, ok := d.GetOk("auto_suspend_secs"); ok {
		createRequest.WithAutoSuspendSecs(sdk.Int(v.(int)))
	}

	if v, ok := d.GetOk("comment"); ok {
		createRequest.WithComment(sdk.String(v.(string)))
	}

	// should consider if "if not exists" makes sense in terraform - methinks not, since that's the point lol

	if err := client.ComputePools.Create(ctx, createRequest); err != nil {
		return err
	}

	d.SetId(helpers.EncodeSnowflakeID(objectIdentifier))
	return ReadComputePool(d, meta)
}

// ReadComputePool implements schema.ReadFunc.
func ReadComputePool(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.AccountObjectIdentifier)

	c, err := client.ComputePools.ShowByID(ctx, id)
	if err != nil {
		return err
	}

	if err = d.Set("name", c.Name); err != nil {
		return err
	}
	// if err = d.Set("state", c.State); err != nil {
	// 	return err
	// }
	if err = d.Set("min_nodes", c.MinNodes); err != nil {
		return err
	}
	if err = d.Set("max_nodes", c.MaxNodes); err != nil {
		return err
	}
	if err = d.Set("instance_family", c.InstanceFamily); err != nil {
		return err
	}
	// if err = d.Set("num_services", c.NumServices); err != nil {
	// 	return err
	// }
	// if err = d.Set("num_jobs", c.NumJobs); err != nil {
	// 	return err
	// }
	if err = d.Set("auto_suspend_secs", c.AutoSuspendSecs); err != nil {
		return err
	}
	if err = d.Set("auto_resume", c.AutoResume); err != nil {
		return err
	}
	// if err = d.Set("active_nodes", c.ActiveNodes); err != nil {
	// 	return err
	// }
	// if err = d.Set("idle_nodes", c.IdleNodes); err != nil {
	// 	return err
	// }
	if err = d.Set("created_on", c.CreatedOn.Format("2006-01-02T16:04:05.000 -0700")); err != nil {
		return err
	}
	// if err = d.Set("resumed_on", c.ResumedOn); err != nil {
	// 	return err
	// }
	// if err = d.Set("updated_on", c.UpdatedOn); err != nil {
	// 	return err
	// }
	if err = d.Set("owner", c.Owner); err != nil {
		return err
	}
	if err = d.Set("comment", c.Comment); err != nil {
		return err
	}

	return nil
}

// UpdateComputePool implements schema.UpdateFunc.
func UpdateComputePool(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.AccountObjectIdentifier)

	alterSetPropsRequest := sdk.NewAlterSetPropertiesRequest()
	needSet := false

	alterUnsetPropsRequest := sdk.NewAlterUnsetPropertiesRequest()
	needUnset := false

	if d.HasChange("min_nodes") {
		minNodes := d.Get("min_nodes").(int)
		needSet = true
		alterSetPropsRequest.WithMinNodes(sdk.Int(minNodes))
	}

	if d.HasChange("max_nodes") {
		maxNodes := d.Get("max_nodes").(int)
		needSet = true
		alterSetPropsRequest.WithMaxNodes(sdk.Int(maxNodes))
	}

	if d.HasChange("auto_resume") {
		if v, ok := d.GetOk("auto_resume"); ok {
			needSet = true
			alterSetPropsRequest.WithAutoResume(sdk.Bool(v.(bool)))
		} else {
			needUnset = true
			alterUnsetPropsRequest.WithAutoResume(sdk.Bool(true))
		}
	}

	if d.HasChange("auto_suspend_secs") {
		if v, ok := d.GetOk("auto_suspend_secs"); ok {
			needSet = true
			alterSetPropsRequest.WithAutoSuspendSecs(sdk.Int(v.(int)))
		} else {
			needUnset = true
			alterUnsetPropsRequest.WithAutoSuspendSecs(sdk.Bool(true))
		}
	}

	if d.HasChange("comment") {
		if v, ok := d.GetOk("comment"); ok {
			needSet = true
			alterSetPropsRequest.WithComment(sdk.String(v.(string)))
		} else {
			needUnset = true
			alterUnsetPropsRequest.WithComment(sdk.Bool(true))
		}
	}

	if needUnset {
		if err := client.ComputePools.Alter(ctx, sdk.NewAlterComputePoolRequest(id).WithUnset(alterUnsetPropsRequest)); err != nil {
			return err
		}
	}

	if needSet {
		if err := client.ComputePools.Alter(ctx, sdk.NewAlterComputePoolRequest(id).WithSet(alterSetPropsRequest)); err != nil {
			return err
		}
	}

	return ReadComputePool(d, meta)
}

// DeleteComputePool implements schema.DeleteFunc.
func DeleteComputePool(d *schema.ResourceData, meta interface{}) error {
	db := meta.(*sql.DB)
	client := sdk.NewClientFromDB(db)
	ctx := context.Background()

	id := helpers.DecodeSnowflakeID(d.Id()).(sdk.AccountObjectIdentifier)

	if err := client.ComputePools.Drop(ctx, sdk.NewDropComputePoolRequest(id)); err != nil {
		return fmt.Errorf("error deleting stream %v err = %w", d.Id(), err)
	}

	d.SetId("")

	return nil
}
