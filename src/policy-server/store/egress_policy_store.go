package store

import (
	"fmt"
	"policy-server/db"
)

//go:generate counterfeiter -o fakes/egress_policy_repo.go --fake-name EgressPolicyRepo . egressPolicyRepo
type egressPolicyRepo interface {
	CreateTerminal(tx db.Transaction) (int64, error)
	CreateApp(tx db.Transaction, sourceTerminalID int64, appGUID string) (int64, error)
	CreateIPRange(tx db.Transaction, destinationTerminalID int64, startIP, endIP, protocol string) (int64, error)
	CreateEgressPolicy(tx db.Transaction, sourceTerminalID, destinationTerminalID int64) (int64, error)
	GetTerminalByAppGUID(tx db.Transaction, appGUID string) (int64, error)
	GetAllPolicies() ([]EgressPolicy, error)
	GetByGuids(ids []string) ([]EgressPolicy, error)
	GetIDsByEgressPolicy(tx db.Transaction, egressPolicy EgressPolicy) (EgressPolicyIDCollection, error)
	DeleteEgressPolicy(tx db.Transaction, egressPolicyID int64) error
	DeleteIPRange(tx db.Transaction, ipRangeID int64) error
	DeleteTerminal(tx db.Transaction, terminalID int64) error
	DeleteApp(tx db.Transaction, appID int64) error
	IsTerminalInUse(tx db.Transaction, terminalID int64) (bool, error)
}

type EgressPolicyStore struct {
	EgressPolicyRepo egressPolicyRepo
}

func (e *EgressPolicyStore) CreateWithTx(tx db.Transaction, policies []EgressPolicy) error {
	for _, policy := range policies {
		sourceTerminalID, err := e.EgressPolicyRepo.GetTerminalByAppGUID(tx, policy.Source.ID)
		if err != nil {
			return fmt.Errorf("failed to get terminal by app guid: %s", err)
		}

		if sourceTerminalID == -1 {
			sourceTerminalID, err = e.EgressPolicyRepo.CreateTerminal(tx)
			if err != nil {
				return fmt.Errorf("failed to create source terminal: %s", err)
			}

			_, err = e.EgressPolicyRepo.CreateApp(tx, sourceTerminalID, policy.Source.ID)
			if err != nil {
				return fmt.Errorf("failed to create source app: %s", err)
			}
		}

		destinationTerminalID, err := e.EgressPolicyRepo.CreateTerminal(tx)
		if err != nil {
			return fmt.Errorf("failed to create destination terminal: %s", err)
		}

		_, err = e.EgressPolicyRepo.CreateIPRange(
			tx,
			destinationTerminalID,
			policy.Destination.IPRanges[0].Start,
			policy.Destination.IPRanges[0].End,
			policy.Destination.Protocol)
		if err != nil {
			return fmt.Errorf("failed to create ip range: %s", err)
		}

		_, err = e.EgressPolicyRepo.CreateEgressPolicy(tx, sourceTerminalID, destinationTerminalID)
		if err != nil {
			return fmt.Errorf("failed to create egress policy: %s", err)
		}
	}
	return nil
}

func (e *EgressPolicyStore) DeleteWithTx(tx db.Transaction, egressPolicies []EgressPolicy) error {
	for _, policy := range egressPolicies {
		egressPolicyIDs, err := e.EgressPolicyRepo.GetIDsByEgressPolicy(tx, policy)
		if err != nil {
			return fmt.Errorf("failed to find egress policy: %s", err)
		}

		err = e.EgressPolicyRepo.DeleteEgressPolicy(tx, egressPolicyIDs.EgressPolicyID)
		if err != nil {
			return fmt.Errorf("failed to delete egress policy: %s", err)
		}

		err = e.EgressPolicyRepo.DeleteIPRange(tx, egressPolicyIDs.DestinationIPRangeID)
		if err != nil {
			return fmt.Errorf("failed to delete destination ip range: %s", err)
		}

		err = e.EgressPolicyRepo.DeleteTerminal(tx, egressPolicyIDs.DestinationTerminalID)
		if err != nil {
			return fmt.Errorf("failed to delete destination terminal: %s", err)
		}

		terminalInUse, err := e.EgressPolicyRepo.IsTerminalInUse(tx, egressPolicyIDs.SourceTerminalID)
		if err != nil {
			return fmt.Errorf("failed to check if source terminal is in use: %s", err)
		}

		if !terminalInUse {
			err = e.EgressPolicyRepo.DeleteApp(tx, egressPolicyIDs.SourceAppID)
			if err != nil {
				return fmt.Errorf("failed to delete source app: %s", err)
			}

			err = e.EgressPolicyRepo.DeleteTerminal(tx, egressPolicyIDs.SourceTerminalID)
			if err != nil {
				return fmt.Errorf("failed to delete source terminal: %s", err)
			}
		}
	}

	return nil
}

func (e *EgressPolicyStore) All() ([]EgressPolicy, error) {
	return e.EgressPolicyRepo.GetAllPolicies()
}

func (e *EgressPolicyStore) ByGuids(ids []string) ([]EgressPolicy, error) {
	policies, err := e.EgressPolicyRepo.GetByGuids(ids)
	if err != nil {
		return []EgressPolicy{}, fmt.Errorf("failed to get policies by guids: %s", err)
	}
	return policies, nil
}
