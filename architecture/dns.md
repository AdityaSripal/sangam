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

The DNS system will consist of two types of entries: `DomainEntry` and `ContentEntry`.

A domain may contain any number of subdomains and content entries. Both domains and entries must
be uniquely namespaced under its subdomain (or globally unique if it is a top-level domain).

A domain is owned by a `DomainOwner` interface, which defines authentication methods 
for adding/updating/removing content and subdomains. Each domain contains a list of subdomains, 
a list of content entries, and a parent domain (top level domains have nil as a domain).

A content entry has a single domain, and an array of content hashes, with each update to the content hash
appended to the array. Thus one can retreive the latest version as well as any previous version from the same entry.

Domains are seperated by the `.` seperator, while content is seperated by the `/` seperator. Thus, a content `blog-post-1` 
stored under the domain hierarchy `domain1 -> domain2 -> domain3` will be stored under the full path: 
`domain1.domain2.domain3/blog-post-1`.

All domain entries are set in store under the key `{DomainPrefix}/{Domain.Path()}`. All Content entries are stored under the 
key `{ContentPrefix}/Content.Path()`.

Each content hash must also have a reverse mapping from `{ReverseContentPrefix}/{ContentHash} -> {Content.Path()}`.


```go
// exported.go
type DomainEntry interface {
    Owner() DomainOwner
    SubDomains() []DomainEntry
    Contents() []ContentEntry
    // returns parent domain, nil if top-level domain
    Parent() DomainEntry

    // returns the string identifier of this domain
    String() string
    // returns the full path string representation
    // up to the global top-level domain
    Path() string

    // Returns the subdomain under this domain as specified by the given string
    GetSubDomain(path string) DomainEntry
    
    // Returns the content under this domain as specified by path String
    // Content may either be in this domain or in a subdomain
    // Returns latest version by default
    GetContent(path string) ContentEntry

    // Returns content at a given sequence
    GetContentAtSequenct(path string, seq uint64) ContentEntry

    // Methods to add/update/delete subdomains and content
    AddSubDomain(domain DomainEntry) error
    AddPrecommit(name string, precommit []byte) error
    AddContent(c ContentEntry, reveal uint64) error
    UpdateSubDomain(name string, domain DomainEntry) error
    UpdateContent(newHash []byte) error
    DeleteSubDomain(name string) error
    DeleteContent(name string)
}

// exported.go
type DomainOwner interface {
    // authenticates add/update/deletes to subdomains of this domain
    AuthenticateDomainChanges(sdk.Msg) error
    // authenticates add/update/deletes to direct content in this domain
    AuthenticateContentChanges(sdk.Msg) error
}

// exported.go
type ContentEntry interface {
    Name() string
    Path() string
    Parent() DomainEntry
    GetContentHashes() [][]byte

    // returns content at latest Version
    GetContent() []byte
    GetContentAtVersion(seq uint64) []byte
}
```

In order to prevent frontrunning of content registration, we implement a simple commit-reveal scheme.

Before adding content under a given domain and name, the domain owner must authenticate a precommit of the content, 
which will get stored under `{PrecommitPrefix}/{Domain.Path()}/Name} -> Precommit`. Only one precommit may be stored for a 
given content path.

```go
type Precommit struct {
    Name string

    Precommit []byte
}
```

## Msgs

We will define three Msg interfaces, `MsgRegisterDomain`, `MsgPreCommit` and `MsgCommitContent`.
Updates to entries will use `MsgPreCommitEntry` and `MsgCommitContent`. Deletions of content/subdomains will include 
the path to be deleted. All concrete msg types may define additional fields to pass domain authentication.

```go
// Register a domain under a given path.
type MsgRegisterDomain interface {
    Domain() string
    ParentPath() string // return full path of parent. Empty string if registering top-level domain
    DomainOwner() DomainOwner
}

// A pre-commit entry will submit a precommit under a domain with a given name
// Hash = hash(content_hash + random_nonce)
type MsgPreCommit interface {
    DomainPath() string // return full path of domain
    Name() string // name of content
    Hash []byte // precomit of content
}

// a commit will be accepted if the the pre-commit hash == hash(content_hash + nonce)
// if successful, the pre-commit is removed and the contenthash is appended to the contenthashes array
// If the content entry does not exist, a new one is created.
type MsgCommitEntry struct {
    Nonce uint64
    ContentHash() []byte
    DomainPath() string
    Name() string
}
```

## Ante

The ante handler for the dns module will assert that the minumum fee for registering domains and content is submitted.

For all registrations/updates/deletions of domains, it will call `parent.DomainOwner().AuthenticateDomainChanges(msg)` which 
will perform arbitrary authentication checks on the msg before allowing the msg to pass.

For precommits/commits/deletions of content, antehandler will call `domain.DomainOwner().AuthenticateContentChanges(msg)`, which 
will perform arbitrary authentication checks on the msg before allowing the msg to pass. For commit msgs, the antehandler will also 
verify that the reveal is valid for the precommit.

## Handler

The domain handler will construct the domain entry and add it under the parent domain.

The precommit handler will construct the precommit entry and store it under the procommit key, replacing the previous precommit 
if it exists. 

The commit handler will append the content hash to the contenthashes array, and construct a new contententry if it does not already 
exist.

Updates and deletions for domains and content are straightforward in handler, once domain-owner authentication passes in antehandler.

## Keeper

We will use byte prefixes to separate storage of pre-commits, entries, reverse entry, and domain mappings.
They will be assinged as follows:

```
byte(0) - pre-commits
byte(1) - regular entries
byte(2) - reverse mapping entries
byte(3) - domains
```

A **domain** mapping looks as follows:

 `{byte(3){domaim.Path()} -> DomainEntry`

A **pre-commit** mapping looks as follows:

`byte(0){content.Path()} -> PreCommit`

An **entry** mapping looks as follows:

`byte(1){content.Path} -> ContentEntry`

A **reverse entry** mapping looks as follows:

`byte(2){contentHash} -> content.Path()`

## Future 

Add governance to update owners of prefixes
