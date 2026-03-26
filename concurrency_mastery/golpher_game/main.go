package main

import (
	"context"
	"fmt"
	"maps"
	"math/rand"
	"time"
)

type ActionType int

const (
	Move ActionType = iota
	Attack
)

type Command struct {
	PlayerID int
	Type     ActionType
	Value    int
}

type GameState struct {
	Positions map[int]int
}

func golpher(id int, cmdChan chan<- Command, stateChan <-chan GameState, ctx context.Context) {
	currentPosition := -1
	initialized := false
	for {
		select {
		case <-ctx.Done():
			return
		case world, ok := <-stateChan:
			if !ok {
				return
			}
			if pos, ok := world.Positions[id]; ok {
				currentPosition = pos
				initialized = true
			}
			if !initialized {
				continue
			}
			fmt.Printf("Golhper %d received the map: %v\n", id, world)

			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			direction := 1
			if rand.Intn(2) == 0 {
				direction = -1
			}
			newPosition := currentPosition + direction
			select {
			case <-ctx.Done():
				return
			case cmdChan <- Command{
				PlayerID: id,
				Type:     Move,
				Value:    newPosition,
			}:
			default:
			}

		}
	}
}

func coordinator(cmdChan <-chan Command, golphers []chan GameState, initialPositions map[int]int, ctx context.Context) {
	gameState := GameState{
		Positions: initialPositions,
	}
	dirty := false
	timer := time.NewTicker(16 * time.Millisecond)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case cmd := <-cmdChan:
			targetPosition := cmd.Value
			alreadyTaken := false
			for _, pos := range gameState.Positions {
				if pos == targetPosition {
					alreadyTaken = true
					break
				}
			}
			if !alreadyTaken {
				gameState.Positions[cmd.PlayerID] = targetPosition
				dirty = true
				fmt.Printf("Golpher %d moved to %d\n", cmd.PlayerID, targetPosition)

			} else {
				fmt.Printf("❌ Move request to %d by Golpher %d is rejected\n", cmd.Value, cmd.PlayerID)
			}
		case <-timer.C:
			if !dirty {
				continue
			}
			snapshot := make(map[int]int)
			maps.Copy(snapshot, gameState.Positions)
			snapState := GameState{Positions: snapshot}

			for _, golpher := range golphers {
				select {
				case <-ctx.Done():
					return
				case golpher <- snapState:
				default:
					// if golpher is busy, just skip this tick
				}
			}
			dirty = false
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var command chan Command = make(chan Command, 5000)
	var golphers []chan GameState
	var golphersCount int = 10

	var initialPositions map[int]int = make(map[int]int)
	for i := 0; i < golphersCount; i++ {
		initialPositions[i] = i * 2
	}

	for i := 0; i < golphersCount; i++ {
		var gCh chan GameState = make(chan GameState, 1)
		golphers = append(golphers, gCh)
		id := i
		go golpher(id, command, gCh, ctx)
	}
	go coordinator(command, golphers, initialPositions, ctx)

	<-ctx.Done()
}
