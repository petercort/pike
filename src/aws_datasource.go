package pike

import (
	"fmt"
)

// GetAWSDataPermissions gets permissions required for datasource's.
//
//goland:noinspection GoLinter,GoLinter
//nolint:funlen
func GetAWSDataPermissions(result ResourceV2) ([]string, error) { //nolint:maintidx
	TFLookup := map[string]interface{}{
		"aws_acm_certificate":                                  dataAwsAcmCertificate,
		"aws_acmpca_certificate":                               dataAwsAcmpcaCertificate,
		"aws_acmpca_certificate_authority":                     dataAwsAcmpcaCertificateAuthority,
		"aws_alb":                                              dataAwsLb,
		"aws_ami":                                              dataAwsAmi,
		"aws_ami_ids":                                          dataAwsAmiIds,
		"aws_api_gateway_api_key":                              dataAwsAPIGateway,
		"aws_api_gateway_authorizer":                           dataAwsAPIGateway,
		"aws_api_gateway_authorizers":                          dataAwsAPIGateway,
		"aws_api_gateway_domain_name":                          dataAwsAPIGateway,
		"aws_api_gateway_export":                               dataAwsAPIGateway,
		"aws_api_gateway_resource":                             dataAwsAPIGateway,
		"aws_api_gateway_rest_api":                             dataAwsAPIGateway,
		"aws_api_gateway_sdk":                                  dataAwsAPIGateway,
		"aws_api_gateway_vpc_link":                             dataAwsAPIGateway,
		"aws_apigatewayv2_api":                                 dataAwsAPIGateway,
		"aws_apigatewayv2_apis":                                dataAwsAPIGateway,
		"aws_apigatewayv2_export":                              dataAwsAPIGateway,
		"aws_appconfig_configuration_profile":                  dataAwsAppconfigConfigurationProfile,
		"aws_appconfig_configuration_profiles":                 dataAwsAppconfigConfigurationProfiles,
		"aws_appconfig_environment":                            dataAwsAppconfigEnvironment,
		"aws_appconfig_environments":                           dataAwsAppconfigEnvironments,
		"aws_appintegrations_event_integration":                dataAwsAppintergrationsEventIntegration,
		"aws_appmesh_gateway_route":                            dataAwsAppmeshGatewayRoute,
		"aws_appmesh_mesh":                                     dataAwsAppmeshMesh,
		"aws_appmesh_route":                                    dataAwsAppmeshRoute,
		"aws_appmesh_virtual_gateway":                          dataAwsAppmeshVirtualGateway,
		"aws_appmesh_virtual_node":                             dataAwsAppmeshVirtualNode,
		"aws_appmesh_virtual_router":                           dataAwsAppmeshVirtualRouter,
		"aws_appmesh_virtual_service":                          dataAppmeshVirtualService,
		"aws_arn":                                              placeholder,
		"aws_auditmanager_control":                             dataAwsAuditmanagerControl,
		"aws_auditmanager_framework":                           dataAwsAuditmanagerFramework,
		"aws_autoscaling_group":                                dataAwsAutoscalingGroup,
		"aws_autoscaling_groups":                               dataAwsAutoscalingGroups,
		"aws_availability_zone":                                dataAwsAvailabilityZone,
		"aws_availability_zones":                               dataAwsAvailabilityZones,
		"aws_backup_framework":                                 dataBackupFramework,
		"aws_backup_plan":                                      dataBackupPlan,
		"aws_backup_report_plan":                               dataBackupReportPlan,
		"aws_backup_selection":                                 dataBackupSelection,
		"aws_backup_vault":                                     dataAwsBackupVault,
		"aws_batch_compute_environment":                        dataAwsBatchComputeEnvironment,
		"aws_batch_job_queue":                                  dataAwsBatchJobQueue,
		"aws_batch_scheduling_policy":                          dataAwsBatchSchedulingPolicy,
		"aws_billing_service_account":                          placeholder,
		"aws_budgets_budget":                                   dataAwsBudgetsBudget,
		"aws_caller_identity":                                  placeholder,
		"aws_canonical_user_id":                                placeholder,
		"aws_ce_cost_category":                                 dataAwsCeCostCategory,
		"aws_ce_tags":                                          dataAwsCeTags,
		"aws_cloudcontrolapi_resource":                         dataAwsCloudcontrolapiResource,
		"aws_cloudformation_export":                            dataAwsCloudformationExport,
		"aws_cloudformation_stack":                             dataAwsCloudformationStack,
		"aws_cloudformation_type":                              dataAwsCloudformationType,
		"aws_cloudfront_cache_policy":                          dataAwsCloudfrontCachePolicy,
		"aws_cloudfront_distribution":                          dataAwsCloudfrontDistribution,
		"aws_cloudfront_function":                              dataAwsCloudfrontFunction,
		"aws_cloudfront_log_delivery_canonical_user_id":        placeholder,
		"aws_cloudfront_origin_access_identities":              dataAwsCloudfrontOriginAccessIdentities,
		"aws_cloudfront_origin_access_identity":                dataAwsCloudfrontOriginAccessIdentity,
		"aws_cloudfront_origin_request_policy":                 dataAwsCloudfrontOriginRequestPolicy,
		"aws_cloudfront_realtime_log_config":                   dataAwsCloudfrontRealtimeLogConfig,
		"aws_cloudfront_response_headers_policy":               dataAwsCloudfrontResponseHeadersPolicy,
		"aws_cloudhsm_v2_cluster":                              dataAwsCloudhsmV2Cluster,
		"aws_cloudtrail_service_account":                       placeholder,
		"aws_cloudwatch_event_bus":                             dataCloudwatchEventBus,
		"aws_cloudwatch_event_connection":                      dataCloudwatchEventConnection,
		"aws_cloudwatch_event_source":                          dataCloudwatchEventSource,
		"aws_cloudwatch_log_data_protection_policy_document":   placeholder,
		"aws_cloudwatch_log_group":                             dataAwsCloudwatchLogGroup,
		"aws_cloudwatch_log_groups":                            dataAwsCloudwatchLogGroups,
		"aws_codeartifact_authorization_token":                 datAwsCodeartifactAutorization,
		"aws_codeartifact_repository_endpoint":                 dataAwsCodeartifactRepositoryEndpoint,
		"aws_codecommit_approval_rule_template":                dataAwsCodecommitApprovalRuleTemplate,
		"aws_codecommit_repository":                            dataAwsCodecommitRepository,
		"aws_codestarconnections_connection":                   dataAwsCodestarconnectionsConnection,
		"aws_cognito_user_pool_client":                         dataAwsCognitoUserPoolClient,
		"aws_cognito_user_pool_clients":                        dataAwsCognitoUserPoolClients,
		"aws_cognito_user_pool_signing_certificate":            dataAwsCognitoUserPoolSigningCertificate,
		"aws_cognito_user_pools":                               dataAwsCognitoUserPools,
		"aws_connect_bot_association":                          dataAwsConnectBotAssociation,
		"aws_connect_contact_flow":                             dataAwsConnectContactFlow,
		"aws_connect_contact_flow_module":                      dataAwsConnectContactFlowModule,
		"aws_connect_hours_of_operation":                       dataAwsConnectHoursOfOperation,
		"aws_connect_instance":                                 dataAwsConnectInstance,
		"aws_connect_instance_storage_config":                  dataAwsConnectInstanceStorageConfig,
		"aws_connect_lambda_function_association":              dataAwsConnectLambdaFunctionAssociation,
		"aws_connect_prompt":                                   dataAwsConnectPrompt,
		"aws_connect_queue":                                    dataAwsConnectQueue,
		"aws_connect_quick_connect":                            dataAwsConnectQuickConnect,
		"aws_connect_routing_profile":                          dataAwsConnectRoutingProfile,
		"aws_connect_security_profile":                         dataAwsConnectSecurityProfile,
		"aws_connect_user":                                     dataAwsConnectUser,
		"aws_connect_user_hierarchy_group":                     dataAwsConnectUserHierarchyGroup,
		"aws_connect_user_hierarchy_structure":                 dataAwsConnectUserHierarchyStructure,
		"aws_connect_vocabulary":                               dataAwsConnectVocabulary,
		"aws_controltower_controls":                            dataAwsControltowerControls,
		"aws_cur_report_definition":                            placeholder,
		"aws_customer_gateway":                                 dataAwsCustomerGateway,
		"aws_datapipeline_pipeline":                            dataAwsDatapipelinePipeline,
		"aws_datapipeline_pipeline_definition":                 dataAwsDatapipelinePipelineDefinition,
		"aws_db_cluster_snapshot":                              dataAwsDBClusterSnapshot,
		"aws_db_event_categories":                              dataAwsDBEventCategories,
		"aws_db_instance":                                      dataAwsDBInstance,
		"aws_db_instances":                                     dataAwsDbInstances,
		"aws_db_proxy":                                         placeholder,
		"aws_db_snapshot":                                      dataAwsDBSnapshot,
		"aws_db_subnet_group":                                  dataAwsDBSubnetGroup,
		"aws_default_tags":                                     placeholder,
		"aws_directory_service_directory":                      dataAwsDirectoryServiceDirectory,
		"aws_dms_endpoint":                                     dataAwsDmsEndpoint,
		"aws_dms_replication_instance":                         dataAwsDmsReplicationInstance,
		"aws_dms_replication_subnet_group":                     dataAwsDmsReplicationSubnetGroup,
		"aws_dms_replication_task":                             dataAwsDmsReplicationTask,
		"aws_docdb_engine_version":                             dataAwsDocDBEngineVersion,
		"aws_docdb_orderable_db_instance":                      dataAwsDocDBOrderableDBInstance,
		"aws_dx_connection":                                    dataAwsDxConnection,
		"aws_dx_gateway":                                       dataAwsDxGateway,
		"aws_dx_location":                                      dataAwsDxLocation,
		"aws_dx_locations":                                     dataAwsDxLocations,
		"aws_dx_router_configuration":                          dataAwsDxRouterConfiguration,
		"aws_dynamodb_table":                                   dataAwsDynamodbTable,
		"aws_dynamodb_table_item":                              dataAwsDynamodbTableItem,
		"aws_ebs_default_kms_key":                              dataAwsEbsDefaultKmsKey,
		"aws_ebs_encryption_by_default":                        dataAwsEbsEncryptionByDefault,
		"aws_ebs_snapshot":                                     dataAwsEbsSnapshot,
		"aws_ebs_snapshot_ids":                                 dataAwsEbsSnapshotIds,
		"aws_ebs_volume":                                       dataAwsEbsVolume,
		"aws_ebs_volumes":                                      dataAwsEbsVolumes,
		"aws_ec2_client_vpn_endpoint":                          dataAwsEc2ClientVpnEndpoint,
		"aws_ec2_coip_pool":                                    dataAwsEc2CoipPool,
		"aws_ec2_coip_pools":                                   dataAwsEc2CoipPools,
		"aws_ec2_host":                                         dataAwsEc2Host,
		"aws_ec2_instance_type":                                dataAwsEc2InstanceType,
		"aws_ec2_instance_type_offering":                       dataAwsEc2InstanceTypeOffering,
		"aws_ec2_instance_type_offerings":                      dataAwsEc2InstanceTypeOfferings,
		"aws_ec2_instance_types":                               dataAwsEc2InstanceTypes,
		"aws_ec2_local_gateway":                                dataAwsEc2LocalGateway,
		"aws_ec2_local_gateway_route_table":                    dataAwsEc2LocalGatewayRouteTable,
		"aws_ec2_local_gateway_route_tables":                   dataAwsEc2LocalGatewayRouteTables,
		"aws_ec2_local_gateway_virtual_interface":              dataAwsEc2LocalGatewayVirtualInterface,
		"aws_ec2_local_gateway_virtual_interface_group":        dataAwsEc2LocalGatewayVirtualInterfaceGroup,
		"aws_ec2_local_gateway_virtual_interface_groups":       dataAwsEc2LocalGatewayVirtualInterfaceGroups,
		"aws_ec2_local_gateways":                               dataAwsEc2LocalGateways,
		"aws_ec2_managed_prefix_list":                          dataAwsEc2ManagedPrefixList,
		"aws_ec2_managed_prefix_lists":                         dataAwsEc2ManagedPrefixLists,
		"aws_ec2_network_insights_analysis":                    dataAwsEc2NetworkInsightsAnalysis,
		"aws_ec2_network_insights_path":                        dataAwsEc2NetworkInsightsPath,
		"aws_ec2_public_ipv4_pool":                             dataAwsEc2PublicIpv4Pool,
		"aws_ec2_public_ipv4_pools":                            dataAwsEc2PublicIpv4Pools,
		"aws_ec2_serial_console_access":                        dataAwsEc2SerialConsoleAccess,
		"aws_ec2_spot_price":                                   dataAwsEc2SpotPrice,
		"aws_ec2_transit_gateway":                              dataAwsEc2Transitgateway,
		"aws_ec2_transit_gateway_attachment":                   dataAwsEc2TransitGatewayAttachment,
		"aws_ec2_transit_gateway_attachments":                  dataAwsEc2TransitGatewayAttachments,
		"aws_ec2_transit_gateway_connect":                      dataAwsEc2TransitGatewayConnect,
		"aws_ec2_transit_gateway_connect_peer":                 dataAwsEc2TransitGatewayConnectPeer,
		"aws_ec2_transit_gateway_dx_gateway_attachment":        dataAwsEc2TransitGatewayDxGatewayAttachment,
		"aws_ec2_transit_gateway_multicast_domain":             dataAwsEc2TransitGatewayMulticastDomain,
		"aws_ec2_transit_gateway_peering_attachment":           dataAwsEc2TransitGatewayPeeringAttachment,
		"aws_ec2_transit_gateway_route_table":                  dataAwsEc2TransitGatewayRouteTable,
		"aws_ec2_transit_gateway_route_table_associations":     dataAwsEc2TransitGatewayRouteTableAssociations,
		"aws_ec2_transit_gateway_route_table_propagations":     dataAwsEc2TransitGatewayRouteTablePropagations,
		"aws_ec2_transit_gateway_route_table_routes":           dataAwsEc2TransitGatewayRouteTableRoutes,
		"aws_ec2_transit_gateway_route_tables":                 dataAwsEc2TransitGatewayRouteTables,
		"aws_ec2_transit_gateway_vpc_attachment":               dataAwsEc2TransitGatewayVpcAttachment,
		"aws_ec2_transit_gateway_vpc_attachments":              dataAwsEc2TransitGatewayVpcAttachments,
		"aws_ec2_transit_gateway_vpn_attachment":               dataAwsEc2TransitGatewayVPNAttachment,
		"aws_ecr_authorization":                                dataAwsEcrAuthorization,
		"aws_ecr_authorization_token":                          dataAwsEcrAuthorizationToken,
		"aws_ecr_image":                                        dataAwsEcrImage,
		"aws_ecr_pull_through_cache_rule":                      dataAwsEcrPullThroughCacheRule,
		"aws_ecr_repository":                                   dataAwsEcrRepository,
		"aws_ecrpublic_authorization_token":                    dataEcrpublicAuthorizationToken,
		"aws_ecs_cluster":                                      dataAwsEcsCluster,
		"aws_ecs_container_definition":                         dataAwsEcsContainerDefinition,
		"aws_ecs_service":                                      dataDataEcsService,
		"aws_ecs_task_definition":                              dataAwsEcsTaskDefinition,
		"aws_ecs_task_execution":                               dataAwsEcsTaskExecution,
		"aws_efs_access_point":                                 dataAwsEfsAccessPoint,
		"aws_efs_access_points":                                dataAwsEfsAccessPoints,
		"aws_efs_file_system":                                  dataAwsEfsFileSystem,
		"aws_efs_mount_target":                                 dataAwsEfsMountTarget,
		"aws_eip":                                              dataAwsEip,
		"aws_eips":                                             dataAwsEips,
		"aws_eks_addon":                                        dataAwsEksAddon,
		"aws_eks_addon_version":                                dataAwsEksAddonVersion,
		"aws_eks_cluster":                                      dataAwsEksCluster,
		"aws_eks_cluster_auth":                                 placeholder,
		"aws_eks_clusters":                                     dataAwsEksClusters,
		"aws_eks_node_group":                                   dataAwsEksNodeGroup,
		"aws_eks_node_groups":                                  dataAwsEksNodeGroups,
		"aws_elastic_beanstalk_application":                    dataAwsElasticBeanstalkApplication,
		"aws_elastic_beanstalk_hosted_zone":                    placeholder,
		"aws_elastic_beanstalk_solution_stack":                 dataAwsElasticBeanstalkSolutionStack,
		"aws_elasticache_cluster":                              dataAwsElasticacheCluster,
		"aws_elasticache_replication_group":                    dataAwsElasticacheReplicationGroup,
		"aws_elasticache_subnet_group":                         dataAwsElasticacheSubnetGroup,
		"aws_elasticache_user":                                 dataAwsElasticacheUser,
		"aws_elasticsearch_domain":                             dataAwsElasticsearchDomain,
		"aws_elb_hosted_zone_id":                               placeholder,
		"aws_elb_service_account":                              placeholder,
		"aws_emr_release_labels":                               dataAwsEmrReleaseLabels,
		"aws_emrcontainers_virtual_cluster":                    placeholder,
		"aws_fsx_openzfs_snapshot":                             dataAwsFsxOpenzfsSnapshot,
		"aws_fsx_windows_file_system":                          dataAwsFsxWindowsFileSystem,
		"aws_globalaccelerator_accelerator":                    dataAwsGlobalAccelerator,
		"aws_globalaccelerator_custom_routing_accelerator":     dataAwsGlobalAcceleratorCustomRoutingAccelerator,
		"aws_glue_catalog_table":                               dataAwsGlueCatalogTable,
		"aws_glue_connection":                                  dataAwsGlueConnection,
		"aws_glue_data_catalog_encryption_settings":            dataAwsDataCatalogEncryptionSettings,
		"aws_glue_script":                                      dataAwsGlueScript,
		"aws_grafana_workspace":                                dataAwsGrafanaWorkspace,
		"aws_guardduty_detector":                               dataAwsGuarddutyDetector,
		"aws_guardduty_finding_ids":                            dataAwsGuarddutyFindingIds,
		"aws_iam_access_keys":                                  dataAwsIamAccessKeys,
		"aws_iam_account_alias":                                dataAwsIamAccountAlias,
		"aws_iam_group":                                        dataAwsIamGroup,
		"aws_iam_instance_profile":                             dataAwsIamInstanceProfile,
		"aws_iam_instance_profiles":                            dataAwsIamInstanceProfiles,
		"aws_iam_openid_connect_provider":                      dataAwsIamOpenIDConnectProvider,
		"aws_iam_policy":                                       dataAwsIamPolicy,
		"aws_iam_policy_document":                              placeholder,
		"aws_iam_principal_policy_simulation":                  dataAwsIamPrincipalPolicySimulation,
		"aws_iam_role":                                         dataAwsIamRole,
		"aws_iam_roles":                                        dataAwsIamRoles,
		"aws_iam_saml_provider":                                dataAwsIamSamlProvider,
		"aws_iam_server_certificate":                           dataAwsIamServerCertificate,
		"aws_iam_session_context":                              placeholder,
		"aws_iam_user":                                         dataAwsIamUser,
		"aws_iam_user_ssh_key":                                 dataAwsIamUserSSHKey,
		"aws_iam_users":                                        dataAwsIamUsers,
		"aws_identitystore_group":                              dataAwsIdentitystoreGroup,
		"aws_identitystore_user":                               dataAwsIdentitystoreUser,
		"aws_inspector_rules_packages":                         dataAwsInspectorRulesPackages,
		"aws_instance":                                         dataAwsInstance,
		"aws_instances":                                        dataAwsInstances,
		"aws_internet_gateway":                                 dataAwsInternetGateway,
		"aws_iot_endpoint":                                     dataAwsIotEndpoint,
		"aws_ip_ranges":                                        placeholder,
		"aws_ivs_stream_key":                                   dataAwsIvsStreamKey,
		"aws_key_pair":                                         dataAwsKeyPair,
		"aws_kinesis_firehose_delivery_stream":                 dataAwsKinesisFirehoseDeliveryStream,
		"aws_kinesis_stream":                                   dataAwsKinesisStream,
		"aws_kinesis_stream_consumer":                          dataAwsKinesisStreamConsumer,
		"aws_kms_alias":                                        dataAwsKmsAlias,
		"aws_kms_ciphertext":                                   dataAwsKmsCiphertext,
		"aws_kms_custom_key_store":                             dataAwsKmsCustomKeyStore,
		"aws_kms_key":                                          dataAwsKmsKey,
		"aws_kms_secrets":                                      dataAwsKmsSecrets,
		"aws_lambda_code_signing_config":                       dataAwsLambdaCodeSigningConfig,
		"aws_lambda_function":                                  dataAwsLambdaFunction,
		"aws_lambda_invocation":                                dataAwsLambdaInvocation,
		"aws_launch_configuration":                             dataAwsLaunchConfiguration,
		"aws_launch_template":                                  dataAwsLaunchTemplate,
		"aws_lb":                                               dataAwsLb,
		"aws_lb_hosted_zone_id":                                placeholder,
		"aws_lb_listener":                                      dataAwsLbListener,
		"aws_lb_target_group":                                  dataAwsLbTargetGroup,
		"aws_lbs":                                              dataAwsLbs,
		"aws_location_tracker_association":                     dataAwsLocationTrackerAssociation,
		"aws_location_tracker_associations":                    dataAwsLocationTrackerAssociations,
		"aws_mq_broker":                                        dataAwsMqBroker,
		"aws_msk_broker_nodes":                                 dataAwsBrokerNodes,
		"aws_msk_cluster":                                      dataAwsMskCluster,
		"aws_msk_configuration":                                dataAwsMskConfiguration,
		"aws_msk_kafka_version":                                dataAwsMskKafkaVersion,
		"aws_nat_gateway":                                      dataAwsNatGateway,
		"aws_nat_gateways":                                     dataAwsNatGateways,
		"aws_neptune_engine_version":                           dataAwsNeptuneEngineVersion,
		"aws_neptune_orderable_db_instance":                    dataAwsNeptuneOrderableDBInstance,
		"aws_network_acls":                                     dataAwsNetworkAcls,
		"aws_network_interfaces":                               dataAwsNetworkInterfaces,
		"aws_opensearch_domain":                                placeholder,
		"aws_organizations_organization":                       dataAwsOrganizationsOrganization,
		"aws_outposts_outpost":                                 dataAwsOutpostsOutpost,
		"aws_partition":                                        placeholder,
		"aws_prefix_list":                                      dataAwsPrefixList,
		"aws_pricing_product":                                  placeholder,
		"aws_prometheus_workspace":                             dataAwsPrometheusWorkspace,
		"aws_prometheus_workspaces":                            dataAwsPrometheusWorkspaces,
		"aws_quicksight_data_set":                              placeholder,
		"aws_quicksight_group":                                 dataAwsQuicksightGroup,
		"aws_quicksight_theme":                                 dataAwsQuicksightTheme,
		"aws_quicksight_user":                                  dataAwsQuicksightUser,
		"aws_rds_certificate":                                  dataAwsRdsCertificate,
		"aws_rds_cluster":                                      dataAwsRdsCluster,
		"aws_rds_clusters":                                     dataAwsRdsClusters,
		"aws_rds_engine_version":                               dataAwsRdsEngineVersion,
		"aws_rds_orderable_db_instance":                        dataAwsRdsOrderableDBInstance,
		"aws_rds_reserved_instance_offering":                   dataAwsRdsReservedInstanceOffering,
		"aws_redshift_cluster":                                 dataAwsRedshiftCluster,
		"aws_redshift_cluster_credentials":                     dataAwsRedshiftClusterCredentials,
		"aws_redshift_orderable_cluster":                       dataAwsRedshiftOrderableCluster,
		"aws_redshift_service_account":                         placeholder,
		"aws_redshift_subnet_group":                            dataAwsRedshiftSubnetGroup,
		"aws_redshiftserverless_credentials":                   placeholder,
		"aws_region":                                           placeholder,
		"aws_regions":                                          dataAwsRegions,
		"aws_route":                                            dataAwsRoute,
		"aws_route53_resolver_firewall_config":                 dataAwsRoute53ResolverFirewallConfig,
		"aws_route53_resolver_firewall_domain_list":            dataAwsRoute53ResolverFirewallDomainList,
		"aws_route53_resolver_firewall_rule_group":             dataAwsRoute53ResolverFirewallRuleGroup,
		"aws_route53_resolver_firewall_rule_group_association": dataAwsRoute53ResolverFirewallGroupAssociation,
		"aws_route53_resolver_firewall_rules":                  dataAwsRoute53ResolverFirewallRules,
		"aws_route53_resolver_rule":                            dataAwsRoute53ResolverRule,
		"aws_route53_traffic_policy_document":                  placeholder,
		"aws_route53_zone":                                     dataAwsRoute53Zone,
		"aws_route_table":                                      dataAwsRouteTable,
		"aws_route_tables":                                     dataAwsRouteTables,
		"aws_s3_account_public_access_block":                   dataAwsS3AccountPublicAccessBlock,
		"aws_s3_bucket":                                        dataAwsS3Bucket,
		"aws_s3_bucket_object":                                 placeholder,
		"aws_s3_bucket_policy":                                 dataAwsS3BucketPolicy,
		"aws_s3_object":                                        placeholder,
		"aws_s3_objects":                                       dataAwsS3Objects,
		"aws_s3control_multi_region_access_point":              dataAwsS3controlMultiRegionAccessPoint,
		"aws_sagemaker_prebuilt_ecr_image":                     placeholder,
		"aws_secretsmanager_secret":                            dataAwsSecretsmanagerSecret,
		"aws_secretsmanager_secret_version":                    dataAwsSecretsmanagerSecretVersion,
		"aws_security_group":                                   dataAwsSecurityGroup,
		"aws_security_groups":                                  dataAwsSecurityGroup,
		"aws_service":                                          placeholder,
		"aws_service_discovery_dns_namespace":                  dataAwsServiceDiscoveryDNSNamespace,
		"aws_service_discovery_http_namespace":                 dataAwsServiceDiscoveryHTTPNamespace,
		"aws_service_discovery_service":                        dataAwsServiceDiscoveryService,
		"aws_servicequotas_service":                            dataAwsServicequotasService,
		"aws_servicequotas_service_quota":                      dataAwsServicequotaServiceQuota,
		"aws_sesv2_dedicated_ip_pool":                          dataAwsSesv2DedicatedIPPool,
		"aws_sfn_activity":                                     dataAwsSfnActivity,
		"aws_sfn_alias":                                        dataAwsSfnAlias,
		"aws_sfn_state_machine":                                dataAwsSfnStateMachine,
		"aws_sfn_state_machine_versions":                       dataAwsSfnStateMachineVersion,
		"aws_sns_topic":                                        dataAwsSnsTopic,
		"aws_sqs_queue":                                        dataAwsSqsQueue,
		"aws_sqs_queues":                                       dataAwsSqsQueues,
		"aws_ssm_document":                                     dataAwsSsmDocument,
		"aws_ssm_instances":                                    dataAwsSsmInstances,
		"aws_ssm_maintenance_windows":                          dataAwsSsmMaintenanceWindows,
		"aws_ssm_parameter":                                    dataAwsSsmParameter,
		"aws_ssm_parameters_by_path":                           dataAwsSsmParametersByPath,
		"aws_ssm_patch_baseline":                               dataAwsSsmPatchBaseline,
		"aws_ssoadmin_instances":                               dataAwsSsoadminInstances,
		"aws_subnet":                                           dataAwsSubnetIds,
		"aws_subnet_ids":                                       dataAwsSubnetIds,
		"aws_subnets":                                          dataAwsSubnetIds,
		"aws_transfer_server":                                  dataAwsTransferServer,
		"aws_vpc":                                              dataAwsVpc,
		"aws_vpc_dhcp_options":                                 dataAwsVpcDhcpOptions,
		"aws_vpc_endpoint":                                     dataAwsVpcEndpoint,
		"aws_vpc_endpoint_service":                             dataAwsVpcEndpointService,
		"aws_vpc_ipam_pool":                                    dataAwsVpcIpamPoolCidrs,
		"aws_vpc_ipam_pool_cidrs":                              dataAwsVpcIpamPoolCidrs,
		"aws_vpc_ipam_pools":                                   dataAwsVpcIpamPools,
		"aws_vpcs":                                             dataAwsVpcs,
		"aws_vpn_gateway":                                      dataAwsVpnGateway,
		"aws_waf_ipset":                                        dataAwsWafIpset,
		"aws_waf_rate_based_rule":                              dataAwsWafRateBasedRule,
		"aws_waf_rule":                                         dataAwsWafRule,
		"aws_waf_web_acl":                                      dataAwsWafWebAcl,
		"aws_wafregional_ipset":                                dataAwsWafregionalIpset,
		"aws_wafregional_rate_based_rule":                      dataAwsWafregionalRateBasedRule,
		"aws_wafregional_rule":                                 dataAwsWafregionalRule,
		"aws_wafregional_web_acl":                              dataAwsWafregionalWebAcl,
		"aws_wafv2_ip_set":                                     dataAwsWafv2IpSet,
		"aws_wafv2_regex_pattern_set":                          dataAwsWafv2RegexPatternSet,
		"aws_wafv2_rule_group":                                 dataAwsWafv2RuleGroup,
		"aws_wafv2_web_acl":                                    dataAwsWafv2WebACL,
		"aws_workspaces_bundle":                                dataAwsWorkspacesBundle,
		"aws_workspaces_directory":                             dataAwsWorkspaceDirectory,
		"aws_workspaces_image":                                 dataAwsWorkspaceImage,
		"aws_lakeformation_data_lake_settings":                 dataAwsLakeformationDataLakeSettings,
		"aws_lakeformation_permissions":                        dataAwsLakeformationPermissions,
		"aws_lakeformation_resource":                           dataAwsLakeformationResource,
		"aws_mq_broker_instance_type_offerings":                dataAwsMqBrokerInstanceTypeOfferings,
		"aws_secretsmanager_random_password":                   dataAwsSecretsmanagerRandomPassword,
		"aws_secretsmanager_secret_rotation":                   dataAwsSecretsmanagerSecretRotation,
		"aws_secretsmanager_secrets":                           dataAwsSecretsmanagerSecrets,
		"aws_ssoadmin_permission_set":                          dataAwsSsoadminPermissionSet,
		"aws_storagegateway_local_disk":                        placeholder,
		"aws_vpc_ipam_preview_next_cidr":                       placeholder,
		"aws_vpc_peering_connection":                           dataAwsVpcPeeringConnection,
		"aws_vpc_peering_connections":                          dataAwsVpcPeeringConnections,
		"aws_vpc_security_group_rule":                          dataAwsVpcSecurityGroupRule,
		"aws_vpc_security_group_rules":                         dataAwsVpcSecurityGroupRules,
	}

	var (
		Permissions []string
		err         error
	)

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions, err = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		//goland:noinspection GoLinter
		return nil, fmt.Errorf("data.%s not implemented", result.Name)
	}

	return Permissions, err
}
