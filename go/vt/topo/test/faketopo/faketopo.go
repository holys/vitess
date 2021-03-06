// Package faketopo contains utitlities for tests that have to interact with a
// Vitess topology.
package faketopo

import (
	"errors"

	"github.com/youtube/vitess/go/vt/topo"
	"golang.org/x/net/context"
)

var errNotImplemented = errors.New("Not implemented")

// FakeTopo is a topo.Server implementation that always returns errNotImplemented errors.
type FakeTopo struct{}

// GetSrvKeyspaceNames implements topo.Server.
func (ft FakeTopo) GetSrvKeyspaceNames(ctx context.Context, cell string) ([]string, error) {
	return nil, errNotImplemented
}

// GetSrvKeyspace implements topo.Server.
func (ft FakeTopo) GetSrvKeyspace(ctx context.Context, cell, keyspace string) (*topo.SrvKeyspace, error) {
	return nil, errNotImplemented
}

// GetEndPoints implements topo.Server.
func (ft FakeTopo) GetEndPoints(ctx context.Context, cell, keyspace, shard string, tabletType topo.TabletType) (*topo.EndPoints, int64, error) {
	return nil, -1, errNotImplemented
}

// Close implements topo.Server.
func (ft FakeTopo) Close() {}

// GetKnownCells implements topo.Server.
func (ft FakeTopo) GetKnownCells(ctx context.Context) ([]string, error) {
	return nil, errNotImplemented
}

// CreateKeyspace implements topo.Server.
func (ft FakeTopo) CreateKeyspace(ctx context.Context, keyspace string, value *topo.Keyspace) error {
	return errNotImplemented
}

// UpdateKeyspace implements topo.Server.
func (ft FakeTopo) UpdateKeyspace(ctx context.Context, ki *topo.KeyspaceInfo, existingVersion int64) (int64, error) {
	return 0, errNotImplemented
}

// GetKeyspace implements topo.Server.
func (ft FakeTopo) GetKeyspace(ctx context.Context, keyspace string) (*topo.KeyspaceInfo, error) {
	return nil, errNotImplemented
}

// GetKeyspaces implements topo.Server.
func (ft FakeTopo) GetKeyspaces(ctx context.Context) ([]string, error) {
	return nil, errNotImplemented
}

// DeleteKeyspaceShards implements topo.Server.
func (ft FakeTopo) DeleteKeyspaceShards(ctx context.Context, keyspace string) error {
	return errNotImplemented
}

// CreateShard implements topo.Server.
func (ft FakeTopo) CreateShard(ctx context.Context, keyspace, shard string, value *topo.Shard) error {
	return errNotImplemented
}

// UpdateShard implements topo.Server.
func (ft FakeTopo) UpdateShard(ctx context.Context, si *topo.ShardInfo, existingVersion int64) (int64, error) {
	return 0, errNotImplemented
}

// ValidateShard implements topo.Server.
func (ft FakeTopo) ValidateShard(ctx context.Context, keyspace, shard string) error {
	return errNotImplemented
}

// GetShard implements topo.Server.
func (ft FakeTopo) GetShard(ctx context.Context, keyspace, shard string) (*topo.ShardInfo, error) {
	return nil, errNotImplemented
}

// GetShardNames implements topo.Server.
func (ft FakeTopo) GetShardNames(ctx context.Context, keyspace string) ([]string, error) {
	return nil, errNotImplemented
}

// DeleteShard implements topo.Server.
func (ft FakeTopo) DeleteShard(ctx context.Context, keyspace, shard string) error {
	return errNotImplemented
}

// CreateTablet implements topo.Server.
func (ft FakeTopo) CreateTablet(ctx context.Context, tablet *topo.Tablet) error {
	return errNotImplemented
}

// UpdateTablet implements topo.Server.
func (ft FakeTopo) UpdateTablet(ctx context.Context, tablet *topo.TabletInfo, existingVersion int64) (newVersion int64, err error) {
	return 0, errNotImplemented
}

// UpdateTabletFields implements topo.Server.
func (ft FakeTopo) UpdateTabletFields(ctx context.Context, tabletAlias topo.TabletAlias, update func(*topo.Tablet) error) error {
	return errNotImplemented
}

// DeleteTablet implements topo.Server.
func (ft FakeTopo) DeleteTablet(ctx context.Context, alias topo.TabletAlias) error {
	return errNotImplemented
}

