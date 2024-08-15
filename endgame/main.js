const ctx = document.getElementById('myChart');

const teamOneName = document.querySelector('#team_one .name.one');
const teamTwoName = document.querySelector('#team_two .name.two');

const teamOneLogo = document.querySelector('#team_one .logo');
const teamTwoLogo = document.querySelector('#team_two .logo');

const teamOneKDA = document.querySelector('.teamStats .kda.one');
const teamTwoKDA = document.querySelector('.teamStats .kda.two');

const teamOneGold = document.querySelector('.teamStats .gold.one');
const teamTwoGold = document.querySelector('.teamStats .gold.two');

const teamOneDragons = document.querySelector('.teamStats .dragons.one');
const teamTwoDragons = document.querySelector('.teamStats .dragons.two');

const teamOneBarons = document.querySelector('.teamStats .barons.one');
const teamTwoBarons = document.querySelector('.teamStats .barons.two');

const teamOneTowers = document.querySelector('.teamStats .tower.one');
const teamTwoTowers = document.querySelector('.teamStats .tower.two');

const teamOneElder = document.querySelector('.teamStats .elder.one');
const teamTwoElder = document.querySelector('.teamStats .elder.two');

const teamOnebans = document.querySelectorAll('.one .ban');
const teamTwobans = document.querySelectorAll('.two .ban');

const teamOnePlayersIcons = document.querySelectorAll('.playerStats .one .player_icon');
const teamTwoPlayersIcons = document.querySelectorAll('.playerStats .two .player_icon');

const teamOnePlayersBars = document.querySelectorAll('.playerStats .bar .bar_one');
const teamTwoPlayersBars = document.querySelectorAll('.playerStats .bar .bar_two');

const teamOnePlayersDamage = document.querySelectorAll('.playerStats .bar .left p');
const teamTwoPlayersDamage = document.querySelectorAll('.playerStats .bar .right p');

const time = document.querySelector('.time p')

function setKDA(k, d, a, element) {
    element.innerHTML = `${k}/${d}/${a}`;
}

function setGold(gold, element) {
    if (gold > 1000) {
        gold = (gold / 1000).toFixed(1) + 'k';
    }
    element.innerHTML = gold;
}

function setTime(sec, el) {
    let m = Math.floor(sec / 60)
    let s = sec % 60
    m = m < 10 ? `0${m}` : `${m}`
    s = s < 10 ? `0${s}` : `${s}`
    el.innerHTML = `${m}:${s}`
}
function lerp(start, end, amt) {
    return (1 - amt) * start + amt * end
}

function setDamage(el_text, el_bar, value, max) {
    el_bar.style.width = `${lerp(0, 180, value / max)}px`;
    if (value > 1000) {
        value = (value / 1000).toFixed(1) + 'k';
    }
    el_text.innerHTML = value;
}
async function getPlayerInfo(champID) {
    const resp = await fetch(`http://localhost/riot/champ/${champID}/`)
    const data = await resp.json();
    return data;
}

async function setupTeam() {
    let selected = await getTeamInfo();
    teamOneName.innerHTML = selected[0].tag;
    teamTwoName.innerHTML = selected[1].tag;
    teamOneLogo.style.backgroundImage = `url(${selected[0]["logoUrl"]})`;
    teamTwoLogo.style.backgroundImage = `url(${selected[1]["logoUrl"]})`;
    if (match.teamStats[0].win) {
        teamOneName.innerHTML += " WIN";
        teamTwoName.innerHTML += " LOSS";
    }
    else {
        teamOneName.innerHTML += " WIN";
        teamTwoName.innerHTML += " LOSS";
    }
}

async function getTeamInfo() {
    let resp = await fetch("http://localhost/admin/teams/full")
    if (resp.status != 200) {
        return;
    }
    let data = await resp.json();
    const teams = data;
    resp = await fetch("http://localhost/admin/teams/selected")
    if (resp.status != 200) {
        return;
    }
    data = await resp.json();
    const selectedTeam = data;
    let selected = []
    for (let i = 0; i < selectedTeam.length; i++) {
        for (let j = 0; j < teams.length; j++) {
            if (teams[j]["name"] == selectedTeam[i]) {
                selected[i] = teams[j];
            }
        }
    }
    return selected;
}


let resp = await fetch('/admin/match/id');
let gameID = await resp.text();
if (gameID == "Internal Server Error" || gameID == "") {
    setInterval(() => {
        // if game id change 
        fetch('/admin/match/id').then(res => res.text()).then(id => {
            if (id != gameID) {
                window.location.reload();
            }
        });
    }, 5000);
    throw new Error("Internal Server Error, Match is not 5v5 bans or picks are not done");
}
console.log(gameID);
let match = await fetch(`/riot/match/${gameID}/endgame`).then(res => res.json());
let teamOneTotal = 0;
let teamTwoTotal = 0;
let label = []
let gold_diff = match.teamStats[0].goldFrames.map((frame, index) => {
    teamOneTotal += frame
    teamTwoTotal += match.teamStats[1].goldFrames[index]
    label.push(index)
    return frame - match.teamStats[1].goldFrames[index]
});
setupTeam();
setTime(match.duration, time);

