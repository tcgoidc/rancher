package crds

import (
	"github.com/rancher/rancher/pkg/features"
)

// RequiredCRDs returns a list of CRD to install based on enabled features.
func RequiredCRDs() []string {
	requiredCRDS := BasicCRDs()
	if features.ProvisioningV2.Enabled() {
		requiredCRDS = append(requiredCRDS, ProvisioningV2CRDs()...)
		if features.RKE2.Enabled() {
			requiredCRDS = append(requiredCRDS, RKE2CRDs()...)
		}
		if features.EmbeddedClusterAPI.Enabled() {
			requiredCRDS = append(requiredCRDS, CAPICRDs()...)
		}
		if features.Fleet.Enabled() {
			requiredCRDS = append(requiredCRDS, "managedcharts.management.cattle.io")
		}
	}
	if features.Fleet.Enabled() {
		requiredCRDS = append(requiredCRDS, FleetCRDs()...)
	}
	if features.MCM.Enabled() {
		requiredCRDS = append(requiredCRDS, MCMCRDs()...)
	}
	if features.Auth.Enabled() {
		requiredCRDS = append(requiredCRDS, AuthCRDs()...)
	}
	if features.UIExtension.Enabled() {
		requiredCRDS = append(requiredCRDS, UIPluginsCRD()...)
	}
	if features.OIDCProvider.Enabled() {
		requiredCRDS = append(requiredCRDS, OIDCClientCRD()...)
	}
	if features.RancherSCCRegistrationExtension.Enabled() {
		requiredCRDS = append(requiredCRDS, SCCRegistrationCRDs()...)
	}

	// get unique CRDs so they aren't registered twice
	uniqueCRDs := make([]string, 0, len(requiredCRDS))
	keys := map[string]struct{}{}
	for _, crdName := range requiredCRDS {
		if _, ok := keys[crdName]; ok {
			continue
		}

		keys[crdName] = struct{}{}
		uniqueCRDs = append(uniqueCRDs, crdName)
	}

	return uniqueCRDs
}

// BasicCRDs returns a list of CRD names needed to run rancher.
func BasicCRDs() []string {
	return []string{
		"apiservices.management.cattle.io",
		"apps.catalog.cattle.io",
		"auditpolicies.auditlog.cattle.io",
		"clusterregistrationtokens.management.cattle.io",
		"clusterrepos.catalog.cattle.io",
		"clusters.management.cattle.io",
		"features.management.cattle.io",
		"navlinks.ui.cattle.io",
		"operations.catalog.cattle.io",
		"podsecurityadmissionconfigurationtemplates.management.cattle.io",
		"preferences.management.cattle.io",
		"settings.management.cattle.io",
	}
}

// ProvisioningV2CRDs returns a list of CRD names needed for ProvisioningV2.
func ProvisioningV2CRDs() []string {
	return []string{
		"clusters.provisioning.cattle.io",
	}
}

// CAPICRDs returns a list of CRD names needed for CAPI.
func CAPICRDs() []string {
	return []string{
		"machines.cluster.x-k8s.io",
		"machinehealthchecks.cluster.x-k8s.io",
		"machinedeployments.cluster.x-k8s.io",
		"machinesets.cluster.x-k8s.io",
		"clusters.cluster.x-k8s.io",
		"machinedrainrules.cluster.x-k8s.io",
	}
}

// RKE2CRDs returns a list of CRD names needed for RKE2.
func RKE2CRDs() []string {
	return []string{
		"clusters.provisioning.cattle.io",
		"custommachines.rke.cattle.io",
		"etcdsnapshots.rke.cattle.io",
		"rkebootstraps.rke.cattle.io",
		"rkebootstraptemplates.rke.cattle.io",
		"rkeclusters.rke.cattle.io",
		"rkecontrolplanes.rke.cattle.io",
	}
}

