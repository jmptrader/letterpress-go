package main

import (
	"testing"
	"fmt"
)

func TestGameCreation(t *testing.T) {
	var game *Game = Make_empty_game("abcdeabcdeabcdeabcdeabcde", "")
	if game.state.mask[12] != EMPTY {
		fmt.Println(game.state.mask)
		t.Errorf("mask should be empty")
	}

	if game.board[12] != 'c' {
		t.Errorf("game not initialized correctly")
	}
}


func TestBestSubset(t *testing.T) {
	var game *Game = Make_empty_game("abcdeabcdeabcdeabcdeabcjy", "")
	if string(game.possible_words[0].word) != "deejayed" || string(game.possible_words[1].word) != "acceded" {
		t.Errorf("Should have start by deejayed and game acceded ")
	}
	game.sort_possible_words_by_letter_subset(word("cjy"))
	if string(game.possible_words[0].word) != "jaycee" || string(game.possible_words[1].word) != "deejayed" {
		t.Errorf("game not initialized correctly")
	}
}

func TestInterestingLetterset(t *testing.T) {
	var game *Game = Make_empty_game("abcxxdexxxfxxxxxxxxxxxxxx", "rrrr rrr  rr   r         ")
	r := game.interesting_letterset()
	if r[0] != 'x' || len(r) != 1 {
		t.Errorf("something is wrong")
	}
}

func TestEvaluation(t *testing.T) {
	var state GameState
	state.mask = make_mask("rrrr " +
			"rrr  " +
			"rr   " +
			"r    " +
			"     ")
	if state.evaluate() != -16 {
		t.Errorf("something is wrong in eval")
	}
}

func TestWordGen(t *testing.T) {
	var game *Game = Make_empty_game("abcdeabcdeabcdeabcdeabcjy", "")
	var state GameState
	state.played_moves = make([]signedword, 0)

	sw1, sw2, sw3 := game.possible_words[0],game.possible_words[1],game.possible_words[2]

	state.played_moves = append(state.played_moves, sw2)
	var wi worditerator
	first := wi.Begin(game, &state)
	if !first.Equal(&sw1) {
		t.Errorf("Should have been ", sw1.word)
	}
	state.played_moves = append(state.played_moves, sw1)
	first = wi.Begin(game, &state)
	if !first.Equal(&sw3) {
		t.Errorf("Should have been ", sw3.word)
	}

}