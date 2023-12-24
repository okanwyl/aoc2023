#!/usr/bin/python3

import re

digit = r"\d+"

TR = 12
TG = 13
TB = 14

from dataclasses import dataclass


@dataclass
class GameHand:
    belong: int = 0
    r: int = 0
    g: int = 0
    b: int = 0


possible_game_ids = []

with open("input.txt", "r") as f:
    for line in f:
        game_split = line.split(";")
        game_id = re.findall(digit, game_split[0].rsplit(":")[0])[0]

        game_split[0] = str(game_split[0]).rsplit(":")[1]
        parsed_games = []
        for game in game_split:
            expect_word = False
            count = 0
            hand = GameHand()
            for word in game.split():
                if expect_word:
                    expect_word = False

                    match word.removesuffix(","):
                        case "red":
                            hand.r = count
                        case "blue":
                            hand.b = count
                        case "green":
                            hand.g = count
                    hand.belong = game_id

                if word.isdigit():
                    count = int(word)
                    expect_word = True

            parsed_games.append(hand)

        mr = 0
        mg = 0
        mb = 0
        for game_hand in parsed_games:
            if game_hand.r > mr:
                mr = game_hand.r
            if game_hand.g > mg:
                mg = game_hand.g
            if game_hand.b > mb:
                mb = game_hand.b

        possible_game_ids.append(mr * mg * mb)


print(sum(list(map(int, possible_game_ids))))
