# server
A simple game server

# State
State keeps data in memory, saving it to a JSON file as needed.

# Components
The engine needs to be support multiple games.

Users should be able to join multiple games concurrently.
When a user joins a game, they create a Player.

(Seems fiddly, but Player is an instance of a User in a particular Game.)

The Player is assigned a race and homeworld.
They can submit and amend orders in their game.

At the the end of each Turn, all orders are combined and sorted.
The sort is to ensure that the engine processes them in the correct order.
Sequence is important for some steps, so we need a stable sort on the orders.

After the engine processes all the orders for a Turn,
it saves the State back out to that JSON file.
Players are notified and use the front-end to run queries
against the new State.

Permissions seem basic - every Player has access to their data.
That data would include their sensor probes and whatever
their Ships were able to see in a System.

# Mock Data
The mock data uses character names from
[Stan Sakai](https://stansakai.com/)'s
[Usagi Yojimbo](http://www.usagiyojimbo.com/).

From the [Wikipedia](https://en.wikipedia.org/wiki/Usagi_Yojimbo) article:

    Usagi Yojimbo (兎用心棒, Usagi Yōjinbō, "rabbit bodyguard") is a
    comic book series created by Stan Sakai. It is set primarily at
    the beginning of the Edo period of Japanese history and features
    anthropomorphic animals replacing humans. The main character is
    a rabbit rōnin, Miyamoto Usagi, whom Sakai based partially on
    the famous swordsman Miyamoto Musashi. Usagi wanders the land on
    a musha shugyō (warrior's pilgrimage), occasionally selling his
    services as a bodyguard. 