// FleetCRDs returns a list of CRD names needed for Fleet.
func FleetCRDs() []string {
	return append(
		bootstrapFleet(),
		"fleetworkspaces.management.cattle.io",
	)
}

func bootstrapFleet() []string {
	return []string{
		"bundles.fleet.cattle.io",
		"clusters.fleet.cattle.io",
		"clustergroups.fleet.cattle.io",
	}
}

// AuthCRDs returns a list of CRD names needed for authentication.
func AuthCRDs() []string {
	return []string{
		"authconfigs.management.cattle.io",
		"groups.management.cattle.io",
		"groupmembers.management.cattle.io",
		"tokens.management.cattle.io",
		"users.management.cattle.io",
		"userattributes.management.cattle.io",
		"clusterproxyconfigs.management.cattle.io",
	}
}

// ClusterAuthCRDs returns a list of CRD names needed for ACE.
func ClusterAuthCRDs() []string {
	return []string{
		"clusterauthtokens.cluster.cattle.io",
		"clusteruserattributes.cluster.cattle.io",
	}
}

// MCMCRDs returns a list of CRD names needed for Multi Cluster Management.
func MCMCRDs() []string {
	return []string{
		"authconfigs.management.cattle.io",
		"clusters.management.cattle.io",
		"clusterregistrationtokens.management.cattle.io",
		"clusterroletemplatebindings.management.cattle.io",
		"composeconfigs.management.cattle.io",
		"dynamicschemas.management.cattle.io",
		"features.management.cattle.io",
		"fleetworkspaces.management.cattle.io",
		"globalroles.management.cattle.io",
		"globalrolebindings.management.cattle.io",
		"groups.management.cattle.io",
		"groupmembers.management.cattle.io",
		"kontainerdrivers.management.cattle.io",
		"monitormetrics.management.cattle.io",
		"nodes.management.cattle.io",
		"nodedrivers.management.cattle.io",
		"nodepools.management.cattle.io",
		"podsecurityadmissionconfigurationtemplates.management.cattle.io",
		"preferences.management.cattle.io",
		"projects.management.cattle.io",
		"projectnetworkpolicys.management.cattle.io",
		"projectroletemplatebindings.management.cattle.io",
		"rancherusernotificationtypes.management.cattle.io",
		"roletemplates.management.cattle.io",
		"samltokens.management.cattle.io",
		"settings.management.cattle.io",
		"templates.management.cattle.io",
		"templatecontents.management.cattle.io",
		"templateversions.management.cattle.io",
		"tokens.management.cattle.io",
		"users.management.cattle.io",
		"userattributes.management.cattle.io",
	}
}

// UIPluginsCRD returns a list of CRD names needed to enable UIPlugins
func UIPluginsCRD() []string {
	return []string{
		"uiplugins.catalog.cattle.io",
	}
}

// OIDCClientCRD returns a list of CRD names needed to enable the OIDC provider
func OIDCClientCRD() []string {
	return []string{
		"oidcclients.management.cattle.io",
	}
}

// SCCRegistrationCRDs returns a list of required CRD names needed for SCC Registration features
func SCCRegistrationCRDs() []string {
	return []string{
		"registrations.scc.cattle.io",
	}
}

