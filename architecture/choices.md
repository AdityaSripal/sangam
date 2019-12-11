# Choices

### Why IPFS?

IPFS provides a content storage and distribution layer that is decentralized and censorship-resistant. Using a blockchain to serve as a storage layer would lead to immediate scalability issues since every node must store all content from every user. Systems like IPFS are ideal for content distribution since nodes only store and host content that they themselves are interested in.

IPFS also offers advantages over other decentralized storage systems for its flexibility. While other systems might enforce charging for storage at the protocol level, IPFS does not. While Filecoin does exist as a way for data creators to incentivize data storers, Sangam exists as a way for data storers to incentivize data creators. The fact that both can exist on IPFS and can interact with each other is a big advantage.

### Why Cosmos

The Cosmos-SDK with its use of modules in a singular baseapp allows for very quick independent iteration of the different components of the Sangam blockchain (governance, routing, payments). 
IBC allows payments from any currency. Any IBC-compatible blockchain can send TokenTransfer IBC packets to the Sangam blockchain where they can be used in donations. Non-IBC compatible blockchains can send payments through an IBC-compatible pegzone. Token-agnosticism with IBC is an important feature of the Sangam donation model.

The SDK and Cosmos ecosystem will increasingly devote resources towards solving the problems of decentralized governance and liquid democracy. Sangam will need a robust governance system in order to scale beyond small communities while remaining a safe place for both creators and users. The liquid democracy work will benefit both Sangam's own use of liquid democracy and can be coopted to build the liquid donation system since their implementations will both face similar challenges. Being able to piggyback off of the development work on projects like Virgo will allow Sangam to reach its milestones more quickly.
