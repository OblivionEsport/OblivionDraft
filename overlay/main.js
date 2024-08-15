let team1 = [].slice.call(document.getElementById("team1").children).slice(2);
let team2 = [].slice.call(document.getElementById("team2").children).slice(2);
let bansTeam1 = document.getElementById("bansTeam1");
let bansTeam2 = document.getElementById("bansTeam2");
let timer = document.getElementById("timer");

let players = document.getElementsByClassName("player");

let is_connected = false;

let ref_timer = 0;
let current_timer = 0;


async function setup() {
    let rawData = await fetch("http://localhost/draft/full")
    if (rawData.status != 200) {
        is_connected = false;
        return;
    }
    let data = await rawData.json();
    is_connected = true;
    // Players
    for (let i = 0; i < 10; i++) {
        if (i < 5) {
            team1[i].dataset.summonerId = data["myTeam"][i]["summonerId"];
            team1[i].dataset.championId = data["myTeam"][i]["championId"];
        } else {
            team2[i - 5].dataset.summonerId = data["theirTeam"][i - 5]["summonerId"];
            team2[i - 5].dataset.championId = data["theirTeam"][i - 5]["championId"];
        }
    }
    showInfo()
    // Bans
    while (bansTeam1.firstChild) {
        bansTeam1.removeChild(bansTeam1.firstChild);
    }
    while (bansTeam2.firstChild) {
        bansTeam2.removeChild(bansTeam2.firstChild);
    }
    for (let i = 0; i < data["bans"]["numBans"]; i++) {
        let div = document.createElement("div");
        div.classList.add("bans");
        if (i < (data["bans"]["numBans"] / 2)) {
            bansTeam1.appendChild(div);
        } else {
            bansTeam2.appendChild(div);
        }
    }
    // Timer 
    sec = data["timer"]["AdjustedTimeLeftInPhase"] / 1000
    timer.innerText = formatTime(sec);
}

async function actions() {
    let rawData = await fetch("http://localhost/draft/actions")
    if (rawData.status != 200) {
        is_connected = false;
        return;
    }
    let data = await rawData.json();
    is_connected = true;
    progress = false;
    // data is an array of array of actions, get the last one that isInProgress
    for (i = 0; i < data.length; i++) {
        for (j = 0; j < data[i].length; j++) {
            let action = data[i][j];
            if (action["isInProgress"] && action["type"] == "pick") {
                progress = true
                let id = action["actorCellId"];
                otherTeam = (id < 5) ? 2 : 1;
                resetTeam(otherTeam);
                setActive(id);
                return;
            }
        }
    }
    if (!progress) {
        resetTeam(1);
        resetTeam(2);
    }
}

async function updateCounter() {
    let rawData = await fetch("http://localhost/draft/timer")
    if (rawData.status != 200) {
        is_connected = false;
        return;
    }
    let data = await rawData.json();
    is_connected = true;
    if (data["timeLeft"] == ref_timer) {
        if (current_timer > 0) {
            current_timer -= 1000;
        }
    } else {
        ref_timer = data["timeLeft"];
        current_timer = ref_timer;
    }
    sec = current_timer / 1000;
    timer.innerText = formatTime(sec);
}

function formatTime(sec) {
    sec_ = Math.floor(sec % 60);
    if (sec_ < 10)
        sec_ = `0${sec_}`;
    min = Math.floor(sec / 60);
    if (min > 0)
        return `${min}:${sec_}`;
    else
        return `${sec_}`;
}



async function showInfo() {
    for (let i = 0; i < players.length; i++) {
        let id = players[i].dataset.summonerId;
        let champID = players[i].dataset.championId;
        let data = await getPlayerInfo(id, champID);

        players[i].style.backgroundImage = `url('${data['splashUrl']}')`;
        players[i].children[0].innerText = data['summonerName'];
        //players[i].children[1].innerText = data['championName'];
    }
}
function setActive(id) {
    let activePlayer = players[id];
    let activePlayerTeam = activePlayer.parentElement;
    sizeArray = ""
    if (id > 4) id -= 5;
    for (let i = 0; i < activePlayerTeam.children.length - 2; i++) {
        if (i == id) {
            sizeArray += "1.8fr ";
        } else {
            sizeArray += "1fr ";
        }
    }
    activePlayerTeam.style.gridTemplateRows = `.5fr ${sizeArray} .5fr`;
}
function resetTeam(team) {
    if (team == 1)
        team1[0].parentElement.style.gridTemplateRows = ".5fr 1fr 1fr 1fr 1fr 1fr .5fr";
    else
        team2[0].parentElement.style.gridTemplateRows = ".5fr 1fr 1fr 1fr 1fr 1fr .5fr";
}