// MigratedResources map list of resource that have been migrated after all resource have a CRD this can be removed.
var MigratedResources = map[string]bool{
	"activedirectoryproviders.management.cattle.io":                   false,
	"apiservices.management.cattle.io":                                false,
	"apps.catalog.cattle.io":                                          false,
	"auditpolicies.auditlog.cattle.io":                                true,
	"authconfigs.management.cattle.io":                                false,
	"authproviders.management.cattle.io":                              false,
	"authtokens.management.cattle.io":                                 false,
	"azureadproviders.management.cattle.io":                           false,
	"basicauths.project.cattle.io":                                    false,
	"certificates.project.cattle.io":                                  false,
	"cloudcredentials.management.cattle.io":                           false,
	"clusterauthtokens.cluster.cattle.io":                             false,
	"clusterclasses.cluster.x-k8s.io":                                 false,
	"clusterproxyconfigs.management.cattle.io":                        true,
	"clusterregistrationtokens.management.cattle.io":                  false,
	"clusterrepos.catalog.cattle.io":                                  true,
	"clusterresourcesetbindings.addons.cluster.x-k8s.io":              false,
	"clusterroletemplatebindings.management.cattle.io":                true,
	"clusters.cluster.x-k8s.io":                                       false,
	"clusters.management.cattle.io":                                   false,
	"clusters.provisioning.cattle.io":                                 true,
	"clusteruserattributes.cluster.cattle.io":                         false,
	"composeconfigs.management.cattle.io":                             false,
	"custommachines.rke.cattle.io":                                    true,
	"dockercredentials.project.cattle.io":                             false,
	"dynamicschemas.management.cattle.io":                             true,
	"etcdsnapshots.rke.cattle.io":                                     true,
	"extensionconfigs.runtime.cluster.x-k8s.io":                       false,
	"features.management.cattle.io":                                   false,
	"fleetworkspaces.management.cattle.io":                            false,
	"freeipaproviders.management.cattle.io":                           false,
	"githubproviders.management.cattle.io":                            false,
	"globalrolebindings.management.cattle.io":                         true,
	"globalroles.management.cattle.io":                                true,
	"googleoauthproviders.management.cattle.io":                       false,
	"groupmembers.management.cattle.io":                               false,
	"groups.management.cattle.io":                                     false,
	"ipaddressclaims.ipam.cluster.x-k8s.io":                           false,
	"kontainerdrivers.management.cattle.io":                           false,
	"localproviders.management.cattle.io":                             false,
	"machinedeployments.cluster.x-k8s.io":                             false,
	"machinepools.cluster.x-k8s.io":                                   false,
	"machines.cluster.x-k8s.io":                                       false,
	"machinesets.cluster.x-k8s.io":                                    false,
	"managedcharts.management.cattle.io":                              false,
	"monitormetrics.management.cattle.io":                             false,
	"navlinks.ui.cattle.io":                                           false,
	"nodedrivers.management.cattle.io":                                true,
	"nodepools.management.cattle.io":                                  false,
	"nodes.management.cattle.io":                                      false,
	"oidcclients.management.cattle.io":                                true,
	"oidcproviders.management.cattle.io":                              false,
	"openldapproviders.management.cattle.io":                          false,
	"operations.catalog.cattle.io":                                    false,
	"podsecurityadmissionconfigurationtemplates.management.cattle.io": false,
	"preferences.management.cattle.io":                                false,
	"principals.management.cattle.io":                                 false,
	"projectnetworkpolicies.management.cattle.io":                     false,
	"projectroletemplatebindings.management.cattle.io":                true,
	"projects.management.cattle.io":                                   true,
	"rancherusernotifications.management.cattle.io":                   false,
	"registrations.scc.cattle.io":                                     true,
	"rkebootstraps.rke.cattle.io":                                     true,
	"rkebootstraptemplates.rke.cattle.io":                             true,
	"rkeclusters.rke.cattle.io":                                       true,
	"rkecontrolplanes.rke.cattle.io":                                  true,
	"roletemplates.management.cattle.io":                              true,
	"samlproviders.management.cattle.io":                              false,
	"samltokens.management.cattle.io":                                 false,
	"serviceaccounttokens.project.cattle.io":                          false,
	"settings.management.cattle.io":                                   false,
	"sshauths.project.cattle.io":                                      false,
	"templatecontents.management.cattle.io":                           false,
	"templates.management.cattle.io":                                  false,
	"templateversions.management.cattle.io":                           false,
	"tokens.management.cattle.io":                                     false,
	"userattributes.management.cattle.io":                             false,
	"users.management.cattle.io":                                      false,
	"uiplugins.catalog.cattle.io":                                     true,
	"workloads.project.cattle.io":                                     false,
}