// GetTablet implements topo.Server.
func (ft FakeTopo) GetTablet(ctx context.Context, alias topo.TabletAlias) (*topo.TabletInfo, error) {
	return nil, errNotImplemented
}

// GetTabletsByCell implements topo.Server.
func (ft FakeTopo) GetTabletsByCell(ctx context.Context, cell string) ([]topo.TabletAlias, error) {
	return nil, errNotImplemented
}

// UpdateShardReplicationFields implements topo.Server.
func (ft FakeTopo) UpdateShardReplicationFields(ctx context.Context, cell, keyspace, shard string, update func(*topo.ShardReplication) error) error {
	return errNotImplemented
}

// GetShardReplication implements topo.Server.
func (ft FakeTopo) GetShardReplication(ctx context.Context, cell, keyspace, shard string) (*topo.ShardReplicationInfo, error) {
	return nil, errNotImplemented
}

// DeleteShardReplication implements topo.Server.
func (ft FakeTopo) DeleteShardReplication(ctx context.Context, cell, keyspace, shard string) error {
	return errNotImplemented
}

// LockSrvShardForAction implements topo.Server.
func (ft FakeTopo) LockSrvShardForAction(ctx context.Context, cell, keyspace, shard, contents string) (string, error) {
	return "", errNotImplemented
}

// UnlockSrvShardForAction implements topo.Server.
func (ft FakeTopo) UnlockSrvShardForAction(ctx context.Context, cell, keyspace, shard, lockPath, results string) error {
	return errNotImplemented
}

// GetSrvTabletTypesPerShard implements topo.Server.
func (ft FakeTopo) GetSrvTabletTypesPerShard(ctx context.Context, cell, keyspace, shard string) ([]topo.TabletType, error) {
	return nil, errNotImplemented
}

// CreateEndPoints implements topo.Server.
func (ft FakeTopo) CreateEndPoints(ctx context.Context, cell, keyspace, shard string, tabletType topo.TabletType, addrs *topo.EndPoints) error {
	return errNotImplemented
}

// UpdateEndPoints implements topo.Server.
func (ft FakeTopo) UpdateEndPoints(ctx context.Context, cell, keyspace, shard string, tabletType topo.TabletType, addrs *topo.EndPoints, existingVersion int64) error {
	return errNotImplemented
}

// DeleteEndPoints implements topo.Server.
func (ft FakeTopo) DeleteEndPoints(ctx context.Context, cell, keyspace, shard string, tabletType topo.TabletType, existingVersion int64) error {
	return errNotImplemented
}

// WatchEndPoints implements topo.Server.
func (ft FakeTopo) WatchEndPoints(ctx context.Context, cell, keyspace, shard string, tabletType topo.TabletType) (<-chan *topo.EndPoints, chan<- struct{}, error) {
	return nil, nil, errNotImplemented
}

// UpdateSrvShard implements topo.Server.
func (ft FakeTopo) UpdateSrvShard(ctx context.Context, cell, keyspace, shard string, srvShard *topo.SrvShard) error {
	return errNotImplemented
}

// GetSrvShard implements topo.Server.
func (ft FakeTopo) GetSrvShard(ctx context.Context, cell, keyspace, shard string) (*topo.SrvShard, error) {
	return nil, errNotImplemented
}

// DeleteSrvShard implements topo.Server.
func (ft FakeTopo) DeleteSrvShard(ctx context.Context, cell, keyspace, shard string) error {
	return errNotImplemented
}

// UpdateSrvKeyspace implements topo.Server.
func (ft FakeTopo) UpdateSrvKeyspace(ctx context.Context, cell, keyspace string, srvKeyspace *topo.SrvKeyspace) error {
	return errNotImplemented
}

// LockKeyspaceForAction implements topo.Server.
func (ft FakeTopo) LockKeyspaceForAction(ctx context.Context, keyspace, contents string) (string, error) {
	return "", errNotImplemented
}

// UnlockKeyspaceForAction implements topo.Server.
func (ft FakeTopo) UnlockKeyspaceForAction(ctx context.Context, keyspace, lockPath, results string) error {
	return errNotImplemented
}

// LockShardForAction implements topo.Server.
func (ft FakeTopo) LockShardForAction(ctx context.Context, keyspace, shard, contents string) (string, error) {
	return "", errNotImplemented
}

// UnlockShardForAction implements topo.Server.
func (ft FakeTopo) UnlockShardForAction(ctx context.Context, keyspace, shard, lockPath, results string) error {
	return errNotImplemented
}
