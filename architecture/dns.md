## Context

Sangam requires a decentralized DNS secured via a blockchain to function. This 
feature can be implemented as a module so other applications can make use of it
and support it's maintenance. The DNS will be fairly simple, mapping some key to
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
least one owner and a non-empty byte array for content. The key for the DNS mapping
will simply be a byte array. 

```go
type Entry interface {
    Owners()  [][]byte // TODO: how general should the owner addresses be?
    Content() []byte
}
```

## Msgs

We will define two Msg types, `MsgPreCommitEntry` and `MsgCommitEntry`. A pre-commit
entry will submit a key and a hash. The hash is a hash of the encoded entry with a 
random nonce. The commit entry message will verify that the hash registered at
key is equal to the hash of the encoded entry.

```go
// Hash is the hash of a random nonce and the encoded entry
type MsgPreCommitEntry struct {
    Key  []byte
    Hash []byte
}

// Key must be equal to the Key used in the pre-commit message
type MsgCommitEntry struct {
    Nonce uint64
    Key   []byte
    Entry Entry
    
}
```

Both ante handlers for pre-commit and commit messages will verify that if the
key is a prefixed domain then the sender must be the owner of that prefix. It
will also check that the amount spent in the transaction is equal to the amount
required to register the entry.

We will also need to define a message to handle update functionality. The update
message does not need to use a commit-reveal scheme since the owners were defined 
in the original registration. Future versions should support updating of ownership
based on a preset percentage of signatures provided with the message. For now
it will remain impossible to update the owners of an entry. Updates can only be made
to the key and the content in an entry.

```go
type MsgUpdateEntry struct {
    Key []byte
    Entry Entry
}
```


