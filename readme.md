# Lem-in

- Lem-in is a program that simulates the movement of ants across a given ant farm (colony) from a starting room to an end room. The objective is to get all the ants from the starting room to the end room in the least number of moves. The challenge is that each room (except the start and end) can only accommodate one ant at a time and each tunnel can only be used by one ant per turn.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Input File Format](#input-file-format)
- [Output](#output)
- [Contributors](#contributors)


## Installation

1. Clone the repository to your local machine.
```bash
git clone [https://learn.reboot01.com/git/aalali/lem-in.git]
```

2. Navigate to the cloned directory.
```bash
cd lem-in
```

## Usage

1. Navigate to location of main.
```bash
cd main
```

2. To run Lem-in, use the following command:
```bash
go run . [input_file.txt]
```

Example:
```bash
go run . input.txt
```

- Replace [input_file.txt] with the path to your ant farm description file. (below is what you should have inside the .txt file)

## Input File Format

The input file describes the ant farm. It begins with the number of ants, followed by rooms and then tunnels. Here's the format:

    number_of_ants: An integer that denotes the number of ants.
    the_rooms: Defined by "name coord_x coord_y". For example: "nameoftheroom 1 6".
    the_links: Defined by "name1-name2", for instance "1-2", "2-5".

Special commands:

    ##start: Indicates the next room is the starting room.
    ##end: Indicates the next room is the ending room.
    Comments starting with # are ignored.

Example of input file content:

    3
    ##start
    1 23 3
    2 16 7
    3 16 3
    4 16 5
    5 9 3
    6 1 5
    7 4 8
    ##end
    0 9 5
    0-4
    0-6
    1-3
    4-3
    5-2
    3-5
    4-2
    2-1
    7-6
    7-2
    7-4
    6-5

## Output

The output displays the original content of the input file followed by the movement of ants in the format:

```
Lx-y Lz-w Lr-o ...
```
- Where x, z, r represents the ants numbers and y, w, o represents the rooms names.

Example Output when running above example input:

```bash
$ go run . text.txt

L1-3 L2-2
L1-4 L2-5 L3-3
L1-0 L2-6 L3-4
L2-0 L3-0
```

## How Our program works

1- Validate the data given is as required.

2- Seperate the data to: a-number of ants, b- the start room, c- the end room, d- the rooms and the links between the rooms.

3- each room has a boolean which if true shows that the room is occupied.

4- Find all the paths from start to the end and then arrange them from shortest to longest.

5- Add the links to a room struct starting from the one with the shortest way to the end.

6- Create an Ant tracker as a an array which contains the number of the ant and which room its in.

7- Then loop over the ant tracker and check the room its in and what it links to, if the link is not occupied move the ant.

8- if the number of ants in the end room matches the number of ants given the program well stop and you well get the result.

## Contributors

- Ahmed Abdeen
- Ahmed AlAli
- Adnan Jaberi
