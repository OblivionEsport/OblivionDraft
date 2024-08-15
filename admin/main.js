const teamNameSelectOne = document.getElementById('team1');
const teamNameSelectTwo = document.getElementById('team2');
const btnSelect = document.getElementById('btnSelect');

const matchIdInput = document.getElementById('matchId');
fetch('/admin/match/id')
    .then(resp => resp.text())
    .then(text => matchIdInput.value = text);

function changeSelect(select) {
    if (select.value === 'New') {
        emptyTeamSettings(select);
        btnSelect.disabled = true;
        return;
    }
    if (teamNameSelectOne.selectedIndex !== 0 && teamNameSelectTwo.selectedIndex !== 0 && teamNameSelectOne.selectedIndex !== teamNameSelectTwo.selectedIndex) {
        btnSelect.disabled = false;
    }
    loadTeamSettings(select, select.selectedIndex - 1);
}

fetch('/admin/teams/full')
    .then(response => response.json())
    .then(data => {
        data.forEach(team => {
            const option = document.createElement('option');
            option.value = team.name;
            option.text = team.name;
            teamNameSelectOne.appendChild(option);
            teamNameSelectTwo.appendChild(option.cloneNode(true));
        });
    });

async function saveTeamSettings(el) {
    const teamDiv = el.parentElement.parentElement.parentElement;
    const inputs = teamDiv.querySelectorAll('input');

    const newTeamData = {
        name: inputs[0].value,
        score: inputs[1].value,
        tag: inputs[2].value.toUpperCase(),
        color: inputs[3].value,
        logoUrl: "",
    };
    let response;
    try {
        response = await fetch('/admin/teams/add/', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(newTeamData)
        });
    } catch (error) {
        console.error('An error occurred:', error);
    }
    if (response.ok) {
        console.log('Team added successfully!');
        // add an option to the select
        console.log(response.status);
        if (response.status == 201) {
            // new team
            const option = document.createElement('option');
            option.value = newTeamData.name;
            option.text = newTeamData.name;
            teamNameSelectOne.appendChild(option);
            teamNameSelectTwo.appendChild(option.cloneNode(true));
            el.value = newTeamData.name;
        }
    } else {
        console.error('Failed to add team.');
    }
}

function RemoveTeamSettings(el) {
    let selectObj = el.parentElement.parentElement.children[0];
    if (selectObj.value === 'New') {
        return;
    }
    fetch('/admin/teams/delete/' + selectObj.value)
        .then(resp => {
            if (resp.ok) {
                console.log('Team deleted successfully!');
                // remove the option from the select
                if (selectObj == teamNameSelectOne) {
                    teamNameSelectTwo.removeChild(teamNameSelectTwo.children[selectObj.selectedIndex]);
                    teamNameSelectOne.removeChild(teamNameSelectOne.children[selectObj.selectedIndex]);
                } else {
                    teamNameSelectOne.removeChild(teamNameSelectOne.children[selectObj.selectedIndex]);
                    teamNameSelectTwo.removeChild(teamNameSelectTwo.children[selectObj.selectedIndex]);
                }
                emptyTeamSettings(selectObj);
            } else {
                console.error('Failed to delete team.');
            }
        });
}

async function select() {
    // POST to /admin/teams/selected with the selected team
    resp = await fetch('/admin/teams/selected', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify([teamNameSelectOne.value, teamNameSelectTwo.value])
    });
    if (resp.ok) {
        console.log('Team selected successfully!');
    } else {
        console.error('Failed to select team.');
    }
}

function emptyTeamSettings(el) {
    parent = el.parentElement.parentElement;
    if (parent.className === "team-controls") {
        parent = parent.parentElement;
    }
    const inputs = parent.querySelectorAll('input');
    inputs.forEach(input => {
        input.value = '';
    });
}

function loadTeamSettings(el, i) {
    fetch('/admin/teams/full')
        .then(response => response.json())
        .then(data => {
            savedSettings = data[i]
            const inputs = el.parentElement.parentElement.querySelectorAll('input');
            inputs.forEach(input => {
                const id = input.id;
                if (savedSettings[id]) {
                    input.value = savedSettings[id];
                }
            });
        });

}

async function loadTeamFromDB() {
    // create a popup asking for tournament name with a select box
    // fetch the tournaments from the server

    // fetch the teams from the server
    let response = await fetch('/db/tournaments');
    let tournaments = await response.json();

    // create the popup
    let popupContainer = document.createElement('div')
    popupContainer.classList = "popupContainer"
    let popup = document.createElement('div');
    popup.className = 'popup';
    let select = document.createElement('select');
    select.id = 'tournamentSelect';
    select.className = "team-name";
    let option = document.createElement('option');
    select.appendChild(option);
    tournaments.forEach(tournament => {
        let option = document.createElement('option');
        option.value = tournament.id;
        option.text = tournament.title;
        select.appendChild(option);
    });
    popup.appendChild(select);
    let button = document.createElement('button');
    button.innerText = 'Load';
    button.classList = "btn";
    button.onclick = loadTeams;
    popup.appendChild(button);
    popupContainer.append(popup)
    document.body.appendChild(popupContainer);
}

async function loadTeams() {
    // get the popup select value
    let select = document.getElementById('tournamentSelect');
    let tournamentId = select.value;
    // fetch the teams from the server
    let response = await fetch('/db/teams/' + tournamentId);
    let teams = await response.json();

    let body = [];
    // add the teams to the server
    teams.forEach((team) => {
        body.push({
            name: team.name,
            score: "0",
            tag: team.tag,
            color: "#000000",
            logoUrl: team.logo_url
        })
    });

    await fetch("/admin/teams/add/?many=true", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    }
    )

    // remove the popup
    let popup = document.getElementsByClassName('popupContainer')[0];
    document.body.removeChild(popup);
    // reload the select
    teamNameSelectOne.innerHTML = '<option value="New">New</option>';
    teamNameSelectTwo.innerHTML = '<option value="New">New</option>';
    fetch('/admin/teams/full')
        .then(response => response.json())
        .then(data => {
            data.forEach(team => {
                const option = document.createElement('option');
                option.value = team.name;
                option.text = team.name;
                teamNameSelectOne.appendChild(option);
                teamNameSelectTwo.appendChild(option.cloneNode(true));
            });
        });
}

async function updateMatchId(el) {
    let matchId = el.value;
    let resp = await fetch('/admin/match/id', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: matchId
    })
    if (resp.ok) {
        console.log('Match ID updated successfully!');
    } else {
        console.error('Failed to update match ID.');
    }
}