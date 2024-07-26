package cw_logParser

import (
	cw_logreader "cloudwalk-assessment/cw-logReader"
	quake3 "cloudwalk-assessment/quake3"
	"strconv"
	"strings"
)

const filepath string = "assets\\qgames.log.txt";

type Game struct {
	Count int
	Players map[int]Player
	Kills []Kill
}

type Player struct {
	Name string 
}

type Kill struct {
	KillerId int
	KilledId int
	MeanOfDeath quake3.MeanOfDeath
}

type searchingKeywords string
const(
	SK_InitGame searchingKeywords = "InitGame:"
	SK_ClientUserinfoChanged searchingKeywords = "ClientUserinfoChanged:"
	SK_Kill searchingKeywords = "Kill:"
)

func ParseLog() []Game {
	games := make([]Game, 0)
	fileContent := cw_logreader.ReadLogFile(filepath)
	lines := strings.SplitN(fileContent, "\n", -1)

	for _, line := range lines{
		//TODO: Use some better way to compare keywords as switch extracting keywords with regex maybe
		if(searchingKeywordExists(line, SK_InitGame)){
			currentGame := Game{
				Count: len(games) + 1,
				Players: make(map[int]Player, 0),
			}
			games = append(games, currentGame)
		} else if(searchingKeywordExists(line, SK_ClientUserinfoChanged)){
			player, playerId := parsePlayerLine(line)
			games[len(games)-1].Players[playerId] = player
		} else if(searchingKeywordExists(line, SK_Kill)) {
			kill := parseKillLine(line)
			games[len(games)-1].Kills = append(games[len(games)-1].Kills, kill)
		}
	}

	return games
}

func searchingKeywordExists(line string, searchingKeyword searchingKeywords) bool {
	contains := strings.Contains(line, string(searchingKeyword))
	return contains
}

func parsePlayerLine(line string) (Player, int) {
	before, _, _ := strings.Cut(line, "\\t")
	before, playerName, _ := strings.Cut(before,"n\\")
	_, playerId, _ := strings.Cut(before, string(SK_ClientUserinfoChanged))
	playerIdInt, _ := strconv.Atoi(strings.TrimSpace(playerId))

	player := Player {
		Name: playerName,
	}

	return player, playerIdInt;	
}

func parseKillLine(line string) Kill {
	_, after, _ := strings.Cut(line, string(SK_Kill))
	before, _, _ := strings.Cut(after, ":")
	infos := strings.Split(strings.TrimSpace(before), " ")

	killerId, _ := strconv.Atoi(infos[0])
	killedId, _ := strconv.Atoi(infos[1])
	meanOfDeath, _ := strconv.Atoi(infos[2])

	kill := Kill{
		KillerId: killerId,
		MeanOfDeath: quake3.MeanOfDeath(meanOfDeath),
		KilledId: killedId,
	}

	return kill;
}