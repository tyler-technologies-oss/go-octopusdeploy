package octopusdeploy

const (
	TestURISelf                              string = "/api"
	TestURIAccounts                          string = "/api/Spaces-1/accounts{/id}{?skip,take,ids,partialName,accountType}"
	TestURIActionTemplateLogo                string = "/api/Spaces-1/actiontemplates/{typeOrId}/logo{?cb}"
	TestURIActionTemplates                   string = "/api/Spaces-1/actiontemplates{/id}{?skip,take,ids,partialName}"
	TestURIActionTemplatesCategories         string = "/api/Spaces-1/actiontemplates/categories"
	TestURIActionTemplatesLogo               string = "/api/Spaces-1/actiontemplates/{typeOrId}/logo{?cb}"
	TestURIActionTemplatesSearch             string = "/api/Spaces-1/actiontemplates/search"
	TestURIActionTemplateVersionedLogo       string = "/api/Spaces-1/actiontemplates/{typeOrId}/versions/{version}/logo{?cb}"
	TestURIAPIKeys                           string = "/api/users/{userId}/apikeys"
	TestURIArtifacts                         string = "/api/Spaces-1/artifacts{/id}{?skip,take,regarding,ids,partialName,order}"
	TestURIAuthenticateOctopusID             string = "/users/authenticate/OctopusID{?returnUrl}"
	TestURIAuthentication                    string = "/api/authentication"
	TestURIAzureDevOpsConnectivityCheck      string = "/api/azuredevopsissuetracker/connectivitycheck"
	TestURIAzureEnvironments                 string = "/api/accounts/azureenvironments"
	TestURIBuildInformation                  string = "/api/Spaces-1/build-information{/id}{?filter,packageId,latest,skip,take,overwriteMode}"
	TestURIBuildInformationBulk              string = "/api/Spaces-1/build-information/bulk{?ids}"
	TestURIBuiltInFeedStats                  string = "/api/feeds/stats"
	TestURICertificateConfiguration          string = "/api/configuration/certificates{/id}{?skip,take,ids,partialName}"
	TestURICertificates                      string = "/api/Spaces-1/certificates{/id}{?skip,take,search,archived,tenant,firstResult,orderBy,ids,partialName}"
	TestURIChannels                          string = "/api/Spaces-1/channels{/id}{?skip,take,ids,partialName}"
	TestURICloudTemplate                     string = "/api/cloudtemplate/{id}/metadata{?packageId,feedId}"
	TestURICommunityActionTemplates          string = "/api/communityactiontemplates{/id}{?skip,take,ids}"
	TestURIConfiguration                     string = "/api/configuration{/id}"
	TestURICurrentLicense                    string = "/api/licenses/licenses-current"
	TestURICurrentLicenseStatus              string = "/api/licenses/licenses-current-status"
	TestURICurrentUser                       string = "/api/users/me"
	TestURIDashboard                         string = "/api/Spaces-1/dashboard{?projectId,releaseId,selectedTenants,selectedTags,showAll,highestLatestVersionPerProjectAndEnvironment}"
	TestURIDashboardConfiguration            string = "/api/Spaces-1/dashboardconfiguration"
	TestURIDashboardDynamic                  string = "/api/Spaces-1/dashboard/dynamic{?projects,environments,includePrevious}"
	TestURIDeploymentProcesses               string = "/api/Spaces-1/deploymentprocesses{/id}{?skip,take,ids}"
	TestURIDeployments                       string = "/api/Spaces-1/deployments{/id}{?skip,take,ids,projects,environments,tenants,channels,taskState,partialName}"
	TestURIDiscoverMachine                   string = "/api/Spaces-1/machines/discover{?host,port,type,proxyId}"
	TestURIDiscoverWorker                    string = "/api/Spaces-1/workers/discover{?host,port,type,proxyId}"
	TestURIDynamicExtensionsFeaturesMetadata string = "/api/dynamic-extensions/features/metadata"
	TestURIDynamicExtensionsFeaturesValues   string = "/api/dynamic-extensions/features/values"
	TestURIDynamicExtensionsScripts          string = "/api/dynamic-extensions/scripts"
	TestURIEnvironments                      string = "/api/Spaces-1/environments{/id}{?name,skip,ids,take,partialName}"
	TestURIEnvironmentSortOrder              string = "/api/Spaces-1/environments/sortorder"
	TestURIEnvironmentsSummary               string = "/api/Spaces-1/environments/summary{?ids,partialName,machinePartialName,roles,isDisabled,healthStatuses,commStyles,tenantIds,tenantTags,hideEmptyEnvironments,shellNames}"
	TestURIEventAgents                       string = "/api/events/agents"
	TestURIEventCategories                   string = "/api/events/categories{?appliesTo}"
	TestURIEventDocumentTypes                string = "/api/events/documenttypes"
	TestURIEventGroups                       string = "/api/events/groups{?appliesTo}"
	TestURIEvents                            string = "/api/events{/id}{?skip,regarding,regardingAny,user,users,projects,projectGroups,environments,eventGroups,eventCategories,eventAgents,tags,tenants,from,to,internal,fromAutoId,toAutoId,documentTypes,asCsv,take,ids,spaces,includeSystem,excludeDifference}"
	TestURIExtensionStats                    string = "/api/serverstatus/extensions"
	TestURIExternalSecurityGroupProviders    string = "/api/externalsecuritygroupproviders"
	TestURIExternalUserSearch                string = "/api/users/external-search{?partialName}"
	TestURIFeaturesConfiguration             string = "/api/featuresconfiguration"
	TestURIFeeds                             string = "/api/feeds{/id}{?skip,take,ids,partialName,feedType}"
	TestURIInterruptions                     string = "/api/Spaces-1/interruptions{/id}{?skip,take,regarding,pendingOnly,ids}"
	TestURIInvitations                       string = "/api/users/invitations"
	TestURIIssueTrackers                     string = "/api/issuetrackers{?skip,take,ids,partialName}"
	TestURIJiraConnectAppCredentialsTest     string = "/api/jiraintegration/connectivitycheck/connectapp"
	TestURIJiraCredentialsTest               string = "/api/jiraintegration/connectivitycheck/jira"
	TestURILetsEncryptConfiguration          string = "/api/letsencryptconfiguration"
	TestURILibraryVariables                  string = "/api/Spaces-1/libraryvariablesets{/id}{?skip,contentType,take,ids,partialName}"
	TestURILifecycles                        string = "/api/Spaces-1/lifecycles{/id}{?skip,take,ids,partialName}"
	TestURILoginInitiated                    string = "/api/authentication/checklogininitiated"
	TestURIMachineOperatingSystems           string = "/api/Spaces-1/machines/operatingsystem/names/all"
	TestURIMachinePolicies                   string = "/api/Spaces-1/machinepolicies{/id}{?skip,take,ids,partialName}"
	TestURIMachinePolicyTemplate             string = "/api/Spaces-1/machinepolicies/template"
	TestURIMachineRoles                      string = "/api/Spaces-1/machineroles/all"
	TestURIMachines                          string = "/api/Spaces-1/machines{/id}{?skip,take,name,ids,partialName,roles,isDisabled,healthStatuses,commStyles,tenantIds,tenantTags,environmentIds,thumbprint,deploymentId,shellNames}"
	TestURIMachineShells                     string = "/api/Spaces-1/machines/operatingsystem/shells/all"
	TestURIMaintenanceConfiguration          string = "/api/maintenanceconfiguration"
	TestURIMigrationsImport                  string = "/api/migrations/import"
	TestURIMigrationsPartialExport           string = "/api/migrations/partialexport"
	TestURIOctopusServerClusterSummary       string = "/api/octopusservernodes/summary"
	TestURIOctopusServerNodes                string = "/api/octopusservernodes{/id}{?skip,take,ids,partialName}"
	TestURIPackageDeltaSignature             string = "/api/Spaces-1/packages/{packageId}/{version}/delta-signature"
	TestURIPackageDeltaUpload                string = "/api/Spaces-1/packages/{packageId}/{baseVersion}/delta{?replace,overwriteMode}"
	TestURIPackageMetadata                   string = "/api/Spaces-1/package-metadata{/id}{?filter,latest,skip,take,replace,overwriteMode}"
	TestURIPackageNotesList                  string = "/api/Spaces-1/packages/notes{?packageIds}"
	TestURIPackages                          string = "/api/Spaces-1/packages{/id}{?nuGetPackageId,filter,latest,skip,take,includeNotes}"
	TestURIPackagesBulk                      string = "/api/Spaces-1/packages/bulk{?ids}"
	TestURIPackageUpload                     string = "/api/Spaces-1/packages/raw{?replace,overwriteMode}"
	TestURIPerformanceConfiguration          string = "/api/performanceconfiguration"
	TestURIPermissionDescriptions            string = "/api/permissions/all"
	TestURIProjectGroups                     string = "/api/Spaces-1/projectgroups{/id}{?skip,take,ids,partialName}"
	TestURIProjectPulse                      string = "/api/Spaces-1/projects/pulse{?projectIds}"
	TestURIProjects                          string = "/api/Spaces-1/projects{/id}{?name,skip,ids,clone,take,partialName,clonedFromProjectId}"
	TestURIProjectsExperimentalSummaries     string = "/api/Spaces-1/projects/experimental/summaries{?ids}"
	TestURIProjectTriggers                   string = "/api/Spaces-1/projecttriggers{/id}{?skip,take,ids,runbooks}"
	TestURIProxies                           string = "/api/Spaces-1/proxies{/id}{?skip,take,ids,partialName}"
	TestURIRegister                          string = "/api/users/register"
	TestURIReleases                          string = "/api/Spaces-1/releases{/id}{?skip,ignoreChannelRules,take,ids}"
	TestURIReportingDeploymentsCountedByWeek string = "/api/Spaces-1/reporting/deployments-counted-by-week{?projectIds}"
	TestURIRoot                              string = "/api"
	TestURIRunbookProcesses                  string = "/api/Spaces-1/runbookProcesses{/id}{?skip,take,ids}"
	TestURIRunbookRuns                       string = "/api/Spaces-1/runbookRuns{/id}{?skip,take,ids,projects,environments,tenants,runbooks,taskState,partialName}"
	TestURIRunbooks                          string = "/api/Spaces-1/runbooks{/id}{?skip,take,ids,partialName,clone,projectIds}"
	TestURIRunbookSnapshots                  string = "/api/Spaces-1/runbookSnapshots{/id}{?skip,take,ids,publish}"
	TestURIScheduledProjectTriggers          string = "/api/Spaces-1/scheduledprojecttriggers{/id}{?skip,take,ids}"
	TestURIScheduler                         string = "/api/scheduler/{name}/logs{?verbose,tail}"
	TestURIScopedUserRoles                   string = "/api/scopeduserroles{/id}{?skip,take,ids,partialName,spaces,includeSystem}"
	TestURIServerConfiguration               string = "/api/serverconfiguration"
	TestURIServerConfigurationSettings       string = "/api/serverconfiguration/settings"
	TestURIServerHealthStatus                string = "/api/serverstatus/health"
	TestURIServerStatus                      string = "/api/serverstatus"
	TestURISignIn                            string = "/api/users/login{?returnUrl}"
	TestURISignOut                           string = "/api/users/logout"
	TestURISmtpConfiguration                 string = "/api/smtpconfiguration"
	TestURISmtpIsConfigured                  string = "/api/smtpconfiguration/isconfigured"
	TestURISpaceHome                         string = "/api/{spaceId}"
	TestURISpaces                            string = "/api/spaces{/id}{?name,skip,ids,take,partialName}"
	TestURISubscriptions                     string = "/api/Spaces-1/subscriptions{/id}{?skip,take,ids,partialName,spaces}"
	TestURITagSets                           string = "/api/Spaces-1/tagsets{/id}{?skip,take,ids,partialName}"
	TestURITagSetSortOrder                   string = "/api/Spaces-1/tagsets/sortorder"
	TestURITasks                             string = "/api/tasks{/id}{?skip,active,environment,tenant,runbook,project,name,node,running,states,hasPendingInterruptions,hasWarningsOrErrors,take,ids,partialName,spaces,includeSystem}"
	TestURITaskTypes                         string = "/api/tasks/taskTypes"
	TestURITeamMembership                    string = "/api/teammembership{?userId,spaces,includeSystem}"
	TestURITeamMembershipPreviewTeam         string = "/api/teammembership/previewteam"
	TestURITeams                             string = "/api/teams{/id}{?skip,take,ids,partialName,spaces,includeSystem}"
	TestURITenants                           string = "/api/Spaces-1/tenants{/id}{?skip,projectId,name,tags,take,ids,clone,partialName,clonedFromTenantId}"
	TestURITenantsMissingVariables           string = "/api/Spaces-1/tenants/variables-missing{?tenantId,projectId,environmentId,includeDetails}"
	TestURITenantsStatus                     string = "/api/Spaces-1/tenants/status"
	TestURITenantTagTest                     string = "/api/Spaces-1/tenants/tag-test{?tenantIds,tags}"
	TestURITenantVariables                   string = "/api/Spaces-1/tenantvariables/all{?projectId}"
	TestURITimezones                         string = "/api/serverstatus/timezones"
	TestURIUpgradeConfiguration              string = "/api/upgradeconfiguration"
	TestURIUserAuthentication                string = "/api/users/authentication{/userId}"
	TestURIUserIdentityMetadata              string = "/api/users/identity-metadata"
	TestURIUserOnboarding                    string = "/api/Spaces-1/useronboarding"
	TestURIUserRoles                         string = "/api/userroles{/id}{?skip,take,ids,partialName}"
	TestURIUsers                             string = "/api/users{/id}{?skip,take,ids,filter}"
	TestURIVariableNames                     string = "/api/Spaces-1/variables/names{?project,runbook,projectEnvironmentsFilter}"
	TestURIVariablePreview                   string = "/api/Spaces-1/variables/preview{?project,runbook,environment,channel,tenant,action,machine,role}"
	TestURIVariables                         string = "/api/Spaces-1/variables{/id}{?ids}"
	TestURIVersionControlClearCache          string = "/api/configuration/versioncontrol/clear-cache"
	TestURIVersionRuleTest                   string = "/api/Spaces-1/channels/rule-test{?version,versionRange,preReleaseTag,feetType}"
	TestURIWeb                               string = "/app"
	TestURIWorkerOperatingSystems            string = "/api/Spaces-1/workers/operatingsystem/names/all"
	TestURIWorkerPools                       string = "/api/Spaces-1/workerpools{/id}{?name,skip,ids,take,partialName}"
	TestURIWorkerPoolsDynamicWorkerTypes     string = "/api/Spaces-1/workerpools/dynamicworkertypes"
	TestURIWorkerPoolsSortOrder              string = "/api/Spaces-1/workerpools/sortorder"
	TestURIWorkerPoolsSummary                string = "/api/Spaces-1/workerpools/summary{?ids,partialName,machinePartialName,isDisabled,healthStatuses,commStyles,hideEmptyWorkerPools,shellNames}"
	TestURIWorkerPoolsSupportedTypes         string = "/api/Spaces-1/workerpools/supportedtypes"
	TestURIWorkers                           string = "/api/Spaces-1/workers{/id}{?skip,take,name,ids,partialName,isDisabled,healthStatuses,commStyles,workerPoolIds,thumbprint,shellNames}"
	TestURIWorkerShells                      string = "/api/Spaces-1/workers/operatingsystem/shells/all"
	TestURIWorkerToolsLatestImages           string = "/api/workertoolslatestimages"
)
