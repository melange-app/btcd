package btcjson

import (
	"encoding/json"
	"fmt"
)

var _ Cmd = &NameCleanCmd{}

// NameCleanCmd implements "name_clean"
//
// Clean unsatisfiable transactions from the wallet - including name_update on an already taken name
type NameCleanCmd struct {
	id interface{}
}

func NewNameCleanCmd(id interface{}) (*NameCleanCmd, error) {
	return &NameCleanCmd{
		id: id,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameCleanCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameCleanCmd) Method() string {
	return "name_clean"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameCleanCmd) MarshalJSON() ([]byte, error) {
	raw, err := NewRawCmd(n.id, n.Method(), []interface{}{})
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameCleanCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 0 {
		return ErrWrongNumberOfParams
	}

	newCmd, err := NewNameCleanCmd(r.Id)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameDebugCmd{}

// NameDebugCmd implements "name_debug"
//
// Dump pending transactions id in the debug file.
type NameDebugCmd struct {
	id interface{}
}

func NewNameDebugCmd(id interface{}) (*NameDebugCmd, error) {
	return &NameDebugCmd{
		id: id,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameDebugCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameDebugCmd) Method() string {
	return "name_debug"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameDebugCmd) MarshalJSON() ([]byte, error) {
	raw, err := NewRawCmd(n.id, n.Method(), []interface{}{})
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameDebugCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 0 {
		return ErrWrongNumberOfParams
	}

	newCmd, err := NewNameDebugCmd(r.Id)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameDebug1Cmd{}

// NameDebug1Cmd implements "name_debug1"
//
// Dump name blocks number and transactions id in the debug file.
type NameDebug1Cmd struct {
	id   interface{}
	Name string
}

func NewNameDebug1Cmd(id interface{}, name string) (*NameDebug1Cmd, error) {
	return &NameDebug1Cmd{
		id:   id,
		Name: name,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameDebug1Cmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameDebug1Cmd) Method() string {
	return "name_debug1"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameDebug1Cmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameDebug1Cmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	newCmd, err := NewNameDebug1Cmd(r.Id, name)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameFilterCmd{}

// NameFilterCmd implements "name_filter"
//
// name_filter [[[[[regexp] maxage=36000] from=0] nb=0] stat]
// scan and filter names
// [regexp] : apply [regexp] on names, empty means all names
// [maxage] : look in last [maxage] blocks
// [from] : show results from number [from]
// [nb] : show [nb] results, 0 means all
// [stats] : show some stats instead of results
// name_filter "" 5 # list names updated in last 5 blocks
// name_filter "^id/" # list all names from the "id" namespace
// name_filter "^id/" 36000 0 0 stat # display stats (number of names) on active names from the "id" namespace
type NameFilterCmd struct {
	id     interface{}
	Regexp string
	MaxAge int
	From   int
	NB     int
	Stat   bool
}

func NewNameFilterCmd(id interface{}, regexp string, maxage int, from int, nb int, stats bool) (*NameFilterCmd, error) {
	return &NameFilterCmd{
		id:     id,
		Regexp: regexp,
		MaxAge: maxage,
		From:   from,
		NB:     nb,
		Stat:   stats,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameFilterCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameFilterCmd) Method() string {
	return "name_filter"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameFilterCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{}

	if n.Regexp != "" {
		params = append(params, n.Regexp)

		if n.MaxAge != 0 {
			params = append(params, n.MaxAge)

			if n.From != 0 {
				params = append(params, n.From)

				if n.NB != 0 {
					params = append(params, n.NB)

					if n.Stat {
						// Based on the documentation, it looks like we pass
						// the literal string "stat".
						params = append(params, "stat")
					}
				}
			}
		}
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameFilterCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 5 {
		return ErrWrongNumberOfParams
	}

	var (
		regexp string
		maxage int
		from   int
		nb     int
		stat   string
	)
	if len(r.Params) > 0 {
		if err := json.Unmarshal(r.Params[0], &regexp); err != nil {
			return fmt.Errorf("first parameter 'regexp' must be a string: %v", err)
		}

		if len(r.Params) > 1 {
			if err := json.Unmarshal(r.Params[1], &maxage); err != nil {
				return fmt.Errorf("second parameter 'maxage' must be an int: %v", err)
			}

			if len(r.Params) > 2 {
				if err := json.Unmarshal(r.Params[2], &from); err != nil {
					return fmt.Errorf("third parameter 'from' must be an int: %v", err)
				}

				if len(r.Params) > 3 {
					if err := json.Unmarshal(r.Params[3], &nb); err != nil {
						return fmt.Errorf("fourth parameter 'nb' must be an int: %v", err)
					}

					if len(r.Params) > 4 {
						if err := json.Unmarshal(r.Params[0], &stat); err != nil {
							return fmt.Errorf("fifth parameter 'stat' must be a string: %v", err)
						}
					}
				}
			}
		}
	}

	statBool := stat != ""

	newCmd, err := NewNameFilterCmd(r.Id, regexp, maxage, from, nb, statBool)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameFirstUpdateCmd{}

// NameFirstUpdateCmd implements "name_firstupdate"
//
// name_firstupdate <name> <rand> [<tx>] <value> [<toaddress>]
// Perform a first update after a name_new reservation.
// Note that the first update will go into a block 12 blocks after the name_new, at the soonest.
type NameFirstUpdateCmd struct {
	id        interface{}
	Name      string
	Rand      string
	TX        string
	Value     string
	ToAddress string
}

func NewNameFirstUpdateCmd(id interface{}, name string, rand string, tx string, value string, toAddress string) (*NameFirstUpdateCmd, error) {
	return &NameFirstUpdateCmd{
		id:        id,
		Name:      name,
		Rand:      rand,
		TX:        tx,
		Value:     value,
		ToAddress: toAddress,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameFirstUpdateCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameFirstUpdateCmd) Method() string {
	return "name_firstupdate"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameFirstUpdateCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
		n.Rand,
	}

	if n.TX != "" {
		params = append(params, n.TX)
	}

	params = append(params, n.Value)

	if n.ToAddress != "" {
		params = append(params, n.ToAddress)
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameFirstUpdateCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) > 5 || len(r.Params) < 3 {
		return ErrWrongNumberOfParams
	}

	var (
		name      string
		rand      string
		tx        string
		value     string
		toAddress string
	)

	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	if err := json.Unmarshal(r.Params[1], &rand); err != nil {
		return fmt.Errorf("second parameter 'rand' must be a string: %v", err)
	}

	if len(r.Params) == 3 {
		if err := json.Unmarshal(r.Params[2], &value); err != nil {
			return fmt.Errorf("third parameter 'value' must be a string: %v", err)
		}
	} else {
		if err := json.Unmarshal(r.Params[3], &value); err != nil {
			return fmt.Errorf("fourth parameter 'value' must be a string: %v", err)
		}

		if err := json.Unmarshal(r.Params[2], &tx); err != nil {
			return fmt.Errorf("third parameter 'tx' must be a string: %v", err)
		}

		if len(r.Params) == 5 {
			if err := json.Unmarshal(r.Params[4], &toAddress); err != nil {
				return fmt.Errorf("fifth parameter 'toAddress' must be a string: %v", err)
			}
		}
	}

	newCmd, err := NewNameFirstUpdateCmd(r.Id, name, rand, tx, value, toAddress)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameHistoryCmd{}

// NameHistoryCmd implements "name_history"
//
// name_history <name>
// List all name values of a name.
type NameHistoryCmd struct {
	id   interface{}
	Name string
}

func NewNameHistoryCmd(id interface{}, name string) (*NameHistoryCmd, error) {
	return &NameHistoryCmd{
		id:   id,
		Name: name,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameHistoryCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameHistoryCmd) Method() string {
	return "name_history"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameHistoryCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameHistoryCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	newCmd, err := NewNameHistoryCmd(r.Id, name)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

type NameScanResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	Expired   int    `json:"expired"`
	ExpiresIn int    `json:"expires_in"`
}

type NameInfoResult struct {
	Name      string
	Value     string
	TX        string
	Address   string
	ExpiresIn int
	Expired   bool
}

// nameInfoResult is used as a temporary store for the JSON unmarshalling of
// NameInfoResult (capital N).
type nameInfoResult struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	TX        string `json:"txid"`
	Address   string `json:"address"`
	ExpiresIn int    `json:"expires_in"`
	Expired   int    `json:"expired"`
}

func (n *NameInfoResult) UnmarshalJSON(b []byte) error {
	var res *nameInfoResult
	err := json.Unmarshal(b, &res)
	if err != nil {
		return err
	}

	*n = NameInfoResult{
		Name:      res.Name,
		Value:     res.Value,
		TX:        res.TX,
		Address:   res.Address,
		ExpiresIn: res.ExpiresIn,
		Expired:   res.Expired == 1,
	}

	return nil
}

var _ Cmd = &NameListCmd{}

// NameListCmd implements "name_list"
//
// name_list [<name>]
// list my own names
type NameListCmd struct {
	id   interface{}
	Name string
}

func NewNameListCmd(id interface{}, name string) (*NameListCmd, error) {
	return &NameListCmd{
		id:   id,
		Name: name,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameListCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameListCmd) Method() string {
	return "name_list"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameListCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameListCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	newCmd, err := NewNameListCmd(r.Id, name)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameNewCmd{}

// NameNewCmd implements "name_new"
//
// name_new <name>
type NameNewCmd struct {
	id   interface{}
	Name string
}

func NewNameNewCmd(id interface{}, name string) (*NameNewCmd, error) {
	return &NameNewCmd{
		id:   id,
		Name: name,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameNewCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameNewCmd) Method() string {
	return "name_new"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameNewCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameNewCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	newCmd, err := NewNameNewCmd(r.Id, name)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NamePendingCmd{}

// NamePendingCmd implements "name_pending"
//
// List all pending name operations known of.
type NamePendingCmd struct {
	id interface{}
}

func NewNamePendingCmd(id interface{}) (*NamePendingCmd, error) {
	return &NamePendingCmd{
		id: id,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NamePendingCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NamePendingCmd) Method() string {
	return "name_pending"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NamePendingCmd) MarshalJSON() ([]byte, error) {
	raw, err := NewRawCmd(n.id, n.Method(), []interface{}{})
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NamePendingCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 0 {
		return ErrWrongNumberOfParams
	}

	newCmd, err := NewNamePendingCmd(r.Id)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameScanCmd{}

// NameScanCmd implements "name_scan"
//
// name_scan [<start-name>] [<max-returned>]
// scan all names, starting at start-name and returning a maximum number of entries (default 500)
type NameScanCmd struct {
	id          interface{}
	StartName   string
	MaxReturned int
}

func NewNameScanCmd(id interface{}, startName string, maxReturned int) (*NameScanCmd, error) {
	return &NameScanCmd{
		id:          id,
		StartName:   startName,
		MaxReturned: maxReturned,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameScanCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameScanCmd) Method() string {
	return "name_scan"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameScanCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{}

	if n.StartName != "" {
		params = append(params, n.StartName)

		if n.MaxReturned != 0 {
			params = append(params, n.MaxReturned)
		}
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameScanCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 0 && len(r.Params) != 1 && len(r.Params) != 2 {
		return ErrWrongNumberOfParams
	}

	var startName string
	var maxReturned int
	if len(r.Params) > 0 {
		if err := json.Unmarshal(r.Params[0], &startName); err != nil {
			return fmt.Errorf("first parameter 'startname' must be a string: %v", err)
		}

		if len(r.Params) > 1 {
			if err := json.Unmarshal(r.Params[1], &maxReturned); err != nil {
				return fmt.Errorf("second parameter 'maxreturned' must be an int: %v", err)
			}
		}
	}

	newCmd, err := NewNameScanCmd(r.Id, startName, maxReturned)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameShowCmd{}

// NameShowCmd implements "name_show"
//
// Show values of a name.
type NameShowCmd struct {
	id   interface{}
	Name string
}

func NewNameShowCmd(id interface{}, name string) (*NameShowCmd, error) {
	return &NameShowCmd{
		id:   id,
		Name: name,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameShowCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameShowCmd) Method() string {
	return "name_show"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameShowCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameShowCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 1 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	newCmd, err := NewNameShowCmd(r.Id, name)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}

var _ Cmd = &NameUpdateCmd{}

// NameUpdateCmd implements "name_update"
//
// name_update <name> <value> [<toaddress>]
// Update and possibly transfer a name
type NameUpdateCmd struct {
	id        interface{}
	Name      string
	Value     string
	ToAddress string
}

func NewNameUpdateCmd(id interface{}, name string, value string, toAddress string) (*NameUpdateCmd, error) {
	return &NameUpdateCmd{
		id:        id,
		Name:      name,
		Value:     value,
		ToAddress: toAddress,
	}, nil
}

// Id satisfies the Cmd interface by returning the id of the command.
func (n *NameUpdateCmd) Id() interface{} {
	return n.id
}

// Method satisfies the Cmd interface by returning the json method.
func (n *NameUpdateCmd) Method() string {
	return "name_update"
}

// MarshalJSON return the JSON encoding of cmd. Part of the Cmd interface.
func (n *NameUpdateCmd) MarshalJSON() ([]byte, error) {
	params := []interface{}{
		n.Name,
		n.Value,
	}

	if n.ToAddress != "" {
		params = append(params, n.ToAddress)
	}

	raw, err := NewRawCmd(n.id, n.Method(), params)
	if err != nil {
		return nil, err
	}
	return json.Marshal(raw)
}

// UnmarshalJSON unmarshals the JSON encoding of cmd into cmd. Part of the Cmd interface.
func (n *NameUpdateCmd) UnmarshalJSON(b []byte) error {
	// Unmarshal into a RawCmd
	var r RawCmd
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	if len(r.Params) != 2 && len(r.Params) != 3 {
		return ErrWrongNumberOfParams
	}

	var name string
	if err := json.Unmarshal(r.Params[0], &name); err != nil {
		return fmt.Errorf("first parameter 'name' must be a string: %v", err)
	}

	var value string
	if err := json.Unmarshal(r.Params[1], &value); err != nil {
		return fmt.Errorf("second parameter 'value' must be a string: %v", err)
	}

	var toAddress string
	if len(r.Params) == 3 {
		if err := json.Unmarshal(r.Params[2], &toAddress); err != nil {
			return fmt.Errorf("third parameter 'toaddress' must be a string: %v", err)
		}
	}

	newCmd, err := NewNameUpdateCmd(r.Id, name, value, toAddress)
	if err != nil {
		return err
	}

	*n = *newCmd
	return nil
}