setGold(teamOneTotal, teamOneGold);
setGold(teamTwoTotal, teamTwoGold);

setKDA(match.teamStats[0].kills, match.teamStats[0].deaths, match.teamStats[0].assists, teamOneKDA);
setKDA(match.teamStats[1].kills, match.teamStats[1].deaths, match.teamStats[1].assists, teamTwoKDA);

teamOneTowers.innerHTML = match.teamStats[0].towerKills;
teamTwoTowers.innerHTML = match.teamStats[1].towerKills;

teamOneDragons.innerHTML = match.teamStats[0].dragonKills;
teamTwoDragons.innerHTML = match.teamStats[1].dragonKills;

teamOneElder.innerHTML = match.teamStats[0].elderDragonKills;
teamTwoElder.innerHTML = match.teamStats[1].elderDragonKills;

teamOneBarons.innerHTML = match.teamStats[0].baronsKills;
teamTwoBarons.innerHTML = match.teamStats[1].baronsKills;

let max = 0;
match.individualStats.forEach(el => {
    if (el.stats.damageDealtToChampions > max) {
        max = el.stats.damageDealtToChampions;
    }
});

match.individualStats.forEach(async (el) => {
    if (el.participantId < 6) {
        setDamage(teamOnePlayersDamage[el.participantId - 1], teamOnePlayersBars[el.participantId - 1], el.stats.damageDealtToChampions, max);
        let playerInfo = await getPlayerInfo(el.championId);
        teamOnePlayersIcons[el.participantId - 1].style.backgroundImage = `url('${playerInfo.iconURL}')`;
    }
    else {
        setDamage(teamTwoPlayersDamage[el.participantId - 6], teamTwoPlayersBars[el.participantId - 6], el.stats.damageDealtToChampions, max);
        let playerInfo = await getPlayerInfo(el.championId);
        teamTwoPlayersIcons[el.participantId - 6].style.backgroundImage = `url('${playerInfo.iconURL}')`;
    }
});

match.teamStats[0].bans.forEach(async (el, index) => {
    let playerInfo = await getPlayerInfo(el);
    teamOnebans[index].style.backgroundImage = `url('${playerInfo.iconURL}')`;
});
match.teamStats[1].bans.forEach(async (el, index) => {
    let playerInfo = await getPlayerInfo(el);
    teamTwobans[index].style.backgroundImage = `url('${playerInfo.iconURL}')`;
});

new Chart(ctx, {
    type: 'line',
    data: {
        labels: label,
        datasets: [{
            data: gold_diff,
            tension: 0.4,
            cubicInterpolationMode: 'monotone',
            pointStyle: false,
            borderColor: 'black',
            fill: { "above": "#6586ffa0", "below": "#ff8553a0", "target": { "value": 0 } },
            borderWidth: 1,
        }]

    },
    options: {
        aspectRatio: 5,
        legend: {
            display: false
        },
        responsive: true,
        plugins: {
            tooltip: {
                enabled: false
            },
            legend: {
                display: false
            }
        },
        interaction: {
            intersect: false,
        },
        scales: {
            x: {
                display: true,
                suggestedMax: 30,
                suggestedMin: 0,
                ticks: {
                    stepSize: 5,
                    callback: function (value) {
                        if (value === 0) {
                            return value;
                        };
                        if ((value + 1) % 5 === 0) {
                            return value + 1;
                        };
                    },
                },
                grid: {
                    display: true,
                    color: function (context) {
                        if (context.tick.value == 30 || context.tick.value == 0) return "#fff";
                        return '#eee';
                    },

                },
            },
            y: {
                display: true,
                grid: {
                    display: true,
                    color: function (context) {
                        if (context.tick.value == 10000 || context.tick.value == -2000) return "#fff";
                        return '#eee';
                    },
                },
                ticks: {
                    stepSize: 2000,
                    callback: function (value) {
                        var ranges = [
                            { divider: 1e3, suffix: 'k' }
                        ];
                        function formatNumber(n) {
                            for (var i = 0; i < ranges.length; i++) {
                                if (Math.abs(n) >= ranges[i].divider) {
                                    return (n / ranges[i].divider).toString() + ranges[i].suffix;
                                }
                            }
                            return n;
                        }
                        return formatNumber(value);
                    }
                }
            }
        }
    },
});


setInterval(() => {
    // if game id change 
    fetch('/admin/match/id').then(res => res.text()).then(id => {
        if (id != gameID) {
            window.location.reload();
        }
    });
}, 5000);