# Core Architecture

Sangam will integrate IPFS to store and distribute content and Cosmos blockchain(s) to facilitate donations, content moderation, authorship, and content addressing. This document is a brief overview of how these components will fit together to achieve the necessary functionality to build Sangam.

### Submitting Content

Creators register profiles on the Sangam blockchain to reserve a unique namespace and register donation addresses:

```go
type Creator struct {
    Name       string
    SangamAddr address
    BtcAddr    address
    EthAddr    address
    AtomAddr   address
    // more crypto addresses ...
}
```

A mapping will be stored from `name => Creator` on the Sangam blockchain.

Registering as a Creator on Sangam will require burning a fixed amount of a native token. This serves two purposes: it reduces the threat of spammers flooding the network, and it allows Sangam to be seeded by a close-knit community. The genesis distribution of the native tokens will largely be the contributors to the Sangam codebase; these genesis token holders can then gift their tokens to friends and family who might be interested in participating on Sangam. Thus, the tokens function as a sort of invite voucher to the Sangam network that can help filter out bad actors and spammers in the early operation of the Sangam network and keep the initial community close-knit and passionate about Sangam's success. This system differs from other invite-only mechanisms since the invites are not solely being distributed by the platform maintainers. Anyone who owns tokens can invite anyone else into the network. This avoids a completely permissionless setting that may be too chaotic to handle early on, while staying true to a decentralized ethos. Once the network becomes more robust, these tokens will be more widely distributed and trivially acquirable for anyone who wants to become a Creator. By that point, there will also be a robust governance system in place to remove bad actors from the network. However, the initial invite mechanism will force Sangam to grow organically and sustainably while protecting it from attackers in its nascency.

Once a Creator is registered on Sangam, they can start uploading content to IPFS where it will be stored and addressed by its hash. Since hashes are not human readable, the Sangam blockchain must provide a DNS service for its creators so that their fans can easily route to their content. Creators can set a human readable path to their IPFS content by submitting a `RegisterContent` message:

```go
type RegisterMsg struct {
   Path   string
   Name   string
   Hash   IPFSHash
   Signer address
}
```

This will then store a mapping from the human readable string to the IPFSHash the content is stored under: `name/path => Hash`.
If I have a blog post that hashes to `0x123`, I can create a human readable path by submitting:

```go
RegisterMsg{
   Path: "blog/politics/my_opinion.txt",
   Name: "AdityaSripal",
   Hash: 0x123,
   Signer: {my_sangam_address},
}
```

This will then allow my readers to retrieve the blog from IPFS by first querying `AdityaSripal/blog/politics/my_opinion.txt` on the Sangam blockchain to retrieve the hash, and then querying IPFS with the hash to retrieve the blog post. Of course, a Sangam client will do most of this on behalf of the user.

However, IPFS objects themselves may be `Tree` objects which have in built human readable paths to each node in the tree. Thus, a Sangam client must be able to seamlessly extend the path it retrieved from the Sangam blockchain with the tree path so that it can retrieve content at individual nodes in the tree object.
Suppose I upload this repository to IPFS as a `Tree` object and register it under the path, `AdityaSripal/sangam`, I should be able to retrieve this file by querying:
`AdityaSripal/sangam/architecture/core.md`.

NOTE: Since `Tree` objects can specify their own human-readable paths, a node in a IPFS `Tree` object may have a direct pathname on Sangam AND an indirect pathname (if a parent node has a direct pathname). This can cause an IPFS node to have two valid pathnames on Sangam. It requires further thought whether enforcing path coherence is possible or even desirable.

NOTE: An attack is possible where a hash registration can be "frontrun" by an attacker who wants to plagiarize another creator's content. This attack can be trivially defeated by using a commit-reveal scheme where the creator first makes a blinded commitment to his IPFS hash and later reveals its true value in a second message.

### Donations

Any object in Sangam can implement the `DonatorModel` interface to facilitate receiving donations. The `DonationModel` interface is defined as:

```go
type DonationModel interface {
    GetDonationDenoms() []string
    Donate(Coins) error
}
```

`GetDonationDenoms` returns the list of denominations that this object accepts as donations. The `Donate` method can implement arbitrary donation logic.

Both Creators and individual pieces of content can implement a custom DonationModel. In the beginning, the DonationModel will be restricted to a very simplistic tranfer of all coins to a single address. However future updates will allow for more complex donation models that can send funds to multiple addresses or even recursively funnel donations through multiple `DonationModel` objects.
This will allow for the liquid-donations mechanism that will be a key feature of the Sangam platform.

Since these are Cosmos blockchains, IBC will enable users to donate in whichever cryptocurrency they prefer. Any blockchain that implements IBC can have its cryptocurrencies used in Sangam donations. Cosmos pegzones will allow non-IBC compatible cryptocurrencies to be used in donations as well. Thus, Sangam makes no restrictions on which cryptocurrencies users can donate with; the only restriction is what creators are willing to accept.

### Governance

The Sangam community will have to come to social consensus on what policies to adopt and who will be in charge of enforcing them.

These policies will include but are not limited to:
- what constitutes plagiarism, and how should it be punished
- what constitutes unacceptable content, how should it be dealt with? How should the author be punished?
- what are the community's collective values and their respective priorities?

The community must also elect committees who will take these adopted policies and enforce them as fairly as possible. Where possible, the evidence and reasoning for a given decision will be publishedoon Sangam itself so that they can be inspected by any interested member of the community. Any decision deemed unfair can be broadcasted to the wider community, and if sufficient consensus exists; the decision can be overruled. If there is persistent abuse of power, the community can vote to disband a committee and re-elect them.

Most users may not be interested in participating in governance, so it is crucially important to implement a liquid democracy system on the Sangam blockchain where users can delegate their votes to more active participants and override with their own preferences when necessary.

Here again, a native token is necessary to tally votes on community proposals. Most blockchain governance systems today are completely plutocratic since the richest users own the most tokens and can tus control the system. Sangam will not ICO its tokens to the highest bidder but instead try to create mechanisms that distribute the token supply to approximately reward the users in proportion to how much they contribute to the network (either through creations or donations); a "proof-of-participation" system rather than a proof-of-wealth system.


