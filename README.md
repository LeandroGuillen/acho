# Acho

A Risk-like country-conquering board game.

## Game flow
Pseudo-code explaining the flow of the game.

    for each player do
        1. Select owned country to deploy troops on
        2. Do one of the following
            2.1 Select non-owned country to attack it
            2.2 Select owned country to build a port
            2.3 Pass turn
        3. Move troops from owned country to owned country

## The map
The map will be represented with a graph. Each node of the graph represents a country, and a link between two nodes represents the adjacency of one country to another.

Every *node* has to contain the following information:

- The number of troops deployed on it. They can be either neutral or belong to a certain player.
- The number of troops that will be generated next turn. For neutral countries troops do not increase.
- Whether the country is a capital or not.
- Connections to other countries represented as links to other nodes.

## Attacking
If a player attacks

## Configuration
Tentative configurable values for a map:

- Width and height.
- % of water.
- Size of country (average and standard deviation).
- Size of water bodies (average and standard deviation).
- Uniformity of country shape (rugged/soft borders).
- % of neutral countries with troops.
- Number of neutral troops (average and standard deviation).


Game configuration:

- Capital countries present (yes/no).
- What happens when a capital country dies? Three options:
	- Player loses, all remaining countries become neutral
	- Player loses, all remaining countries become loyal to the conqueror player
	- Player loses % of countries to conquering player and a new capital is chosen from remaining countries randomly
- Number of troops generated by country each turn.