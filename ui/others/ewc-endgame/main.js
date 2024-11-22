const usernameEl = document.querySelectorAll("thead th:not(.spacer)");
const scoreEl = document.querySelectorAll("tbody tr:nth-child(1) td:not(.spacer)");
const butEl = document.querySelectorAll("tbody tr:nth-child(2) td:not(.spacer)");
const passDEl = document.querySelectorAll("tbody tr:nth-child(3) td:not(.spacer)");
const tirEl = document.querySelectorAll("tbody tr:nth-child(4) td:not(.spacer)");
const arretEl = document.querySelectorAll("tbody tr:nth-child(5) td:not(.spacer)");
const demoEl = document.querySelectorAll("tbody tr:nth-child(6) td:not(.spacer)");
const toucheEl = document.querySelectorAll("tbody tr:nth-child(7) td:not(.spacer)");

const teamImgEl = document.querySelectorAll(".team img");
const teamNameEl = document.querySelectorAll(".team h1");
const teamScoreEl = document.querySelectorAll("#teamScore h1:not(.timer)");


const statsEls = { score: scoreEl, but: butEl, tir: tirEl, arret: arretEl, demo: demoEl, touche: toucheEl, passD: passDEl };

function setPlayerStats(playerId = 0, stats = {}) {
    // Set the player stats
    for (let stat in stats) {
        statsEls[stat][playerId].textContent = stats[stat];
    }
}

function setPlayerName(playerId = 0, name = "") {
    // Set the player name
    usernameEl[playerId].textContent = name;
}

function setTeams(teams = [], logo, score = []) {
    // Set the teams
    for (let i = 0; i < 2; i++) {
        teamImgEl[i].src = logo[i];
        teamNameEl[i].textContent = teams[i];
        teamScoreEl[i].textContent = score[i];
    }
}

async function main() {
    const data = await fetch("http://localhost:80/api/db/ewc/stats")
    let jsonData = await data.json();
    jsonData = jsonData[0];

    const teamsName = jsonData.stats.data.game.teams.map(team => team.name);
    const teamsLogo = [jsonData.match_id.team_one.logo_url, jsonData.match_id.team_two.logo_url];
    const teamsScore = jsonData.stats.data.game.teams.map(team => team.score);
    setTeams(teamsName, teamsLogo, teamsScore);

    let players = jsonData.stats.data.players;
    let pStats = [];
    for ( let i = 0; i < Object.keys(players).length; i++) {
        const element = players[Object.keys(players)[i]];
        const stats = {
            score: element.score,
            but: element.goals,
            passD: element.assists,
            tir: element.shots,
            arret: element.saves,
            demo: element.demos,
            touche: element.touches,
            team: element.team,
            name: element.name
        }
        pStats.push(stats);
    };
    pStats.sort((a, b) => b.team - a.team);

    const names = pStats.map(player => player.name);
    
    // remove team and name items from the array
    pStats = pStats.map(player => {
        delete player.team;
        delete player.name;
        return player;
    });

    for (let i = 0; i < pStats.length; i++) {
        setPlayerName(i, names[i]);
        setPlayerStats(i, pStats[i]);
    }

}

await main();