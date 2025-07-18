package management

import (
	"context"

	"github.com/rancher/rancher/pkg/clustermanager"
	"github.com/rancher/rancher/pkg/controllers/management/agentupgrade"
	"github.com/rancher/rancher/pkg/controllers/management/auth"
	"github.com/rancher/rancher/pkg/controllers/management/certsexpiration"
	"github.com/rancher/rancher/pkg/controllers/management/cloudcredential"
	"github.com/rancher/rancher/pkg/controllers/management/cluster"
	"github.com/rancher/rancher/pkg/controllers/management/clusterdeploy"
	"github.com/rancher/rancher/pkg/controllers/management/clustergc"
	"github.com/rancher/rancher/pkg/controllers/management/clusterprovisioner"
	"github.com/rancher/rancher/pkg/controllers/management/clusterstats"
	"github.com/rancher/rancher/pkg/controllers/management/clusterstatus"
	"github.com/rancher/rancher/pkg/controllers/management/drivers/kontainerdriver"
	"github.com/rancher/rancher/pkg/controllers/management/drivers/nodedriver"
	"github.com/rancher/rancher/pkg/controllers/management/node"
	"github.com/rancher/rancher/pkg/controllers/management/secretmigrator"
	"github.com/rancher/rancher/pkg/controllers/management/settings"
	"github.com/rancher/rancher/pkg/controllers/management/usercontrollers"
	"github.com/rancher/rancher/pkg/controllers/managementlegacy"
	"github.com/rancher/rancher/pkg/types/config"
	"github.com/rancher/rancher/pkg/wrangler"
)

func Register(ctx context.Context, management *config.ManagementContext, manager *clustermanager.Manager, wrangler *wrangler.Context) {
	// auth handlers need to run early to create namespaces that back clusters and projects
	// also, these handlers are purely in the mgmt plane, so they are lightweight compared to those that interact with machines and clusters
	auth.RegisterEarly(ctx, management, manager)
	usercontrollers.RegisterEarly(ctx, management, manager)

	// a-z
	agentupgrade.Register(ctx, management)
	certsexpiration.Register(ctx, management)
	cluster.Register(ctx, management)
	clusterdeploy.Register(ctx, management, manager)
	clustergc.Register(ctx, management)
	clusterprovisioner.Register(ctx, management)
	clusterstats.Register(ctx, management, manager)
	clusterstatus.Register(ctx, management)
	kontainerdriver.Register(ctx, management)
	nodedriver.Register(ctx, management)
	cloudcredential.Register(ctx, management, wrangler)
	node.Register(ctx, management, manager)

	secretmigrator.Register(ctx, management)
	settings.Register(ctx, management)
	managementlegacy.Register(ctx, management, manager)

	// Register last
	auth.RegisterLate(ctx, management)
}
