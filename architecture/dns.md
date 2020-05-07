## Context

Sangam requires a decentralized DNS secured via a blockchain to function. This 
feature can be implemented as a module so other applications can make use of it
and support its maintenance. The DNS will be fairly simple, mapping some key to
some encoded value. The DNS will support payments for registration and updates
as well as security checks to prevent malicious behavior.  

## Decision

We will implement a DNS module supporting the following functionality:

- basic KVStore functionality for any type of DNS entry
- required payment amounts for setting and updating entries
- ownership verification for entry updates
- commit-reveal scheme to prevent validators from stealing registration rights
- optional domain ownership over a certain key prefix
- [future] support governance of DNS entries

## Types

First, we define the `Entry` interface type. An `Entry` should always have at
least one owner and a non-empty byte slice representing the hash of the content. 
The key for the DNS mapping will simply be a `domain/{sequence}`. 

```go
// exported.go
type Entry interface {
    GetOwners()  []sdk.AccAddress // sdk reference the cosmos-sdk
    GetContentHash() []byte
}

// entry.go
type Entry struct {
	// Owners represent the addresses which have authority over this entry.
	Owners []sdk.AccAddress `json:"owners" yaml:"owners"`

	// ContentHash represents the hash of the entry's data.
	ContentHash []byte `json:"content_hash" yaml:"content_hash"`

	// Sequence
	Sequence uint64 `json:"sequence" yaml:"sequence"`
}

// GetOwners returns the owners of the entry.
func (e Entry) GetOwners() []sdk.AccAddress {
	return e.Owners
}

// GetContentHash returns the hash of the content.
func (e Entry) GetContentHash() []byte {
	return e.ContentHash
}

// prefix.go
type PrefixEntry struct {
    Owners []sdk.AccAddress
}

func (pe PrefixEntry) GetOwners() []sdk.AccAddress {
    return pe.PrefixEntry
}

func (pe PrefixEntry) GetContentHash() []byte {
    return nil
}

// domain.go
type Domain struct {
    GetPrefix() string
    GetContentName() string
    GetBytes() []byte // []byte(prefix + content name)
}

// EntryOwner is owner over a specific entry. An address may own many entries.
type EntryOwner {
    Address sdk.AccAddress // owner address
    PrefixOwnershipIndex uint64 // index within the set of prefix owners
}
```

## Msgs

We will define three Msg types, `MsgRegisterPrefix`, `MsgPreCommitEntry` and `MsgCommitEntry`.
Updates to entries will use `MsgPreCommitEntry` and `MsgCommitEntry`

```go
// Register a prefix with a set of owners. This message will fails for already registered prefixes.
type MsgRegisterPrefix struct {
    Prefix string
    Owners []sdk.AccAddress
}

// A pre-commit entry will submit a domain, a hash, and a set of owners.
// Hash = hash(domain + encoded_entry + random_nonce)
type MsgPreCommitEntry struct {
    Domain Domain
    Hash []byte
    Owners EntryOwners // first address must be signer of the message
}

// a commit will be accepted if the the pre-commit hash == hash(domain + encoded_entry + nonce)
// if successful, the pre-commit is removed and the sequence number for the entry is incremented.
// Initial commits have a sequence value of 0
type MsgCommitEntry struct {
    Nonce uint64
    Domain Domain
    Entry exported.Entry
    Signer EntryOwner
}
```

Both ante handlers for pre-commit and commit messages will verify that the sender must be a owner
of the prefix. It will also deduct the registration fee required to register the entry.

## Keeper

We will use byte prefixes to separate storage of pre-commits, entries, and reverse entry mappings.
They will be assinged as follows:

```
byte(0) - pre-commits
byte(1) - regular entries
byte(2) - reverse mapping entries
```

A **domain** is `{prefix}/{contentName}`

A **pre-commit** mapping looks as follows:

`byte(0){domain} -> PreCommit{Owners: []Owner, Hash}`

A **prefix** mapping looks as follows:

`byte(1){prefix} -> PrefixEntry{Owners: []sdk.AccAddress}`

An **entry** mapping looks as follows:

`byte(1){domain}/{sequence} -> Entry{Owners, ContentHash}`

A **reverse entry** mapping looks as follows:

`byte(2){contentHash} -> Entry{Owners, Domain, LatestSequence}`

The **latest entry** can be stored at:

`byte(1){domain}/0`

## Future 

Add governance to update owners of prefixes
