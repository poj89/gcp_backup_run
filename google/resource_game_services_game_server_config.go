// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGameServicesGameServerConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceGameServicesGameServerConfigCreate,
		Read:   resourceGameServicesGameServerConfigRead,
		Delete: resourceGameServicesGameServerConfigDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGameServicesGameServerConfigImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"config_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `A unique id for the deployment config.`,
			},
			"deployment_id": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A unique id for the deployment.`,
			},
			"fleet_configs": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Description: `The fleet config contains list of fleet specs. In the Single Cloud, there
will be only one.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fleet_spec": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The fleet spec, which is sent to Agones to configure fleet.
The spec can be passed as inline json but it is recommended to use a file reference
instead. File references can contain the json or yaml format of the fleet spec. Eg:

* fleet_spec = jsonencode(yamldecode(file("fleet_configs.yaml")))
* fleet_spec = file("fleet_configs.json")

The format of the spec can be found :
'https://agones.dev/site/docs/reference/fleet/'.`,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Optional:    true,
							ForceNew:    true,
							Description: `The name of the FleetConfig.`,
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `The description of the game server config.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Description: `The labels associated with this game server config. Each label is a
key-value pair.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"location": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `Location of the Deployment.`,
				Default:     "global",
			},
			"scaling_configs": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Optional. This contains the autoscaling settings.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fleet_autoscaler_spec": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `Fleet autoscaler spec, which is sent to Agones.
Example spec can be found :
https://agones.dev/site/docs/reference/fleetautoscaler/`,
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: `The name of the ScalingConfig`,
						},
						"schedules": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `The schedules to which this scaling config applies.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cron_job_duration": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Description: `The duration for the cron job event. The duration of the event is effective
after the cron job's start time.

A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".`,
									},
									"cron_spec": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Description: `The cron definition of the scheduled event. See
https://en.wikipedia.org/wiki/Cron. Cron spec specifies the local time as
defined by the realm.`,
									},
									"end_time": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Description: `The end time of the event.

A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
									},
									"start_time": {
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Description: `The start time of the event.

