package mitreattack

import (
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

type Tactic int

var tactics = []string{
	"Unknown",
	"Initial Access",
	"Execution",
	"Persistence",
	"Privilege Escalation",
	"Defense Evasion",
	"Credential Access",
	"Discovery",
	"Lateral Movement",
	"Collection",
	"Exfiltration",
	"Impact",
}

const (
	UNSPECIFIED Tactic = iota
	InitialAccess
	Execution
	Persistence
	PrivilegeEscalation
	DefenseEvasion
	CredentialAccess
	Discovery
	LateralMovement
	Collection
	Exfiltration
	Impact
)

func AttackTacticFromString(name string) (Tactic, error) {
	lowerName := strings.ToLower(name)
	for i := range tactics {
		if strings.ToLower(tactics[i]) == lowerName {
			return Tactic(i), nil
		}
	}
	return -1, errors.New("unknown MITRE ATT&CK tactic: " + name)
}

func AttackTacticToString(tactic Tactic) string {
	return tactics[tactic]
}

// MarshalYAML implements the Marshaler interface from "gopkg.in/yaml.v3".
// This method makes Tactic type to return a string rather than an int when marshalling to YAML.
func (t Tactic) MarshalYAML() (interface{}, error) {
	return tactics[t], nil
}

// UnmarshalYAML implements the Marshaler interface from "gopkg.in/yaml.v3".
// This method moes the reverse of MarshalYAML, it gets a string with the tactic name mutates into an int.
func (t *Tactic) UnmarshalYAML(node *yaml.Node) error {
	value := node.Value
	newTactic, err := AttackTacticFromString(value)
	if err != nil {
		return err
	}
	*t = newTactic
	return nil
}

func GetAllMitreAttackTactics() []Tactic {
	allTactics := make([]Tactic, 0, len(tactics)-1)
	// Start with '1' to skip the 'Unspecified' tactic
	for i := 1; i < len(tactics); i++ {
		allTactics = append(allTactics, Tactic(i))
	}
	return allTactics
}