async function getTeamInfo() {
    resp = await fetch("http://localhost/admin/teams/full")
    if (resp.status != 200) {
        return;
    }
    data = await resp.json();
    teams = data;
    resp = await fetch("http://localhost/admin/teams/selected")
    if (resp.status != 200) {
        return;
    }
    data = await resp.json();
    selectedTeam = data;
    selected = []
    for (i = 0; i < selectedTeam.length; i++) {
        for (j = 0; j < teams.length; j++) {
            if (teams[j]["name"] == selectedTeam[i]) {
                selected[i] = teams[j];
            }
        }
    }
    return selected;
}

async function setupTeamInfo() {
    let teamName = document.getElementsByClassName("teamName");
    let blob = document.getElementsByClassName("blob");
    let score = document.getElementsByClassName("score");

    selected = await getTeamInfo();
    for (let i = 0; i < teamName.length; i++) {
        teamName[i].getElementsByTagName("h1")[0].innerText = selected[i].tag;

        teamName[i].getElementsByTagName("img")[0].src = `${selected[i]["logoUrl"]}`;


        score[i].innerText = selected[i]["score"];
        blob[i].style.backgroundColor = `${selected[i]["color"]}33`;
        
    }
}

async function updateInfo() {
    for (let i = 0; i < players.length; i++) {
        let id = players[i].dataset.summonerId;
        let champID = players[i].dataset.championId;
        resp = await fetch(`http://localhost/draft/summoner/${i}`)
        if (resp.status != 200) {
            is_connected = false;
            return;
        }
        data = await resp.json();
        if (data["summonerId"] != id || data["championId"] != champID) {
            players[i].dataset.summonerId = data["summonerId"];
            players[i].dataset.championId = data["championId"];
            showInfo();
        }
    }
    resp = await fetch("http://localhost/draft/bans/");
    data = await resp.json();
    let bansTeam1 = document.getElementById("bansTeam1").children;
    let bansTeam2 = document.getElementById("bansTeam2").children;
    for (let i = 0; i < bansTeam1.length; i++) {
        if (data["myTeamBans"][i] != bansTeam1[i].dataset.championId) {
            bansTeam1[i].dataset.championId = data["myTeamBans"][i];
        }
        if (data["theirTeamBans"][i] != bansTeam2[i].dataset.championId) {
            bansTeam2[i].dataset.championId = data["theirTeamBans"][i];
        }
    }
    showBan();
}

async function showBan() {
    let bansTeam1 = document.getElementById("bansTeam1").children;
    let bansTeam2 = document.getElementById("bansTeam2").children;

    for (let i = 0; i < bansTeam1.length; i++) {
        let champID = bansTeam1[i].dataset.championId;
        data = await getPlayerInfo(0, champID);
        if (bansTeam1[i].style.backgroundImage != `url('${data['iconURL']}')`)
            bansTeam1[i].style.backgroundImage = `url('${data['iconURL']}')`;
    }
    for (let i = 0; i < bansTeam2.length; i++) {
        let champID = bansTeam2[i].dataset.championId;
        data = await getPlayerInfo(0, champID);
        if (bansTeam2[i].style.backgroundImage != `url('${data['iconURL']}')`)
            bansTeam2[i].style.backgroundImage = `url('${data['iconURL']}')`;
    }

}

async function getPlayerInfo(id, champID) {
    resp = await fetch(`http://localhost/draft/summoner/info/?summonerID=${id}&championID=${champID}`)
    data = await resp.json();
    return data;
}

setInterval(() => {
    if (is_connected) {
        updateInfo();
        actions();
        updateCounter();
    } else {
        setup();
        setupTeamInfo();
        checkImg();
    }
    if (document.getElementsByClassName("teamName")[0].getElementsByTagName("h1")[0].innerText == "TEAM 1") {
        setupTeamInfo();
    }
}, 1000);