A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".`,
									},
								},
							},
						},
						"selectors": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `Labels used to identify the clusters to which this scaling config
applies. A cluster is subject to this scaling config if its labels match
any of the selector entries.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"labels": {
										Type:        schema.TypeMap,
										Optional:    true,
										ForceNew:    true,
										Description: `Set of labels to group by.`,
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the game server config, in the form:

'projects/{project_id}/locations/{location}/gameServerDeployments/{deployment_id}/configs/{config_id}'.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceGameServicesGameServerConfigCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandGameServicesGameServerConfigDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandGameServicesGameServerConfigLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	fleetConfigsProp, err := expandGameServicesGameServerConfigFleetConfigs(d.Get("fleet_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fleet_configs"); !isEmptyValue(reflect.ValueOf(fleetConfigsProp)) && (ok || !reflect.DeepEqual(v, fleetConfigsProp)) {
		obj["fleetConfigs"] = fleetConfigsProp
	}
	scalingConfigsProp, err := expandGameServicesGameServerConfigScalingConfigs(d.Get("scaling_configs"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scaling_configs"); !isEmptyValue(reflect.ValueOf(scalingConfigsProp)) && (ok || !reflect.DeepEqual(v, scalingConfigsProp)) {
		obj["scalingConfigs"] = scalingConfigsProp
	}

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs?configId={{config_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new GameServerConfig: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating GameServerConfig: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = gameServicesOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating GameServerConfig", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create GameServerConfig: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerConfigName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating GameServerConfig %q: %#v", d.Id(), res)

	return resourceGameServicesGameServerConfigRead(d, meta)
}

func resourceGameServicesGameServerConfigRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerConfig: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("GameServicesGameServerConfig %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}

	if err := d.Set("name", flattenGameServicesGameServerConfigName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}
	if err := d.Set("description", flattenGameServicesGameServerConfigDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}
	if err := d.Set("labels", flattenGameServicesGameServerConfigLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}
	if err := d.Set("fleet_configs", flattenGameServicesGameServerConfigFleetConfigs(res["fleetConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}
	if err := d.Set("scaling_configs", flattenGameServicesGameServerConfigScalingConfigs(res["scalingConfigs"], d, config)); err != nil {
		return fmt.Errorf("Error reading GameServerConfig: %s", err)
	}

	return nil
}

func resourceGameServicesGameServerConfigDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for GameServerConfig: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting GameServerConfig %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "GameServerConfig")
	}

	err = gameServicesOperationWaitTime(
		config, res, project, "Deleting GameServerConfig", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting GameServerConfig %q: %#v", d.Id(), res)
	return nil
}

func resourceGameServicesGameServerConfigImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/gameServerDeployments/(?P<deployment_id>[^/]+)/configs/(?P<config_id>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<deployment_id>[^/]+)/(?P<config_id>[^/]+)",
		"(?P<location>[^/]+)/(?P<deployment_id>[^/]+)/(?P<config_id>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/gameServerDeployments/{{deployment_id}}/configs/{{config_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenGameServicesGameServerConfigName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigFleetConfigs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"fleet_spec": flattenGameServicesGameServerConfigFleetConfigsFleetSpec(original["fleetSpec"], d, config),
			"name":       flattenGameServicesGameServerConfigFleetConfigsName(original["name"], d, config),
		})
	}
	return transformed
}
func flattenGameServicesGameServerConfigFleetConfigsFleetSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigFleetConfigsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigs(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"name":                  flattenGameServicesGameServerConfigScalingConfigsName(original["name"], d, config),
			"fleet_autoscaler_spec": flattenGameServicesGameServerConfigScalingConfigsFleetAutoscalerSpec(original["fleetAutoscalerSpec"], d, config),
			"selectors":             flattenGameServicesGameServerConfigScalingConfigsSelectors(original["selectors"], d, config),
			"schedules":             flattenGameServicesGameServerConfigScalingConfigsSchedules(original["schedules"], d, config),
		})
	}
	return transformed
}
func flattenGameServicesGameServerConfigScalingConfigsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsFleetAutoscalerSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsSelectors(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"labels": flattenGameServicesGameServerConfigScalingConfigsSelectorsLabels(original["labels"], d, config),
		})
	}
	return transformed
}
func flattenGameServicesGameServerConfigScalingConfigsSelectorsLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsSchedules(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"start_time":        flattenGameServicesGameServerConfigScalingConfigsSchedulesStartTime(original["startTime"], d, config),
			"end_time":          flattenGameServicesGameServerConfigScalingConfigsSchedulesEndTime(original["endTime"], d, config),
			"cron_job_duration": flattenGameServicesGameServerConfigScalingConfigsSchedulesCronJobDuration(original["cronJobDuration"], d, config),
			"cron_spec":         flattenGameServicesGameServerConfigScalingConfigsSchedulesCronSpec(original["cronSpec"], d, config),
		})
	}
	return transformed
}
func flattenGameServicesGameServerConfigScalingConfigsSchedulesStartTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsSchedulesEndTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsSchedulesCronJobDuration(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenGameServicesGameServerConfigScalingConfigsSchedulesCronSpec(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandGameServicesGameServerConfigDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGameServicesGameServerConfigFleetConfigs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedFleetSpec, err := expandGameServicesGameServerConfigFleetConfigsFleetSpec(original["fleet_spec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFleetSpec); val.IsValid() && !isEmptyValue(val) {
			transformed["fleetSpec"] = transformedFleetSpec
		}

		transformedName, err := expandGameServicesGameServerConfigFleetConfigsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerConfigFleetConfigsFleetSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigFleetConfigsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigs(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedName, err := expandGameServicesGameServerConfigScalingConfigsName(original["name"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
			transformed["name"] = transformedName
		}

		transformedFleetAutoscalerSpec, err := expandGameServicesGameServerConfigScalingConfigsFleetAutoscalerSpec(original["fleet_autoscaler_spec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFleetAutoscalerSpec); val.IsValid() && !isEmptyValue(val) {
			transformed["fleetAutoscalerSpec"] = transformedFleetAutoscalerSpec
		}

		transformedSelectors, err := expandGameServicesGameServerConfigScalingConfigsSelectors(original["selectors"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSelectors); val.IsValid() && !isEmptyValue(val) {
			transformed["selectors"] = transformedSelectors
		}

		transformedSchedules, err := expandGameServicesGameServerConfigScalingConfigsSchedules(original["schedules"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedSchedules); val.IsValid() && !isEmptyValue(val) {
			transformed["schedules"] = transformedSchedules
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerConfigScalingConfigsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigsFleetAutoscalerSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigsSelectors(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedLabels, err := expandGameServicesGameServerConfigScalingConfigsSelectorsLabels(original["labels"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedLabels); val.IsValid() && !isEmptyValue(val) {
			transformed["labels"] = transformedLabels
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerConfigScalingConfigsSelectorsLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandGameServicesGameServerConfigScalingConfigsSchedules(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedStartTime, err := expandGameServicesGameServerConfigScalingConfigsSchedulesStartTime(original["start_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
			transformed["startTime"] = transformedStartTime
		}

		transformedEndTime, err := expandGameServicesGameServerConfigScalingConfigsSchedulesEndTime(original["end_time"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !isEmptyValue(val) {
			transformed["endTime"] = transformedEndTime
		}

		transformedCronJobDuration, err := expandGameServicesGameServerConfigScalingConfigsSchedulesCronJobDuration(original["cron_job_duration"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCronJobDuration); val.IsValid() && !isEmptyValue(val) {
			transformed["cronJobDuration"] = transformedCronJobDuration
		}

		transformedCronSpec, err := expandGameServicesGameServerConfigScalingConfigsSchedulesCronSpec(original["cron_spec"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedCronSpec); val.IsValid() && !isEmptyValue(val) {
			transformed["cronSpec"] = transformedCronSpec
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandGameServicesGameServerConfigScalingConfigsSchedulesStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigsSchedulesEndTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigsSchedulesCronJobDuration(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandGameServicesGameServerConfigScalingConfigsSchedulesCronSpec(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
